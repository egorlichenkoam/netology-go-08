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
		Balance:  2_000_00,
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

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int
		wantOk    bool
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0001",
				to:     "0002",
				amount: 1_000_00,
			},
			wantTotal: 1_000_00,
			wantOk:    true,
		}, {
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0003",
				to:     "0002",
				amount: 10_000_00,
			},
			wantTotal: 10_000_00,
			wantOk:    false,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0004",
				to:     "0007",
				amount: 1_500_00,
			},
			wantTotal: 1_510_00,
			wantOk:    true,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0006",
				to:     "0009",
				amount: 1_500_00,
			},
			wantTotal: 1_510_00,
			wantOk:    false,
		}, {
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0010",
				to:     "0005",
				amount: 1_500_00,
			},
			wantTotal: 1_500_00,
			wantOk:    true,
		}, {
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc:      cardService,
				Transactions: nil,
			},
			args: args{
				from:   "0016",
				to:     "0011",
				amount: 1_500_00,
			},
			wantTotal: 1_500_00,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Transactions: tt.fields.Transactions,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
