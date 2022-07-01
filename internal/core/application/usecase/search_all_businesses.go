package usecase

import (
	"context"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/business"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/port"
)

type SearchAllBusinessesHandler interface {
	Handle(ctx context.Context) ([]business.Business, error)
}

type searchBusinessesHandler struct {
	repository port.BusinessPort
}

func NewSearchAllBusinessesHandler(businessRepository port.BusinessPort) SearchAllBusinessesHandler {
	if businessRepository == nil {
		panic("BusinessPort is nil")
	}
	return searchBusinessesHandler{businessRepository}
}

func (h searchBusinessesHandler) Handle(ctx context.Context) ([]business.Business, error) {
	return h.repository.SearchAllBusinesses(ctx)
}
