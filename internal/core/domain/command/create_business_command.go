package command

import "github.com/jonathanch7/go-modern-api/internal/core/domain/business"

type CreateBusinessesCommand struct {
	Business business.Business
}
