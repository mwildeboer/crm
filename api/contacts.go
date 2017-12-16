package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type Contacts struct{}

func (c Contacts) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(contactCtx)

	r.Get("/", c.List)
	r.Post("/", c.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", c.Get)
		r.Put("/", c.Update)
		r.Delete("/", c.Delete)
	})

	return r
}

////////////////////////////////////////////////////////////////
// Middleware
////////////////////////////////////////////////////////////////

func contactCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "segment", 1)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

////////////////////////////////////////////////////////////////
// Routes
////////////////////////////////////////////////////////////////

func (c Contacts) List(w http.ResponseWriter, r *http.Request) {

}

func (c Contacts) Create(w http.ResponseWriter, r *http.Request) {

}

func (c Contacts) Get(w http.ResponseWriter, r *http.Request) {

}

func (c Contacts) Update(w http.ResponseWriter, r *http.Request) {

}

func (c Contacts) Delete(w http.ResponseWriter, r *http.Request) {

}
