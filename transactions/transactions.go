package transactions

import (
	"fetch/conf"
	"fetch/types"
	"os"

	"github.com/gocarina/gocsv"
)

func FetchCsvFileAsStruct() (*[]types.Transactions, error) {
	conf := conf.GetConf()
	transactionsFileName := conf.TransactionsFileName
	reader, err := os.Open(transactionsFileName)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	transactions := []*types.Transactions{}

	if err := gocsv.UnmarshalFile(reader, &transactions); err != nil {
		return nil, err
	}

	finalTransactionsArray := []types.Transactions{}
	for i := 0; i < len(transactions); i++ {
		finalTransactionsArray = append(finalTransactionsArray, *transactions[i])
	}

	return &finalTransactionsArray, nil
}
