package main

import (
	"fmt"
	"homework/pkg/card"
	"time"
)

func main() {

	master := &card.Card{
		Issuer:   "Gipermarket",
		Balance:  1000000,
		Currency: "",
		Number:   "7643-1237-2383-9386",
		Icon:     "masterCard.png",
		Transactions: []card.Transaction{
			{
				Id:       1,
				Sum:      -73555,
				Datetime: time.Now().Unix(),
				Mcc:      "5411",
				Status:   "DONE",
			},
			{
				Id:       2,
				Sum:      200000,
				Datetime: time.Now().Unix(),
				Mcc:      "0000",
				Status:   "DONE",
			},
			{
				Id:       3,
				Sum:      -120391,
				Datetime: time.Now().Unix(),
				Mcc:      "5812",
				Status:   "DONE",
			}},
	}

	fmt.Println(master)

	transaction := &card.Transaction{
		Id:       4,
		Sum:      99900,
		Datetime: time.Now().Unix(),
		Mcc:      "5555",
		Status:   "DONE",
	}

	fmt.Println(transaction)

	card.AddTransaction(master, transaction)

	fmt.Println(master)

	fmt.Println(card.SumByMMC(master.Transactions, []string{"5411", "0000"}))

	fmt.Println(card.TranslateMCC(transaction.Mcc))

	fmt.Println(card.LastNTransactions(master, 2))
}
