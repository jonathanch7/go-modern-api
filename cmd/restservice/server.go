package restservice

import (
	"github.com/gin-gonic/gin"
	docs "github.com/jonathanch7/go-modern-api/api/rest"
	coreCtrs "github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/controller"
	"github.com/jonathanch7/go-modern-api/internal/shared/application"
	"github.com/jonathanch7/go-modern-api/internal/shared/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

const basePath = "/api"

// @title Modern API
// @version 1.0.0
// @description The purpose of this application is be modern API.
// @host localhost:8080
// @BasePath /api
type RestService struct {
	appRest     *gin.Engine
	controllers Controllers
}

type Controllers struct {
	Businesses coreCtrs.BusinessesController
}

func InitRestService(app application.Application) RestService {
	log.Println("initializing REST service")
	appRest := gin.Default()

	apiRoot := appRest.Group(basePath)

	log.Println("configuring controllers")

	ctrs := Controllers{
		Businesses: coreCtrs.NewBusinessesController(apiRoot, app),
	}
	swaggerRoute(apiRoot)

	return RestService{
		appRest:     appRest,
		controllers: ctrs,
	}
}

// Run the rest service
func (r *RestService) RunRestService() error {
	configEnv := config.InstanceConfig()
	log.Println("configuring port ", configEnv.ListenPort)
	return r.appRest.Run(":" + configEnv.ListenPort)
}

func swaggerRoute(apiRoot *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = basePath
	apiRoot.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
