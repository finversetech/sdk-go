/*
Finverse Public

Documentation of the early finverse services

API version: 0.0.1
Contact: info@finverse.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finverse

import (
	"encoding/json"
)

// CreatePaymentLinkCardPaymentResponse struct for CreatePaymentLinkCardPaymentResponse
type CreatePaymentLinkCardPaymentResponse struct {
	// URL to redirect to for making the card payment (e.g. Stripe)
	CardProcessorRedirectUri string `json:"card_processor_redirect_uri"`
	// Finverse Payment ID
	PaymentId *string `json:"payment_id,omitempty"`
}

// NewCreatePaymentLinkCardPaymentResponse instantiates a new CreatePaymentLinkCardPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentLinkCardPaymentResponse(cardProcessorRedirectUri string) *CreatePaymentLinkCardPaymentResponse {
	this := CreatePaymentLinkCardPaymentResponse{}
	this.CardProcessorRedirectUri = cardProcessorRedirectUri
	return &this
}

// NewCreatePaymentLinkCardPaymentResponseWithDefaults instantiates a new CreatePaymentLinkCardPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentLinkCardPaymentResponseWithDefaults() *CreatePaymentLinkCardPaymentResponse {
	this := CreatePaymentLinkCardPaymentResponse{}
	return &this
}

// GetCardProcessorRedirectUri returns the CardProcessorRedirectUri field value
func (o *CreatePaymentLinkCardPaymentResponse) GetCardProcessorRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProcessorRedirectUri
}

// GetCardProcessorRedirectUriOk returns a tuple with the CardProcessorRedirectUri field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentLinkCardPaymentResponse) GetCardProcessorRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProcessorRedirectUri, true
}

// SetCardProcessorRedirectUri sets field value
func (o *CreatePaymentLinkCardPaymentResponse) SetCardProcessorRedirectUri(v string) {
	o.CardProcessorRedirectUri = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *CreatePaymentLinkCardPaymentResponse) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentLinkCardPaymentResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *CreatePaymentLinkCardPaymentResponse) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *CreatePaymentLinkCardPaymentResponse) SetPaymentId(v string) {
	o.PaymentId = &v
}

func (o CreatePaymentLinkCardPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_processor_redirect_uri"] = o.CardProcessorRedirectUri
	}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePaymentLinkCardPaymentResponse struct {
	value *CreatePaymentLinkCardPaymentResponse
	isSet bool
}

func (v NullableCreatePaymentLinkCardPaymentResponse) Get() *CreatePaymentLinkCardPaymentResponse {
	return v.value
}

func (v *NullableCreatePaymentLinkCardPaymentResponse) Set(val *CreatePaymentLinkCardPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentLinkCardPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentLinkCardPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentLinkCardPaymentResponse(val *CreatePaymentLinkCardPaymentResponse) *NullableCreatePaymentLinkCardPaymentResponse {
	return &NullableCreatePaymentLinkCardPaymentResponse{value: val, isSet: true}
}

func (v NullableCreatePaymentLinkCardPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentLinkCardPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
