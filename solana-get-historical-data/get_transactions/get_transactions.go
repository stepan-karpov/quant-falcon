package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	ID      string        `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	Signature          string      `json:"signature"`
	Slot               int         `json:"slot"`
	Err                interface{} `json:"err"`
	Memo               *string     `json:"memo"`
	BlockTime          *int64      `json:"blockTime"`
	ConfirmationStatus string      `json:"confirmationStatus"`
}

const url string = "https://mainnet.helius-rpc.com/?api-key=<tour-api-key>"
const address string = "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // TRUMP
const filePath string = "/Users/stepan-karpov/Desktop/quant-falcon/solana-get-historical-data/get_transactions/transactions.json"

func ExecuteRequest(lastTransaction string) (string, error) {

	payload := []byte(fmt.Sprintf(`{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["%s"]}`, address))

	if lastTransaction != "" {
		payload = []byte(fmt.Sprintf(`{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["%s", {"limit": 1000, "before": "%s"}]}`, address, lastTransaction))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return "", err
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return "", err
	}
	fmt.Println("Ответ добавлен в файл:", filePath)

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return "", err
	}

	txCount := len(response.Result)
	fmt.Println("Количество транзакций:", txCount)

	return response.Result[len(response.Result)-1].Signature, nil
}

func main() {
	lastTransaction := ""

	for range 3 {
		var err error
		lastTransaction, err = ExecuteRequest(lastTransaction)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}
	
		if lastTransaction != "" {
			fmt.Printf("Сигнатура первой успешной транзакции: %s\n", lastTransaction)
		} else {
			fmt.Println("Нет успешных транзакций.")
		}
	}
}

// test run is:
// stepan-karpov@i113819829 get_transactions % go run "/Users/stepan-karpov/quant-falcon/solana-get-historical-data/get_transaction
// s/get_transactions.go"
// Ответ добавлен в файл: /Users/stepan-karpov/quant-falcon/solana-get-historical-data/get_transactions/transactions.json
// Количество транзакций: 1000
// Сигнатура первой успешной транзакции: 2bs5uS26dmK7eU9tD3V2khpwrd4KVT8cSfsBWR5TzYiLxbLKEBwRUgnK53vwDTUcWeXhZXsBYw4XbZUkELHJF9eD
// Ответ добавлен в файл: /Users/stepan-karpov/quant-falcon/solana-get-historical-data/get_transactions/transactions.json
// Количество транзакций: 1000
// Сигнатура первой успешной транзакции: mHk7zQ7G721dYLmkA8Cg5NVcNPmmiJsjhUioUqaUrMY4KYReDSH5rq9XTmFZ7A9CDpsXRUNjz1e5UCkahhHi6ii
// Ответ добавлен в файл: /Users/stepan-karpov/quant-falcon/solana-get-historical-data/get_transactions/transactions.json
// Количество транзакций: 1000
// Сигнатура первой успешной транзакции: 123q8LmjJr2ybG7McenD853sG6ZJ3X3TZZqa7NS3ZLsojY6xubb7UKFzsbJrNMqxGBUuUC5vyaRVWVi5RwwLjqZC