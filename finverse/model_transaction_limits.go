/*
Finverse Public

Documentation of the early finverse services

API version: 0.0.1
Contact: devs@finverse.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finverse

import (
	"encoding/json"
)

// TransactionLimits struct for TransactionLimits
type TransactionLimits struct {
	// Maximum amount of money that can be paid during the reference period (across any number of transactions). Expressed in currency’s smallest unit or “minor unit”, as defined in ISO 4217.
	MaxPeriodAmount *int32 `json:"max_period_amount,omitempty"`
	// Maximum number of transactions (of any amount) that can be executed during the reference period.
	MaxPeriodCount *int32 `json:"max_period_count,omitempty"`
	// The maximum amount of money that can be transferred in a single transaction under this mandate. Expressed in currency’s smallest unit or “minor unit”, as defined in ISO 4217.
	MaxTransactionAmount *int32 `json:"max_transaction_amount,omitempty"`
	// Reference calendar periods for the payment limits. Possible values (DAILY, WEEKLY, MONTHLY, QUARTERLY, YEARLY)
	Period string `json:"period"`
}

// NewTransactionLimits instantiates a new TransactionLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionLimits(period string) *TransactionLimits {
	this := TransactionLimits{}
	this.Period = period
	return &this
}

// NewTransactionLimitsWithDefaults instantiates a new TransactionLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionLimitsWithDefaults() *TransactionLimits {
	this := TransactionLimits{}
	return &this
}

// GetMaxPeriodAmount returns the MaxPeriodAmount field value if set, zero value otherwise.
func (o *TransactionLimits) GetMaxPeriodAmount() int32 {
	if o == nil || o.MaxPeriodAmount == nil {
		var ret int32
		return ret
	}
	return *o.MaxPeriodAmount
}

// GetMaxPeriodAmountOk returns a tuple with the MaxPeriodAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionLimits) GetMaxPeriodAmountOk() (*int32, bool) {
	if o == nil || o.MaxPeriodAmount == nil {
		return nil, false
	}
	return o.MaxPeriodAmount, true
}

// HasMaxPeriodAmount returns a boolean if a field has been set.
func (o *TransactionLimits) HasMaxPeriodAmount() bool {
	if o != nil && o.MaxPeriodAmount != nil {
		return true
	}

	return false
}

// SetMaxPeriodAmount gets a reference to the given int32 and assigns it to the MaxPeriodAmount field.
func (o *TransactionLimits) SetMaxPeriodAmount(v int32) {
	o.MaxPeriodAmount = &v
}

// GetMaxPeriodCount returns the MaxPeriodCount field value if set, zero value otherwise.
func (o *TransactionLimits) GetMaxPeriodCount() int32 {
	if o == nil || o.MaxPeriodCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxPeriodCount
}

// GetMaxPeriodCountOk returns a tuple with the MaxPeriodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionLimits) GetMaxPeriodCountOk() (*int32, bool) {
	if o == nil || o.MaxPeriodCount == nil {
		return nil, false
	}
	return o.MaxPeriodCount, true
}

// HasMaxPeriodCount returns a boolean if a field has been set.
func (o *TransactionLimits) HasMaxPeriodCount() bool {
	if o != nil && o.MaxPeriodCount != nil {
		return true
	}

	return false
}

// SetMaxPeriodCount gets a reference to the given int32 and assigns it to the MaxPeriodCount field.
func (o *TransactionLimits) SetMaxPeriodCount(v int32) {
	o.MaxPeriodCount = &v
}

// GetMaxTransactionAmount returns the MaxTransactionAmount field value if set, zero value otherwise.
func (o *TransactionLimits) GetMaxTransactionAmount() int32 {
	if o == nil || o.MaxTransactionAmount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTransactionAmount
}

// GetMaxTransactionAmountOk returns a tuple with the MaxTransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionLimits) GetMaxTransactionAmountOk() (*int32, bool) {
	if o == nil || o.MaxTransactionAmount == nil {
		return nil, false
	}
	return o.MaxTransactionAmount, true
}

// HasMaxTransactionAmount returns a boolean if a field has been set.
func (o *TransactionLimits) HasMaxTransactionAmount() bool {
	if o != nil && o.MaxTransactionAmount != nil {
		return true
	}

	return false
}

// SetMaxTransactionAmount gets a reference to the given int32 and assigns it to the MaxTransactionAmount field.
func (o *TransactionLimits) SetMaxTransactionAmount(v int32) {
	o.MaxTransactionAmount = &v
}

// GetPeriod returns the Period field value
func (o *TransactionLimits) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *TransactionLimits) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *TransactionLimits) SetPeriod(v string) {
	o.Period = v
}

func (o TransactionLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxPeriodAmount != nil {
		toSerialize["max_period_amount"] = o.MaxPeriodAmount
	}
	if o.MaxPeriodCount != nil {
		toSerialize["max_period_count"] = o.MaxPeriodCount
	}
	if o.MaxTransactionAmount != nil {
		toSerialize["max_transaction_amount"] = o.MaxTransactionAmount
	}
	if true {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionLimits struct {
	value *TransactionLimits
	isSet bool
}

func (v NullableTransactionLimits) Get() *TransactionLimits {
	return v.value
}

func (v *NullableTransactionLimits) Set(val *TransactionLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionLimits(val *TransactionLimits) *NullableTransactionLimits {
	return &NullableTransactionLimits{value: val, isSet: true}
}

func (v NullableTransactionLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}