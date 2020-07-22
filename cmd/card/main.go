package main

import (
	"fmt"
	"homework/pkg/card"
	"homework/pkg/transaction"
	"homework/pkg/transfer"
	"sort"
	"sync"
	"time"
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

	transactions := []*transaction.Transaction{
		{
			Id:            0,
			Amount:        100_00,
			Datetime:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            1,
			Amount:        400_00,
			Datetime:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        600_00,
			Datetime:      time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        1000_00,
			Datetime:      time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        4500_00,
			Datetime:      time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        1000_00,
			Datetime:      time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
	}

	transferService := transfer.NewServiceWithTransactions(cardService, transactions)

	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 1_500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 2_500_00)
	_ = transferService.Card2Card(cardOne.Number, cardTwo.Number, 300_00)

	startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	endTime := time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)

	groupedTransactions := transferService.GetTransactionsGroupedByMonths(&cardOne, startTime.Unix(), endTime.Unix())

	fmt.Println("")
	fmt.Println("-------------------------------------")

	result := sumConcurrently(groupedTransactions)

	keys := make([]string, 0)

	result.Range(func(key, value interface{}) bool {

		k, _ := key.(string)

		keys = append(keys, k)

		return true
	})

	sort.Strings(keys)

	for _, key := range keys {

		value, _ := result.Load(key)

		fmt.Println(key, " - ", value)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("")
}

// суммирует расходы по транзацкиям по месяцам и выводит в терминал
func sumConcurrently(groupedTransactions map[string][]*transaction.Transaction) (result sync.Map) {

	monthsCount := len(groupedTransactions)

	wg := sync.WaitGroup{}
	wg.Add(monthsCount)

	for key, value := range groupedTransactions {

		transactions := value

		mark := key

		go func() {

			sum := 0

			for i := range transactions {

				sum += transactions[i].Amount
			}

			result.Store(mark, sum)

			wg.Done()
		}()

	}

	wg.Wait()

	return result
}
