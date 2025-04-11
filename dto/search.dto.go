package dto

type SearchRequest struct {
	BusinessType      string   `json:"businessType" form:"businessType" binding:"required"`
	Location          string   `json:"location" form:"location" binding:"required"`
	SearchOffset      int      `json:"searchOffset" form:"searchOffset"`
	NumberOfLeads     int      `json:"numberOfLeads" form:"numberOfLeads"`
	NextPageToken     string   `json:"nextPageToken" form:"nextPageToken,omitempty"`
	RemainingPlaceIds []string `json:"remainingPlaceIds" form:"remainingPlaceIds,omitempty"`
}

type PlaceDetailRequest struct {
	PlaceID string `json:"placeId" form:"placeId" binding:"required"`
}
