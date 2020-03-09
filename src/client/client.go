package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ADYEN_API is the API URL to use
const ADYEN_API = "https://checkout-test.adyen.com/v51/"

// Init creates a new CheckoutAPI client if it doesn't exist already
func (api *CheckoutAPI) Init(config ClientConfig) (*CheckoutAPI, error) {
	if api == nil {
		api = &CheckoutAPI{}
	}
	if config == (ClientConfig{}) {
		return nil, fmt.Errorf("A configuration is required")
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
	return api, nil
}

// PaymentMethods will fetch all availabel payment methods supported by the Adyen API
func (api *CheckoutAPI) PaymentMethods(req PaymentMethodsReq) (PaymentMethodsRes, error) {
	req.MerchantAccount = api.config.MerchantAccount

	paymentRes := PaymentMethodsRes{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return paymentRes, err
	}

	httpReq, err := http.NewRequest("POST", ADYEN_API+"/paymentMethods", bytes.NewBuffer(reqBody))
	if err != nil {
		return paymentRes, err
	}

	updateHeaders(httpReq, api.config.APIKey)

	resp, err := api.client.Do(httpReq)
	if err != nil {
		return paymentRes, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return paymentRes, err
	}

	json.Unmarshal(body, &paymentRes)

	return paymentRes, nil
}

func updateHeaders(req *http.Request, APIKey string) {
	req.Header.Add("x-API-key", APIKey)
	req.Header.Add("content-type", "application/json")
}
