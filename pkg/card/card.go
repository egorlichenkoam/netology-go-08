package card

// карта
type Card struct {
	id       int
	issuer   string
	Balance  int
	Currency string
	Number   string
	Icon     string
	external bool
}

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
func (s *Service) FindCardByNumber(number string) (card *Card, ok bool) {

	for _, c := range s.Cards {
		if c.Number == number {
			return c, true
		}
	}

	return nil, false
}

// спивывает с карты средства и возвращает текущий баланс и метку выполнена операция или нет
func (s *Service) TransferFromCard(cardFrom *Card, amount int) (balance int, ok bool) {

	result := false

	if cardFrom.Balance >= amount || cardFrom.external {

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

// возвращает наименование банка
func getBankName(card *Card) (string, bool) {

	if len(card.issuer) > 0 {

		return card.issuer, true
	}

	return "", false
}

// задает метку карты внещнего банка
func SetExternalBank(card *Card, external bool) {

	card.external = external
}
