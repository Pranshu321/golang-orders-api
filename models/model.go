package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerID  uuid.UUID  `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	LineItems   []LineItem `json:"line_items,omitempty" bson:"line_items,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	ShippedAt   *time.Time `json:"shipped_at,omitempty" bson:"shipped_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty" bson:"completed_at,omitempty"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id,omitempty" bson:"item_id,omitempty"`
	ItemName string    `json:"item_name,omitempty" bson:"item_name,omitempty"`
	Quantity uint      `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price    uint      `json:"price,omitempty" bson:"price,omitempty"`
}
