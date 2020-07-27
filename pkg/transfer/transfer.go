package transfer

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"homework/pkg/card"
	"homework/pkg/transaction"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

// сервис
type Service struct {
	CardSvc      *card.Service
	Transactions []*transaction.Transaction
}

var (
	ErrFromCardNotFound       = errors.New("Source card not found")
	ErrFromCardNumberNotValid = errors.New("Source card number not valid")

	ErrToCardNotFound       = errors.New("Target card not found")
	ErrToCardNumberNotValid = errors.New("Target card number not valid")
)

// конструктор сервиса
func NewService(cardSvc *card.Service) *Service {

	return &Service{CardSvc: cardSvc}
}

func NewServiceWithTransactions(cardSvc *card.Service, transactions []*transaction.Transaction) *Service {

	return &Service{CardSvc: cardSvc, Transactions: transactions}
}

// возвращает ключ сформированный из времени в виде ГОД_МЕСЯЦ
func makeYearMonthKey(unixTime int64) string {

	t := time.Unix(unixTime, 0)

	startYear := strconv.Itoa(t.Year())

	startMonth := strconv.Itoa(int(t.Month()))

	if len(startMonth) == 1 {

		startMonth = "0" + startMonth
	}

	return startYear + "_" + startMonth
}

// возвращает карту mcc кодов c суммами затрат по ним используя mutex и пишет прямо в результатирующую карту
func (s *Service) GetMccTransactionsSumAmountMapWithMutexStraightToMap(transactions []*transaction.Transaction) (result map[int]int) {

	partCount := 10

	wg := sync.WaitGroup{}

	wg.Add(partCount)

	mu := sync.Mutex{}

	result = make(map[int]int)

	partSize := len(transactions) / partCount

	for i := 0; i < partCount; i++ {

		part := transactions[i*partSize : (i+1)*partSize]

		if i == partCount-1 {

			for _, value := range transactions[(i+1)*partSize:] {

				part = append(part, value)
			}
		}

		go func() {

			for _, tx := range part {

				mu.Lock()

				result[tx.Mcc] += tx.Amount

				mu.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return result
}

// возвращает карту mcc кодов c суммами затрат по ним используя Channels
func (s *Service) GetMccTransactionsSumAmountMapWithChannels(transactions []*transaction.Transaction) (result map[int]int) {

	partCount := 10

	result = make(map[int]int)

	chMap := make(chan map[int]int)

	partSize := len(transactions) / partCount

	for i := 0; i < partCount; i++ {

		part := transactions[i*partSize : (i+1)*partSize]

		if i == partCount-1 {

			for _, value := range transactions[(i+1)*partSize:] {

				part = append(part, value)
			}
		}
		go func(chMap chan<- map[int]int) {

			chMap <- s.GetMccTransactionsSumAmountMap(part)

		}(chMap)
	}

	finished := 0

	for value := range chMap {

		for key, value := range value {

			result[key] += value
		}

		finished++

		if finished == partCount {
			break
		}
	}
	return result
}

// возвращает карту mcc кодов c суммами затрат по ним используя mutex
func (s *Service) GetMccTransactionsSumAmountMapWithMutex(transactions []*transaction.Transaction) (result map[int]int) {

	partCount := 10

	wg := sync.WaitGroup{}

	wg.Add(partCount)

	mu := sync.Mutex{}

	result = make(map[int]int)

	partSize := len(transactions) / partCount

	for i := 0; i < partCount; i++ {

		part := transactions[i*partSize : (i+1)*partSize]

		if i == partCount-1 {

			for _, value := range transactions[(i+1)*partSize:] {

				part = append(part, value)
			}
		}
		go func() {

			m := s.GetMccTransactionsSumAmountMap(part)

			mu.Lock()

			for key, value := range m {

				result[key] += value
			}

			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	return result
}

// возвращает карту mcc кодов c суммами затрат по ним
func (s *Service) GetMccTransactionsSumAmountMap(transactions []*transaction.Transaction) (result map[int]int) {

	result = make(map[int]int)

	for _, tx := range transactions {

		result[tx.Mcc] += tx.Amount
	}

	return result
}

// возвращает карту транхакций группированных по месяцу и году
func (s *Service) GetTransactionsGroupedByMonths(card *card.Card, start, end int64) (result map[string][]*transaction.Transaction) {

	if start < end {

		groupedTransactions := make(map[string][]*transaction.Transaction, 0)

		next := time.Unix(start, 0)

		for next.Before(time.Unix(end, 0)) {

			groupedTransactions[makeYearMonthKey(next.Unix())] = make([]*transaction.Transaction, 0)

			next = next.AddDate(0, 1, 0)
		}

		groupedTransactions[makeYearMonthKey(end)] = make([]*transaction.Transaction, 0)

		transactions := s.GetTransactionsByType(card, "from")

		for n := range transactions {

			tx := transactions[n]

			mapKey := makeYearMonthKey(tx.Datetime)

			for key, value := range groupedTransactions {

				if key == mapKey {

					groupedTransactions[key] = append(value, tx)
				}
			}
		}

		return groupedTransactions
	}

	return nil
}

// возвращает список транзакций карты по типу (from - расход, to - приход)
func (s *Service) GetTransactionsByTypeAndOwner(owner string, operationType string) (transactions []*transaction.Transaction) {

	result := make([]*transaction.Transaction, 0)

	for _, c := range s.CardSvc.OwnerCards(owner) {

		if c.Owner == owner {

			result = append(result, s.GetTransactionsByType(c, operationType)...)
		}
	}

	return result
}

// возвращает список транзакций карты по типу (from - расход, to - приход)
func (s *Service) GetTransactionsByType(card *card.Card, operationType string) (transactions []*transaction.Transaction) {

	result := make([]*transaction.Transaction, 0)

	for n := range s.Transactions {

		tx := s.Transactions[n]

		if tx.CardFrom == card.Number && tx.OperationType == operationType {

			result = append(result, tx)
		}
	}

	return result
}

// возвращает сортированный по amount список транзакций карты по типу (from - расход, to - приход)
func (s *Service) GetSortedTransactionsByType(card *card.Card, operationType string) (transactions []*transaction.Transaction) {

	result := s.GetTransactionsByType(card, operationType)

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Amount > result[j].Amount
	})

	return result
}

// добавляет транзакцию
func (s *Service) addTransaction(transaction *transaction.Transaction) {

	transaction.Datetime = time.Now().Unix()

	s.Transactions = append(s.Transactions, transaction)
}

// валидирует номера карт
func (s *Service) validateCards(from, to string) (err error) {

	if !s.isValid(from) {

		return ErrFromCardNumberNotValid
	}

	if !s.isValid(to) {

		return ErrToCardNumberNotValid
	}

	return nil
}

// ищем карты в сервисе
func (s *Service) searchCards(from, to string) (err error, fromCard, toCard *card.Card) {

	// ищем карту с которой будем преводить
	toCard = s.CardSvc.FindCardByNumber(to)

	if toCard == nil {

		err = ErrToCardNotFound
	}

	// ищем карту с которой будем преводить
	fromCard = s.CardSvc.FindCardByNumber(from)

	if fromCard == nil {

		err = ErrFromCardNotFound
	}

	return err, fromCard, toCard
}

// перевод с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (err error) {

	if err = s.validateCards(from, to); err != nil {

		return err
	}

	err, fromCard, toCard := s.searchCards(from, to)

	if err != nil {

		return err
	}

	// процент за перевод и минимальная коммисия
	transferFeePercentage, transferFeeMin := transferFeePercentageAndMinimum(fromCard, toCard)

	totalToTransfer := amountPlusCommission(amount, transferFeePercentage, transferFeeMin)

	if err = s.CardSvc.Transfer(fromCard, totalToTransfer, true); err != nil {

		return err
	}

	// транзакция для списания
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Amount:        totalToTransfer,
		OperationType: "from",
		Status:        true,
		CardFrom:      fromCard.Number,
		CardTo:        toCard.Number,
	})

	_ = s.CardSvc.Transfer(toCard, amount, false)

	// транзакция для зачисления
	s.addTransaction(&transaction.Transaction{
		Id:            0,
		Amount:        amount,
		OperationType: "to",
		Status:        true,
		CardFrom:      fromCard.Number,
		CardTo:        toCard.Number,
	})

	return nil
}

