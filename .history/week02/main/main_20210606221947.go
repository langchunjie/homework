package main

import "homework/week02/dao"

func main() {

	dao := dao.NewDao()
	user, err := dao.GetUserName(1)


}
