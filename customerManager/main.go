package main

import (
	"go/customerManager/view"
)

func main() {
	// 启动客户管理系统
	view.NewCustomerView().MainMenu()
}
