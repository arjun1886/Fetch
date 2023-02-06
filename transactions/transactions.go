package transactions

import (
	"fetch/conf"
	"fetch/types"
	"os"

	"github.com/gocarina/gocsv"
)

// FetchCsvFileAsStruct - Reads csv file as input and returns a pointer to an array of structs which represents the input data
func FetchCsvFileAsStruct() (*[]types.Transactions, error) {

	conf := conf.GetConf()
	transactionsFileName := conf.TransactionsFileName
	reader, err := os.Open(transactionsFileName)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	finalTransactionsArray, err := getTransactionsArrayFromFile(reader)
	if err != nil {
		return nil, err

	}
	return finalTransactionsArray, nil
}

// Converts os file reader to array of transactions struct
func getTransactionsArrayFromFile(reader *os.File) (*[]types.Transactions, error) {

	transactions := []*types.Transactions{}

	if err := gocsv.UnmarshalFile(reader, &transactions); err != nil {
		return nil, err
	}

	// Converting Array of pointers to a pointer to an array of structs for ease of use
	finalTransactionsArray := []types.Transactions{}
	for i := 0; i < len(transactions); i++ {
		finalTransactionsArray = append(finalTransactionsArray, *transactions[i])
	}

	return &finalTransactionsArray, nil
}
