package controller

import (
	"log"
	"net/http"
	"pagi/models"
	"pagi/repository"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PR repository.PaymentRepo
}

func (c *PaymentController) PayOrder(e echo.Context) error {
	var req models.PaymentRequest
	err := e.Bind(&req)
	if err != nil {
		log.Println("BIND ERR:", err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	order_id := e.Param("id")

	if req.Channel == "" || !(req.Channel == "akulaku" || req.Channel == "shopeepaylater") || req.UserID == "" || req.PaymentDate == "" {
		return e.JSON(http.StatusBadRequest, "error or missing param")
	}

	res, err := c.PR.Pay(order_id, &req)
	if err != nil {
		log.Println("Pay ERR", err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusCreated, res)
}
