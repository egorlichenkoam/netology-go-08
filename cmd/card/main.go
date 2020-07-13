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
		Balance:  1_000_00,
		Currency: "RUB",
		Number:   "5106210001",
		Icon:     "card.png",
	}

	cardTwo := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "5106210002",
		Icon:     "card.png",
	}

	cardThree := card.Card{
		Balance:  100_00,
		Currency: "RUB",
		Number:   "5106210003",
		Icon:     "card.png",
	}

	cardFour := card.Card{
		Balance:  1_000_00,
		Currency: "RUB",
		Number:   "5106210004",
		Icon:     "card.png",
	}

	cardFive := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "5106210005",
		Icon:     "card.png",
	}

	cardSix := card.Card{
		Balance:  100_00,
		Currency: "RUB",
		Number:   "5106210006",
		Icon:     "card.png",
	}

	cardService.AddCard(&cardOne)
	cardService.AddCard(&cardTwo)
	cardService.AddCard(&cardThree)
	cardService.AddCard(&cardFour)
	cardService.AddCard(&cardFive)
	cardService.AddCard(&cardSix)

	transferService := transfer.NewService(cardService)

	fmt.Println(transferService.Card2Card("5106210006", "0009", 1_500_00))
}
