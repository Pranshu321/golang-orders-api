package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderName        string             `json:"orderName,omitempty" bson:"orderName,omitempty"`
	OrderDescription bool               `json:"orderDesc" bson:"orderDesc"`
}
