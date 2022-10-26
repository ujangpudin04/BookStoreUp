package routes

import (
	"nutech/handlers"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", h.FindCategories).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", h.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE")
}
