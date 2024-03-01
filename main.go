package main

import (
	"context"
	"fmt"

	"github.com/Pranshu321/orders-api.git/application"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting the application...")
	enverr := godotenv.Load()
	if enverr != nil {
		fmt.Println("Error loading .env file")
	}
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting the application: %s\n", err)
	}
}
