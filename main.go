package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type LogEntry struct {
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	ResourceID string    `json:"resourceID"`
	Timestamp  time.Time `json:"timestamp"`
	TraceID    string    `json:"traceID"`
	SpanID     string    `json:"spanID"`
	Commit     string    `json:"commit"`
	Metadata   struct {
		ParentResourceID string `json:"parentResourceId"`
	} `json:"metadata"`
}

func main() {
	// Create logs directory if it doesn't exist
	err := os.Mkdir("logs", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Error creating logs directory: %s", err)
	}

	logFile, err := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err)
	}
	//defer logFile.Close()

	log.SetOutput(logFile)

	http.HandleFunc("/api", handleAPIRequest)
	http.HandleFunc("/", handleDefault)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	logEntry := LogEntry{
		Level:      "error",
		Message:    "Failed to connect to DB",
		ResourceID: "server-1234",
		Timestamp:  time.Now().UTC(),
		TraceID:    "abc-xyz-123",
		SpanID:     "span-456",
		Commit:     "5e5342f",
	}
	logEntry.Metadata.ParentResourceID = "server-0987"
	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		log.Println("Error marshaling log entry:", err)
		return
	}
	log.Println(string(logJSON))
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "response sent")
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the application!")
}
