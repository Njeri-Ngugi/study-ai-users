package auth

import (
	externalCfg "github.com/Njeri-Ngugi/toolbox/config"
	"github.com/gorilla/mux"
	"net/http"
	"users/internal/auth/handlers"
)

type Router struct {
	Router *mux.Router
}

func InitializeRoutes(cfg *externalCfg.GlobalConfig, r *Router) {
	authRoutes := r.Router.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", handlers.AuthLoginHandler).Methods(http.MethodPost)
}
