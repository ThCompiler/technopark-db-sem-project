package server

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/factories/handler_factory"
	"tech-db-forum/internal/app/factories/repository_factory"
)

type Server struct {
	config      *internal.Config
	logger      *log.Logger
	connections app.ExpectedConnections
}

func New(config *internal.Config, connections app.ExpectedConnections, logger *log.Logger) *Server {
	return &Server{
		config:      config,
		logger:      logger,
		connections: connections,
	}
}

func (s *Server) checkConnection() error {
	if err := s.connections.SqlConnection.Ping(); err != nil {
		return fmt.Errorf("Can't check connection to sql with error %v ", err)
	}

	s.logger.Info("Success check connection to sql db")

	return nil
}

// @title Patreon
// @version 1.0
// @description Server for Patreon application.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1/

// @x-extension-openapi {"example": "value on a json format"}

func (s *Server) Start(config *internal.Config) error {
	if err := s.checkConnection(); err != nil {
		return err
	}

	router := mux.NewRouter()
	//router.Get("/debug/pprof/<profile>", handler_interfaces.FastHTTPFunc(pprofhandler.PprofHandler).ServeHTTP)

	routerApi := router.PathPrefix("/api").Subrouter()

	repositoryFactory := repository_factory.NewRepositoryFactory(s.logger, s.connections)
	factory := handler_factory.NewFactory(s.logger, repositoryFactory)
	hs := factory.GetHandleUrls()

	//utilitsMiddleware := middleware.NewUtilitiesMiddleware(s.logger)
	//routerApi.Use(utilitsMiddleware.UpgradeLogger().ServeHTTP, utilitsMiddleware.CheckPanic().ServeHTTP)

	for apiUrl, h := range *hs {
		h.Connect(routerApi.Path(apiUrl))
	}

	//s.logger.Info("start no http server")
	return http.ListenAndServe(s.config.BindAddr, router)
}
