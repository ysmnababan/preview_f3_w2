package main

import (
	"context"
	"log"
	"pagi/config"
	"pagi/controller"
	"pagi/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, db := config.Connect(context.TODO(), "preview_f3_w2")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	repo := &repository.Repo{DB: db}
	pc := &controller.PaymentController{PR: repo}

	e.POST("/pay-order/:id", pc.PayOrder)

	// cron job
	c := cron.New()
	// c.AddFunc("* * * * *", func() {
	c.AddFunc("0 13 * * *", func() {
		coll := db.Collection("payments")

		filter := bson.M{
			"$or": []bson.M{
				{"status": nil},
				{"status": "pending"},
			},
		}
		res, err := coll.UpdateMany(
			context.TODO(),
			filter,
			bson.M{"$set": bson.M{"status": "settlement"}},
		)
		if err != nil {
			log.Println("ERR CRON:")
			return
		}
		log.Println(res)
	})
	c.Start()
	e.Logger.Fatal(e.Start(":" + config.Port))
}
