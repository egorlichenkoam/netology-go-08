package main

import (
	"homework/pkg/card"
	"homework/pkg/testdata"
	"homework/pkg/transaction"
	"homework/pkg/transfer"
	"log"
	"os"
)

func main() {

	cards := testdata.MakeCards()

	transactions := testdata.MakeTransactions(cards)

	cardService := card.NewService("ТРУ ЛЯ ЛЯ")

	for _, c := range cards {

		cardService.AddCard(c)
	}

	transferService := transfer.NewServiceWithTransactions(cardService, transactions)

	_ = transferService.ExportTransactions()

	transferService.Transactions = make([]*transaction.Transaction, 0)

	dir, _ := os.Getwd()

	_ = transferService.ImportTransactions(dir + "/exports.csv")

	for _, tx := range transferService.Transactions {

		log.Println(tx.String())
	}
}
