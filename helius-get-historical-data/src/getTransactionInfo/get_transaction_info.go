package getTransactionInfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const TransactionsInfoFile string = "/Users/stepan-karpov/Desktop/quant-falcon/helius-get-historical-data/data_files/transactionsInfo.json"

type Response struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

type Result struct {
	Slot        int         `json:"slot"`
	Meta        Meta        `json:"meta"`
	Transaction Transaction `json:"transaction"`
}

type Meta struct {
	Err               interface{} `json:"err"`
	Fee               int         `json:"fee"`
	InnerInstructions []struct{}  `json:"innerInstructions"`
	PostBalances      []int64     `json:"postBalances"`
	PreBalances       []int64     `json:"preBalances"`
	Rewards           []Reward    `json:"rewards"`
}

type Reward struct {
	Pubkey      string `json:"pubkey"`
	Lamports    int    `json:"lamports"`
	PostBalance int    `json:"postBalance"`
	RewardType  string `json:"rewardType"`
}

type Transaction struct {
	Signatures []string `json:"signatures"`
	Message    Message  `json:"message"`
}

type Message struct {
	RecentBlockhash string        `json:"recentBlockhash"`
	AccountKeys     []string      `json:"accountKeys"`
	Header          Header        `json:"header"`
	Instructions    []Instruction `json:"instructions"`
}

type Header struct {
	NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int `json:"numRequiredSignatures"`
}

type Instruction struct {
	Data           string `json:"data"`
	ProgramIDIndex int    `json:"programIdIndex"`
	Accounts       []int  `json:"accounts"`
}

func GetTransactionInfo(url string, signature string) (Response, error) {
	payload := []byte(fmt.Sprintf(`{"id":"1","jsonrpc":"2.0","method":"getTransaction","params":["%s", {"maxSupportedTransactionVersion": 0}]}`, signature))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return Response{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return Response{}, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return Response{}, err
	}

	return response, nil
}

func WriteTransactionInfoToFile(txInfo Response) error {
	var transactions []Response

	fileData, err := os.ReadFile(TransactionsInfoFile)
	if err == nil {
		if len(fileData) > 0 {
			if err := json.Unmarshal(fileData, &transactions); err != nil {
				fmt.Println("Ошибка при разборе JSON из файла:", err)
				return err
			}
		}
	}

	transactions = append(transactions, txInfo)

	jsonData, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return err
	}

	if err := os.WriteFile(TransactionsInfoFile, jsonData, 0644); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return err
	}

	fmt.Println("Транзакция успешно добавлена в файл:", TransactionsInfoFile)
	return nil
}

