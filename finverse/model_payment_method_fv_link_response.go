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

// PaymentMethodFvLinkResponse struct for PaymentMethodFvLinkResponse
type PaymentMethodFvLinkResponse struct {
	PaymentMethodId   *string                `json:"payment_method_id,omitempty"`
	PaymentMethodType *string                `json:"payment_method_type,omitempty"`
	Mandate           *MandateFvLinkResponse `json:"mandate,omitempty"`
}

// NewPaymentMethodFvLinkResponse instantiates a new PaymentMethodFvLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodFvLinkResponse() *PaymentMethodFvLinkResponse {
	this := PaymentMethodFvLinkResponse{}
	return &this
}

// NewPaymentMethodFvLinkResponseWithDefaults instantiates a new PaymentMethodFvLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodFvLinkResponseWithDefaults() *PaymentMethodFvLinkResponse {
	this := PaymentMethodFvLinkResponse{}
	return &this
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *PaymentMethodFvLinkResponse) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodFvLinkResponse) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PaymentMethodFvLinkResponse) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *PaymentMethodFvLinkResponse) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *PaymentMethodFvLinkResponse) GetPaymentMethodType() string {
	if o == nil || o.PaymentMethodType == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodFvLinkResponse) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || o.PaymentMethodType == nil {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *PaymentMethodFvLinkResponse) HasPaymentMethodType() bool {
	if o != nil && o.PaymentMethodType != nil {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *PaymentMethodFvLinkResponse) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetMandate returns the Mandate field value if set, zero value otherwise.
func (o *PaymentMethodFvLinkResponse) GetMandate() MandateFvLinkResponse {
	if o == nil || o.Mandate == nil {
		var ret MandateFvLinkResponse
		return ret
	}
	return *o.Mandate
}

// GetMandateOk returns a tuple with the Mandate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodFvLinkResponse) GetMandateOk() (*MandateFvLinkResponse, bool) {
	if o == nil || o.Mandate == nil {
		return nil, false
	}
	return o.Mandate, true
}

// HasMandate returns a boolean if a field has been set.
func (o *PaymentMethodFvLinkResponse) HasMandate() bool {
	if o != nil && o.Mandate != nil {
		return true
	}

	return false
}

// SetMandate gets a reference to the given MandateFvLinkResponse and assigns it to the Mandate field.
func (o *PaymentMethodFvLinkResponse) SetMandate(v MandateFvLinkResponse) {
	o.Mandate = &v
}

func (o PaymentMethodFvLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if o.PaymentMethodType != nil {
		toSerialize["payment_method_type"] = o.PaymentMethodType
	}
	if o.Mandate != nil {
		toSerialize["mandate"] = o.Mandate
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodFvLinkResponse struct {
	value *PaymentMethodFvLinkResponse
	isSet bool
}

func (v NullablePaymentMethodFvLinkResponse) Get() *PaymentMethodFvLinkResponse {
	return v.value
}

func (v *NullablePaymentMethodFvLinkResponse) Set(val *PaymentMethodFvLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodFvLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodFvLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodFvLinkResponse(val *PaymentMethodFvLinkResponse) *NullablePaymentMethodFvLinkResponse {
	return &NullablePaymentMethodFvLinkResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodFvLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodFvLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}