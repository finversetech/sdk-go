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

// OtherInfo struct for OtherInfo
type OtherInfo struct {
	BankCode *string `json:"bank_code,omitempty"`
}

// NewOtherInfo instantiates a new OtherInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherInfo() *OtherInfo {
	this := OtherInfo{}
	return &this
}

// NewOtherInfoWithDefaults instantiates a new OtherInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherInfoWithDefaults() *OtherInfo {
	this := OtherInfo{}
	return &this
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *OtherInfo) GetBankCode() string {
	if o == nil || o.BankCode == nil {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherInfo) GetBankCodeOk() (*string, bool) {
	if o == nil || o.BankCode == nil {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *OtherInfo) HasBankCode() bool {
	if o != nil && o.BankCode != nil {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *OtherInfo) SetBankCode(v string) {
	o.BankCode = &v
}

func (o OtherInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankCode != nil {
		toSerialize["bank_code"] = o.BankCode
	}
	return json.Marshal(toSerialize)
}

type NullableOtherInfo struct {
	value *OtherInfo
	isSet bool
}

func (v NullableOtherInfo) Get() *OtherInfo {
	return v.value
}

func (v *NullableOtherInfo) Set(val *OtherInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherInfo(val *OtherInfo) *NullableOtherInfo {
	return &NullableOtherInfo{value: val, isSet: true}
}

func (v NullableOtherInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
