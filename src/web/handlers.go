package web

import (
	"fmt"
	"go-client/src/client"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PaymentDataCookie = "paymentData"

func ClientIP(c *gin.Context) {
	c.Header("Content-Type", "application/text")
	c.JSON(http.StatusOK, c.ClientIP())
	return
}

func handleError(method string, c *gin.Context, err error) {
	log.Printf("Error in %s: %s\n", method, err.Error())
	c.JSON(http.StatusBadRequest, err.Error())
}

// PaymentMethodsHandler retrieves a list of available payment methods from Adyen API
func PaymentMethodsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req client.PaymentMethodsReq

	if err := c.BindJSON(&req); err != nil {
		handleError("PaymentMethodsHandler", c, err)
		return
	}
	res, err := checkoutAPI.PaymentMethods(req)
	if err != nil {
		handleError("PaymentMethodsHandler", c, err)
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

// PaymentsHandler makes payment using Adyen API
func PaymentsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req client.PaymentsReq

	if err := c.BindJSON(&req); err != nil {
		handleError("PaymentsHandler", c, err)
		return
	}
	res, err := checkoutAPI.Payments(req)
	if err != nil {
		handleError("PaymentsHandler", c, err)
		return
	}
	if res.Action != nil {
		if action, ok := res.Action.(map[string]interface{}); ok {
			c.SetCookie(PaymentDataCookie, action["paymentData"].(string), 3600, "", "localhost", false, true)
		}
	}
	c.JSON(http.StatusOK, res)
	return
}

// PaymentDetailsHandler gets payment details using Adyen API
func PaymentDetailsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req client.PaymentDetailsReq

	if err := c.BindJSON(&req); err != nil {
		handleError("PaymentDetailsHandler", c, err)
		return
	}
	res, err := checkoutAPI.PaymentDetails(req)
	if err != nil {
		handleError("PaymentDetailsHandler", c, err)
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func RedirectHandler(c *gin.Context) {
	var redirect client.Redirect
	log.Println("Redirect received")

	if err := c.ShouldBind(&redirect); err != nil {
		handleError("RedirectHandler", c, err)
		return
	}
	paymentData, err := c.Cookie(PaymentDataCookie)
	if err != nil {
		handleError("RedirectHandler", c, err)
		return
	}
	var details map[string]string
	if redirect.Payload != "" {
		details = map[string]string{
			"payload": redirect.Payload,
		}
	} else {
		details = map[string]string{
			"MD":    redirect.MD,
			"PaRes": redirect.PaRes,
		}
	}
	res, err := checkoutAPI.PaymentDetails(client.PaymentDetailsReq{
		PaymentData: paymentData,
		Details:     details,
	})
	if err != nil {
		handleError("RedirectHandler", c, err)
		return
	}
	if res.PspReference != "" {
    // TODO encode params
		c.Redirect(http.StatusFound, fmt.Sprintf("/?PspReference=%s&ResultCode=%s&RefusalReason=%s", res.PspReference, res.ResultCode, res.RefusalReason))
		return
	}
	c.JSON(res.Status, res)
	return
}
