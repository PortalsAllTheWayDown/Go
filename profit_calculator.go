package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Revenue: ")
	if revenue <= 0 {
		panic("Revenue can not be less than zero")
	}
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := profitCalc(revenue, expenses, taxRate)
	fmt.Printf(`Future Value: %.1f
EBT: %.1f
Ratio: %.3f`, profit, ebt, ratio)
	writeResultsToFile(profit, ebt, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func profitCalc(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func writeResultsToFile(profit, ebt, ratio float64) {
	results := fmt.Sprint("Profit: ", profit, "\nEBT: ", ebt, "\nRatio: ", ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
