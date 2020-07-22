package transfer

import (
	"homework/pkg/card"
	"homework/pkg/transaction"
	"reflect"
	"testing"
)

func TestService_Card2Card(t *testing.T) {

	type fields struct {
		CardSvc      *card.Service
		Transactions []*transaction.Transaction
	}

	type args struct {
		from string
		to   string
	}

	cardService := card.NewService("БАНК БАБАБАНК")

	cardOne := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "5106 2184 1644 4735",
		Icon:     "card.png",
	}

	cardTwo := card.Card{
		Balance:  10_000_00,
		Currency: "RUB",
		Number:   "5106 2132 1882 2113",
		Icon:     "card.png",
	}

	cardService.AddCard(&cardOne)
	cardService.AddCard(&cardTwo)

	transactions := []*transaction.Transaction{

		{
			Id:            0,
			Amount:        1000_00,
			Datetime:      0,
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        900_00,
			Datetime:      0,
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        800_00,
			Datetime:      0,
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        700_00,
			Datetime:      0,
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
		{
			Id:            0,
			Amount:        600_00,
			Datetime:      0,
			OperationType: "from",
			Status:        true,
			CardFrom:      &cardOne,
			CardTo:        &cardTwo,
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*transaction.Transaction
	}{
		{
			name: "Сортировка транзакций",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from: "5106 2184 1644 4735",
				to:   "5106 2132 1882 2113",
			},
			want: transactions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Transactions: tt.fields.Transactions,
			}

			errors := make([]error, 0)
			errors = append(errors, s.Card2Card(tt.args.from, tt.args.to, 600_00))
			errors = append(errors, s.Card2Card(tt.args.from, tt.args.to, 900_00))
			errors = append(errors, s.Card2Card(tt.args.from, tt.args.to, 700_00))
			errors = append(errors, s.Card2Card(tt.args.from, tt.args.to, 1000_00))
			errors = append(errors, s.Card2Card(tt.args.from, tt.args.to, 800_00))

			for n := range errors {

				t.Log(errors[n])
			}

			cardFrom := s.CardSvc.FindCardByNumber(tt.args.from)

			// проверка совпадает ли сортированый слайс транзакций с требуемым
			if got := s.GetSortedTransactionsByType(cardFrom, "from"); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
