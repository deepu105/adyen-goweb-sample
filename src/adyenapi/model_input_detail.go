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
// InputDetail struct for InputDetail
type InputDetail struct {
	// Configuration parameters for the required input.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// Input details can also be provided recursively.
	Details []SubInputDetail `json:"details,omitempty"`
	// Input details can also be provided recursively (deprecated).
	InputDetails []SubInputDetail `json:"inputDetails,omitempty"`
	// In case of a select, the URL from which to query the items.
	ItemSearchUrl string `json:"itemSearchUrl,omitempty"`
	// In case of a select, the items to choose from.
	Items []Item `json:"items,omitempty"`
	// The value to provide in the result.
	Key string `json:"key,omitempty"`
	// True if this input value is optional.
	Optional bool `json:"optional,omitempty"`
	// The type of the required input.
	Type string `json:"type,omitempty"`
	// The value can be pre-filled, if available.
	Value string `json:"value,omitempty"`
}
