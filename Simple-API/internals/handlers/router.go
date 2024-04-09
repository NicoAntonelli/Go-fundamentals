package handlers

import (
	"simple-api/internals/middleware"

	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
)

// Routes handling
func Handler(router *mux.Router) {
	// Global middlewares
	router.Use(chiMiddleware.StripSlashes)

	// User module
	router.HandleFunc("/user", UserGet).Methods("GET")
	router.HandleFunc("/user", UserPost).Methods("POST")

	// Wallet balance module
	routerBalance := router.PathPrefix("/balance").Subrouter()
	routerBalance.Use(middleware.Authorization)
	routerBalance.HandleFunc("", WalletBalanceGetOne).Methods("GET")
}
