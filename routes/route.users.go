package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func UsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {

		r.Post("/login", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Post("/register", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Get("/activation", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Post("/forgot-password", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Put("/reset-password", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Put("/change-password", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Get("/profile/:id", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})

		r.Put("/profile/:id", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})
	})
}
