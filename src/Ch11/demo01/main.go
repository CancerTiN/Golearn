package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposit(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("The password you entered is incorrect.")
		return
	}

	if money <= 0 {
		fmt.Println("The amount you entered is incorrect.")
		return
	}

	account.Balance += money
	fmt.Println("Deposit successful.")
}

func (account *Account) Withdraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("The password you entered is incorrect.")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("The amount you entered is incorrect.")
		return
	}

	account.Balance -= money
	fmt.Println("Withdrawal succeeded.")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("The password you entered is incorrect.")
		return
	}

	fmt.Printf("Your account number is %v and your balance is %v.\n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		"gs1111111",
		"666666",
		100.0,
	}

	account.Query("666666")
	account.Deposit(200.0, "666666")
	account.Query("666666")
	account.Withdraw(150.0, "666666")
	account.Query("666666")
}
