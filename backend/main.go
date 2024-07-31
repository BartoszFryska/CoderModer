package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BartoszFryska/CoderModer/handlers"
)

func main() {
	http.HandleFunc("/submit", handlers.SubmitJobHandler)
	http.HandleFunc("/status", handlers.JobStatusHandler)
	http.HandleFunc("/list", handlers.ListJobsHandler)
	http.HandleFunc("/cancel", handlers.CancelJobHandler)
	http.HandleFunc("/execute", handlers.ExecuteJobHandler)
	http.HandleFunc("/health", handlers.HealthCheckHandler)

	http.HandleFunc("/ws", handlers.WebSocketHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
