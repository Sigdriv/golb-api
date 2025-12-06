package handler

import (
	"golb-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleGetTags(c *gin.Context) {
	log := srv.getLog(c)

	tags, err := srv.DB.GetTags()
	if err != nil {
		log.Errorf("Error retrieving tags >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(tags) == 0 {
		tags = []model.Tag{}
	}

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}
