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
	"time"
)

// checks if the CardAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardAccount{}

// CardAccount struct for CardAccount
type CardAccount struct {
	// Account this card is associated with
	AccountId   *string `json:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	// Masked Account number of the card account
	AccountNumberMasked *string      `json:"account_number_masked,omitempty"`
	AccountType         *AccountType `json:"account_type,omitempty"`
	// The statement payment due date
	StatementPaymentDueDate *string `json:"statement_payment_due_date,omitempty"`
	// The next payment due date
	NextPaymentDueDate *string `json:"next_payment_due_date,omitempty"`
	// The statement date
	StatementDate *string `json:"statement_date,omitempty"`
	// The date of the last payment
	LastPaymentDate      *string         `json:"last_payment_date,omitempty"`
	LastPaymentAmount    *CurrencyAmount `json:"last_payment_amount,omitempty"`
	CurrentBalance       *CurrencyAmount `json:"current_balance,omitempty"`
	PaymentDueAmount     *CurrencyAmount `json:"payment_due_amount,omitempty"`
	StatementDueAmount   *CurrencyAmount `json:"statement_due_amount,omitempty"`
	TotalCreditLimit     *CurrencyAmount `json:"total_credit_limit,omitempty"`
	AvailableCreditLimit *CurrencyAmount `json:"available_credit_limit,omitempty"`
	MinimumPaymentDue    *CurrencyAmount `json:"minimum_payment_due,omitempty"`
	RewardsBalances      []GenericAmount `json:"rewards_balances,omitempty"`
	UpdatedAt            *time.Time      `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CardAccount CardAccount

// NewCardAccount instantiates a new CardAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAccount() *CardAccount {
	this := CardAccount{}
	return &this
}

// NewCardAccountWithDefaults instantiates a new CardAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAccountWithDefaults() *CardAccount {
	this := CardAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CardAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CardAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CardAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *CardAccount) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *CardAccount) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *CardAccount) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *CardAccount) GetAccountNumberMasked() string {
	if o == nil || IsNil(o.AccountNumberMasked) {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumberMasked) {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *CardAccount) HasAccountNumberMasked() bool {
	if o != nil && !IsNil(o.AccountNumberMasked) {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *CardAccount) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *CardAccount) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *CardAccount) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *CardAccount) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetStatementPaymentDueDate returns the StatementPaymentDueDate field value if set, zero value otherwise.
func (o *CardAccount) GetStatementPaymentDueDate() string {
	if o == nil || IsNil(o.StatementPaymentDueDate) {
		var ret string
		return ret
	}
	return *o.StatementPaymentDueDate
}

// GetStatementPaymentDueDateOk returns a tuple with the StatementPaymentDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetStatementPaymentDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatementPaymentDueDate) {
		return nil, false
	}
	return o.StatementPaymentDueDate, true
}

// HasStatementPaymentDueDate returns a boolean if a field has been set.
func (o *CardAccount) HasStatementPaymentDueDate() bool {
	if o != nil && !IsNil(o.StatementPaymentDueDate) {
		return true
	}

	return false
}

// SetStatementPaymentDueDate gets a reference to the given string and assigns it to the StatementPaymentDueDate field.
func (o *CardAccount) SetStatementPaymentDueDate(v string) {
	o.StatementPaymentDueDate = &v
}

// GetNextPaymentDueDate returns the NextPaymentDueDate field value if set, zero value otherwise.
func (o *CardAccount) GetNextPaymentDueDate() string {
	if o == nil || IsNil(o.NextPaymentDueDate) {
		var ret string
		return ret
	}
	return *o.NextPaymentDueDate
}

// GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetNextPaymentDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextPaymentDueDate) {
		return nil, false
	}
	return o.NextPaymentDueDate, true
}

// HasNextPaymentDueDate returns a boolean if a field has been set.
func (o *CardAccount) HasNextPaymentDueDate() bool {
	if o != nil && !IsNil(o.NextPaymentDueDate) {
		return true
	}

	return false
}

// SetNextPaymentDueDate gets a reference to the given string and assigns it to the NextPaymentDueDate field.
func (o *CardAccount) SetNextPaymentDueDate(v string) {
	o.NextPaymentDueDate = &v
}

