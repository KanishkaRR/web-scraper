package handler

import (
	"net/http"

	dto "github.com/KanishkaRR/web-scraper/dtos"
	processer "github.com/KanishkaRR/web-scraper/processes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/")
	api.POST("getUrlInformation", getAnalysisInformation)
}

func getAnalysisInformation(c *gin.Context) {
	req := &dto.AnalysesRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req.WebUrl == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	process := processer.NewProcessor()
	res, err := process.ProcessWebPage(req.WebUrl)
	if err != nil {
		logrus.Errorf("error while trying to process the page, %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, res)
}
