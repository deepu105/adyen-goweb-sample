package web

import (
	"fmt"
	"go-client/src/client"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var checkoutAPI *client.CheckoutAPI

func Init() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	var err error
	// Init the Adyen checkout API
	godotenv.Load("./.env")
	checkoutAPI, err = checkoutAPI.Init(client.ClientConfig{
		MerchantAccount: os.Getenv("ADYEN_MERCHANT"),
		APIKey:          os.Getenv("ADYEN_API_KEY"),
	})

	if err != nil {
		fmt.Printf("Error initializing API client: %s", err.Error())
		os.Exit(1)
	}

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
