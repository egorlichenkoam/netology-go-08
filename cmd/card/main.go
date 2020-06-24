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
		Number:   "0001",
		Icon:     "card.png",
	}

	cardTwo := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "0002",
		Icon:     "card.png",
	}

	cardThree := card.Card{
		Balance:  100_00,
		Currency: "RUB",
		Number:   "0003",
		Icon:     "card.png",
	}

	cardFour := card.Card{
		Balance:  1_000_00,
		Currency: "RUB",
		Number:   "0004",
		Icon:     "card.png",
	}

	cardFive := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "0005",
		Icon:     "card.png",
	}

	cardSix := card.Card{
		Balance:  100_00,
		Currency: "RUB",
		Number:   "0006",
		Icon:     "card.png",
	}

	cardService.AddCard(&cardOne)
	cardService.AddCard(&cardTwo)
	cardService.AddCard(&cardThree)
	cardService.AddCard(&cardFour)
	cardService.AddCard(&cardFive)
	cardService.AddCard(&cardSix)

	transferService := transfer.NewService(cardService)

	fmt.Println(transferService.Card2Card("0006", "0009", 1_500_00))

	//fmt.Println(transferService.Card2Card("0007", "0003", 5_000_00))
	//
	//fmt.Println(transferService.Card2Card("0002", "0009", 7_000_00))
	//
	//fmt.Println(transferService.Card2Card("0007", "0009", 1_000_00))

}
