package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Endpoint phục vụ Stage 5 kiểm tra Health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok","pipeline":"6-stage-zero-trust"}`))
}

// Endpoint chính có kèm các Security Headers bắt buộc
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Zero-Trust DevSecOps Service is LIVE & SECURE!")
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

	// Khắc phục triệt để lỗi G114: Cài đặt Timeout chống DoS
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Printf("Server đang chạy tại cổng :%d...", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server gặp lỗi: %v", err)
	}
}
