package main

import (
	"encoding/json"
	"fmt"
	"main/getTransactionInfo"
	"main/getTransactions"
	"os"
)

const url string = "https://mainnet.helius-rpc.com/?api-key=<your-api-key>"
const address string = "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // TRUMP

func main() {
	skipGetTransactionsForAccount := true
	skipGetTransactionsInfo := false

	if !skipGetTransactionsForAccount {
		lastTransaction := ""

		for range 3 {
			transactions, err := getTransactions.GetTransactionsForAccount(url, address, lastTransaction, skipGetTransactionsForAccount)
			if err != nil {
				fmt.Println("Ошибка выполнения запроса:", err)
				return
			}
			err = getTransactions.WriteSuccessfulTransactionsToFile(transactions)
			if err != nil {
				fmt.Println("Ошибка записи успешных транзакций в файл:", err)
				return
			}
			lastTransaction = transactions[len(transactions)-1].Signature

		}
	}

	if !skipGetTransactionsInfo {
		var existingTxs []getTransactions.Transaction
		file, err := os.Open(getTransactions.TransactionsForAccountFile)
		if err == nil {
			defer file.Close()
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&existingTxs); err != nil {
				fmt.Println("Ошибка при разборе существующего JSON:", err)
			}
		}

		for _, tx := range existingTxs {
			resp, err := getTransactionInfo.GetTransactionInfo(url, tx.Signature)
			if err != nil {
				fmt.Println("Ошибка при выполнении запроса:", err)
				return
			}
			getTransactionInfo.WriteTransactionInfoToFile(resp)
		}
	}

}
