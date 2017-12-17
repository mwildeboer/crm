package main

import (
	"crm/api"
	"crm/views"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"},
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)

	router.Route("/v1", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.Render(w, r, views.NewRootResponse(router))
		})

		router.Group(func(router chi.Router) {
			router.Mount("/contacts", api.Contacts{}.Routes())
			router.Mount("/segments", api.Segments{}.Routes())
			router.Mount("/segments/{segmentId}/contacts", api.Contacts{}.Routes())
		})
	})
	http.ListenAndServe(":3000", router)
}