// GetStatementDate returns the StatementDate field value if set, zero value otherwise.
func (o *CardAccount) GetStatementDate() string {
	if o == nil || IsNil(o.StatementDate) {
		var ret string
		return ret
	}
	return *o.StatementDate
}

// GetStatementDateOk returns a tuple with the StatementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetStatementDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDate) {
		return nil, false
	}
	return o.StatementDate, true
}

// HasStatementDate returns a boolean if a field has been set.
func (o *CardAccount) HasStatementDate() bool {
	if o != nil && !IsNil(o.StatementDate) {
		return true
	}

	return false
}

// SetStatementDate gets a reference to the given string and assigns it to the StatementDate field.
func (o *CardAccount) SetStatementDate(v string) {
	o.StatementDate = &v
}

// GetLastPaymentDate returns the LastPaymentDate field value if set, zero value otherwise.
func (o *CardAccount) GetLastPaymentDate() string {
	if o == nil || IsNil(o.LastPaymentDate) {
		var ret string
		return ret
	}
	return *o.LastPaymentDate
}

// GetLastPaymentDateOk returns a tuple with the LastPaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetLastPaymentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastPaymentDate) {
		return nil, false
	}
	return o.LastPaymentDate, true
}

// HasLastPaymentDate returns a boolean if a field has been set.
func (o *CardAccount) HasLastPaymentDate() bool {
	if o != nil && !IsNil(o.LastPaymentDate) {
		return true
	}

	return false
}

// SetLastPaymentDate gets a reference to the given string and assigns it to the LastPaymentDate field.
func (o *CardAccount) SetLastPaymentDate(v string) {
	o.LastPaymentDate = &v
}

// GetLastPaymentAmount returns the LastPaymentAmount field value if set, zero value otherwise.
func (o *CardAccount) GetLastPaymentAmount() CurrencyAmount {
	if o == nil || IsNil(o.LastPaymentAmount) {
		var ret CurrencyAmount
		return ret
	}
	return *o.LastPaymentAmount
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetLastPaymentAmountOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.LastPaymentAmount) {
		return nil, false
	}
	return o.LastPaymentAmount, true
}

// HasLastPaymentAmount returns a boolean if a field has been set.
func (o *CardAccount) HasLastPaymentAmount() bool {
	if o != nil && !IsNil(o.LastPaymentAmount) {
		return true
	}

	return false
}

// SetLastPaymentAmount gets a reference to the given CurrencyAmount and assigns it to the LastPaymentAmount field.
func (o *CardAccount) SetLastPaymentAmount(v CurrencyAmount) {
	o.LastPaymentAmount = &v
}

// GetCurrentBalance returns the CurrentBalance field value if set, zero value otherwise.
func (o *CardAccount) GetCurrentBalance() CurrencyAmount {
	if o == nil || IsNil(o.CurrentBalance) {
		var ret CurrencyAmount
		return ret
	}
	return *o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetCurrentBalanceOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.CurrentBalance) {
		return nil, false
	}
	return o.CurrentBalance, true
}

// HasCurrentBalance returns a boolean if a field has been set.
func (o *CardAccount) HasCurrentBalance() bool {
	if o != nil && !IsNil(o.CurrentBalance) {
		return true
	}

	return false
}

// SetCurrentBalance gets a reference to the given CurrencyAmount and assigns it to the CurrentBalance field.
func (o *CardAccount) SetCurrentBalance(v CurrencyAmount) {
	o.CurrentBalance = &v
}

// GetPaymentDueAmount returns the PaymentDueAmount field value if set, zero value otherwise.
func (o *CardAccount) GetPaymentDueAmount() CurrencyAmount {
	if o == nil || IsNil(o.PaymentDueAmount) {
		var ret CurrencyAmount
		return ret
	}
	return *o.PaymentDueAmount
}

// GetPaymentDueAmountOk returns a tuple with the PaymentDueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetPaymentDueAmountOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.PaymentDueAmount) {
		return nil, false
	}
	return o.PaymentDueAmount, true
}

// HasPaymentDueAmount returns a boolean if a field has been set.
func (o *CardAccount) HasPaymentDueAmount() bool {
	if o != nil && !IsNil(o.PaymentDueAmount) {
		return true
	}

	return false
}

