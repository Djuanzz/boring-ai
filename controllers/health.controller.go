package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/services"
	"github.com/Djuanzz/boring-ai/utils"
	"github.com/gin-gonic/gin"
)

type HealthController interface {
	CheckPing(ctx *gin.Context)
	CheckResponseSuccess(ctx *gin.Context)
	CheckResponseFailed(ctx *gin.Context)
}

type healthController struct {
	healthService services.HealthService
}

func NewHealthController(hs services.HealthService) HealthController {
	return &healthController{
		healthService: hs,
	}
}

func (h *healthController) CheckPing(ctx *gin.Context) {
	pingMsg, err := h.healthService.CheckPing()

	if err != nil {
		res := utils.ResponseFailed(nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	result := map[string]any{
		"message": pingMsg,
	}

	res := utils.ResponseSuccess(nil, result, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *healthController) CheckResponseSuccess(ctx *gin.Context) {
	pingMsg, err := h.healthService.CheckResponseSuccess()

	if err != nil {
		res := utils.ResponseFailed(nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	result := map[string]any{
		"message": pingMsg,
	}

	res := utils.ResponseSuccess(nil, result, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *healthController) CheckResponseFailed(ctx *gin.Context) {
	pingMsg, err := h.healthService.CheckResponseFailed()

	if err != nil {
		res := utils.ResponseFailed(nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	result := map[string]any{
		"message": pingMsg,
	}

	res := utils.ResponseSuccess(nil, result, nil)
	ctx.JSON(http.StatusOK, res)
}
