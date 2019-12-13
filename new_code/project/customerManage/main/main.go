package main

import (
	"go_dev/new_code/project/customerManage/view"
)

func main()  {


	customerView := view.NewCustomerView()
	customerView.ShowMainMenu()
	/*customerView.Key = ""
	customerView.IsExist = false

	customerView.CustomerService = service.NewCustomerService()*/




}
