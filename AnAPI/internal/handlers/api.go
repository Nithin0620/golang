package handlers

import (
	"net/http"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/avukadin/goapi/internal/middleware"
)

// creating a function with uppercase starting letter to make it public and accessible from other packages
// else if we make it start with lowercase letter it will be private and not accessible from other packages
func Handler(r *chi.Mux) {
	// Global Middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		
		//Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins",GetCoinBalance)
	})
}