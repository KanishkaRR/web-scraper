package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type scraperHandler struct {
}

func New() *scraperHandler {
	return &scraperHandler{}
}

func (h *scraperHandler) RegisterRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1/")
	apiV1.GET("analyse", h.getAnalysisInformation)
}

func (h *scraperHandler) getAnalysisInformation(c *gin.Context) {
	log.Infof("Retrieving analysed Informations")
	//res, err := h.processor.GetProcessResults(c)
	// if err != nil {
	// 	logrus.Error(err)
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }

	// c.JSON(http.StatusOK, res)
}
