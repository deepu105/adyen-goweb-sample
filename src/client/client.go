package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Init creates a new CheckoutAPI client if it doesn't exist already
func (api *CheckoutAPI) Init(config ClientConfig) (*CheckoutAPI, error) {
	if api == nil {
		api = &CheckoutAPI{}
	}
	if config == (ClientConfig{}) {
		return nil, fmt.Errorf("A valid configuration is required")
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("Adyen API key is required in the configuration")
	}
	if config.MerchantAccount == "" {
		return nil, fmt.Errorf("Adyen Merchant Account is required in the configuration")
	}
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	api.client = httpClient
	api.config = config
	// Set the API URL to use according to environment
	if config.Environment == "live" {
		api.BaseURL = "https://checkout.adyen.com/v51/" // TODO use correct URL
	} else {
		api.BaseURL = "https://checkout-test.adyen.com/v51/"
	}
	return api, nil
}

// MakeHTTPRequest wraps the common behavior of HTTP requests
func (api *CheckoutAPI) MakeHTTPRequest(req interface{}, httpMethod, url string) ([]byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(httpMethod, api.BaseURL+url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	updateHeaders(httpReq, api.config.APIKey)

	log.Printf("Request for %s API: %s\n", url, string(reqBody))

	resp, err := api.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// PaymentMethods will fetch all availabel payment methods supported by the Adyen API
func (api *CheckoutAPI) PaymentMethods(req PaymentMethodsReq) (PaymentMethodsRes, error) {
	req.MerchantAccount = api.config.MerchantAccount
	paymentRes := PaymentMethodsRes{}
	res, err := api.MakeHTTPRequest(&req, http.MethodPost, "paymentMethods")
	if err != nil {
		return paymentRes, err
	}
	json.Unmarshal(res, &paymentRes)
	return paymentRes, nil
}

// Payments will submit a payment to the Adyen API
func (api *CheckoutAPI) Payments(req PaymentsReq) (PaymentsRes, error) {
	req.MerchantAccount = api.config.MerchantAccount
	paymentRes := PaymentsRes{}
	res, err := api.MakeHTTPRequest(&req, http.MethodPost, "payments")
	if err != nil {
		return paymentRes, err
	}
	json.Unmarshal(res, &paymentRes)
	return paymentRes, nil
}

// PaymentDetails will submit a payment to the Adyen API
func (api *CheckoutAPI) PaymentDetails(req PaymentDetailsReq) (PaymentsRes, error) {
	paymentRes := PaymentsRes{}
	res, err := api.MakeHTTPRequest(&req, http.MethodPost, "payments/details")
	if err != nil {
		return paymentRes, err
	}
	json.Unmarshal(res, &paymentRes)
	return paymentRes, nil
}

func updateHeaders(req *http.Request, APIKey string) {
	req.Header.Add("x-API-key", APIKey)
	req.Header.Add("content-type", "application/json")
}
