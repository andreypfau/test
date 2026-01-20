package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-service/handlers"
	"go-service/services"
)

func main() {
	port := os.Getenv("GO_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	backendURL := os.Getenv("NODE_BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://localhost:5007"
	}

	apiClient := services.NewAPIClient(backendURL)
	pdfGen := services.NewPDFGenerator()
	studentHandler := handlers.NewStudentHandler(apiClient, pdfGen)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1/students", func(r chi.Router) {
		r.Get("/{id}/report", studentHandler.GetStudentReport)
	})

	log.Printf("Starting PDF service on :%s (backend: %s)", port, backendURL)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
