package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	defer logFile.Close()

	log.SetOutput(logFile)

	http.HandleFunc("/api", handleAPIRequest)
	http.HandleFunc("/", handleDefault)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("API request received")
	// Simulating an error
	log.Println("Error occurred: Something went wrong")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error occurred")
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the application!")
}
