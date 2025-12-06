package handler

import (
	"golb-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleGetAuthors(c *gin.Context) {
	log := srv.getLog(c)

	authors, err := srv.DB.GetAuthors()
	if err != nil {
		log.Errorf("Error retrieving authors >> %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(authors) == 0 {
		authors = []model.Author{}
	}

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}
