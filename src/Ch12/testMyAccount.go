package main

import (
	"fmt"
	"strings"
)

func main() {

	var loop bool = true
	var key string
	var details string = "Type\tBalance\tMoney\tNote\n"
	var balance float64 = 10000.0
	var money float64
	var note string
	var flag bool

	for {

		fmt.Println(strings.Repeat("-", 64))
		fmt.Println("Household income and expenditure accounting software")
		fmt.Println("\t1 Income and expenditure details")
		fmt.Println("\t2 Register income")
		fmt.Println("\t3 Register expenditure")
		fmt.Println("\t4 Exit the software")
		fmt.Printf("Please choose (1-4): ")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("Current income and expenditure details")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("There are currently no income and expenditure details ... come here!")
			}
		case "2":
			fmt.Println("Amount of this income:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("Description of this income:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("Income\t%v\t%v\t%v\n", balance, money, note)
			flag = true
		case "3":
			fmt.Println("Amount of this expenditure:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("Insufficient balance")
				break
			}
			balance -= money
			fmt.Println("Expenditure description:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("Expenditure\t%v\t%v\t%v\n", balance, money, note)
			flag = true
		case "4":
			fmt.Println("Are you sure you want to quit? (y/n)")
			var choice string
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("Your input is incorrect. Please try again. (y/n)")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("Please enter the correct options ...")
		}

		if !loop {
			break
		}

		fmt.Println(strings.Repeat("-", 64))
	}
	fmt.Println("You have exited the family accounting software.")
	fmt.Println(strings.Repeat("-", 64))
}
