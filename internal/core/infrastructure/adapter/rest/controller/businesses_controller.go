package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/command"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/dto/dtorequest"
	_ "github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/dto/dtorequest"
	_ "github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/dto/dtoresponse"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/rest/mapper"
	"github.com/jonathanch7/go-modern-api/internal/shared/application"
	"github.com/jonathanch7/go-modern-api/internal/shared/infrastructure/adapter/rest/response"
	"log"
)

type BusinessesController struct {
	app application.Application
}

func NewBusinessesController(root *gin.RouterGroup, app application.Application) BusinessesController {
	ctr := BusinessesController{app}
	ctr.routing(root)
	return ctr
}

func (ctr *BusinessesController) routing(root *gin.RouterGroup) {
	route := root.Group("/businesses")
	route.POST("/", ctr.Save)
	route.GET("/", ctr.SearchAll)
}

// Create a new business
// @Summary Create a new business
// @Description Create a new business
// @Tags Businesses
// @Produce application/json
// @Accept application/json
// @Success 200 {object} dtoresponse.BusinessResponse
// @Failure 500 {object} dtoresponse.BusinessResponse
// @Failure 400 {object} dtoresponse.BusinessResponse
// @Param Business body dtorequest.BusinessRequest true "Business"
// @Router /businesses [post]
func (ctr *BusinessesController) Save(c *gin.Context) {
	dataPersist := dtorequest.BusinessRequest{}
	if err := c.ShouldBindJSON(&dataPersist); err != nil {
		response.NewResponseBuilder(c, nil).Error(err).Build()
		return
	}

	cmd := command.CreateBusinessesCommand{
		Business: mapper.Business.RequestToDomain(dataPersist),
	}
	if err := ctr.app.Core.Commands.CreateBusinesses.Handle(context.Background(), &cmd); err != nil {
		response.NewResponseBuilder(c, nil).Error(err).Build()
		return
	}

	res := mapper.Business.ToDTOResponse(cmd.Business)
	response.NewResponseBuilder(c, &res).Build()
}

// Search all businesses
// @Summary Search all businesses
// @Description Search all businesses
// @Tags Businesses
// @Produce application/json
// @Success 200 {array} dtoresponse.BusinessResponse
// @Failure 500 {object} dtoresponse.BusinessResponse
// @Failure 400 {object} dtoresponse.BusinessResponse
// @Router /businesses [get]
func (ctr *BusinessesController) SearchAll(c *gin.Context) {
	log.Printf("1 ")
	businesses, err := ctr.app.Core.Queries.SearchAllBusinesses.Handle(context.Background())
	if err != nil {
		log.Printf("2 ")
		response.NewResponseBuilder(c, nil).Error(err).Build()
		return
	}
	log.Printf("3 ")
	if len(businesses) == 0 {
		log.Printf("4 ")
		response.NewResponseBuilder(c, nil).Build()
		return
	}
	res := mapper.Business.ToDTOsResponse(businesses)
	log.Printf("businesses %v", businesses)
	log.Printf("res %v", res)
	response.NewResponseBuilder(c, res).Build()
}
