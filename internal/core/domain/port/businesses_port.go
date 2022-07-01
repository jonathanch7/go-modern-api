package port

import (
	"context"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/business"
)

type BusinessPort interface {
	SaveBusinesses(ctx context.Context, business business.Business) (business.Business, error)

	PartialUpdateBusiness(ctx context.Context, business business.Business) (business.Business, error)

	RemoveBusiness(ctx context.Context, business business.Business) (business.Business, error)

	SearchAllBusinesses(ctx context.Context) ([]business.Business, error)
}
