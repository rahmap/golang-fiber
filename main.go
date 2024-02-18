package main

import (
	"PZN/config"
	"PZN/handlers"
	"github.com/gofiber/fiber/v2"
)

//func main() {
//	fmt.Println("rahmap")
//	fmt.Println(sum(1, 3))
//}
//
//func sum(number1 int, number2 int) int {
//	return number1 + number2
//}

func main() {
	app := fiber.New()

	var err error

	err = config.Connect()
	if err != nil {
		return
	}

	app.Get("/user", handlers.GetUser)
	app.Post("/user", handlers.AddUser)
	app.Get("/user/:id", handlers.GetUserById)
	app.Delete("/user/:id", handlers.DeleteUserById)
	app.Put("/user/:id", handlers.UpdateUser)

	err = app.Listen(": 3000")
	if err != nil {
		return
	}
}
