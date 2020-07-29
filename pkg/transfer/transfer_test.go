package transfer

import (
	"homework/pkg/card"
	"homework/pkg/testdata"
	"homework/pkg/transaction"
	"os"
	"reflect"
	"testing"
)

// тестируем импорт
func TestService_ImportTransactionsFromCsv(t *testing.T) {
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
			name:            "Импорт-тест-csv",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToCsv(tt.transferService.Transactions); err != tt.want {
				t.Errorf("ExportTransactionsToCsv() gotErr = %v, want %v", err, tt.want)
			}
			dir, _ := os.Getwd()
			if testTransactions, err := transaction.ImportTransactionsFromCsv(dir + "/exports.csv"); err != tt.want && testTransactions == nil {
				t.Errorf("ImportTransactionsFromCsv() gotErr = %v, want %v, gotTransactions = %v", err, tt.want, testTransactions)
			}
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
			name:            "Экспорт-тест-csv",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToCsv(tt.transferService.Transactions); err != nil {
				t.Errorf("ExportTransactionsToCsv() gotErr = %v, want %v", err, tt.want)
			}
		})
	}
}

// тестируем импорт транзакций из json файла
func TestService_ImportTransactionsFromJson(t *testing.T) {
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
			name:            "Импорт-тест-json",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToJson(tt.transferService.Transactions); err != tt.want {
				t.Errorf("ExportTransactionsToJson() gotErr = %v, want %v", err, tt.want)
			}
			dir, _ := os.Getwd()
			if testTransactions, err := transaction.ImportTransactionsFromJson(dir + "/exports.json"); err != tt.want && testTransactions == nil {
				t.Errorf("ImportTransactionsFromJson() gotErr = %v, want %v, gotTransactions = %v", err, tt.want, testTransactions)
			}
		})
	}
}

// тестируем экпорт транзакций в json
func TestService_ExportTransactionsToJson(t *testing.T) {
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
			name:            "Экспорт-тест-json",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToJson(tt.transferService.Transactions); err != nil {
				t.Errorf("ExportTransactionsToJson() gotErr = %v, want %v", err, tt.want)
			}
		})
	}
}

// тестируем импорт транзакций из Xml файла
func TestService_ImportTransactionsFromXml(t *testing.T) {
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
			name:            "Импорт-тест-Xml",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToXml(tt.transferService.Transactions); err != tt.want {
				t.Errorf("ExportTransactionsToXml() gotErr = %v, want %v", err, tt.want)
			}
			dir, _ := os.Getwd()
			if testTransactions, err := transaction.ImportTransactionsFromXml(dir + "/exports.xml"); err != tt.want && testTransactions == nil {
				t.Errorf("ImportTransactionsFromXml() gotErr = %v, want %v, gotTransactions = %v", err, tt.want, testTransactions)
			}
		})
	}
}

// тестируем экпорт транзакций в xml
func TestService_ExportTransactionsToXml(t *testing.T) {
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
			name:            "Экспорт-тест-xml",
			transferService: tSvc,
			want:            nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transaction.ExportTransactionsToXml(tt.transferService.Transactions); err != nil {
				t.Errorf("ExportTransactionsToXml() gotErr = %v, want %v", err, tt.want)
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