// SetPaymentDueAmount gets a reference to the given CurrencyAmount and assigns it to the PaymentDueAmount field.
func (o *CardAccount) SetPaymentDueAmount(v CurrencyAmount) {
	o.PaymentDueAmount = &v
}

// GetStatementDueAmount returns the StatementDueAmount field value if set, zero value otherwise.
func (o *CardAccount) GetStatementDueAmount() CurrencyAmount {
	if o == nil || IsNil(o.StatementDueAmount) {
		var ret CurrencyAmount
		return ret
	}
	return *o.StatementDueAmount
}

// GetStatementDueAmountOk returns a tuple with the StatementDueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetStatementDueAmountOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.StatementDueAmount) {
		return nil, false
	}
	return o.StatementDueAmount, true
}

// HasStatementDueAmount returns a boolean if a field has been set.
func (o *CardAccount) HasStatementDueAmount() bool {
	if o != nil && !IsNil(o.StatementDueAmount) {
		return true
	}

	return false
}

// SetStatementDueAmount gets a reference to the given CurrencyAmount and assigns it to the StatementDueAmount field.
func (o *CardAccount) SetStatementDueAmount(v CurrencyAmount) {
	o.StatementDueAmount = &v
}

// GetTotalCreditLimit returns the TotalCreditLimit field value if set, zero value otherwise.
func (o *CardAccount) GetTotalCreditLimit() CurrencyAmount {
	if o == nil || IsNil(o.TotalCreditLimit) {
		var ret CurrencyAmount
		return ret
	}
	return *o.TotalCreditLimit
}

// GetTotalCreditLimitOk returns a tuple with the TotalCreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetTotalCreditLimitOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.TotalCreditLimit) {
		return nil, false
	}
	return o.TotalCreditLimit, true
}

// HasTotalCreditLimit returns a boolean if a field has been set.
func (o *CardAccount) HasTotalCreditLimit() bool {
	if o != nil && !IsNil(o.TotalCreditLimit) {
		return true
	}

	return false
}

// SetTotalCreditLimit gets a reference to the given CurrencyAmount and assigns it to the TotalCreditLimit field.
func (o *CardAccount) SetTotalCreditLimit(v CurrencyAmount) {
	o.TotalCreditLimit = &v
}

// GetAvailableCreditLimit returns the AvailableCreditLimit field value if set, zero value otherwise.
func (o *CardAccount) GetAvailableCreditLimit() CurrencyAmount {
	if o == nil || IsNil(o.AvailableCreditLimit) {
		var ret CurrencyAmount
		return ret
	}
	return *o.AvailableCreditLimit
}

// GetAvailableCreditLimitOk returns a tuple with the AvailableCreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetAvailableCreditLimitOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.AvailableCreditLimit) {
		return nil, false
	}
	return o.AvailableCreditLimit, true
}

// HasAvailableCreditLimit returns a boolean if a field has been set.
func (o *CardAccount) HasAvailableCreditLimit() bool {
	if o != nil && !IsNil(o.AvailableCreditLimit) {
		return true
	}

	return false
}

// SetAvailableCreditLimit gets a reference to the given CurrencyAmount and assigns it to the AvailableCreditLimit field.
func (o *CardAccount) SetAvailableCreditLimit(v CurrencyAmount) {
	o.AvailableCreditLimit = &v
}

// GetMinimumPaymentDue returns the MinimumPaymentDue field value if set, zero value otherwise.
func (o *CardAccount) GetMinimumPaymentDue() CurrencyAmount {
	if o == nil || IsNil(o.MinimumPaymentDue) {
		var ret CurrencyAmount
		return ret
	}
	return *o.MinimumPaymentDue
}

// GetMinimumPaymentDueOk returns a tuple with the MinimumPaymentDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetMinimumPaymentDueOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.MinimumPaymentDue) {
		return nil, false
	}
	return o.MinimumPaymentDue, true
}

// HasMinimumPaymentDue returns a boolean if a field has been set.
func (o *CardAccount) HasMinimumPaymentDue() bool {
	if o != nil && !IsNil(o.MinimumPaymentDue) {
		return true
	}

	return false
}

