package main

import (
	"fmt"
	"go/google-forms-webhook/internal/webhook"
	"net/http"
)

func App() *http.ServeMux {
	router := http.NewServeMux()

	// Handler
	webhook.NewWebhookHandler(router)

	return router
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
