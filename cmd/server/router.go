package main

import (
	"net/http"

	"go_webserv/internal/handlers"

	"github.com/gorilla/mux"
)

type router struct {
	handlers *handlers.Handlers
}

// corsMiddleware adds common CORS headers and handles OPTIONS preflight requests.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from the frontend origin; adjust as needed.
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		// Allow common headers plus Access-Control-Allow-Origin which some clients mistakenly send.
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Access-Control-Allow-Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			// Preflight request - return immediately.
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// New builds the router with CORS middleware applied.
func (r *router) New(handlers *handlers.Handlers) http.Handler {
	muxr := mux.NewRouter()
	muxr.HandleFunc("/", handlers.QueryAllHandler)

	muxr.HandleFunc("/blog", handlers.CreateNewBlogHandler).Methods("POST")
	muxr.HandleFunc("/blog", handlers.QueryAllBlogHandler)
	muxr.HandleFunc("/blog/{id:[0-9]+}", handlers.QueryBlogHandler)
	muxr.HandleFunc("/blog/upload-link", handlers.RequestUploadLinkHandler).Methods("POST")

	muxr.HandleFunc("/project", handlers.CreateNewProjectHandler).Methods("POST")
	muxr.HandleFunc("/project", handlers.QueryAllProjectHandler)
	muxr.HandleFunc("/project/{id:[0-9]+}", handlers.QueryProjectHandler)

	return corsMiddleware(muxr)
}
