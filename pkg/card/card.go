package card

import (
	"strconv"
	"strings"
)

// карта
type Card struct {
	id       int
	issuer   string
	Balance  int
	Currency string
	Number   string
	Icon     string
}

// сервис
type Service struct {
	BankName string
	Cards    []*Card
}

var (
	cardnumber_prefix = "510621"
)

// конструктор
func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

// добавляет карту
func (s *Service) AddCard(card *Card) {

	SetBankName(card, s.BankName)

	s.Cards = append(s.Cards, card)
}

// возвращает карту по номеру карты или nil
func (s *Service) FindCardByNumber(number string) (card *Card, ok bool) {

	if isCardInternal(number) {

		var cardInternal *Card = nil

		for _, c := range s.Cards {

			if c.Number == number {
				cardInternal = c
			}
		}

		if cardInternal == nil {

			cardInternal := &Card{
				Balance:  0,
				Currency: "RUB",
				Number:   number,
				Icon:     "card.png",
			}

			s.AddCard(cardInternal)
		}

		return cardInternal, true
	}

	return nil, false
}

// спивывает с карты средства и возвращает текущий баланс и метку выполнена операция или нет
func (s *Service) TransferFromCard(cardFrom *Card, amount int) (balance int, ok bool) {

	result := false

	if cardFrom.Balance >= amount {

		cardFrom.Balance -= amount
		result = true
	}

	return cardFrom.Balance, result
}

// зачисляет на карту средства и возвращает текущий баланс и метку выполнена операция или нет
func (s *Service) TransferToCard(cardTo *Card, amount int) (balance int, ok bool) {

	cardTo.Balance += amount

	return cardTo.Balance, true
}

// устанаваливает наименование банка
func SetBankName(card *Card, bankName string) {

	card.issuer = bankName
}

// возвращает метку принадлежит карта нашему банку или нет
func isCardInternal(number string) bool {

	if strings.HasPrefix(number, "510621") {

		return true
	}

	return false
}

// возвращает метку валидности номера карты поалгоритмы Луна (упрощенному)
func (s *Service) CheckCardNumberByLuna(number string) bool {

	number = strings.ReplaceAll(number, " ", "")

	numberInString := strings.Split(number, "")

	numberInNumders := make([]int, 0)

	for s := range numberInString {

		if n, err := strconv.Atoi(numberInString[s]); err == nil {

			numberInNumders = append(numberInNumders, n)
		} else {

			return false
		}
	}

	sum := 0

	for n := range numberInNumders {

		if (n+1)%2 > 0 {

			numberInNumders[n] = numberInNumders[n] * 2

			if numberInNumders[n] > 9 {

				numberInNumders[n] = numberInNumders[n] - 9
			}
		}

		sum += numberInNumders[n]
	}

	if (((sum % 10) - 10) * -1) == 10 {

		return true
	}

	return false
}