// возвращает процент коммисии за перевод и минимальную коммисию за перевод
func transferFeePercentageAndMinimum(cardFrom, cardTo *card.Card) (transferFeePercentage float64, transferFeeMin int) {

	if cardFrom == nil && cardTo == nil {

		return 1.5, 3000
	}

	if cardFrom != nil && cardTo == nil {

		return 0.5, 1000
	}

	return 0.0, 0
}

// возвращает сумму для списания с карты с учтом комиссии
func amountPlusCommission(amount int, transferFeePercentage float64, transferFeeMin int) (total int) {

	internalCommission := int(float64(amount) / 100 * transferFeePercentage)

	if internalCommission < transferFeeMin {

		internalCommission = transferFeeMin
	}

	return amount + internalCommission
}

func (s *Service) isValid(number string) bool {

	return s.CardSvc.CheckCardNumberByLuna(number)
}

// экспортируем транзакции в json файл
func (s *Service) ExportTransactionsToJson() (err error) {

	err = os.Chdir("temp")

	if err != nil {

		log.Println("Can not open temp catalog", err)

		err = os.Mkdir("temp", os.ModePerm)

		if err != nil {

			log.Println("Can not create temp directory", err)

			return err
		} else {

			err = os.Chdir("temp")
		}
	}

	file, err := os.Create("exports.json")

	if err != nil {

		log.Println("Can not create file", err)

		return err
	}

	defer func(c io.Closer) {

		if cerr := c.Close(); cerr != nil {

			log.Println(cerr)

			err = cerr
		}
	}(file)

	encoder := json.NewEncoder(file)

	err = encoder.Encode(s.Transactions)

	if err != nil {

		log.Println("Can not write to file", err)
	}

	return err
}

