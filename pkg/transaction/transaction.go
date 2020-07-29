package transaction

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
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

		switch key {
		case 0:
			t.Id, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
			break
		case 1:
			t.Amount, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
			break
		case 2:
			t.Mcc, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
			break
		case 3:
			t.OperationType = value
			break
		case 4:
			t.CardFrom = value
			break
		case 5:
			t.CardTo = value
			break
		case 6:
			val, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			t.Datetime = int64(val)
			break
		}
	}
	return nil
}

// конвертация транзакции в строку
func (t Transaction) AsStrings() (result []string) {
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

// экспортируем транзакции в json файл
func ExportTransactionsToJson(transactions []*Transaction) (err error) {
	file, err := os.Create("exports.json")
	if err != nil {
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}(file)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(transactions)
	if err != nil {
		return err
	}
	return nil
}

// импортирует транзакции из json файла
func ImportTransactionsFromJson(filePath string) (transactions []*Transaction, err error) {
	transactions = make([]*Transaction, 0)
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}(reader)
	err = json.NewDecoder(reader).Decode(&transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// экспортируем транзакции в xml файл
func ExportTransactionsToXml(transactions []*Transaction) (err error) {
	file, err := os.Create("exports.xml")
	if err != nil {
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}(file)
	encoder := xml.NewEncoder(file)
	internalTransactions := Transactions{
		Transactions: transactions,
	}
	err = encoder.Encode(&internalTransactions)
	if err != nil {
		return err
	}
	return nil
}

// импортирует транзакции из xml файла
func ImportTransactionsFromXml(filePath string) (transactions []*Transaction, err error) {
	transactions = make([]*Transaction, 0)
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}(reader)
	internalTransactions := Transactions{}
	err = xml.NewDecoder(reader).Decode(&internalTransactions)
	if err != nil {
		return nil, err
	}
	transactions = internalTransactions.Transactions
	return transactions, nil
}

// экспортирует транзакции в csv файл
func ExportTransactionsToCsv(transactions []*Transaction) (err error) {
	file, err := os.Create("exports.csv")
	if err != nil {
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, tx := range transactions {
		err = writer.Write(tx.AsStrings())
		if err != nil {
			return err
		}
	}
	return nil
}

// импортирует транзакции из csv файла
func ImportTransactionsFromCsv(filePath string) (transactions []*Transaction, err error) {
	transactions = make([]*Transaction, 0)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, content := range records {
		tx := Transaction{}
		err = tx.MapRowToTransaction(content)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &tx)
	}
	return transactions, nil
}
