package utils

import (
	"fmt"
	"strings"
)

type familyAccount struct {
	loop    bool
	key     string
	details string
	balance float64
	money   float64
	note    string
	flag    bool
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		true,
		"",
		"Type\tBalance\tMoney\tNote\n",
		10000.0,
		0,
		"",
		false,
	}
}

func (f *familyAccount) showDetails() {
	fmt.Println("Current income and expenditure details")
	if f.flag {
		fmt.Println(f.details)
	} else {
		fmt.Println("There are currently no income and expenditure details ... come here!")
	}
}

func (f *familyAccount) income() {
	fmt.Println("Amount of this income:")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("Income description:")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("Income\t%v\t%v\t%v\n", f.balance, f.money, f.note)
	f.flag = true
}

func (f *familyAccount) pay() {
	fmt.Println("Amount of this expenditure:")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("Insufficient balance")
	} else {
		f.balance -= f.money
		fmt.Println("Expenditure description:")
		fmt.Scanln(&f.note)
		f.details += fmt.Sprintf("Expenditure\t%v\t%v\t%v\n", f.balance, f.money, f.note)
		f.flag = true
	}
}

func (f *familyAccount) exit() {
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
		f.loop = false
	}
}

func (f *familyAccount) MainMenu() {
	for {
		fmt.Println(strings.Repeat("-", 64))

		fmt.Println("Household income and expenditure accounting software")
		fmt.Println("\t1 Income and expenditure details")
		fmt.Println("\t2 Register income")
		fmt.Println("\t3 Register expenditure")
		fmt.Println("\t4 Exit the software")
		fmt.Printf("Please choose (1-4): ")
		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		default:
			fmt.Println("Please enter the correct options ...")
		}

		if !f.loop {
			break
		}

		fmt.Println(strings.Repeat("-", 64))
	}
	fmt.Println("You have exited the family accounting software.")
	fmt.Println(strings.Repeat("-", 64))
}
