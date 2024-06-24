package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repo) IsProductExist(product_id primitive.ObjectID) (bool, error) {
	var result bson.M
	err := r.DB.Collection("products").FindOne(context.TODO(), bson.M{"_id": product_id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Println("ERR:", err)
		return false, err
	}
	return true, nil
}
