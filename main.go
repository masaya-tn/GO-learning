package main

import (
	"fmt"

	"github.com/masaya-tn/GO-learning/app/controllers"
	"github.com/masaya-tn/GO-learning/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	// user, _ := models.GetUserByEmail("f@f.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
