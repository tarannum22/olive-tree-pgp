package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

func StartHTTPServer(port int) {
	http.HandleFunc("/", handleStatusMessage) // Handle requests to the root path ("/")

	// Start the server on port 8080 (or any other available port)
	slog.Info(fmt.Sprintf("Server listening on port : %v", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		slog.Error("Error starting server:", err)
		return
	}
}

func handleStatusMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Watermelon PGP is running.")
}
