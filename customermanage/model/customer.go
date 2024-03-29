package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int, name, gender string, age int,  phone, email string) Customer {
	return Customer{
		Id: id,
		Name:name,
		Gender:gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

func NewCustomer2(name, gender string, age int, phone, email string) Customer {
	return Customer{
		Name:name,
		Gender:gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

func(this Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
}
