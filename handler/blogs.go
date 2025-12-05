package handler

import (
	"golb-api/model"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleGetBlogs(c *gin.Context) {
	log := srv.getLog(c)

	blogs, err := srv.DB.GetBlogs()
	if err != nil {
		log.Errorf("Error retrieving blogs >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (srv *Handler) HandleGetBlog(c *gin.Context) {
	log := srv.getLog(c)

	id := c.Param("id")

	blog, err := srv.DB.GetBlog(id)
	if err != nil {
		if err.Error() == "no blog found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
			return
		}
		log.Errorf("Error retrieving blog >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (srv *Handler) HandleCreateBlog(c *gin.Context) {
	log := srv.getLog(c)

	var blog model.Blog
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		log.Errorf("Error binding JSON >> %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if blog.Title == "" || blog.Content == "" {
		log.Error("Title and content are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	blogID, err := srv.DB.CreateBlog(blog)
	if err != nil {
		log.Errorf("Error creating blog >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"blogId": blogID})
}

func (srv *Handler) HandleUploadBlogPhoto(c *gin.Context) {
	log := srv.getLog(c)

	id := c.Param("blogId")

	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("Error retrieving file >> %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Errorf("Error opening file >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer src.Close()

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		log.Errorf("Error reading file >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = srv.DB.UploadBlogPhoto(id, imageBytes)
	if err != nil {
		log.Errorf("Error uploading blog photo >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
