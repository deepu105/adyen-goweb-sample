package client

import "net/http"

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
	AdditionalData        interface{} `json:"additionalData,omitempty"`
	AllowedPaymentMethods []string    `json:"allowedPaymentMethods,omitempty"`
	BlockedPaymentMethods []string    `json:"blockedPaymentMethods,omitempty"`
	Amount                struct {
		Currency string `json:"currency,omitempty"`
		Value    int64  `json:"value,omitempty"`
	} `json:"amount,omitempty"`
	Channel                   string `json:"channel,omitempty"`
	CountryCode               string `json:"countryCode,omitempty"`
	MerchantAccount           string `json:"merchantAccount,omitempty"`
	ShopperLocale             string `json:"shopperLocale,omitempty"`
	ShopperReference          string `json:"shopperReference,omitempty"`
	ThreeDSAuthenticationOnly bool   `json:"threeDSAuthenticationOnly,omitempty"`
}

type PaymentMethodsRes struct {
	Groups                 []interface{} `json:"groups,omitempty"`
	OneClickPaymentMethods []interface{} `json:"oneClickPaymentMethods,omitempty"`
	PaymentMethods         []interface{} `json:"paymentMethods,omitempty"`
	StoredPaymentMethods   []interface{} `json:"storedPaymentMethods,omitempty"`
}
