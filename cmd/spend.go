package main

import (
	"fetch/points"
	"fetch/transactions"
	"fmt"
	"log"
	"strconv"
)

func main() {

	var pointsToSpend int
	fmt.Println("Enter the number of points you want to spend below :")
	fmt.Scanf("%d", &pointsToSpend)

	transactionsStruct, err := transactions.FetchCsvFileAsStruct()
	if err != nil {
		log.Fatal(err)
	}

	finalPointsMap, err := points.FetchFinalPointsMap(transactionsStruct, pointsToSpend)

	if err != nil {
		log.Fatal(err)
	}

	formattedOutput := formatOutput(finalPointsMap)
	fmt.Println("The final output is : " + formattedOutput)
}

// Formats the final payer to points mapping according to the requirement
func formatOutput(finalPointsMap map[string]int) string {
	outputString := "{" + "\n"
	for payer, points := range finalPointsMap {
		outputString += payer + " : " + strconv.Itoa(points) + "\n"
	}
	outputString += "}"
	return outputString
}
