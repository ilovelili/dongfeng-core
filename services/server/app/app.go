package app

import (
	api "github.com/ilovelili/dongfeng-core/services/proto"
	handlers "github.com/ilovelili/dongfeng-core/services/server/handlers"

	"github.com/ilovelili/dongfeng-core/services/utils"
	micro "github.com/micro/go-micro"	
)

// App app. They call me God Object so I guess I am cool
type App struct {
	Service micro.Service
}

// Bootstarp Bootstarp the service
func (app *App) Bootstarp() error {
	myapp, err := app.init()
	if err != nil {
		return err
	}

	return myapp.Service.Run()
}

// init init the app
func (app *App) init() (application *App, err error) {
	if application, err = app.initializeWebService(); err != nil {
		return app, err
	}

	return application, err
}

// initializeProxyService init reverse proxy service with router
func (app *App) initializeWebService() (*App, error) {
	config := utils.GetConfig()
	service := micro.NewService(
		micro.Name(config.ServiceNames.CoreServer),
		micro.RegisterTTL(config.ServiceMeta.GetRegistryTTL()),
		micro.RegisterInterval(config.ServiceMeta.GetRegistryHeartbeat()),
		micro.Version(config.ServiceMeta.GetVersion()),
	)

	service.Init()
	app.Service = service

	// Register a facade handler
	api.RegisterApiHandler(app.Service.Server(), handlers.NewFacade())

	return app, nil
}
