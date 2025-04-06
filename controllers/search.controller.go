package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/Djuanzz/boring-ai/utils"
	"github.com/gin-gonic/gin"
)

type SearchController interface {
	HandleSearch(ctx *gin.Context)
	PlaceDetail(ctx *gin.Context)
}

type searchController struct {
	searchService services.SearchService
}

func NewSearchController(service services.SearchService) SearchController {
	return &searchController{
		searchService: service,
	}
}

func (sc *searchController) HandleSearch(c *gin.Context) {
	var req dto.SearchRequest
	if err := c.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed(nil, "invalid request body: "+err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := sc.searchService.SearchBusiness(req)
	if err != nil {
		res := utils.ResponseFailed(nil, err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	state := result["state"].(map[string]any)
	next := result["next"].(*utils.NextTask)

	res := utils.ResponseSuccess(state, nil, next)
	c.JSON(http.StatusOK, res)
}

func (sc *searchController) PlaceDetail(c *gin.Context) {
	var req dto.PlaceDetailRequest
	if err := c.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed(nil, "invalid request body: "+err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := sc.searchService.GetPlaceDetail(req.PlaceID)
	if err != nil {
		res := utils.ResponseFailed(nil, err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(nil, result, nil)
	c.JSON(http.StatusOK, res)
}
