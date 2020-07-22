package transfer

import (
	"fmt"
	"homework/pkg/card"
	"homework/pkg/transaction"
	"math/rand"
	"testing"
	"time"
)

// создает карты
func makeCards() (cards []*card.Card) {

	cardsDataMap := map[string]string{
		"5106 2184 1644 4735": "mister red",
		"5106 2132 1882 2113": "mister blue",
		"5106 2128 6659 6714": "mister green",
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
func makeTransactions(cards []*card.Card) (transactions []*transaction.Transaction) {

	cardsCount := len(cards)

	transactions = make([]*transaction.Transaction, 10000)

	transactionAmount := 100_00

	for i := range transactions {

		tx := transaction.Transaction{
			Id:            0,
			Amount:        rand.Intn(transactionAmount),
			Datetime:      time.Now().Unix(),
			OperationType: "from",
			Status:        true,
			Mcc:           0,
			CardFrom:      cards[rand.Intn(cardsCount)],
			CardTo:        cards[rand.Intn(cardsCount)],
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

func TestService_Card2Card(t *testing.T) {

	type fields struct {
		CardSvc      *card.Service
		Transactions []*transaction.Transaction
	}

	type args struct {
		cardNumber string
	}

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := makeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := makeTransactions(cards)

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Вывод группированных по MCC затрат без ",
			fields: fields{
				CardSvc:      cardService,
				Transactions: transactions,
			},
			args: args{
				cardNumber: cards[rand.Intn(len(cards))].Number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Transactions: tt.fields.Transactions,
			}

			for _, c := range s.CardSvc.Cards {

				fmt.Println("")
				fmt.Println("CARD NUMBER : ", c.Number)
				fmt.Println("")

				mccTransactionSumAmountMap := s.GetMccTransactionsSumAmountMapByCard(s.GetTransactionsByType(c, "from"))

				fmt.Println("SIMPLE")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapByCardWithMutex(s.GetTransactionsByType(c, "from"))

				fmt.Println("MUTEX")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapByCardWithChannels(s.GetTransactionsByType(c, "from"))

				fmt.Println("CHANNELS")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapByCardWithMutexStraightToMap(s.GetTransactionsByType(c, "from"))

				fmt.Println("MUTEXSTRAIGHTTOMAP")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")
			}
		})
	}
}
