package db

import (
	"fmt"
	"golb-api/model"
)

func (db *DB) GetBlogs() (blogs []model.Blog, err error) {
	query := `
	select b.id, b.title, b.content, b.created_at, t."name" as "tagName", a."name" as "authorName", count(s.id) as "views"
	from blogs b 
	left join blog_tags bt on b.id = bt.blog_id
	left join tag t on t.id = bt.tag_id
	left join author a on b.author_id = a.id
	left join "statistics" s on s.blog_id = b.id
	group by b.id, b.title, b.content, b.created_at, t.id, t."name", a."name"
	`

	exisitngBlogs, err := Query[Blog](db, query, map[string]any{})
	if err != nil {
		err = fmt.Errorf("error getting blogs >> %w", err)
		return
	}

	blogs = mapBlog(exisitngBlogs)

	return
}

func (db *DB) GetBlog(id string) (blog model.Blog, err error) {
	query := `
	select b.id, b.title, b.content, b.created_at, t."name" as "tagName", a."name" as "authorName", count(s.id) as "views"
	from blogs b 
	left join blog_tags bt on b.id = bt.blog_id
	left join tag t on t.id = bt.tag_id
	left join author a on b.author_id = a.id
	left join "statistics" s on s.blog_id = b.id
	where b.id = :id
	group by b.id, b.title, b.content, b.created_at, t.id, t."name", a."name"
	`

	args := map[string]any{
		"id": id,
	}

	blogs, err := Query[Blog](db, query, args)
	if err != nil {
		err = fmt.Errorf("error getting blog >> %w", err)
		return
	}

	if len(blogs) == 0 {
		err = fmt.Errorf("no blog found")
		return
	}

	mappedBlogs := mapBlog(blogs)

	return mappedBlogs[0], nil
}

func mapBlog(blogs []Blog) []model.Blog {
	var newBlogs []model.Blog
	for _, blog := range blogs {
		found := false
		for i := range newBlogs {
			if newBlogs[i].ID == blog.ID {
				newBlogs[i].Tags = append(newBlogs[i].Tags, blog.Tag)
				found = true
				break
			}
		}

		if !found {
			newBlog := model.Blog{
				ID:        blog.ID,
				Title:     blog.Title,
				Content:   blog.Content,
				Author:    blog.Author,
				CreatedAt: blog.CreatedAt,
				Tags:      []string{blog.Tag},
				Views:     blog.Views,
			}
			newBlogs = append(newBlogs, newBlog)
		}
	}

	return newBlogs
}

func (db *DB) CreateBlog(blog model.Blog) (err error) {
	var authorID string
	var tagIDs []string

	authorID, err = getAuthorID(db, blog.Author)
	if err != nil {
		err = fmt.Errorf("error getting author ID >> %w", err)
		return
	}

	tagIDs, err = getTagIDs(db, blog.Tags)
	if err != nil {
		err = fmt.Errorf("error getting tag IDs >> %w", err)
		return
	}

	blogID, err := insertBlog(db, blog, authorID, tagIDs)
	if err != nil {
		err = fmt.Errorf("error inserting blog >> %w", err)
		return
	}

	err = createStatistics(db, blogID)
	if err != nil {
		err = fmt.Errorf("error creating statistics >> %w", err)
		return
	}

	return
}

func getAuthorID(db *DB, authorName string) (authorID string, err error) {
	authorQuery := `SELECT id FROM author WHERE name = :authorName`
	authorArgs := map[string]any{
		"authorName": authorName,
	}
	author, err := Query[Author](db, authorQuery, authorArgs)
	if err != nil {
		err = fmt.Errorf("error querying author ID >> %w", err)
		return
	}

	if len(author) == 0 {
		insertAuthorQuery := `INSERT INTO author (name) VALUES (:authorName) RETURNING id`
		newAuthorID, errExec := Exec(db, insertAuthorQuery, authorArgs)
		if errExec != nil {
			errExec = fmt.Errorf("error inserting author >> %w", errExec)
			return "", errExec
		}

		authorID = newAuthorID
	}

	if len(author) > 0 {
		authorID = author[0].ID
	}

	return
}

func getTagIDs(db *DB, tags []string) (tagIDs []string, err error) {
	var insertingTags []string

	tagsQuery := `SELECT id, name FROM tag WHERE name IN (:tagNames)`
	tagsArgs := map[string]any{
		"tagNames": tags,
	}

	query, args := In(tagsQuery, tagsArgs)
	existingTags, err := Query[Tag](db, query, args)
	if err != nil {
		err = fmt.Errorf("error querying tags >> %w", err)
		return
	}

	existingTagNames := make(map[string]string)
	for _, tag := range existingTags {
		existingTagNames[tag.Name] = tag.ID
		tagIDs = append(tagIDs, tag.ID)
	}

	for _, tagName := range tags {
		_, exists := existingTagNames[tagName]
		if !exists {
			insertingTags = append(insertingTags, tagName)
		}
	}

	if len(insertingTags) > 0 {
		for _, tagName := range insertingTags {
			insertagQuery := `INSERT INTO tag (name) VALUES (:tagName) RETURNING id`
			tagArgs := map[string]any{
				"tagName": tagName,
			}
			tagID, errExec := Exec(db, insertagQuery, tagArgs)
			if errExec != nil {
				errExec = fmt.Errorf("error inserting tag >> %w", errExec)
				return nil, errExec
			}

			tagIDs = append(tagIDs, tagID)
		}
	}

	return
}

func insertBlog(db *DB, blog model.Blog, authorID string, tagIDs []string) (blogID string, err error) {
	insertBlogQuery := `
	INSERT INTO blogs (title, content, author_id)
	VALUES (:title, :content, :authorID)
	RETURNING id
	`
	blogArgs := map[string]any{
		"title":    blog.Title,
		"content":  blog.Content,
		"authorID": authorID,
	}
	blogID, err = Exec(db, insertBlogQuery, blogArgs)
	if err != nil {
		err = fmt.Errorf("error inserting blog >> %w", err)
		return
	}

	for _, tagID := range tagIDs {
		insertBlogTagQuery := `
		INSERT INTO blog_tags (blog_id, tag_id)
		VALUES (:blogID, :tagID)
		ON CONFLICT DO NOTHING;
		`
		blogTagArgs := map[string]any{
			"blogID": blogID,
			"tagID":  tagID,
		}
		_, err = Exec(db, insertBlogTagQuery, blogTagArgs)
		if err != nil {
			err = fmt.Errorf("error inserting blog_tag >> %w", err)
			return
		}
	}

	return
}
