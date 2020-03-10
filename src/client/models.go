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
		Currency string  `json:"currency,omitempty"`
		Value    float64 `json:"value,omitempty"`
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
	Status                 interface{}   `json:"status,omitempty"`
	ErrorCode              string        `json:"errorCode,omitempty"`
	Message                string        `json:"message,omitempty"`
}

type PaymentsReq struct {
	AccountInfo    interface{} `json:"accountInfo,omitempty"`
	AdditionalData interface{} `json:"additionalData,omitempty"`
	Amount         struct {
		Currency string  `json:"currency,omitempty"`
		Value    float64 `json:"value,omitempty"`
	} `json:"amount,omitempty"`
	Reference       string      `json:"reference,omitempty"`
	PaymentMethod   interface{} `json:"paymentMethod,omitempty"`
	ReturnURL       string      `json:"returnUrl,omitempty"`
	Channel         string      `json:"channel,omitempty"`
	CountryCode     string      `json:"countryCode,omitempty"`
	MerchantAccount string      `json:"merchantAccount,omitempty"`
	BrowserInfo     interface{} `json:"browserInfo,omitempty"`
	BillingAddress  interface{} `json:"billingAddress,omitempty"`
	ShopperIP       string      `json:"shopperIP,omitempty"`
	Origin          string      `json:"origin,omitempty"`
}

type PaymentsRes struct {
	ResultCode    string        `json:"resultCode,omitempty"`
	RefusalReason string        `json:"refusalReason,omitempty"`
	PspReference  interface{}   `json:"pspReference,omitempty"`
	Action        interface{}   `json:"action,omitempty"`
	Details       []interface{} `json:"details,omitempty"`
	Status        int           `json:"status,omitempty"`
	ErrorCode     string        `json:"errorCode,omitempty"`
	Message       string        `json:"message,omitempty"`
}

type PaymentDetailsReq struct {
	Details                   interface{} `json:"details,omitempty"`
	PaymentData               string      `json:"paymentData,omitempty"`
	ThreeDSAuthenticationOnly bool        `json:"threeDSAuthenticationOnly,omitempty"`
}

type Redirect struct {
	MD      string
	PaRes   string
	Payload string `form:"payload"`
}
