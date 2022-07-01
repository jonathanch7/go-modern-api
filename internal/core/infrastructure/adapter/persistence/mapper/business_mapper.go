package mapper

import (
	"database/sql"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/business"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/entity"
	"time"
)

var Business BusinessMapper = BusinessMapper{}

type BusinessMapper struct{}

func (m BusinessMapper) ToEntity(dataFrom business.Business) entity.Business {
	return entity.Business{
		BusinessId:  dataFrom.BusinessId(),
		Name:        dataFrom.Name(),
		Description: dataFrom.Description(),
		PublicCode:  sql.NullString{String: dataFrom.PublicCode(), Valid: true},
		RegDate:     sql.NullTime{Time: *dataFrom.RegDate(), Valid: true},
	}
}

func (m BusinessMapper) ToEntities(dataFrom []business.Business) []entity.Business {
	if len(dataFrom) > 0 {
		dataTo := make([]entity.Business, len(dataFrom))
		for ix, data := range dataFrom {
			dataTo[ix] = m.ToEntity(data)
		}
		return dataTo
	}
	return nil
}

func (m BusinessMapper) ToDomain(dataFrom entity.Business) business.Business {
	// Remove zero values of mapper
	var regDate *time.Time
	if !dataFrom.RegDate.Time.IsZero() {
		regDate = &dataFrom.RegDate.Time
	}
	return business.NewBusiness(
		dataFrom.BusinessId,
		dataFrom.Name,
		dataFrom.Description,
		dataFrom.PublicCode.String,
		regDate,
	)
}

func (m BusinessMapper) ToDomains(dataFrom []entity.Business) (dataTo []business.Business) {
	if len(dataFrom) > 0 {
		dataTo := make([]business.Business, len(dataFrom))
		for ix, data := range dataFrom {
			dataTo[ix] = m.ToDomain(data)
		}
		return dataTo
	}
	return nil
}
