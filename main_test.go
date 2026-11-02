package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// Test mô phỏng 50 luồng truy cập đồng thời để thỏa mãn cờ 'go test -race'
func TestConcurrentHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w := httptest.NewRecorder()
			healthHandler(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("Kỳ vọng mã 200, thực tế nhận %d", w.Code)
			}
		}()
	}
	wg.Wait()
}