package main

import (
	"fmt"

	"github.com/masaya-tn/GO-learning/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")
	fmt.Println(models.Db)

	// u := &models.User{}
	// // fmt.Println(u)
	// u.Name = "next"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ := models.GetUser(1)
	// u.CreateTodo("first todo")

	// u.DeleteUser()

	// u, _ = models.GetUser(1)

	// fmt.Println(u)

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
