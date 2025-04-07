package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/Djuanzz/boring-ai/utils"
	"github.com/gin-gonic/gin"
)

type ScrapeController interface {
	GetReviews(ctx *gin.Context)
}

type scrapeController struct {
	service services.ScrapeService
}

func NewScrapeController(service services.ScrapeService) ScrapeController {
	return &scrapeController{service: service}
}

func (sc *scrapeController) GetReviews(c *gin.Context) {
	var req dto.ScrapeRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(nil, "place_id is required"))
		return
	}

	result, err := sc.service.GetReviews(req.PlaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(nil, result, nil))
}
