package application

import (
	"github.com/jonathanch7/go-modern-api/internal/core/application/usecase"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateBusinesses usecase.CreateBusinessesHandler
}

type Queries struct {
	SearchAllBusinesses usecase.SearchAllBusinessesHandler
}
