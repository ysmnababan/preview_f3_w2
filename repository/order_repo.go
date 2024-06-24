package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repo) IsOrderExist(order_id, user_id primitive.ObjectID) (bool, error) {
	var result bson.M
	err := r.DB.Collection("orders").FindOne(context.TODO(), bson.M{"_id": order_id, "user_id": user_id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("ISORDER:", order_id, user_id)
			return false, nil
		}
		log.Println("ERR:", err)
		return false, err
	}
	return true, nil
}
