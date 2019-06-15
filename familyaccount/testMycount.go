package main

import (
	"fmt"
)

func main() {
	var key = ""
	var loop = true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支金额\t说明"
	flag := true

	for {
		fmt.Println("\n--------家庭收支记账软件--------")
		fmt.Println("        1.收支明细")
		fmt.Println("        2.登记收入")
		fmt.Println("        3.登记支出")
		fmt.Println("        4.退出软件")

		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("--------当前收支记录明细-------")
				if !flag {
					fmt.Println(details)
				} else {
					fmt.Println("当前没有收支，来一笔吧")
				}

			case "2":
				fmt.Println("本次的收入金额")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次的收入说明")
				fmt.Scanln(&note)

				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = false
			case "3":
				fmt.Println("本次的支出金额")
				fmt.Scanln(&money)

				if money > balance {
					fmt.Println("你的余额不足")
					break
				} else  {
					balance -= money
				}

				fmt.Println("本次的支出说明")
				fmt.Scanln(&note)

				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = false
			case "4":
				fmt.Println("你确定要退出吗？（y/n）")
				res := ""
				for {
					fmt.Scanln(&res)
					if res =="y" || res == "n"{
						break
					} else {
						fmt.Println("输入有误，请重新输入")
					}
				}

				if res == "y" {
					loop =false
				}
			default:
				fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件")
}


