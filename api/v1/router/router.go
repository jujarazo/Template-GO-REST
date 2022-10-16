package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Initialize() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Force the Content-type to be JSON
		middleware.RedirectSlashes,                    // Drop trailing slashes when getting params
		middleware.Recoverer,                          // To recover
		middleware.Heartbeat("/health"),               // Sets /health endpoint for checking API health status
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"}, // Change to allow only specified origins
			AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           150,
		}),
	)

	// Create context for the requests and set the timeout to 15 secs
	router.Use(middleware.Timeout(15 * time.Second))

	router.Route("/v1", func(r chi.Router) {
		// r.Mount("/", handlers.Routes()) // Implements the routes from the handlers
	})

	return router
}
