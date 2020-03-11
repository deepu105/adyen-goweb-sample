package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutAPI_Init(t *testing.T) {
	type args struct {
		config ClientConfig
		api    *CheckoutAPI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should fail when config doesn't exist",
			args{
				ClientConfig{},
				nil,
			},
			true,
		},
		{
			"should fail when APIkey doesn't exist",
			args{
				ClientConfig{MerchantAccount: "123"},
				nil,
			},
			true,
		},
		{
			"should fail when MerchantAccount doesn't exist",
			args{
				ClientConfig{APIKey: "123"},
				nil,
			},
			true,
		},
		{
			"should create and init API instance when doesn't exist",
			args{
				ClientConfig{APIKey: "123", MerchantAccount: "1234"},
				nil,
			},
			false,
		},
		{
			"should init API instance when it exists",
			args{
				ClientConfig{APIKey: "123", MerchantAccount: "1234"},
				&CheckoutAPI{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := tt.args.api
			got, err := api.Init(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckoutAPI.Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("CheckoutAPI.Init() = %v, want non nil", got)
			}
		})
	}
}

func TestCheckoutAPI_MakeHTTPRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	jsonRes := `
    {
      "groups": [
        {
          "name": "Credit Card",
          "types": [
            "mc",
            "visa",
            "amex"
          ]
        }
      ],
      "paymentMethods": [
        {
          "details": [
            {
              "items": [
                {
                  "id": "1121",
                  "name": "Test Issuer"
                },
                {
                  "id": "1154",
                  "name": "Test Issuer 5"
                }
              ],
              "key": "issuer",
              "type": "select"
            }
          ],
          "name": "iDEAL",
          "supportsRecurring": true,
          "type": "ideal"
        },
        {
          "brands": [
            "mc",
            "visa",
            "amex"
          ],
          "details": [
            {
              "key": "number",
              "type": "text"
            },
            {
              "key": "expiryMonth",
              "type": "text"
            },
            {
              "key": "expiryYear",
              "type": "text"
            },
            {
              "key": "cvc",
              "type": "text"
            },
            {
              "key": "holderName",
              "optional": true,
              "type": "text"
            }
          ],
          "name": "Credit Card",
          "type": "scheme"
        }
      ]
    }
  `
	// Mock Exact URL match
	httpmock.RegisterResponder("POST", "https://checkout-test.adyen.com/v51/paymentMethods", httpmock.NewStringResponder(200, jsonRes))

	api, _ := (&CheckoutAPI{}).Init(ClientConfig{APIKey: "123", MerchantAccount: "1234"})
	type args struct {
		req        interface{}
		httpMethod string
		url        string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"make a HTTP POST request for paymentMethods",
			args{
				&PaymentMethodsReq{
					CountryCode:   "NL",
					ShopperLocale: "nl-NL",
					Amount: Amount{
						Currency: "EUR",
						Value:    100,
					},
				},
				http.MethodPost,
				"paymentMethods",
			},
			string(jsonRes),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.MakeHTTPRequest(tt.args.req, tt.args.httpMethod, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckoutAPI.MakeHTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}
