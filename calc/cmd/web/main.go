package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/Wqescs/petPgo/calc/internal/handlers"
)

func main() {
	router := mux.NewRouter()
	
	handler := handlers.NewHTTPHandler()
	handler.RegisterRoutes(router)
	
	// Статические файлы
	fs := http.FileServer(http.Dir("./web/static"))
	router.PathPrefix("/").Handler(fs)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		log.Printf("🚀 Server started on http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	
	<-stop
	
	log.Println("🛑 Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
	
	log.Println("Server stopped gracefully")
}