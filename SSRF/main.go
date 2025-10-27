package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// fetchURLHandler is the vulnerable handler.
// It reads the 'url' query parameter and makes a GET request to it.
func fetchURLHandler(w http.ResponseWriter, r *http.Request) {
	urlToFetch := r.URL.Query().Get("url")

	if urlToFetch == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	// VULNERABLE PART: The application makes a request to an arbitrary URL
	// provided by the user without any validation.
	// An attacker can use this to probe internal services, query cloud
	// metadata endpoints, or scan the local network.
	log.Printf("Attempting to fetch: %s", urlToFetch)

	resp, err := http.Get(urlToFetch)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching URL: %s", err.Error()), http.StatusInternalServerError)
		log.Printf("Error fetching URL: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/fetch", fetchURLHandler)

	// Start the web server
	port := "8080"
	log.Printf("Starting vulnerable SSRF server on :%s", port)
	log.Println("Test with: http://localhost:8080/fetch?url=https://api.example.com")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
