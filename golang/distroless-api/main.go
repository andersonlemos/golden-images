package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type RootResponse struct {
	Message     string `json:"message"`
	GoVersion   string `json:"goVersion"`
	Environment string `json:"environment"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8080")
	env := getEnv("APP_ENV", "undefined")

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(HealthResponse{
			Status:      "ok",
			Environment: env,
		})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Not Found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(RootResponse{
			Message:     "Hello from Go Golden Image",
			GoVersion:   runtime.Version(),
			Environment: env,
		})
	})

	serverAddr := ":" + port
	log.Printf("API running on port %s", port)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}

