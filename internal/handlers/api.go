package handlers

import (
	"github.com/SantGT5/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// By setting Handler with uppercase 'H'
// this tells that this func can be imported to another packages
func Handlers(r *chi.Mux) {
	// StripSlashes will ensure that '/' wont be need
	// at the end of the url api request
	// example:
	// /account/coins -> 200 - OK
	// /account/coins/ -> 404 - OK
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
