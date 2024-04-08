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
	router.Use(middleware.Authorization)

	// Account module
	router.HandleFunc("/balance", GetWalletBalance).Methods("GET")
}
