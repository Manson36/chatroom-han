package service

import "github.com/chatroom-han/customermanage/model"

type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, 22, "张三", "男",
		"22222222", "222@souhu.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}
