package webhook

import (
	"fmt"
	"go/google-forms-webhook/pkg/res"
	"net/http"
	"os"
)

func NewWebhookHandler(router *http.ServeMux) {
	router.Handle("POST /webhook", setWebhook())
}

func setWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := "example.txt"

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
			return
		}
		defer file.Close()

		data := "Привет, это текст, записанный в файл!\n"

		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Ошибка при записи данных:", err)
			return
		}

		fmt.Println("Данные успешно записаны в файл:", fileName)

		res.Json(w, "Webhook set", http.StatusOK)
	}
}
