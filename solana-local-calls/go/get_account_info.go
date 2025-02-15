package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetBalance() {
	// account := "954m3369W174DASH1gyMGFEZaYF74eyxuQbyVMjVatxQ"
	account := "B8sNke1a2sRPhh4fZjMcXvLqskHnNW6jMU45EvH4GWiA"

	// client := rpc.New("https://api.mainnet-beta.solana.com")
	// client := rpc.New("http://localhost:8899")
	client := rpc.New("http://15.236.48.228:8899")

	ctx := context.Background()
	publicKey := solana.MustPublicKeyFromBase58(account)
	commitment := rpc.CommitmentConfirmed

	res, err := client.GetBalance(ctx, publicKey, commitment)
	if err != nil {
		log.Fatal(err)
	}

	balanceInSOL := float64(res.Value) / float64(solana.LAMPORTS_PER_SOL)
	fmt.Printf("Баланс кошелька %s: %.9f SOL\n", account, balanceInSOL)
}

func main() {
	GetBalance()
}