package main

import (
	"context"
	"fmt"

	"github.com/Pranshu321/orders-api.git/application"
)

func main() {
	fmt.Println("Starting the application...")
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting the application: %s\n", err)
	}
}
