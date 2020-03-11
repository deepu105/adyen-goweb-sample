/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including One-Click and 3D Secure), mobile wallets, and local payment methods (e.g. iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v51/payments ```
 *
 * API version: 51
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package adyenapi
import (
	"time"
)
// ForexQuote struct for ForexQuote
type ForexQuote struct {
	// The account name.
	Account string `json:"account,omitempty"`
	// The account type.
	AccountType string `json:"accountType,omitempty"`
	BaseAmount Amount `json:"baseAmount,omitempty"`
	// The base points.
	BasePoints int32 `json:"basePoints"`
	Buy Amount `json:"buy,omitempty"`
	Interbank Amount `json:"interbank,omitempty"`
	// The reference assigned to the forex quote request.
	Reference string `json:"reference,omitempty"`
	Sell Amount `json:"sell,omitempty"`
	// The signature to validate the integrity.
	Signature string `json:"signature,omitempty"`
	// The source of the forex quote.
	Source string `json:"source,omitempty"`
	// The type of forex.
	Type string `json:"type,omitempty"`
	// The date until which the forex quote is valid.
	ValidTill time.Time `json:"validTill"`
}
