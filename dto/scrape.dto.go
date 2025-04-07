package dto

type ScrapeRequest struct {
	PlaceID string `json:"place_id" form:"place_id" binding:"required"`
}
