package transfer

import (
	"homework/pkg/card"
	"homework/pkg/transaction"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc      *card.Service
		Transactions []*transaction.Transaction
	}
	type args struct {
		from   string
		to     string
		amount int
	}

	cardService := card.NewService("БАНК БАБАБАНК")

	cardOne := card.Card{
		Balance:  1_0000_00,
		Currency: "RUB",
		Number:   "4191637314259912",
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
		Balance:  2_000_00,
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

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Карта-имеет-валидный-номер,-но-не-будет-найдена-потому-что-принадлежит-другому-банку",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "4191637314259912",
				to:     "4191637314259912",
				amount: 1_000_00,
			},
			wantErr: ErrFromCardNotFound,
		},
		{
			name: "Карта-имеет-невалидный-номер",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "4191637314259919",
				to:     "4191637314259912",
				amount: 1_000_00,
			},
			wantErr: ErrFromCardNumberNotValid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Transactions: tt.fields.Transactions,
			}

			gotErr := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)

			// отхватили ли ошибку
			if gotErr != tt.wantErr {

				t.Errorf("Card2Card() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
