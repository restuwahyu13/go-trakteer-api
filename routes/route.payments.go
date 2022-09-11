package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func PaymentsRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {
		r.Post("/notification", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})
	})
}
