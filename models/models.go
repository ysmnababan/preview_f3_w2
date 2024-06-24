package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}

type Product struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Price float64            `json:"price" bson:"price"`
}

type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ProductID  primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	TotalPrice float64            `json:"total_price" bson:"total_price"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	Status     string             `json:"status" bson:"status"`
	OrderDate  string             `json:"order_date" bson:"order_date,omitempty"`
}

type Payment struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID     primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	PaymentDate string             `json:"payment_date" bson:"payment_date"`
	TotalPrice  float64            `json:"total_price" bson:"total_price"`
	Channel     string             `json:"channel" bson:"channel"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
}

type PaymentRequest struct {
	UserID      string  `json:"user_id" bson:"user_id"`
	PaymentDate string  `json:"payment_date" bson:"payment_date"`
	Channel     string  `json:"channel" bson:"channel"`
}
