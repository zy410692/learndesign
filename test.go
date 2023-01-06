package main

import (
	"fmt"
	"learndesign/desp/Factory"
	"learndesign/desp/models"
)

func main() {
	user := Factory.CreateUser(Factory.AdminUser)(1223, "zhangyi").(*models.Admin)
	fmt.Println(user)
}
