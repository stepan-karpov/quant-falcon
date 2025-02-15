package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetBalance() {
	account := "GgVXANXY7dkaYmmMwVvnq2CofN3njbuS9Jv1E4uD8M6W" // Замените на свой адрес

	// client := rpc.New("https://api.mainnet-beta.solana.com")
	client := rpc.New("http://localhost:8899")

	// Получаем баланс кошелька
	ctx := context.Background()                          // Контекст запроса
	publicKey := solana.MustPublicKeyFromBase58(account) // Публичный ключ
	commitment := rpc.CommitmentConfirmed                // Тип подтверждения (Confirmed)

	// Запрос баланса
	res, err := client.GetBalance(ctx, publicKey, commitment)
	if err != nil {
		log.Fatal(err)
	}

	// Печатаем баланс в SOL
	balanceInSOL := float64(res.Value) / float64(solana.LAMPORTS_PER_SOL)
	fmt.Printf("Баланс кошелька %s: %.9f SOL\n", account, balanceInSOL)
}

func main() {
	GetBalance()
}