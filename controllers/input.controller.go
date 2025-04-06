package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/Djuanzz/boring-ai/utils"
	"github.com/gin-gonic/gin"
)

type InputController interface {
	HandleInput(ctx *gin.Context)
}

type inputController struct {
	inputService services.InputService
}

func NewInputController(service services.InputService) InputController {
	return &inputController{
		inputService: service,
	}
}

func (c *inputController) HandleInput(ctx *gin.Context) {
	var req dto.InputRequest

	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed(nil, "invalid input: "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	state, nextKey, payload := c.inputService.ProcessInput(req)

	res := utils.ResponseSuccess(
		state,
		nil,
		&utils.NextTask{
			Key:     *nextKey,
			Payload: payload,
		},
	)

	ctx.JSON(http.StatusOK, res)
}
