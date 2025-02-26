package getTransactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


const TransactionsForAccountFile string = "/Users/stepan-karpov/Desktop/quant-falcon/helius-get-historical-data/data_files/transactionsForAccount.json"

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

func GetTransactionsForAccount(url string, address string, lastTransaction string, mock bool) ([]Transaction, error) {
	payload := []byte(fmt.Sprintf(`{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["%s"]}`, address))
	if lastTransaction != "" {
		payload = []byte(fmt.Sprintf(`{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["%s", {"limit": 1000, "before": "%s", "commitment": "finalized"}]}`, address, lastTransaction))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return nil, err
	}

	txCount := len(response.Result)
	fmt.Println("Количество транзакций:", txCount)

	return response.Result, nil
}

func WriteSuccessfulTransactionsToFile(transactions []Transaction) error {
	var successfulTxs []Transaction
	for _, tx := range transactions {
		if tx.Err == nil {
			successfulTxs = append(successfulTxs, tx)
		}
	}

	if len(successfulTxs) == 0 {
		fmt.Println("Нет успешных транзакций для записи.")
		return nil
	}

	var existingTxs []Transaction
	file, err := os.Open(TransactionsForAccountFile)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingTxs); err != nil {
			fmt.Println("Ошибка при разборе существующего JSON:", err)
		}
	}

	existingTxs = append(existingTxs, successfulTxs...)

	file, err = os.OpenFile(TransactionsForAccountFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Красивый формат
	if err := encoder.Encode(existingTxs); err != nil {
		fmt.Println("Ошибка при записи JSON:", err)
		return err
	}

	fmt.Println("Добавлено", len(successfulTxs), "успешных транзакций в", TransactionsForAccountFile)
	return nil
}
