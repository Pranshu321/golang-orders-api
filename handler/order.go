package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating an order")
}

func (o *Order) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting an order")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting an order by id")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Updating an order")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Deleting an order")
}
