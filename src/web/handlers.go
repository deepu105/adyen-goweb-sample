package web

import (
	"go-client/src/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaymentMethodsHandler retrieves a list of available payment methods from Adyen API
func PaymentMethodsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req client.PaymentMethodsReq
	c.BindJSON(&req)
	res, err := checkoutAPI.PaymentMethods(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// PaymentsHandler makes payment using Adyen API
func PaymentsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req client.PaymentsReq
	c.BindJSON(&req)
	res, err := checkoutAPI.Payments(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}
