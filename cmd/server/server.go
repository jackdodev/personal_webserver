package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go_webserv/internal/handlers"

	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	conn *http.Server
	handlers *handlers.Handlers
	router *router
}

func NewServer(db *gorm.DB) *server {
	return &server{
		db: db,
		conn: &http.Server{
			Addr: ":8080",
		},
		handlers: handlers.NewHandlers(db),
	}
}

func (s *server) start() {
	r := s.router.New(s.handlers)
	s.conn.Handler = r
	port := ":8080"

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<- sigChan

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownRelease()

		if err := s.conn.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("HTTP close error: %v", err)
		}

		log.Println("Graceful shutdown complete.")
	}()

	fmt.Printf("Server is running on port %s\n", port)
	if err := s.conn.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stopped serving new connections.")
}

func (s *server) stop() error {
	s.conn.Close()
	fmt.Println("Server stopped gracefully")

	return nil
}
