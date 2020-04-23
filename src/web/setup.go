package web

import (
	"net/http"
	"os"

	"github.com/adyen/adyen-go-api-library/src/adyen"
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	client          *adyen.APIClient
	merchantAccount string
)

func Init() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	godotenv.Load("./.env")

	client = adyen.NewClient(&common.Config{
		ApiKey:      os.Getenv("ADYEN_API_KEY"),
		Environment: common.TestEnv,
	})

	merchantAccount = os.Getenv("ADYEN_MERCHANT")

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	router.GET("/redirect", RedirectHandler)
	router.POST("/redirect", RedirectHandler)

	// Setup route group and routes for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}

	api.GET("/clientIP", ClientIP)
	api.POST("/paymentMethods", PaymentMethodsHandler)
	api.POST("/payments", PaymentsHandler)
	api.POST("/paymentDetails", PaymentDetailsHandler)

	// Start and run the server
	router.Run(":3000")
}
