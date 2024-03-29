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

// PaymentLinkSenderResponse struct for PaymentLinkSenderResponse
type PaymentLinkSenderResponse struct {
	Email *string `json:"email,omitempty"`
	// Customer App's user ID, representing the end-user making the payment.
	ExternalUserId *string `json:"external_user_id,omitempty"`
	// Accountholder name of the sender's account
	Name *string `json:"name,omitempty"`
	// A unique identifier generated after creating sender
	UserId *string `json:"user_id,omitempty"`
}

// NewPaymentLinkSenderResponse instantiates a new PaymentLinkSenderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinkSenderResponse() *PaymentLinkSenderResponse {
	this := PaymentLinkSenderResponse{}
	return &this
}

// NewPaymentLinkSenderResponseWithDefaults instantiates a new PaymentLinkSenderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinkSenderResponseWithDefaults() *PaymentLinkSenderResponse {
	this := PaymentLinkSenderResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PaymentLinkSenderResponse) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkSenderResponse) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PaymentLinkSenderResponse) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PaymentLinkSenderResponse) SetEmail(v string) {
	o.Email = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *PaymentLinkSenderResponse) GetExternalUserId() string {
	if o == nil || o.ExternalUserId == nil {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkSenderResponse) GetExternalUserIdOk() (*string, bool) {
	if o == nil || o.ExternalUserId == nil {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *PaymentLinkSenderResponse) HasExternalUserId() bool {
	if o != nil && o.ExternalUserId != nil {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *PaymentLinkSenderResponse) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentLinkSenderResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkSenderResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentLinkSenderResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentLinkSenderResponse) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PaymentLinkSenderResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkSenderResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PaymentLinkSenderResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PaymentLinkSenderResponse) SetUserId(v string) {
	o.UserId = &v
}

func (o PaymentLinkSenderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.ExternalUserId != nil {
		toSerialize["external_user_id"] = o.ExternalUserId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentLinkSenderResponse struct {
	value *PaymentLinkSenderResponse
	isSet bool
}

func (v NullablePaymentLinkSenderResponse) Get() *PaymentLinkSenderResponse {
	return v.value
}

func (v *NullablePaymentLinkSenderResponse) Set(val *PaymentLinkSenderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinkSenderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinkSenderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinkSenderResponse(val *PaymentLinkSenderResponse) *NullablePaymentLinkSenderResponse {
	return &NullablePaymentLinkSenderResponse{value: val, isSet: true}
}

func (v NullablePaymentLinkSenderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinkSenderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
