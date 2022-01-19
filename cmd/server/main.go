package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"os"
	"tech-db-forum/internal"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/server"
	"tech-db-forum/internal/pkg/utilits"
)

var (
	configPath          string
	useServerRepository bool
	runHttps            bool
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
	flag.BoolVar(&useServerRepository, "server-run", false, "true if it server run, false if it local run")
	flag.BoolVar(&runHttps, "run-https", false, "run https serve with certificates")
}

// @title Patreon
// @version 1.0
// @description Server for Patreon application.

// @tag.name user
// @tag.description "Some methods for work with user"

// @tag.name creators
// @tag.description "Some methods for work with creators"

// @tag.name attaches
// @tag.description "Some methods for work with attaches of post"

// @tag.name posts
// @tag.description "Some methods for work with posts"

// @tag.name awards
// @tag.description "Some methods for work with posts"

// @tag.name payments
// @tag.description "Some methods for work with payments"

// @tag.name comments
// @tag.description "Some methods for work with comments"

// @tag.name utilities
// @tag.description "Some methods for front work"

// @host api.pyaterochka-team.site
// @BasePath /api/v1

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	flag.Parse()
	logrus.Info(os.Args[:])

	config := internal.Config{}

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		logrus.Fatal(err)
	}

	logger, closeResource := utilits.NewLogger(&config, false, "")

	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	db, closeResource := utilits.NewPostgresConnection(config.Repository.DataBaseUrl)

	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	serv := server.New(&config,
		app.ExpectedConnections{
			SqlConnection: db,
		},
		logger,
	)

	if err = serv.Start(&config); err != nil {
		logger.Fatal(err)
	}
	logger.Info("Server was stopped")
}
