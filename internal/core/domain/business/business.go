package business

import "time"

type Business struct {
	businessId  string
	name        string
	description string
	publicCode  string
	regDate     *time.Time
}

func NewBusiness(businessId string, name string, description string, publicCode string, regDate *time.Time) Business {
	return Business{businessId, name, description, publicCode, regDate}
}

func (e Business) BusinessId() string {
	return e.businessId
}

func (e Business) Name() string {
	return e.name
}

func (e Business) Description() string {
	return e.description
}

func (e Business) PublicCode() string {
	return e.publicCode
}

func (e Business) RegDate() *time.Time {
	return e.regDate
}
