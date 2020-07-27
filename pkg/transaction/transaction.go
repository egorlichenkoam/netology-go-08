package transaction

import (
	"strconv"
)

type Transaction struct {
	XMLName       string `xml:"transaction"`
	Id            int    `json:"id" xml:"id"`
	Amount        int    `json:"amount" xml:"amount"`
	Datetime      int64  `json:"datetime" xml:"datetime"`
	OperationType string `json:"operation_type" xml:"operation_type"`
	Status        bool   `json:"status" xml:"status"`
	Mcc           int    `json:"mcc" xml:"mcc"`
	CardFrom      string `json:"card_from" xml:"card_from"`
	CardTo        string `json:"card_to" xml:"card_to"`
}

type Transactions struct {
	XMLName      string `xml:"transactions"`
	Transactions []*Transaction
}

func (t *Transaction) MapRowToTransaction(content []string) (err error) {

	for key, value := range content {

		if key == 0 {

			t.Id, err = strconv.Atoi(value)

			if err != nil {

				return err
			}
		} else if key == 1 {

			t.Amount, err = strconv.Atoi(value)

			if err != nil {

				return err
			}
		} else if key == 2 {

			t.Mcc, err = strconv.Atoi(value)

			if err != nil {

				return err
			}
		} else if key == 3 {

			t.OperationType = value
		} else if key == 4 {

			t.CardFrom = value
		} else if key == 5 {

			t.CardTo = value
		} else if key == 6 {

			val, cerr := strconv.Atoi(value)

			if cerr != nil {

				return cerr
			}

			t.Datetime = int64(val)
		}
	}

	return nil
}

// конвертация транзакции в строку
func (t Transaction) String() (result []string) {

	result = make([]string, 0)

	result = append(result, strconv.Itoa(t.Id))
	result = append(result, strconv.Itoa(t.Amount))
	result = append(result, strconv.Itoa(t.Mcc))
	result = append(result, t.OperationType)
	result = append(result, t.CardFrom)
	result = append(result, t.CardTo)
	result = append(result, strconv.FormatInt(t.Datetime, 10))

	return result
}
