package main

import (
	"fmt"
	"go/google-forms-webhook/internal/webhook"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Handler
	webhook.NewWebhookHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
