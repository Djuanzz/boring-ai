package dto

type SearchRequest struct {
	BusinessType      string   `json:"businessType"`
	Location          string   `json:"location"`
	SearchOffset      int      `json:"searchOffset"`
	NumberOfLeads     int      `json:"numberOfLeads"`
	NextPageToken     string   `json:"nextPageToken,omitempty"`
	RemainingPlaceIds []string `json:"remainingPlaceIds,omitempty"`
}

type PlaceDetailRequest struct {
	PlaceID string `json:"placeId" form:"placeId" binding:"required"`
}
