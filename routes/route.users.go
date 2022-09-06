package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func UsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {

		r.Post("/", func(rw http.ResponseWriter, r *http.Request) {
			res, _ := json.Marshal(map[string]string{"message": "hello wordl"})
			rw.Write(res)

		})
	})
}
