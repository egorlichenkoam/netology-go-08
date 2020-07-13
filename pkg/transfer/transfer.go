package transfer

import (
	"errors"
	"homework/pkg/card"
	"homework/pkg/transaction"
	"time"
)

// сервис
type Service struct {
	CardSvc      *card.Service
	Transactions []*transaction.Transaction
}

var (
	ErrFromCardNoEnoughMoney = errors.New("Source card: not enough money")
)

// конструктор сервиса
func NewService(cardSvc *card.Service) *Service {
	return &Service{CardSvc: cardSvc}
}

// добавляет транзакцию
func (s *Service) addTransaction(transaction *transaction.Transaction) {

	transaction.Datetime = time.Now().Unix()

	s.Transactions = append(s.Transactions, transaction)
}

// перевод с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (total int, err error, ok bool) {

	ok = true

	// ищем карту с которой будем преводить
	cardFrom, ok := s.CardSvc.FindCardByNumber(from)

	// ищем карту с которой будем преводить
	cardTo, ok := s.CardSvc.FindCardByNumber(to)

	// процент за перевод и минимальная коммисия
	transferFeePercentage, transferFeeMin := transferFeePercentageAndMinimum(cardFrom, cardTo)

	totalToTransfer := amountPlusCommission(amount, transferFeePercentage, transferFeeMin)

	if cardFrom == nil {

		cardFrom = &card.Card{
			Balance:  0,
			Currency: "RUB",
			Number:   from,
			Icon:     "card.png",
		}

		card.SetBankName(cardFrom, "ДРУГОЙ БАНК")

		card.SetExternalBank(cardFrom, true)
	}

	_, ok = s.CardSvc.TransferFromCard(cardFrom, totalToTransfer)

	if !ok {

		return totalToTransfer, ErrFromCardNoEnoughMoney, false
	}

	if cardTo == nil {

		cardTo = &card.Card{
			Balance:  0,
			Currency: "RUB",
			Number:   to,
			Icon:     "card.png",
		}

		card.SetBankName(cardTo, "ДРУГОЙ БАНК")
		card.SetExternalBank(cardTo, true)
	}

	s.CardSvc.TransferToCard(cardTo, amount)

	// транзакция для списания
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Sum:           totalToTransfer,
		OperationType: "from",
		Status:        ok,
		CardFrom:      cardFrom,
		CardTo:        cardTo,
	})

	// транзакция для зачисления
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Sum:           amount,
		OperationType: "to",
		Status:        ok,
		CardFrom:      cardFrom,
		CardTo:        cardTo,
	})

	return totalToTransfer, err, ok
}

// возвращает процент коммисии за перевод и минимальную коммисию за перевод
func transferFeePercentageAndMinimum(cardFrom, cardTo *card.Card) (transferFeePercentage float64, transferFeeMin int) {

	if cardFrom == nil && cardTo == nil {

		return 1.5, 3000
	}

	if cardFrom != nil && cardTo == nil {

		return 0.5, 1000
	}

	return 0.0, 0
}

// возвращает сумму для списания с карты с учтом комиссии
func amountPlusCommission(amount int, transferFeePercentage float64, transferFeeMin int) (total int) {

	internalCommission := int(float64(amount) / 100 * transferFeePercentage)

	if internalCommission < transferFeeMin {

		internalCommission = transferFeeMin
	}

	return amount + internalCommission
}
