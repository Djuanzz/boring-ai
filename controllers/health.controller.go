package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/Djuanzz/boring-ai/utils"
	"github.com/gin-gonic/gin"
)

type HealthController interface {
	CheckPing(ctx *gin.Context)
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
	ping, err := h.healthService.CheckPing()

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_PING_FAILED, err)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_PING_SUCCESS, ping)
	ctx.JSON(http.StatusOK, res)
}
