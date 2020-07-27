package testdata

import (
	"homework/pkg/card"
	"homework/pkg/transaction"
	"math/rand"
	"time"
)

// создает карты
func MakeCards() (cards []*card.Card) {

	cardOwnersDatamap := []string{
		"mister green",
		"mister blue",
		"mister grey",
		"mister yellow",
		"mister red",
		"mister gold",
		"mister white",
		"mister black",
		"mister purple",
		"mister multicolor",
		"mister pink",
	}

	cardsDataMap := map[string]string{
		"5106 2184 1644 4735": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2132 1882 2113": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2128 6659 6714": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2176 9107 2252": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2123 5239 5522": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2130 9602 8379": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2121 3543 4895": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2163 9916 2894": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2153 7805 4189": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2120 2303 5804": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2126 1596 2522": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2153 9233 6513": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2166 5150 6119": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2193 5734 7762": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2113 7668 5587": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2174 1863 7700": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
		"5106 2130 9653 1406": cardOwnersDatamap[rand.Intn(len(cardOwnersDatamap))],
	}

	for k, v := range cardsDataMap {

		cards = append(cards, &card.Card{
			Owner:    v,
			Balance:  1000_000_00,
			Currency: "RUB",
			Number:   k,
			Icon:     "/icon.png",
		})

	}

	return cards
}

// генерирует транзакции
func MakeTransactions(cards []*card.Card) (transactions []*transaction.Transaction) {

	cardsCount := len(cards)

	transactions = make([]*transaction.Transaction, 1000)

	transactionAmount := 10_00

	for i := range transactions {

		tx := transaction.Transaction{
			Id:            0,
			Amount:        rand.Intn(transactionAmount),
			Datetime:      time.Now().Unix(),
			OperationType: "from",
			Status:        true,
			Mcc:           0,
			CardFrom:      cards[rand.Intn(cardsCount)].Number,
			CardTo:        cards[rand.Intn(cardsCount)].Number,
		}

		switch i % 10 {

		case 0:
			tx.Mcc = 4112
		case 1:
			tx.Mcc = 4121
		case 2:
			tx.Mcc = 4131
		case 3:
			tx.Mcc = 4225
		case 4:
			tx.Mcc = 4789
		case 5:
			tx.Mcc = 4821
		case 6:
			tx.Mcc = 4899
		case 7:
			tx.Mcc = 5044
		default:
			tx.Mcc = 5013
		}

		transactions[i] = &tx
	}

	return transactions
}
