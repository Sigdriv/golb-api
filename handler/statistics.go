package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleViewBlog(c *gin.Context) {
	log := srv.getLog(c)

	blogID := c.Param("blogId")

	err := srv.DB.IncrementBlogViews(blogID)
	if err != nil {
		log.Errorf("Error incrementing blog views >> %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog view incremented"})
}
