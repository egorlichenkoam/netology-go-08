package transfer

import (
	"errors"
	"homework/pkg/card"
	"homework/pkg/transaction"
	"sort"
)

// сервис
type Service struct {
	CardSvc      *card.Service
	Transactions []*transaction.Transaction
}

var (
	ErrFromCardNotFound       = errors.New("Source card not found")
	ErrFromCardNumberNotValid = errors.New("Source card number not valid")

	ErrToCardNotFound       = errors.New("Target card not found")
	ErrToCardNumberNotValid = errors.New("Target card number not valid")
)

// конструктор сервиса
func NewService(cardSvc *card.Service) *Service {

	return &Service{CardSvc: cardSvc}
}

// возвращает сортированный по amount список транзакций карты по типу (from - расход, to - приход)
func (s *Service) GetSortedTransactionsByType(card *card.Card, operationType string) (transactions []*transaction.Transaction) {

	result := make([]*transaction.Transaction, 0)

	for n := range s.Transactions {

		tx := s.Transactions[n]

		if tx.CardFrom == card && tx.OperationType == operationType {

			result = append(result, tx)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Amount > result[j].Amount
	})

	return result
}

// добавляет транзакцию
func (s *Service) addTransaction(transaction *transaction.Transaction) {

	//transaction.Datetime = time.Now().Unix()
	transaction.Datetime = 0

	s.Transactions = append(s.Transactions, transaction)
}

// валидирует номера карт
func (s *Service) validateCards(from, to string) (err error) {

	if !s.isValid(from) {

		return ErrFromCardNumberNotValid
	}

	if !s.isValid(to) {

		return ErrToCardNumberNotValid
	}

	return nil
}

// ищем карты в сервисе
func (s *Service) searchCards(from, to string) (err error, fromCard, toCard *card.Card) {

	// ищем карту с которой будем преводить
	toCard = s.CardSvc.FindCardByNumber(to)

	if toCard == nil {

		err = ErrToCardNotFound
	}

	// ищем карту с которой будем преводить
	fromCard = s.CardSvc.FindCardByNumber(from)

	if fromCard == nil {

		err = ErrFromCardNotFound
	}

	return err, fromCard, toCard
}

// перевод с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (err error) {

	if err = s.validateCards(from, to); err != nil {

		return err
	}

	err, fromCard, toCard := s.searchCards(from, to)

	if err != nil {

		return err
	}

	// процент за перевод и минимальная коммисия
	transferFeePercentage, transferFeeMin := transferFeePercentageAndMinimum(fromCard, toCard)

	totalToTransfer := amountPlusCommission(amount, transferFeePercentage, transferFeeMin)

	if err = s.CardSvc.Transfer(fromCard, totalToTransfer, true); err != nil {

		return err
	}

	// транзакция для списания
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Amount:        totalToTransfer,
		OperationType: "from",
		Status:        true,
		CardFrom:      fromCard,
		CardTo:        toCard,
	})

	_ = s.CardSvc.Transfer(toCard, amount, false)

	// транзакция для зачисления
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Amount:        amount,
		OperationType: "to",
		Status:        true,
		CardFrom:      fromCard,
		CardTo:        toCard,
	})

	return nil
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

func (s *Service) isValid(number string) bool {

	return s.CardSvc.CheckCardNumberByLuna(number)
}
