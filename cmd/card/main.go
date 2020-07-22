package main

import (
	"fmt"
	"homework/pkg/card"
	"homework/pkg/transfer"
)

// основной метод
func main() {

	cardService := card.NewService("БАНК БАНКОВ")

	cardOne := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "5106 2183 5342 7107",
		Icon:     "card.png",
	}

	cardTwo := card.Card{
		Balance:  1_000_00,
		Currency: "RUB",
		Number:   "5106 2115 1463 2228",
		Icon:     "card.png",
	}

	cardService.AddCard(&cardOne)
	cardService.AddCard(&cardTwo)

	transferService := transfer.NewService(cardService)

	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 1_500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 2_500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 300_00)

	transactions := transferService.GetSortedTransactionsByType(&cardOne, "from")

	fmt.Println(cardOne.Number)

	for n := range transactions {

		tx := transactions[n]

		fmt.Println(tx.Amount, tx.OperationType, tx.Status, tx.CardFrom.Number)
	}

}
