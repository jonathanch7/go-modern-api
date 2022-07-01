package usecase

import (
	"context"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/command"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/port"
	"github.com/jonathanch7/go-modern-api/internal/shared/logger"
)

type CreateBusinessesHandler interface {
	Handle(ctx context.Context, cmd *command.CreateBusinessesCommand) error
}

type createBusinessesHandler struct {
	repository port.BusinessPort
}

func NewCreateBusinessesHandler(businessRepository port.BusinessPort) CreateBusinessesHandler {
	if businessRepository == nil {
		panic("BusinessPort is nil")
	}
	return createBusinessesHandler{businessRepository}
}

func (h createBusinessesHandler) Handle(ctx context.Context, cmd *command.CreateBusinessesCommand) (err error) {
	defer func() { logger.Command("CreateBusinessesHandler", cmd, err) }()

	cmd.Business, err = h.repository.SaveBusinesses(ctx, cmd.Business)
	return
}
