package main

import (
	"fmt"
	"github.com/Njeri-Ngugi/toolbox/config"
	"github.com/Njeri-Ngugi/toolbox/postgres"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"users/internal/routes"
)

type Server struct {
	Configuration *config.GlobalConfig
	Router        *routes.Router
}

// NewServer ...
func NewServer(config *config.GlobalConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        routes.NewRouter(),
	}

	return server
}

func main() {
	// fetch env variables
	localCfg, err := config.FromEnv()
	if err != nil {
		logrus.Error(err)
		return
	}

	// connect DB
	err = postgres.ConnectDB(localCfg.PostgresDSN)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Infoln("Connected to database")

	// Continue with other setups like routes
	server := NewServer(localCfg)
	server.Router.RegisterRoutes(localCfg)

	// add cors configuration
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"tenant", "*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "OPTIONS", "DELETE", "PATCH"},
	})

	var handler http.Handler
	handler = c.Handler(server.Router.Router)

	logrus.Infoln("Starting server on Port", localCfg.Port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", "", localCfg.Port),
		handler)
	if err != nil {
		logrus.Error(err)
		return
	}

}
