package services

import (
	"context"
	"fmt"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/utils"
	"googlemaps.github.io/maps"
)

type SearchService interface {
	SearchBusiness(req dto.SearchRequest) (map[string]any, error)
	GetPlaceDetail(placeID string) (map[string]any, error)
}

type searchService struct {
	Client *maps.Client
}

func NewSearchService(client *maps.Client) SearchService {
	return &searchService{Client: client}
}

func (s *searchService) SearchBusiness(req dto.SearchRequest) (map[string]any, error) {
	ctx := context.Background()

	query := req.BusinessType + " in " + req.Location
	textRequest := &maps.TextSearchRequest{
		Query:  query,
		Radius: 10000,
	}

	resp, err := s.Client.TextSearch(ctx, textRequest)
	if err != nil || len(resp.Results) == 0 {
		return nil, fmt.Errorf("no results found or error occurred")
	}

	start := req.SearchOffset
	end := start + req.NumberOfLeads
	if end > len(resp.Results) {
		end = len(resp.Results)
	}

	placeIDs := []string{}
	for i := start; i < end; i++ {
		placeIDs = append(placeIDs, resp.Results[i].PlaceID)
	}

	if len(placeIDs) == 0 {
		return nil, fmt.Errorf("no place IDs to process")
	}

	state := map[string]any{
		"nextPageToken":     resp.NextPageToken,
		"remainingPlaceIds": placeIDs[1:],
		"searchOffset":      end,
	}

	next := &utils.NextTask{
		Key: "scrape",
		Payload: map[string]any{
			"placeId": placeIDs[0],
		},
	}

	result := map[string]any{
		"state": state,
		"next":  next,
	}
	return result, nil
}

func (s *searchService) GetPlaceDetail(placeID string) (map[string]any, error) {
	ctx := context.Background()

	detailReq := &maps.PlaceDetailsRequest{
		PlaceID: placeID,
		Fields: []maps.PlaceDetailsFieldMask{
			maps.PlaceDetailsFieldMaskName,
			maps.PlaceDetailsFieldMaskFormattedAddress,
			maps.PlaceDetailsFieldMaskFormattedPhoneNumber,
			maps.PlaceDetailsFieldMaskWebsite,
			maps.PlaceDetailsFieldMaskGeometry,
			maps.PlaceDetailsFieldMaskRatings,
			maps.PlaceDetailsFieldMaskUserRatingsTotal,
			maps.PlaceDetailsFieldMaskOpeningHours,
			maps.PlaceDetailsFieldMaskTypes,
		},
	}

	detailResp, err := s.Client.PlaceDetails(ctx, detailReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get place details: %v", err)
	}

	result := map[string]any{
		"placeId":      detailResp.PlaceID,
		"name":         detailResp.Name,
		"address":      detailResp.FormattedAddress,
		"phoneNumber":  detailResp.FormattedPhoneNumber,
		"website":      detailResp.Website,
		"location":     detailResp.Geometry.Location,
		"rating":       detailResp.Rating,
		"userRatings":  detailResp.UserRatingsTotal,
		"openingHours": detailResp.OpeningHours,
		"types":        detailResp.Types,
	}

	return result, nil
}
