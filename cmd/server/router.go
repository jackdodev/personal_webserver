package main

import (
	"go_webserv/internal/handlers"

	"github.com/gorilla/mux"
)

type router struct {
	handlers *handlers.Handlers
}

func (r *router) New(handlers *handlers.Handlers) *mux.Router{
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.QueryAllHandler)

	mux.HandleFunc("/blog", handlers.CreateNewBlogHandler).Methods("POST")
	mux.HandleFunc("/blog", handlers.QueryAllBlogHandler)
	mux.HandleFunc("/blog/{id:[0-9]+}", handlers.QueryBlogHandler)

	mux.HandleFunc("/project", handlers.CreateNewProjectHandler).Methods("POST")
	mux.HandleFunc("/project", handlers.QueryAllProjectHandler)
	mux.HandleFunc("/project/{id:[0-9]+}", handlers.QueryProjectHandler)

	return mux
}
