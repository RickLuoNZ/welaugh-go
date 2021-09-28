package server

import (
	"context"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rickluonz/pawsitive/cmd/api/handler"
	"github.com/rickluonz/pawsitive/pkg/logger"
)

type Server struct {
	Instance *echo.Echo
	Port     string
}

func New(port string) *Server {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Static("/static", "assets")
	e.Use(middleware.Logger())
	// we add our own App-Client-ID here for cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "App-Client-ID"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return &Server{Instance: e, Port: port}
}

func (srv *Server) SetHandler(route string, handler *handler.Handler) {
	g := srv.Instance.Group(route)
    handler.Register(g)
}

func (srv *Server) Start() error {
	logger.Info.Printf("starting server at %s", srv.Port)
	return srv.Instance.Start(srv.Port)
}

func (srv *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Instance.Shutdown(ctx)
}
