package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	// Логируем весь полученный JSON (или любой формат данных)
	log.Println("Received body:", string(body))

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged POST request body.")
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
