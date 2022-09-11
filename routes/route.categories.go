package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func CategoriesRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {
		r.Post("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})

		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})

		r.Get("/:id", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})

		r.Delete("/:id", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})

		r.Put("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("hello wordl"))
		})
	})
}
