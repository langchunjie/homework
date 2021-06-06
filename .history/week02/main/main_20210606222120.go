package main

import (
	"fmt"
	"homework/week02/dao"
	"log"
)

func main() {

	dao := dao.NewDao()
	user, err := dao.GetUserName(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user name: ")
}
