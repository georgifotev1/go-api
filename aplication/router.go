package aplication

import (
	"net/http"

	"github.com/georgifotev1/go-api/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (a *App) newRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/auth", a.loadUsersRoutes)

	a.router = r
}

func (a *App) loadUsersRoutes(r chi.Router) {
	authHandler := &handlers.User{
		Storage: a.db,
	}

	r.Post("/register", a.isGuest(authHandler.Register))
	r.Post("/signin", a.isGuest(authHandler.SignIn))
	r.Get("/signout/{userId}", a.isUser(authHandler.SignOut))
}
