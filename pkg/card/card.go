package card

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

	for _, c := range s.Cards {
		if c.Number == number {
			return c
		}
	}

	return nil
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
