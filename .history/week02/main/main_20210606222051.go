package main

import (
	"homework/week02/dao"
	"log"
)

func main() {

	dao := dao.NewDao()
	user, err := dao.GetUserName(1)
	if err != nil {
		log.Fatal(err)
	}

}
