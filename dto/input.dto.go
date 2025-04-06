package dto

type InputRequest struct {
	BusinessType  string `json:"businessType" form:"businessType" binding:"required"`
	Location      string `json:"location" form:"location" binding:"required"`
	NumberOfLeads int    `json:"numberOfLeads" form:"numberOfLeads" binding:"required"`
}
