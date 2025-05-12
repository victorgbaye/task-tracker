package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)

type application struct{
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1/task", func(r chi.Router) {
		r.Post("/", app.createTask)
		r.Put("/", app.updateTask)
		r.Delete("/", app.deleteTask)
		r.Get("/", app.getAllTask)
		r.Get("/{taskID}", app.getTask)
	})



	return r
}

func (app *application) run(mux *chi.Mux)  {
	srv:= http.Server{
		Addr: app.addr,
		Handler: mux,

	}
	srv.ListenAndServe()
}