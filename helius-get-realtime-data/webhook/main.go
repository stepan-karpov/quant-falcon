package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Константы (замени на свои данные)
const (
	HeliusAPIKey     = "API_KEY"
	HeliusRPCURL     = "https://mainnet.helius-rpc.com/?api-key=" + HeliusAPIKey
)

// WebhookPayload — структура для обработки входящих данных
type WebhookPayload struct {
	Description string `json:"description"`
	Timestamp   int64  `json:"timestamp"`
	Signature   string `json:"signature"`
}

// handleRequest — обработчик POST-запросов
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON
	var payload []WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Проверяем, есть ли данные
	if len(payload) == 0 {
		http.Error(w, "Empty payload", http.StatusBadRequest)
		return
	}

	// Берем первую транзакцию
	tx := payload[0]
	txTimestamp := time.Unix(tx.Timestamp, 0).Format("2006-01-02 15:04:05")
	txLink := fmt.Sprintf("https://xray.helius.xyz/tx/%s", tx.Signature)

	// Формируем сообщение
	message := fmt.Sprintf(
		"----NEW UPDATE---\nDescription:\n%s\nSignature:\n%s\nTimestamp:\n%s",
		tx.Description, txLink, txTimestamp,
	)

	// Логируем (можно заменить на отправку в Telegram)
	log.Println(message)

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged POST request body.")
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
