package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type Segments struct{}

func (c Segments) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(segmentCtx)

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

func segmentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "segment", 1)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

////////////////////////////////////////////////////////////////
// Routes
////////////////////////////////////////////////////////////////

func (c Segments) List(w http.ResponseWriter, r *http.Request) {

}

func (c Segments) Create(w http.ResponseWriter, r *http.Request) {

}

func (c Segments) Get(w http.ResponseWriter, r *http.Request) {

}

func (c Segments) Update(w http.ResponseWriter, r *http.Request) {

}

func (c Segments) Delete(w http.ResponseWriter, r *http.Request) {

}