// SetMinimumPaymentDue gets a reference to the given CurrencyAmount and assigns it to the MinimumPaymentDue field.
func (o *CardAccount) SetMinimumPaymentDue(v CurrencyAmount) {
	o.MinimumPaymentDue = &v
}

// GetRewardsBalances returns the RewardsBalances field value if set, zero value otherwise.
func (o *CardAccount) GetRewardsBalances() []GenericAmount {
	if o == nil || IsNil(o.RewardsBalances) {
		var ret []GenericAmount
		return ret
	}
	return o.RewardsBalances
}

// GetRewardsBalancesOk returns a tuple with the RewardsBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetRewardsBalancesOk() ([]GenericAmount, bool) {
	if o == nil || IsNil(o.RewardsBalances) {
		return nil, false
	}
	return o.RewardsBalances, true
}

// HasRewardsBalances returns a boolean if a field has been set.
func (o *CardAccount) HasRewardsBalances() bool {
	if o != nil && !IsNil(o.RewardsBalances) {
		return true
	}

	return false
}

// SetRewardsBalances gets a reference to the given []GenericAmount and assigns it to the RewardsBalances field.
func (o *CardAccount) SetRewardsBalances(v []GenericAmount) {
	o.RewardsBalances = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CardAccount) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CardAccount) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CardAccount) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CardAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.AccountNumberMasked) {
		toSerialize["account_number_masked"] = o.AccountNumberMasked
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.StatementPaymentDueDate) {
		toSerialize["statement_payment_due_date"] = o.StatementPaymentDueDate
	}
	if !IsNil(o.NextPaymentDueDate) {
		toSerialize["next_payment_due_date"] = o.NextPaymentDueDate
	}
	if !IsNil(o.StatementDate) {
		toSerialize["statement_date"] = o.StatementDate
	}
	if !IsNil(o.LastPaymentDate) {
		toSerialize["last_payment_date"] = o.LastPaymentDate
	}
	if !IsNil(o.LastPaymentAmount) {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount
	}
	if !IsNil(o.CurrentBalance) {
		toSerialize["current_balance"] = o.CurrentBalance
	}
	if !IsNil(o.PaymentDueAmount) {
		toSerialize["payment_due_amount"] = o.PaymentDueAmount
	}
	if !IsNil(o.StatementDueAmount) {
		toSerialize["statement_due_amount"] = o.StatementDueAmount
	}
	if !IsNil(o.TotalCreditLimit) {
		toSerialize["total_credit_limit"] = o.TotalCreditLimit
	}
	if !IsNil(o.AvailableCreditLimit) {
		toSerialize["available_credit_limit"] = o.AvailableCreditLimit
	}
	if !IsNil(o.MinimumPaymentDue) {
		toSerialize["minimum_payment_due"] = o.MinimumPaymentDue
	}
	if !IsNil(o.RewardsBalances) {
		toSerialize["rewards_balances"] = o.RewardsBalances
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardAccount) UnmarshalJSON(data []byte) (err error) {
	varCardAccount := _CardAccount{}

	err = json.Unmarshal(data, &varCardAccount)

	if err != nil {
		return err
	}

	*o = CardAccount(varCardAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "account_number_masked")
		delete(additionalProperties, "account_type")
		delete(additionalProperties, "statement_payment_due_date")
		delete(additionalProperties, "next_payment_due_date")
		delete(additionalProperties, "statement_date")
		delete(additionalProperties, "last_payment_date")
		delete(additionalProperties, "last_payment_amount")
		delete(additionalProperties, "current_balance")
		delete(additionalProperties, "payment_due_amount")
		delete(additionalProperties, "statement_due_amount")
		delete(additionalProperties, "total_credit_limit")
		delete(additionalProperties, "available_credit_limit")
		delete(additionalProperties, "minimum_payment_due")
		delete(additionalProperties, "rewards_balances")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardAccount struct {
	value *CardAccount
	isSet bool
}

func (v NullableCardAccount) Get() *CardAccount {
	return v.value
}

func (v *NullableCardAccount) Set(val *CardAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAccount(val *CardAccount) *NullableCardAccount {
	return &NullableCardAccount{value: val, isSet: true}
}

func (v NullableCardAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
