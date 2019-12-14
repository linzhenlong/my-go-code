package main

import (
	"github.com/linzhenlong/my-go-code/new_code/project/customerManage/view"
)

func main()  {


	customerView := view.NewCustomerView()
	customerView.ShowMainMenu()
	/*customerView.Key = ""
	customerView.IsExist = false

	customerView.CustomerService = service.NewCustomerService()*/




}
