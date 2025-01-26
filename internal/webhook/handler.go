package webhook

import (
	"fmt"
	"go/google-forms-webhook/pkg/req"
	"go/google-forms-webhook/pkg/res"
	"net/http"
)

func NewWebhookHandler(router *http.ServeMux) {
	router.Handle("POST /webhook", setWebhook())
}

func setWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SetWebhookRequest](w, r)

		if err != nil {
			return
		}

		fmt.Println("Received webhook request")

		for _, response := range body.Responses {
			fmt.Println("Question: ", response.Question)
			fmt.Println("Answer: ", response.Answer)
		}

		res.Json(w, "Received webhook request", http.StatusOK)
	}
}
