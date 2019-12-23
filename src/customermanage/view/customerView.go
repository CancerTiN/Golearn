package main

import (
	"customermanage/model"
	"customermanage/service"
	"fmt"
	"strings"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (cv *customerView) list() {
	customers := cv.customerService.List()
	fmt.Println()
	fmt.Println(strings.Repeat("-", 24), "Customer  List", strings.Repeat("-", 24))
	fmt.Println("Id\tName\tGender\tAge\tPhone\tEmail")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println(strings.Repeat("-", 24), "Customer  List", strings.Repeat("-", 24))
	fmt.Println()
}

func (cv *customerView) add() {
	fmt.Println(strings.Repeat("-", 25), "Add customer", strings.Repeat("-", 25))
	fmt.Println("Name: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("Gender: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("Age: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("Phone: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("Email: ")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println(strings.Repeat("-", 24), "Add  completed", strings.Repeat("-", 24))
	}
}

func (cv *customerView) delete() {
	fmt.Println(strings.Repeat("-", 23), "Delete  customer", strings.Repeat("-", 23))
	fmt.Println("Please enter the customer number to be deleted (exit -1): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("Confirm whether to delete (Y/N): ")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N" {
			break
		} else {
			fmt.Println("Your input is incorrect. Please try again (Y/N): ")
		}
	}
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println(strings.Repeat("-", 23), "Delete completed", strings.Repeat("-", 23))
		} else {
			fmt.Println("The deletion failed and the number entered does not exist.")
		}
	}
}

func (cv *customerView) exit() {
	fmt.Println("Confirm whether to exit (Y/N): ")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key == "y" || cv.key == "N" || cv.key == "n" {
			break
		}
		fmt.Println("Your input is incorrect. Please try again (Y/N): ")
	}
	if cv.key == "Y" || cv.key == "y" {
		cv.loop = false
	}
}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println(strings.Repeat("#", 64))

		fmt.Println("Customer Information Management Software")
		fmt.Println("\t1 Add customer")
		fmt.Println("\t2 Modify customer")
		fmt.Println("\t3 Delete customer")
		fmt.Println("\t4 Customer List")
		fmt.Println("\t5 Exit")
		fmt.Print("please choose (1-5): ")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("Modify customer ...")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("Your input is wrong, please re-enter ...")
		}

		if !cv.loop {
			break
		}

	}
	fmt.Println("You exited the customer relationship management system.")
	fmt.Println(strings.Repeat("#", 64))
}

func main() {
	cv := customerView{}
	cv.loop = true
	cv.customerService = service.NewCustomerService()
	cv.mainMenu()
}
