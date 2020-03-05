package main

import (
	"fmt"
	"net/http"
	"time"
)

// ClientConfig is used to setup Adyen API default header values
type ClientConfig struct {
	APIKey          string
	MerchantAccount string
	Environment     string
}

// CheckoutAPI is used to call Adyen checkout API https://docs.adyen.com/api-explorer/#/PaymentSetupAndVerificationService/v51/overview
type CheckoutAPI struct {
	config ClientConfig
	client *http.Client
}

type PaymentMethodsReq struct {
	additionalData        interface{}
	allowedPaymentMethods []string
	blockedPaymentMethods []string
	amount                struct {
		currency string
		value    int64
	}
	channel     string
	countryCode string
	// merchantAccount string
	shopperLocale             string
	shopperReference          string
	threeDSAuthenticationOnly bool
}

type PaymentMethodsRes struct {
	Raw                    string
	Groups                 []interface{}
	OneClickPaymentMethods []interface{}
	PaymentMethods         []interface{}
	StoredPaymentMethods   []interface{}
}

func (api *CheckoutAPI) Init(config ClientConfig) error {
	if config == (ClientConfig{}) {
		return fmt.Errorf("A configuration is required")
	}
	if config.APIKey == "" {
		return fmt.Errorf("Adyen API key is required in the configuration")
	}
	if config.MerchantAccount == "" {
		return fmt.Errorf("Adyen Merchant Account is required in the configuration")
	}
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	api.client = httpClient
	api.config = config
	return nil
}

func (api *CheckoutAPI) PaymentMethods(req PaymentMethodsReq) (PaymentMethodsRes, error) {
	// json.Unmarshal()
	return PaymentMethodsRes{}, nil
}
