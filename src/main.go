package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var checkoutAPI = CheckoutAPI{}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Init the Adyen checkout API
	godotenv.Load("./.env")
	checkoutAPI.Init(ClientConfig{
		MerchantAccount: os.Getenv("ADYEN_MERCHANT"),
		APIKey:          os.Getenv("ADYEN_API_KEY"),
	})

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Setup route group and routes for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}

	api.GET("/paymentMethods", PaymentMethodsHandler)

	// Start and run the server
	router.Run(":3000")
}

// PaymentMethodsHandler retrieves a list of available payment methods from Adyen API
func PaymentMethodsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Jokes handler not implemented yet",
	})
}
