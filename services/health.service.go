package services

import "errors"

type HealthService interface {
	CheckPing() (string, error)
	CheckResponseSuccess() (string, error)
	CheckResponseFailed() (string, error)
}

type healthService struct {
}

func NewHealthService() HealthService {
	return &healthService{}
}

func (h *healthService) CheckPing() (string, error) {
	return "the server is already running", nil
}
func (h *healthService) CheckResponseSuccess() (string, error) {
	return "this is a successful response test", nil
}

func (h *healthService) CheckResponseFailed() (string, error) {
	return "", errors.New("this is a failed response test")
}
