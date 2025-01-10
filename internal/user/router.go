package routes

import (
	externalCfg "github.com/Njeri-Ngugi/toolbox/config"
	"github.com/gorilla/mux"
	"net/http"
	"users/internal/user/handlers"
)

type Router struct {
	Router *mux.Router
}

func InitializeRoutes(cfg *externalCfg.GlobalConfig, r *Router) {
	userRoutes := r.Router.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/create", handlers.CreateUser).Methods(http.MethodPost)
}
