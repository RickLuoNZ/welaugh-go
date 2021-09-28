package application

import (
	"github.com/joho/godotenv"
	"github.com/rickluonz/pawsitive/cmd/api/config"
	"github.com/rickluonz/pawsitive/cmd/api/db"
	"github.com/rickluonz/pawsitive/cmd/api/handler"
	"github.com/rickluonz/pawsitive/cmd/api/server"
	"github.com/rickluonz/pawsitive/pkg/logger"
)

type Application struct {
	Cfg    *config.Config
	DB     *db.DB
	Server *server.Server
}

func New() (*Application, error) {
	cfg, err := loadAppConfiguration()
	if err != nil {
		return nil, err
	}

	// init db connection pool
	db, err := db.New(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	// api server
    handler := handler.New(db)
	server := server.New(cfg.GetApiServerPort())
	server.SetHandler("api/v1", handler)

	return &Application{
		Cfg:    cfg,
		DB:     db,
		Server: server,
	}, nil
}

// load application configuration from .env file
func loadAppConfiguration() (*config.Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Error.Println("Error loading app configuration file .env.")
		return nil, err
	}

	cfg := config.New()
	return cfg, nil
}

// start the api http server
func (app *Application) Start() error {
	return app.Server.Start()
}

// clean resources while the application stops
func (app *Application) Stop() {
	if err := app.DB.Close(); err != nil {
		logger.Error.Println("Error stop application." + err.Error())
	}
	if err := app.Server.Stop(); err != nil {
		logger.Error.Println("Error stop server." + err.Error())
	}
}
