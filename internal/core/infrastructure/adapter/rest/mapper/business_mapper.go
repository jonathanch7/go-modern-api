package mapper

import (
	"github.com/jonathanch7/go-modern-api/internal/core/domain/business"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/dto/dtorequest"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/dto/dtoresponse"
)

var Business BusinessMapper = BusinessMapper{}

type BusinessMapper struct{}

func (m BusinessMapper) ToDTOResponse(dataFrom business.Business) dtoresponse.BusinessResponse {
	return dtoresponse.BusinessResponse{
		BusinessId:  dataFrom.BusinessId(),
		Name:        dataFrom.Name(),
		Description: dataFrom.Description(),
		PublicCode:  dataFrom.PublicCode(),
		RegDate:     dataFrom.RegDate(),
	}
}

func (m BusinessMapper) ToDTOsResponse(dataFrom []business.Business) []dtoresponse.BusinessResponse {
	if len(dataFrom) > 0 {
		dataTo := make([]dtoresponse.BusinessResponse, len(dataFrom))
		for ix, data := range dataFrom {
			dataTo[ix] = m.ToDTOResponse(data)
		}
		return dataTo
	}
	return nil
}

func (m BusinessMapper) RequestToDomain(dataFrom dtorequest.BusinessRequest) business.Business {
	return business.NewBusiness(
		dataFrom.BusinessId,
		dataFrom.Name,
		dataFrom.Description,
		dataFrom.PublicCode,
		dataFrom.RegDate,
	)
}

func (m BusinessMapper) RequestToDomains(dataFrom []dtorequest.BusinessRequest) (dataTo []business.Business) {
	if len(dataFrom) > 0 {
		dataTo := make([]business.Business, len(dataFrom))
		for ix, data := range dataFrom {
			dataTo[ix] = m.RequestToDomain(data)
		}
		return dataTo
	}
	return nil
}
