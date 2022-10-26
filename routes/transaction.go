package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST") // Notification from midtrans route ...
}
