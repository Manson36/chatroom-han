package main

import (
	"fmt"
	"github.com/chatroom-han/customermanage/service"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this customerView)list() {
	customers := this.customerService.List()
	fmt.Println("--------客户列表--------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("--------客户列表完成--------")
}

func (this customerView) mainMenu() {
	for {
		fmt.Println("\n--------客户信息管理软件--------")
		fmt.Println("          1.添加客户")
		fmt.Println("          2.修改客户")
		fmt.Println("          3.删除客户")
		fmt.Println("          4.客户列表")
		fmt.Println("          5.退   出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			//fmt.Println("客户列表")
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户信息管理软件5")
}

func main() {
	customerView := customerView{
		key: "",
		loop: true,
		customerService: service.NewCustomerService(),
	}
	customerView.mainMenu()
}


