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

// RecipientAccount struct for RecipientAccount
type RecipientAccount struct {
	// A unique identifier generated after creating recipient
	RecipientAccountId *string `json:"recipient_account_id,omitempty"`
	// Accountholder name of the recipient's account
	AccountholderName string                 `json:"accountholder_name"`
	AccountNumber     RecipientAccountNumber `json:"account_number"`
	// Type of recipient account.
	AccountType string `json:"account_type"`
	// List of currencies supported by the recipient account
	Currencies []string `json:"currencies"`
	// Finverse Institution ID for the recipient’s institution.
	InstitutionId string `json:"institution_id"`
}

// NewRecipientAccount instantiates a new RecipientAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientAccount(accountholderName string, accountNumber RecipientAccountNumber, accountType string, currencies []string, institutionId string) *RecipientAccount {
	this := RecipientAccount{}
	this.AccountholderName = accountholderName
	this.AccountNumber = accountNumber
	this.AccountType = accountType
	this.Currencies = currencies
	this.InstitutionId = institutionId
	return &this
}

// NewRecipientAccountWithDefaults instantiates a new RecipientAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientAccountWithDefaults() *RecipientAccount {
	this := RecipientAccount{}
	return &this
}

// GetRecipientAccountId returns the RecipientAccountId field value if set, zero value otherwise.
func (o *RecipientAccount) GetRecipientAccountId() string {
	if o == nil || o.RecipientAccountId == nil {
		var ret string
		return ret
	}
	return *o.RecipientAccountId
}

// GetRecipientAccountIdOk returns a tuple with the RecipientAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetRecipientAccountIdOk() (*string, bool) {
	if o == nil || o.RecipientAccountId == nil {
		return nil, false
	}
	return o.RecipientAccountId, true
}

// HasRecipientAccountId returns a boolean if a field has been set.
func (o *RecipientAccount) HasRecipientAccountId() bool {
	if o != nil && o.RecipientAccountId != nil {
		return true
	}

	return false
}

// SetRecipientAccountId gets a reference to the given string and assigns it to the RecipientAccountId field.
func (o *RecipientAccount) SetRecipientAccountId(v string) {
	o.RecipientAccountId = &v
}

// GetAccountholderName returns the AccountholderName field value
func (o *RecipientAccount) GetAccountholderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountholderName
}

// GetAccountholderNameOk returns a tuple with the AccountholderName field value
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetAccountholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountholderName, true
}

// SetAccountholderName sets field value
func (o *RecipientAccount) SetAccountholderName(v string) {
	o.AccountholderName = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *RecipientAccount) GetAccountNumber() RecipientAccountNumber {
	if o == nil {
		var ret RecipientAccountNumber
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetAccountNumberOk() (*RecipientAccountNumber, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *RecipientAccount) SetAccountNumber(v RecipientAccountNumber) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value
func (o *RecipientAccount) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *RecipientAccount) SetAccountType(v string) {
	o.AccountType = v
}

// GetCurrencies returns the Currencies field value
func (o *RecipientAccount) GetCurrencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetCurrenciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currencies, true
}

// SetCurrencies sets field value
func (o *RecipientAccount) SetCurrencies(v []string) {
	o.Currencies = v
}

// GetInstitutionId returns the InstitutionId field value
func (o *RecipientAccount) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *RecipientAccount) GetInstitutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *RecipientAccount) SetInstitutionId(v string) {
	o.InstitutionId = v
}

func (o RecipientAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecipientAccountId != nil {
		toSerialize["recipient_account_id"] = o.RecipientAccountId
	}
	if true {
		toSerialize["accountholder_name"] = o.AccountholderName
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["currencies"] = o.Currencies
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	return json.Marshal(toSerialize)
}

type NullableRecipientAccount struct {
	value *RecipientAccount
	isSet bool
}

func (v NullableRecipientAccount) Get() *RecipientAccount {
	return v.value
}

func (v *NullableRecipientAccount) Set(val *RecipientAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientAccount(val *RecipientAccount) *NullableRecipientAccount {
	return &NullableRecipientAccount{value: val, isSet: true}
}

func (v NullableRecipientAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
