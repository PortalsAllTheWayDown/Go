package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil

}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----")
		panic("Can't continue sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greated than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! new amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Printf("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! new amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

		// wantsCheckBalance := choice == 1

		// if wantsCheckBalance {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greated than 0.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! new amount: ", accountBalance)
		// } else if choice == 3 {
		// 		fmt.Print("Your withdrawal: ")
		// 		var withdrawalAmount float64
		// 		fmt.Scan(&withdrawalAmount)

		// 		if withdrawalAmount <= 0 {
		// 			fmt.Printf("Invalid amount. Must be greater than 0.")
		// 			continue
		// 		}

		// 		if withdrawalAmount > accountBalance {
		// 			fmt.Println("Invalid amount. You cant withdraw more than you have.")
		// 			continue
		// 		}

		// 		accountBalance -= withdrawalAmount
		// 		fmt.Println("Balance updated! new amount: ", accountBalance)
		// 	} else {
		// 		fmt.Println("Goodbye!")
		// 		break
		// 	}
		// }
	}
}
