package repository

import (
	"context"
	"fmt"
	"log"
	"pagi/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepo interface {
	Pay(order_id string, p *models.PaymentRequest) (interface{}, error)
}

type Repo struct {
	DB *mongo.Database
}

func (r *Repo) Pay(order_id string, p *models.PaymentRequest) (interface{}, error) {
	o_id, _ := primitive.ObjectIDFromHex(order_id)
	user_id, _ := primitive.ObjectIDFromHex(p.UserID)

	isUserExist, err := r.IsUserExist(user_id)
	if err != nil || !isUserExist {
		return nil, fmt.Errorf("user not found")
	}

	isOrderExist, err := r.IsOrderExist(o_id, user_id)
	if err != nil || !isOrderExist {
		return nil, fmt.Errorf("order not found")
	}

	var payment models.Payment
	payment.PaymentDate = p.PaymentDate
	payment.Channel = p.Channel
	if p.Channel == "akulaku" {
		payment.TotalPrice = 50000.0
	} else {
		payment.TotalPrice = 75000.0
	}

	payment.OrderID = o_id

	res, err := r.DB.Collection("payments").InsertOne(context.TODO(), payment)
	if err != nil {
		log.Println("ERR :", err)
		return nil, fmt.Errorf("error query")
	}

	return res, nil
}
