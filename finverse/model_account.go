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
	"time"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	AccountId string `json:"account_id"`
	// The SHA3-256 hash of the account number, salted with the loginIdentityId
	GroupId              string            `json:"group_id"`
	AccountHolderName    *string           `json:"account_holder_name,omitempty"`
	AccountName          string            `json:"account_name"`
	AccountNickname      *string           `json:"account_nickname,omitempty"`
	AccountSubType       *string           `json:"account_sub_type,omitempty"`
	AccountNumberMasked  *string           `json:"account_number_masked,omitempty"`
	Country              *string           `json:"country,omitempty"`
	CreatedAt            *time.Time        `json:"created_at,omitempty"`
	UpdatedAt            *time.Time        `json:"updated_at,omitempty"`
	AccountCurrency      *string           `json:"account_currency,omitempty"`
	Balance              *CurrencyAmount   `json:"balance,omitempty"`
	StatementBalance     *CurrencyAmount   `json:"statement_balance,omitempty"`
	IsParent             bool              `json:"is_parent"`
	IsClosed             bool              `json:"is_closed"`
	IsExcluded           bool              `json:"is_excluded"`
	AccountType          *AccountType      `json:"account_type,omitempty"`
	Metadata             map[string]string `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(accountId string, groupId string, accountName string, isParent bool, isClosed bool, isExcluded bool, metadata map[string]string) *Account {
	this := Account{}
	this.AccountId = accountId
	this.GroupId = groupId
	this.AccountName = accountName
	this.IsParent = isParent
	this.IsClosed = isClosed
	this.IsExcluded = isExcluded
	this.Metadata = metadata
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *Account) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Account) SetAccountId(v string) {
	o.AccountId = v
}

// GetGroupId returns the GroupId field value
func (o *Account) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *Account) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *Account) SetGroupId(v string) {
	o.GroupId = v
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise.
func (o *Account) GetAccountHolderName() string {
	if o == nil || IsNil(o.AccountHolderName) {
		var ret string
		return ret
	}
	return *o.AccountHolderName
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountHolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountHolderName) {
		return nil, false
	}
	return o.AccountHolderName, true
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *Account) HasAccountHolderName() bool {
	if o != nil && !IsNil(o.AccountHolderName) {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given string and assigns it to the AccountHolderName field.
func (o *Account) SetAccountHolderName(v string) {
	o.AccountHolderName = &v
}

// GetAccountName returns the AccountName field value
func (o *Account) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *Account) SetAccountName(v string) {
	o.AccountName = v
}

// GetAccountNickname returns the AccountNickname field value if set, zero value otherwise.
func (o *Account) GetAccountNickname() string {
	if o == nil || IsNil(o.AccountNickname) {
		var ret string
		return ret
	}
	return *o.AccountNickname
}

// GetAccountNicknameOk returns a tuple with the AccountNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNickname) {
		return nil, false
	}
	return o.AccountNickname, true
}

// HasAccountNickname returns a boolean if a field has been set.
func (o *Account) HasAccountNickname() bool {
	if o != nil && !IsNil(o.AccountNickname) {
		return true
	}

	return false
}

// SetAccountNickname gets a reference to the given string and assigns it to the AccountNickname field.
func (o *Account) SetAccountNickname(v string) {
	o.AccountNickname = &v
}

// GetAccountSubType returns the AccountSubType field value if set, zero value otherwise.
func (o *Account) GetAccountSubType() string {
	if o == nil || IsNil(o.AccountSubType) {
		var ret string
		return ret
	}
	return *o.AccountSubType
}

// GetAccountSubTypeOk returns a tuple with the AccountSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSubType) {
		return nil, false
	}
	return o.AccountSubType, true
}

// HasAccountSubType returns a boolean if a field has been set.
func (o *Account) HasAccountSubType() bool {
	if o != nil && !IsNil(o.AccountSubType) {
		return true
	}

	return false
}

// SetAccountSubType gets a reference to the given string and assigns it to the AccountSubType field.
func (o *Account) SetAccountSubType(v string) {
	o.AccountSubType = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *Account) GetAccountNumberMasked() string {
	if o == nil || IsNil(o.AccountNumberMasked) {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumberMasked) {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *Account) HasAccountNumberMasked() bool {
	if o != nil && !IsNil(o.AccountNumberMasked) {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *Account) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Account) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Account) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Account) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Account) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Account) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Account) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Account) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Account) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Account) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAccountCurrency returns the AccountCurrency field value if set, zero value otherwise.
func (o *Account) GetAccountCurrency() string {
	if o == nil || IsNil(o.AccountCurrency) {
		var ret string
		return ret
	}
	return *o.AccountCurrency
}

// GetAccountCurrencyOk returns a tuple with the AccountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCurrency) {
		return nil, false
	}
	return o.AccountCurrency, true
}

// HasAccountCurrency returns a boolean if a field has been set.
func (o *Account) HasAccountCurrency() bool {
	if o != nil && !IsNil(o.AccountCurrency) {
		return true
	}

	return false
}

// SetAccountCurrency gets a reference to the given string and assigns it to the AccountCurrency field.
func (o *Account) SetAccountCurrency(v string) {
	o.AccountCurrency = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Account) GetBalance() CurrencyAmount {
	if o == nil || IsNil(o.Balance) {
		var ret CurrencyAmount
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBalanceOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Account) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given CurrencyAmount and assigns it to the Balance field.
func (o *Account) SetBalance(v CurrencyAmount) {
	o.Balance = &v
}

// GetStatementBalance returns the StatementBalance field value if set, zero value otherwise.
func (o *Account) GetStatementBalance() CurrencyAmount {
	if o == nil || IsNil(o.StatementBalance) {
		var ret CurrencyAmount
		return ret
	}
	return *o.StatementBalance
}

// GetStatementBalanceOk returns a tuple with the StatementBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetStatementBalanceOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.StatementBalance) {
		return nil, false
	}
	return o.StatementBalance, true
}

// HasStatementBalance returns a boolean if a field has been set.
func (o *Account) HasStatementBalance() bool {
	if o != nil && !IsNil(o.StatementBalance) {
		return true
	}

	return false
}

// SetStatementBalance gets a reference to the given CurrencyAmount and assigns it to the StatementBalance field.
func (o *Account) SetStatementBalance(v CurrencyAmount) {
	o.StatementBalance = &v
}

// GetIsParent returns the IsParent field value
func (o *Account) GetIsParent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsParent
}

// GetIsParentOk returns a tuple with the IsParent field value
// and a boolean to check if the value has been set.
func (o *Account) GetIsParentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsParent, true
}

// SetIsParent sets field value
func (o *Account) SetIsParent(v bool) {
	o.IsParent = v
}

// GetIsClosed returns the IsClosed field value
func (o *Account) GetIsClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *Account) GetIsClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *Account) SetIsClosed(v bool) {
	o.IsClosed = v
}

// GetIsExcluded returns the IsExcluded field value
func (o *Account) GetIsExcluded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExcluded
}

// GetIsExcludedOk returns a tuple with the IsExcluded field value
// and a boolean to check if the value has been set.
func (o *Account) GetIsExcludedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExcluded, true
}

// SetIsExcluded sets field value
func (o *Account) SetIsExcluded(v bool) {
	o.IsExcluded = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *Account) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *Account) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *Account) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetMetadata returns the Metadata field value
func (o *Account) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Account) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Account) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.AccountHolderName) {
		toSerialize["account_holder_name"] = o.AccountHolderName
	}
	toSerialize["account_name"] = o.AccountName
	if !IsNil(o.AccountNickname) {
		toSerialize["account_nickname"] = o.AccountNickname
	}
	if !IsNil(o.AccountSubType) {
		toSerialize["account_sub_type"] = o.AccountSubType
	}
	if !IsNil(o.AccountNumberMasked) {
		toSerialize["account_number_masked"] = o.AccountNumberMasked
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.AccountCurrency) {
		toSerialize["account_currency"] = o.AccountCurrency
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.StatementBalance) {
		toSerialize["statement_balance"] = o.StatementBalance
	}
	toSerialize["is_parent"] = o.IsParent
	toSerialize["is_closed"] = o.IsClosed
	toSerialize["is_excluded"] = o.IsExcluded
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"group_id",
		"account_name",
		"is_parent",
		"is_closed",
		"is_excluded",
		"metadata",
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

	varAccount := _Account{}

	err = json.Unmarshal(data, &varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "account_holder_name")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "account_nickname")
		delete(additionalProperties, "account_sub_type")
		delete(additionalProperties, "account_number_masked")
		delete(additionalProperties, "country")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "account_currency")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "statement_balance")
		delete(additionalProperties, "is_parent")
		delete(additionalProperties, "is_closed")
		delete(additionalProperties, "is_excluded")
		delete(additionalProperties, "account_type")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
