package service

import "customermanage/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.InitCustomer(1, "ZhangSan", "male", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	} else {
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		return true
	}
}

func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i, v := range cs.customers {
		if v.Id == id {
			index = i
		}
	}
	return index
}
