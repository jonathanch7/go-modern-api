package service

import (
	coreApp "github.com/jonathanch7/go-modern-api/internal/core/application"
	coreUseCase "github.com/jonathanch7/go-modern-api/internal/core/application/usecase"
	app "github.com/jonathanch7/go-modern-api/internal/shared/application"
	"github.com/jonathanch7/go-modern-api/internal/shared/config"
)

func NewApplication() app.Application {
	configuration := config.InstanceConfig()
	providers := InstanceProvidersApp(configuration)
	return NewApplicationWithProviders(providers)
}

func NewApplicationWithProviders(providersApp ProvidersApp) app.Application {
	return app.Application{
		Core: coreApp.Application{
			Commands: coreApp.Commands{
				CreateBusinesses: coreUseCase.NewCreateBusinessesHandler(providersApp.Core.BusinessPort),
			},
			Queries: coreApp.Queries{
				SearchAllBusinesses: coreUseCase.NewSearchAllBusinessesHandler(providersApp.Core.BusinessPort),
			},
		},
	}
}
