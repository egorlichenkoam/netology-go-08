package transfer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"homework/pkg/card"
	"homework/pkg/testdata"
	"homework/pkg/transaction"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

// тестируем импорт
func TestService_ImportTransactions(t *testing.T) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	tSvc := Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	tests := []struct {
		name            string
		transferService Service
		want            error
	}{
		{
			name:            "Импорт-тест",
			transferService: tSvc,
			want:            nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			_ = tt.transferService.ExportTransactions()

			t.Log("Transactions count export : ", len(tt.transferService.Transactions))

			tt.transferService.Transactions = make([]*transaction.Transaction, 0)

			t.Log("Transactions count clear : ", len(tt.transferService.Transactions))

			dir, _ := os.Getwd()

			if err := tt.transferService.ImportTransactions(dir + "/exports.csv"); err != tt.want {

				t.Errorf("ExportTransactions() gotErr = %v, want %v", err, tt.want)
			}

			t.Log("Transactions count import : ", len(tt.transferService.Transactions))

			tt.transferService.Transactions = make([]*transaction.Transaction, 0)

			t.Log("Transactions count clear : ", len(tt.transferService.Transactions))

			data, err := ioutil.ReadFile(dir + "/exports.csv")

			if err != nil {

				log.Println("Can not open import transactions file", err)
			}

			reader := csv.NewReader(bytes.NewReader(data))

			records, err := reader.ReadAll()

			if err != nil {

				log.Println("Can not read data from import file", err)
			}

			for _, content := range records {

				tx := transaction.Transaction{}

				if err := tx.MapRowToTransaction(content); err != tt.want {

					t.Errorf("MapRowToTransaction() gotErr = %v, want %v", err, tt.want)
				}

				if err != nil {

					log.Println("Can not import transaction", err, content)
				} else {

					tt.transferService.Transactions = append(tt.transferService.Transactions, &tx)
				}
			}

			t.Log("Transactions count import : ", len(tt.transferService.Transactions))
		})
	}
}

// тестируем экпорт транзакций
func TestService_ExportTransactions(t *testing.T) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	tSvc := Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	tests := []struct {
		name            string
		transferService Service
		want            error
	}{
		{
			name:            "Экспорт-тест",
			transferService: tSvc,
			want:            nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if err := tt.transferService.ExportTransactions(); err != nil {

				t.Errorf("ExportTransactions() gotErr = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestService_Card2Card(t *testing.T) {

	type fields struct {
		CardSvc      *card.Service
		Transactions []*transaction.Transaction
	}

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Вывод группированных по MCC затрат без ",
			fields: fields{
				CardSvc:      cardService,
				Transactions: transactions,
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
				fmt.Println("OWNER : ", c.Owner)
				fmt.Println("CARD NUMBER : ", c.Number)
				fmt.Println("")

				mccTransactionSumAmountMap := s.GetMccTransactionsSumAmountMap(s.GetTransactionsByType(c, "from"))

				fmt.Println("SIMPLE")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithMutex(s.GetTransactionsByType(c, "from"))

				fmt.Println("MUTEX")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithChannels(s.GetTransactionsByType(c, "from"))

				fmt.Println("CHANNELS")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithMutexStraightToMap(s.GetTransactionsByType(c, "from"))

				fmt.Println("MUTEXSTRAIGHTTOMAP")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")
			}

			for owner, cards := range s.CardSvc.Owners() {

				fmt.Println("")
				fmt.Println("OWNER : ", owner)
				fmt.Println("")

				fmt.Println("")
				fmt.Println("CARDS COUNT : ", len(cards))
				fmt.Println("")

				mccTransactionSumAmountMap := s.GetMccTransactionsSumAmountMap(s.GetTransactionsByTypeAndOwner(owner, "from"))

				fmt.Println("SIMPLE")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithMutex(s.GetTransactionsByTypeAndOwner(owner, "from"))

				fmt.Println("MUTEX")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithChannels(s.GetTransactionsByTypeAndOwner(owner, "from"))

				fmt.Println("CHANNELS")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")

				mccTransactionSumAmountMap = s.GetMccTransactionsSumAmountMapWithMutexStraightToMap(s.GetTransactionsByTypeAndOwner(owner, "from"))

				fmt.Println("MUTEXSTRAIGHTTOMAP")
				fmt.Println("------------------------")
				fmt.Println(mccTransactionSumAmountMap)
				fmt.Println("------------------------")
			}
		})
	}
}

func BenchmarkMutexByCard(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for _, c := range s.CardSvc.Cards {

		want[c.Number] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByType(c, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for _, c := range s.CardSvc.Cards {

			result[c.Number] = s.GetMccTransactionsSumAmountMapWithMutex(s.GetTransactionsByType(c, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}

func BenchmarkMutexByOwner(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for owner := range s.CardSvc.Owners() {

		want[owner] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByTypeAndOwner(owner, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for owner := range s.CardSvc.Owners() {

			result[owner] = s.GetMccTransactionsSumAmountMapWithMutex(s.GetTransactionsByTypeAndOwner(owner, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}

func BenchmarkChannelsByCard(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for _, c := range s.CardSvc.Cards {

		want[c.Number] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByType(c, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for _, c := range s.CardSvc.Cards {

			result[c.Number] = s.GetMccTransactionsSumAmountMapWithChannels(s.GetTransactionsByType(c, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}

func BenchmarkChannelsByOwner(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for owner := range s.CardSvc.Owners() {

		want[owner] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByTypeAndOwner(owner, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for owner := range s.CardSvc.Owners() {

			result[owner] = s.GetMccTransactionsSumAmountMapWithChannels(s.GetTransactionsByTypeAndOwner(owner, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}

func BenchmarkMutexStraightToMapByCard(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for _, c := range s.CardSvc.Cards {

		want[c.Number] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByType(c, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for _, c := range s.CardSvc.Cards {

			result[c.Number] = s.GetMccTransactionsSumAmountMapWithMutexStraightToMap(s.GetTransactionsByType(c, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}

func BenchmarkMutexStraightToMapByOwner(b *testing.B) {

	cardService := card.NewService("БАНК БАБАБАНК")

	cards := testdata.MakeCards()

	for _, v := range cards {

		cardService.AddCard(v)
	}

	transactions := testdata.MakeTransactions(cards)

	s := &Service{
		CardSvc:      cardService,
		Transactions: transactions,
	}

	want := make(map[string]map[int]int)

	for owner := range s.CardSvc.Owners() {

		want[owner] = s.GetMccTransactionsSumAmountMap(s.GetTransactionsByTypeAndOwner(owner, "from"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		result := make(map[string]map[int]int)

		for owner := range s.CardSvc.Owners() {

			result[owner] = s.GetMccTransactionsSumAmountMapWithMutexStraightToMap(s.GetTransactionsByTypeAndOwner(owner, "from"))
		}
		b.StopTimer()

		if !reflect.DeepEqual(result, want) {

			b.Fatalf("invalid result, got %v, want %v", result, want)
		}

		b.StartTimer()
	}
}
