package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok","pipeline":"6-stage-zero-trust"}`))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Zero-Trust DevSecOps Service is LIVE and SECURE!")
}

func main() {
	portStr := os.Getenv("PORT")
	port := 8080
	if p, err := strconv.Atoi(portStr); err == nil && p > 0 && p < 65536 {
		port = p
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/", rootHandler)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Printf("Server dang chay tai cong :%d...", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server gap loi: %v", err)
	}
}
