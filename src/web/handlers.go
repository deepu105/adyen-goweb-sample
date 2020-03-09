package web

import (
	"go-client/src/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaymentMethodsHandler retrieves a list of available payment methods from Adyen API
func PaymentMethodsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	res, err := checkoutAPI.PaymentMethods(client.PaymentMethodsReq{
		CountryCode: "NL",
		Channel:     "Web",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}
