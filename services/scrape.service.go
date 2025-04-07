package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ScrapeService interface {
	GetReviews(placeID string) (map[string]any, error)
}

type scrapeService struct {
	apiKey string
}

func NewScrapeService(apiKey string) ScrapeService {
	return &scrapeService{apiKey: apiKey}
}

func (s *scrapeService) GetReviews(placeID string) (map[string]any, error) {
	baseURL := "https://www.searchapi.io/api/v1/search"
	queryParams := url.Values{
		"engine":   {"google_maps_reviews"},
		"place_id": {placeID},
		"api_key":  {s.apiKey},
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request search API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return result, nil
}
