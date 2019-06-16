package main

import (
	"fmt"
	"github.com/chatroom-han/customermanage/model"
	"github.com/chatroom-han/customermanage/service"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView)list() {
	customers := this.customerService.List()
	fmt.Println("--------客户列表--------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("--------客户列表完成--------")
}

func (this *customerView) add() {
	fmt.Println("--------添加客户列表--------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender,age , phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("--------添加客户列表完成--------")
	} else {
		fmt.Println("--------添加客户列表失败--------")
	}
}

func (this *customerView) delete() {
	fmt.Println("--------删除客户列表--------")
	id := -1
	fmt.Println("请选择客户列表的id")
	fmt.Scanln(&id)

	if id == -1 {
		return
	}
	fmt.Println("你确定要删除吗？（y/n）")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	if choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("--------删除客户列表完成--------")
		} else {
			fmt.Println("--------删除客户列表失败，用户不存在--------")
		}
	}
}

func (this *customerView) updata() {
	fmt.Println("--------修改客户列表--------")
	id := -1
	fmt.Println("请输入要修改的id编号(-1退出)：")
	fmt.Scanln(&id)

	if id == -1 {
		return
	}

	index := -1
	index = this.customerService.FindById(id)
	if index == -1 {
		return
	}

	fmt.Println("姓名",)
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(id, name, gender, age , phone, email)
	if this.customerService.Updata(customer) {
		fmt.Println("修改完成")
	} else {
		fmt.Println("你输入的id不存在")
	}
}

func (this *customerView) exit() {
	fmt.Println("你确定要退出吗？")
	choice := ""
	for {
		fmt.Println("请输入（y/n）：")
		fmt.Scanln(&choice)

		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}
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
			//fmt.Println("添加客户")
			this.add()
		case "2":
			//fmt.Println("修改客户")
			this.updata()
		case "3":
			//fmt.Println("删除客户")
			this.delete()
		case "4":
			//fmt.Println("客户列表")
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户信息管理软件")
}

func main() {
	customerView := customerView{
		key: "",
		loop: true,
		customerService: service.NewCustomerService(),
	}
	customerView.mainMenu()
}


