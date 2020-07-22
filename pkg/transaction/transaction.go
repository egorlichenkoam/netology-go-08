package transaction

import "homework/pkg/card"

type Transaction struct {
	Id            int
	Amount        int
	Datetime      int64
	OperationType string
	Status        bool
	Mcc           int
	CardFrom      *card.Card
	CardTo        *card.Card
}
