package transfer

import "homework/pkg/card"

// сервис
type Service struct {
	CardSvc *card.Service
}

// конструктор сервиса
func NewService(cardSvc *card.Service) *Service {
	return &Service{CardSvc: cardSvc}
}

// перевод с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool) {

	ok = true

	// ищем карту с которой будем преводить
	cardFrom := s.CardSvc.FindCardByNumber(from)

	// ищем карту с которой будем преводить
	cardTo := s.CardSvc.FindCardByNumber(to)

	// процент за перевод и минимальная коммисия
	transferFeePercentage, transferFeeMin := transferFeePercentageAndMinimum(cardFrom, cardTo)

	totalToTransfer := amountPlusCommission(amount, transferFeePercentage, transferFeeMin)

	if cardFrom != nil {

		_, ok = s.CardSvc.TransferFromCard(cardFrom, totalToTransfer)
	}

	if cardTo != nil {

		s.CardSvc.TransferToCard(cardTo, amount)
	}

	return totalToTransfer, ok
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
