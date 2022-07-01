package bootstrap

import (
	"github.com/jonathanch7/go-modern-api/cmd/restservice"
	"github.com/jonathanch7/go-modern-api/internal/service"
	"log"
)

type Bootstrap struct {
	rest restservice.RestService
}

func (m Bootstrap) Run() (err error) {
	m.InitServices()
	err = m.rest.RunRestService()
	return err
}

func (m *Bootstrap) InitServices() {
	log.Println("configuring applications")
	app := service.NewApplication()
	m.rest = restservice.InitRestService(app)
}
