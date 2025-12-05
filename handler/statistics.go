package handler

import (
	"golb-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleViewBlog(c *gin.Context) {
	log := srv.getLog(c)

	blogID := c.Param("blogId")
	var body model.Statistics
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Errorf("Error binding JSON >> %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	mes, err := srv.DB.RegisterViewToDB(blogID, body)
	if err != nil {
		log.Errorf("Error incrementing blog views >> %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	message := mes
	if message == "" {
		message = "Blog view registered"
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
