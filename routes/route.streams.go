package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func StreamsRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {
		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})
	})
}
