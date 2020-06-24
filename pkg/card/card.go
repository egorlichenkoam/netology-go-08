package card

type Card struct {
	id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
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
