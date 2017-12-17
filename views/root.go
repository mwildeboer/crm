package views

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type RootResponse struct {
	Version string   `json:"version"`
	Routes  []string `json:"routes"`
}

func (rd *RootResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewRootResponse(router chi.Router) render.Renderer {
	endpoints := []string{}
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		endpoints = append(endpoints, fmt.Sprintf("%s: %s", method, route))
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return &RootResponse{
		Version: "1.0",
		Routes:  endpoints,
	}
}
