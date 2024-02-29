package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Pranshu321/orders-api.git/controller"
	"github.com/Pranshu321/orders-api.git/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID         `json:"customer_id"`
		LineItems  []models.LineItem `json:"line_items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	now := time.Now().UTC()
	order := models.Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err := controller.Insert(order)
	if err != nil {
		http.Error(w, "failed to insert order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "failed to marshal order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (o *Order) Get(w http.ResponseWriter, r *http.Request) {
	ans, err := controller.FindAll()
	if err != nil {
		http.Error(w, "failed to get orders", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(ans)
	if err != nil {
		http.Error(w, "failed to marshal order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(id, base, bitSize)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ans, err := controller.FindById(orderId)
	if err != nil {
		http.Error(w, "failed to get order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(ans)
	if err != nil {
		http.Error(w, "failed to marshal order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	var body struct {
		LineItems []models.LineItem `json:"line_items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(id, base, bitSize)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prevorder, err := controller.FindById(orderId)
	// now := time.Now().UTC()
	order := models.Order{
		OrderID:    orderId,
		CustomerID: prevorder.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  prevorder.CreatedAt,
	}

	err = controller.Update(order)
	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "failed to marshal order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(id, base, bitSize)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = controller.Delete(orderId)
	if err != nil {
		http.Error(w, "failed to delete order", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Order Deleted")
}
