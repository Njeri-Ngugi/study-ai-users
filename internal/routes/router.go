package routes

import (
	externalCfg "github.com/Njeri-Ngugi/toolbox/config"
	"github.com/gorilla/mux"
	authRoutes "users/internal/auth"
	userRoutes "users/internal/user"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() *Router {
	return &Router{Router: mux.NewRouter()}
}

func (r *Router) RegisterRoutes(cfg *externalCfg.GlobalConfig) {
	userRoutes.InitializeRoutes(cfg, (*userRoutes.Router)(r))
	authRoutes.InitializeRoutes(cfg, (*authRoutes.Router)(r))
}
