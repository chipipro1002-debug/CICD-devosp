package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Endpoint phục vụ Stage 5 kiểm tra Health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","pipeline":"6-stage-zero-trust"}`))
}

// Endpoint chính có kèm các Security Headers bắt buộc
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Zero-Trust DevSecOps Service is LIVE & SECURE!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/", rootHandler)

	log.Printf("Server đang chạy tại cổng :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server gặp lỗi: %v", err)
	}
}