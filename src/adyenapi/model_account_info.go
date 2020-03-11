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
// AccountInfo struct for AccountInfo
type AccountInfo struct {
	// Indicator for the length of time since this shopper account was created in the merchant's environment. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountAgeIndicator string `json:"accountAgeIndicator,omitempty"`
	// Date when the shopper's account was last changed.
	AccountChangeDate time.Time `json:"accountChangeDate,omitempty"`
	// Indicator for the length of time since the shopper's account was last updated. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountChangeIndicator string `json:"accountChangeIndicator,omitempty"`
	// Date when the shopper's account was created.
	AccountCreationDate time.Time `json:"accountCreationDate,omitempty"`
	// Indicates the type of account. For example, for a multi-account card product. Allowed values: * notApplicable * credit * debit
	AccountType string `json:"accountType,omitempty"`
	// Number of attempts the shopper tried to add a card to their account in the last day.
	AddCardAttemptsDay int32 `json:"addCardAttemptsDay,omitempty"`
	// Date the selected delivery address was first used.
	DeliveryAddressUsageDate time.Time `json:"deliveryAddressUsageDate,omitempty"`
	// Indicator for the length of time since this delivery address was first used. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	DeliveryAddressUsageIndicator string `json:"deliveryAddressUsageIndicator,omitempty"`
	// Shopper's home phone number (including the country code).
	HomePhone string `json:"homePhone,omitempty"`
	// Shopper's mobile phone number (including the country code).
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Date when the shopper last changed their password.
	PasswordChangeDate time.Time `json:"passwordChangeDate,omitempty"`
	// Indicator when the shopper has changed their password. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PasswordChangeIndicator string `json:"passwordChangeIndicator,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past 24 hours.
	PastTransactionsDay int32 `json:"pastTransactionsDay,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past year.
	PastTransactionsYear int32 `json:"pastTransactionsYear,omitempty"`
	// Date this payment method was added to the shopper's account.
	PaymentAccountAge time.Time `json:"paymentAccountAge,omitempty"`
	// Indicator for the length of time since this payment method was added to this shopper's account. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PaymentAccountIndicator string `json:"paymentAccountIndicator,omitempty"`
	// Number of successful purchases in the last six months.
	PurchasesLast6Months int32 `json:"purchasesLast6Months,omitempty"`
	// Whether suspicious activity was recorded on this account.
	SuspiciousActivity bool `json:"suspiciousActivity,omitempty"`
	// Shopper's work phone number (including the country code).
	WorkPhone string `json:"workPhone,omitempty"`
}
