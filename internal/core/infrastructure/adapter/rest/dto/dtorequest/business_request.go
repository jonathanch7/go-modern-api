package dtorequest

import "time"

type BusinessRequest struct {
	// Business ID
	BusinessId string `json:"businessId,omitempty" validate:"required" minLength:"1" maxLength:"3"`
	// Business name
	Name string `json:"name,omitempty" validate:"required" minLength:"1" maxLength:"50"`
	// Detailed description of business
	Description string `json:"description,omitempty" validate:"required" minLength:"1" maxLength:"200"`
	// Register date
	RegDate *time.Time `json:"regDate,omitempty" validate:"required"`
	// Public code
	PublicCode string `json:"publicCode,omitempty" maxLength:"36"`
}
