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

// PayoutSender struct for PayoutSender
type PayoutSender struct {
	Name *string `json:"name,omitempty"`
}

// NewPayoutSender instantiates a new PayoutSender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSender() *PayoutSender {
	this := PayoutSender{}
	return &this
}

// NewPayoutSenderWithDefaults instantiates a new PayoutSender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSenderWithDefaults() *PayoutSender {
	this := PayoutSender{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PayoutSender) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSender) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PayoutSender) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PayoutSender) SetName(v string) {
	o.Name = &v
}

func (o PayoutSender) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePayoutSender struct {
	value *PayoutSender
	isSet bool
}

func (v NullablePayoutSender) Get() *PayoutSender {
	return v.value
}

func (v *NullablePayoutSender) Set(val *PayoutSender) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSender) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSender(val *PayoutSender) *NullablePayoutSender {
	return &NullablePayoutSender{value: val, isSet: true}
}

func (v NullablePayoutSender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
