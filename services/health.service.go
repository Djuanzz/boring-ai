package services

type HealthService interface {
	CheckPing() (string, error)
}

type healthService struct {
}

func NewHealthService() HealthService {
	return &healthService{}
}

func (h *healthService) CheckPing() (string, error) {
	return "Server is healthy", nil
}
