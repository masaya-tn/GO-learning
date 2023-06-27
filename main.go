package main

import (
	"fmt"

	"github.com/masaya-tn/GO-learning/app/controllers"
	"github.com/masaya-tn/GO-learning/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
