package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWatchdogTableDriven(t *testing.T) {
	// Testcase
	tests := []struct {
		name string // Name test
		method string // Method
		mockStatusCode int	// Response status
		expectedResult int	// Result status
	}{
		{"Status OK 200", http.MethodPost, 200, 200},
		{"Not Found 404", http.MethodGet, 404, 404},
		{"Internal Server Error 500", http.MethodGet, 500, 500},
	}

	// Run Testcase
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create mock server
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.mockStatusCode);
			}))
			defer mockServer.Close()
			
			// Clear and set data
			result = make(map[string]int);
			target := Target{
				Url: mockServer.URL,
				Method: test.method,
			}
			
			// Run watchdog
			wg.Add(1)
			watchdog(target)
			wg.Wait()
			
			key := test.method + " " + mockServer.URL

			if result[key] != test.mockStatusCode {
				t.Errorf("Expected %d, got %d", test.expectedResult, result[key])
			}
		})
	}
}