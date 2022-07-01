package dtoresponse

import "time"

type BusinessResponse struct {
	// Business ID
	BusinessId string `json:"businessId,omitempty" maxLength:"3"`
	// Business name
	Name string `json:"name,omitempty" maxLength:"50"`
	// Detailed description of business
	Description string `json:"description,omitempty" maxLength:"200"`
	// Register date
	RegDate *time.Time `json:"regDate,omitempty" validate:"required"`
	// Public code
	PublicCode string `json:"publicCode,omitempty" maxLength:"36"`
}
