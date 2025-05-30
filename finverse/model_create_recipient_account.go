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
	"fmt"
)

// checks if the CreateRecipientAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecipientAccount{}

// CreateRecipientAccount struct for CreateRecipientAccount
type CreateRecipientAccount struct {
	// Accountholder name of the recipient's account
	AccountholderName string                 `json:"accountholder_name"`
	AccountNumber     RecipientAccountNumber `json:"account_number"`
	// Type of recipient account.
	AccountType string `json:"account_type"`
	// List of currencies supported by the recipient account
	Currencies []string `json:"currencies"`
	// Finverse Institution ID for the recipient’s institution.
	InstitutionId        string `json:"institution_id"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecipientAccount CreateRecipientAccount

// NewCreateRecipientAccount instantiates a new CreateRecipientAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecipientAccount(accountholderName string, accountNumber RecipientAccountNumber, accountType string, currencies []string, institutionId string) *CreateRecipientAccount {
	this := CreateRecipientAccount{}
	this.AccountholderName = accountholderName
	this.AccountNumber = accountNumber
	this.AccountType = accountType
	this.Currencies = currencies
	this.InstitutionId = institutionId
	return &this
}

// NewCreateRecipientAccountWithDefaults instantiates a new CreateRecipientAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecipientAccountWithDefaults() *CreateRecipientAccount {
	this := CreateRecipientAccount{}
	return &this
}

// GetAccountholderName returns the AccountholderName field value
func (o *CreateRecipientAccount) GetAccountholderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountholderName
}

// GetAccountholderNameOk returns a tuple with the AccountholderName field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientAccount) GetAccountholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountholderName, true
}

// SetAccountholderName sets field value
func (o *CreateRecipientAccount) SetAccountholderName(v string) {
	o.AccountholderName = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *CreateRecipientAccount) GetAccountNumber() RecipientAccountNumber {
	if o == nil {
		var ret RecipientAccountNumber
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientAccount) GetAccountNumberOk() (*RecipientAccountNumber, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *CreateRecipientAccount) SetAccountNumber(v RecipientAccountNumber) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value
func (o *CreateRecipientAccount) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *CreateRecipientAccount) SetAccountType(v string) {
	o.AccountType = v
}

// GetCurrencies returns the Currencies field value
func (o *CreateRecipientAccount) GetCurrencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientAccount) GetCurrenciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currencies, true
}

// SetCurrencies sets field value
func (o *CreateRecipientAccount) SetCurrencies(v []string) {
	o.Currencies = v
}

// GetInstitutionId returns the InstitutionId field value
func (o *CreateRecipientAccount) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientAccount) GetInstitutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *CreateRecipientAccount) SetInstitutionId(v string) {
	o.InstitutionId = v
}

func (o CreateRecipientAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecipientAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountholder_name"] = o.AccountholderName
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["account_type"] = o.AccountType
	toSerialize["currencies"] = o.Currencies
	toSerialize["institution_id"] = o.InstitutionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecipientAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountholder_name",
		"account_number",
		"account_type",
		"currencies",
		"institution_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRecipientAccount := _CreateRecipientAccount{}

	err = json.Unmarshal(data, &varCreateRecipientAccount)

	if err != nil {
		return err
	}

	*o = CreateRecipientAccount(varCreateRecipientAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountholder_name")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "account_type")
		delete(additionalProperties, "currencies")
		delete(additionalProperties, "institution_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecipientAccount struct {
	value *CreateRecipientAccount
	isSet bool
}

func (v NullableCreateRecipientAccount) Get() *CreateRecipientAccount {
	return v.value
}

func (v *NullableCreateRecipientAccount) Set(val *CreateRecipientAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecipientAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecipientAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecipientAccount(val *CreateRecipientAccount) *NullableCreateRecipientAccount {
	return &NullableCreateRecipientAccount{value: val, isSet: true}
}

func (v NullableCreateRecipientAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecipientAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
