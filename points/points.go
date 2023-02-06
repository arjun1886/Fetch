package points

import (
	"errors"
	"fetch/types"
	"sort"
)

func FetchFinalPointsMap(transactionsStruct *[]types.Transactions, pointsToSpend int) (map[string]int, error) {

	updatedTransactions, err := spendPoints(transactionsStruct, pointsToSpend)
	if err != nil {
		return nil, err
	}

	finalPointsMap := buildFinalPointsMap(updatedTransactions)
	return finalPointsMap, nil

}

func buildFinalPointsMap(updatedTransactions *[]types.Transactions) map[string]int {

	finalPointsMap := make(map[string]int)
	for i := 0; i < len(*updatedTransactions); i++ {
		_, ok := finalPointsMap[(*updatedTransactions)[i].Payer]
		if ok {
			finalPointsMap[(*updatedTransactions)[i].Payer] += (*updatedTransactions)[i].Points
		} else {
			finalPointsMap[(*updatedTransactions)[i].Payer] = (*updatedTransactions)[i].Points
		}
	}

	return finalPointsMap
}

func spendPoints(transactionsStruct *[]types.Transactions, pointsToSpend int) (*[]types.Transactions, error) {

	sort.Slice(*transactionsStruct, func(i, j int) bool {
		return (*transactionsStruct)[i].Timestamp < (*transactionsStruct)[j].Timestamp
	})

	for i := 0; i < len(*transactionsStruct); i++ {
		if (*transactionsStruct)[i].Points <= pointsToSpend {
			pointsToSpend -= (*transactionsStruct)[i].Points
			(*transactionsStruct)[i].Points = 0
		} else {
			(*transactionsStruct)[i].Points -= pointsToSpend
			pointsToSpend = 0
			break
		}
	}

	if pointsToSpend > 0 {
		return nil, errors.New("Points to be spent is more than the total points available")
	}

	return transactionsStruct, nil
}
