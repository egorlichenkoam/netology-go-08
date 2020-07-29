package main

import (
	"homework/pkg/card"
	"homework/pkg/testdata"
	"homework/pkg/transfer"
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
}
