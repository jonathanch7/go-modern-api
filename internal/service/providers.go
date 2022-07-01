package service

import (
	"github.com/jonathanch7/go-modern-api/internal/core/domain/port"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/db"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/repository"
	"github.com/jonathanch7/go-modern-api/internal/shared/config"
)

type CoreProvidersApp struct {
	BusinessPort port.BusinessPort
}

// Provider for inject in domain
type ProvidersApp struct {
	Core CoreProvidersApp
}

// Externals Api instances
type Apis struct {
	MySQL *db.Database
}

// Instance the application providers
func InstanceProvidersApp(config config.Config) ProvidersApp {
	apisInstance := InstanceApis(config)
	return InstanceProvidersAppWithApis(config, apisInstance)
}

// Instance the application providers with the provided APIs
func InstanceProvidersAppWithApis(config config.Config, apis Apis) ProvidersApp {
	providersApp := ProvidersApp{
		Core: CoreProvidersApp{
			BusinessPort: repository.NewBusinessRepository(apis.MySQL),
		},
	}

	return providersApp
}

// Instance the external APIs of the application
func InstanceApis(config config.Config) Apis {
	apis := Apis{}
	apis.MySQL = db.ConnectToMysql(config.DB)
	return apis
}
