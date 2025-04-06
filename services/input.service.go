package services

import (
	"github.com/Djuanzz/boring-ai/dto"
)

type InputService interface {
	ProcessInput(req dto.InputRequest) (map[string]any, *string, map[string]any)
}

type inputService struct{}

func NewInputService() InputService {
	return &inputService{}
}

func (s *inputService) ProcessInput(req dto.InputRequest) (map[string]any, *string, map[string]any) {
	state := map[string]any{
		"businessType":      req.BusinessType,
		"location":          req.Location,
		"numberOfLeads":     req.NumberOfLeads,
		"leadCount":         0,
		"searchOffset":      0,
		"remainingPlaceIds": []string{},
	}

	nextKey := "search"
	payload := map[string]any{
		"businessType":  "$state.businessType",
		"location":      "$state.location",
		"searchOffset":  "$state.searchOffset",
		"numberOfLeads": "$state.numberOfLeads",
	}

	return state, &nextKey, payload
}
