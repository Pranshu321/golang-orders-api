package application

import (
	"encoding/json"
	"net/http"

	"github.com/Pranshu321/orders-api.git/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		message := "Welcome to the orders api"
		json.NewEncoder(w).Encode(map[string]string{"message": message})
	})

	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.Get)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.Update)
	router.Delete("/{id}", orderHandler.Delete)

}
