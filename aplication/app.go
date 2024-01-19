package aplication

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/georgifotev1/go-api/database/sqlc"
)

type App struct {
	db     *sqlc.Queries
	router http.Handler
}

func New(conn *sql.DB) *App {
	return &App{
		db:     sqlc.New(conn),
		router: newRouter(),
	}
}

func (a *App) Start() error {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler: a.router,
	}

	fmt.Println("Starting server")
	return server.ListenAndServe()
}
