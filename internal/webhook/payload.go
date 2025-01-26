package webhook

type WebhookRequest struct {
	Timestamp     string         `json:"timestamp"`
	FormResponses []FormResponse `json:"form_responses"`
}

type FormResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
