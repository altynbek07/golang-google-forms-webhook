package webhook

type SetWebhookRequest struct {
	Timestamp string     `json:"timestamp"`
	Responses []Response `json:"responses"`
}

type Response struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
