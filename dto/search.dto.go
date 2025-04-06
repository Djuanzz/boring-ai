package dto

type SearchRequest struct {
	BusinessType  string `json:"businessType" form:"businessType" binding:"required"`
	Location      string `json:"location" form:"location" binding:"required"`
	SearchOffset  int    `json:"searchOffset" form:"searchOffset"`
	NumberOfLeads int    `json:"numberOfLeads" form:"numberOfLeads" binding:"required"`
}

type PlaceDetailRequest struct {
	PlaceID string `json:"placeId" form:"placeId" binding:"required"`
}
