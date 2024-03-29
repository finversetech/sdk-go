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

// MandateRecipientAccount struct for MandateRecipientAccount
type MandateRecipientAccount struct {
	// Merchant account ID assigned by Finverse
	AccountId string `json:"account_id"`
	// Type of recipient account.
	AccountType string `json:"account_type"`
}

// NewMandateRecipientAccount instantiates a new MandateRecipientAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandateRecipientAccount(accountId string, accountType string) *MandateRecipientAccount {
	this := MandateRecipientAccount{}
	this.AccountId = accountId
	this.AccountType = accountType
	return &this
}

// NewMandateRecipientAccountWithDefaults instantiates a new MandateRecipientAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateRecipientAccountWithDefaults() *MandateRecipientAccount {
	this := MandateRecipientAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *MandateRecipientAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MandateRecipientAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MandateRecipientAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountType returns the AccountType field value
func (o *MandateRecipientAccount) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *MandateRecipientAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *MandateRecipientAccount) SetAccountType(v string) {
	o.AccountType = v
}

func (o MandateRecipientAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	return json.Marshal(toSerialize)
}

type NullableMandateRecipientAccount struct {
	value *MandateRecipientAccount
	isSet bool
}

func (v NullableMandateRecipientAccount) Get() *MandateRecipientAccount {
	return v.value
}

func (v *NullableMandateRecipientAccount) Set(val *MandateRecipientAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableMandateRecipientAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableMandateRecipientAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandateRecipientAccount(val *MandateRecipientAccount) *NullableMandateRecipientAccount {
	return &NullableMandateRecipientAccount{value: val, isSet: true}
}

func (v NullableMandateRecipientAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandateRecipientAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
