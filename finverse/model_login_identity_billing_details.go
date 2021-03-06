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

// LoginIdentityBillingDetails struct for LoginIdentityBillingDetails
type LoginIdentityBillingDetails struct {
	BilledProducts []string `json:"billed_products,omitempty"`
}

// NewLoginIdentityBillingDetails instantiates a new LoginIdentityBillingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginIdentityBillingDetails() *LoginIdentityBillingDetails {
	this := LoginIdentityBillingDetails{}
	return &this
}

// NewLoginIdentityBillingDetailsWithDefaults instantiates a new LoginIdentityBillingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginIdentityBillingDetailsWithDefaults() *LoginIdentityBillingDetails {
	this := LoginIdentityBillingDetails{}
	return &this
}

// GetBilledProducts returns the BilledProducts field value if set, zero value otherwise.
func (o *LoginIdentityBillingDetails) GetBilledProducts() []string {
	if o == nil || o.BilledProducts == nil {
		var ret []string
		return ret
	}
	return o.BilledProducts
}

// GetBilledProductsOk returns a tuple with the BilledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityBillingDetails) GetBilledProductsOk() ([]string, bool) {
	if o == nil || o.BilledProducts == nil {
		return nil, false
	}
	return o.BilledProducts, true
}

// HasBilledProducts returns a boolean if a field has been set.
func (o *LoginIdentityBillingDetails) HasBilledProducts() bool {
	if o != nil && o.BilledProducts != nil {
		return true
	}

	return false
}

// SetBilledProducts gets a reference to the given []string and assigns it to the BilledProducts field.
func (o *LoginIdentityBillingDetails) SetBilledProducts(v []string) {
	o.BilledProducts = v
}

func (o LoginIdentityBillingDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BilledProducts != nil {
		toSerialize["billed_products"] = o.BilledProducts
	}
	return json.Marshal(toSerialize)
}

type NullableLoginIdentityBillingDetails struct {
	value *LoginIdentityBillingDetails
	isSet bool
}

func (v NullableLoginIdentityBillingDetails) Get() *LoginIdentityBillingDetails {
	return v.value
}

func (v *NullableLoginIdentityBillingDetails) Set(val *LoginIdentityBillingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginIdentityBillingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginIdentityBillingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginIdentityBillingDetails(val *LoginIdentityBillingDetails) *NullableLoginIdentityBillingDetails {
	return &NullableLoginIdentityBillingDetails{value: val, isSet: true}
}

func (v NullableLoginIdentityBillingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginIdentityBillingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
