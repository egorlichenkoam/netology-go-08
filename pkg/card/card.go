package card

// карта
type Card struct {
	id           int
	issuer       string
	Balance      int
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
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

	card.issuer = s.BankName

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

func AddTransaction(card *Card, transaction *Transaction) {

	if transaction != nil {
		card.Transactions = append(card.Transactions, *transaction)
	}
}

func SumByMMC(transactions []Transaction, mccs []string) int64 {

	var result int64 = 0

	transactions = filterTransactionsByMMC(transactions, mccs)

	for _, transaction := range transactions {
		result = result + transaction.Sum
	}

	return result
}

func filterTransactionsByMMC(transactions []Transaction, mccs []string) []Transaction {

	result := make([]Transaction, 0)

	for _, transaction := range transactions {
		for _, mcc := range mccs {
			if transaction.Mcc == mcc {
				result = append(result, transaction)
			}
		}
	}

	return result
}

func TranslateMCC(code string) string {

	result := "Категория не указана"

	mcc := map[string]string{
		"5411": "Типа магазин",
		"0000": "ОПлата услуг сверхсекретного агента",
		"5812": "Кто девушку платит тот ее и танцует",
		"5555": "Жижино три тотора",
	}

	value, ok := mcc[code]

	if ok {

		result = value
	}

	return result
}

func LastNTransactions(card *Card, n int) []Transaction {

	if len(card.Transactions) < n {
		n = len(card.Transactions)
	}

	nTransactions := make([]Transaction, n)

	n = len(card.Transactions) - n

	copy(nTransactions, card.Transactions[n:len(card.Transactions)])

	for i := len(nTransactions)/2 - 1; i >= 0; i-- {
		opp := len(nTransactions) - 1 - i
		nTransactions[i], nTransactions[opp] = nTransactions[opp], nTransactions[i]
	}

	return nTransactions
}
