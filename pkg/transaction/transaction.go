package transaction

import "homework/pkg/card"

type Transaction struct {
	Id            int
	Sum           int
	Datetime      int64
	OperationType string
	Status        bool
	CardFrom      *card.Card
	CardTo        *card.Card
}
