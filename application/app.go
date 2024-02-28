package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Pranshu321/orders-api.git/controller"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	mongoErr := controller.Connect()
	if mongoErr != nil {
		return fmt.Errorf("error connecting to MongoDB: %s", mongoErr)
	}
	defer func() {
		if err := controller.Connect(); err != nil {
			fmt.Println("Failed to connect from DB: ", err)
		}
	}()
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	ch := make(chan error, 1)
	err := server.ListenAndServe()
	if err != nil {
		close(ch)
		return fmt.Errorf("failed to listen to server: %w", err)
	}
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
