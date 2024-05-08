package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %3.3f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %3.3f\n", futureRealValue)
	//fmt.Printf("Future Value: %3.3f\nFuture Value (adjusted for Inflation): %3.3f\n", futureValue, futureRealValue)
	//fmt.Println("Futre Value Adjusted for inflation:", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

//fmt.Scan() has some limitations when it comes to multi-word input values.

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv  or you can just use return
	return
}

// := inits a variable