// импортирует транзакции из json файла
func (s *Service) ImportTransactionsFromJson(filePath string) (err error) {

	reader, err := os.Open(filePath)

	if err != nil {

		log.Println("Can not open import transactions file", err)

		return err
	}

	defer func(c io.Closer) {

		if cerr := c.Close(); cerr != nil {

			log.Println(cerr)

			err = cerr
		}
	}(reader)

	err = json.NewDecoder(reader).Decode(&s.Transactions)

	if err != nil {

		log.Println("Can not read data from import file", err)

		return err
	}

	return nil
}

// экспортирует транзакции в файл
func (s *Service) ExportTransactions() (err error) {

	err = os.Chdir("temp")

	if err != nil {

		log.Println("Can not open temp catalog", err)

		err = os.Mkdir("temp", os.ModePerm)

		if err != nil {

			log.Println("Can not create temp directory", err)

			return err
		} else {

			err = os.Chdir("temp")
		}
	}

	file, err := os.Create("exports.csv")

	if err != nil {

		log.Println("Can not create file", err)

		return err
	}

	defer func(c io.Closer) {

		if cerr := c.Close(); cerr != nil {

			log.Println(cerr)

			err = cerr
		}
	}(file)

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, tx := range s.Transactions {

		err := writer.Write(tx.String())

		if err != nil {

			log.Println("Can not write to file", err)
		}
	}

	return err
}

// импортирует транзакции из файла
func (s *Service) ImportTransactions(filePath string) (err error) {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {

		log.Println("Can not open import transactions file", err)

		return err
	}

	reader := csv.NewReader(bytes.NewReader(data))

	records, err := reader.ReadAll()

	if err != nil {

		log.Println("Can not read data from import file", err)

		return err
	}

	for _, content := range records {

		tx := transaction.Transaction{}

		err = tx.MapRowToTransaction(content)

		if err != nil {

			log.Println("Can not import transaction", err, content)
		} else {

			s.Transactions = append(s.Transactions, &tx)
		}
	}

	return nil
}
