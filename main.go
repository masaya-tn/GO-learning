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
	// u.Name = "ちょんまげ"
	// u.Email = "mage@example.com"
	// u.PassWord = "magetyon"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ := models.GetUser(1)

	// u.DeleteUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("お仕事")

	// user2, _ := models.GetUser(4)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// fmt.Println(u)

	// t, _ := models.GetTodo(1)
	// t.UpdateTodo("掃除")
	// t, _ = models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
