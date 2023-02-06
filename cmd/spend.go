package main

import (
	"fetch/points"
	"fetch/transactions"
	"fmt"
	"strconv"
)

func main() {

	var pointsToSpend int
	fmt.Scanf("%d", &pointsToSpend)

	transactionsStruct, err := transactions.FetchCsvFileAsStruct()
	if err != nil {
		fmt.Println(err)
	}

	finalPointsMap, err := points.FetchFinalPointsMap(transactionsStruct, pointsToSpend)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(formatOutput(finalPointsMap))
}

func formatOutput(finalPointsMap map[string]int) string {
	outputString := "{" + "\n"
	for payer, points := range finalPointsMap {
		outputString += payer + " : " + strconv.Itoa(points) + "\n"
	}
	outputString += "}"
	return outputString
}
