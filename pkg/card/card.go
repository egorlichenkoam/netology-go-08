package card

import (
	"errors"
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

var (
	ErrFromCardNotEnoughMoney = errors.New("Source card: not enough money")
)

// сервис
type Service struct {
	BankName string
	Cards    []*Card
}

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
func (s *Service) FindCardByNumber(number string) (card *Card) {

	if isCardInternal(number) {

		var cardInternal *Card = nil

		for _, c := range s.Cards {

			if c.Number == number {
				cardInternal = c
			}
		}

		return cardInternal
	}

	return nil
}

// перенос средст: transferFrom = true - с карты, transferFrom = false - на карту
func (s *Service) Transfer(card *Card, amount int, transferFrom bool) (err error) {

	if transferFrom {

		if card.Balance >= amount {

			card.Balance -= amount
		} else {

			err = ErrFromCardNotEnoughMoney
		}
	} else {

		card.Balance += amount
	}

	return err
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

	return true
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
