package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profile", middleware.Auth(h.GetProfile)).Methods("GET")
}
