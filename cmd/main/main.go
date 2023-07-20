package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sumyabazar93/go-bookstore/pkg/routes"
)

func main(){
	r := fiber.New()
	routes.RegisterBookStore(r)

	port := 9010

	fmt.Printf("Server started on http://localhost:%d\n", port)

	if err := r.Listen(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Println("Error:", err)
	}
}