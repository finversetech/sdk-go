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

// checks if the IncomeTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeTotal{}

// IncomeTotal struct for IncomeTotal
type IncomeTotal struct {
	EstimatedMonthlyIncome *IncomeEstimate `json:"estimated_monthly_income,omitempty"`
	// Number of transactions counted towards income
	TransactionCount     float32                 `json:"transaction_count"`
	MonthlyHistory       []MonthlyIncomeEstimate `json:"monthly_history"`
	AdditionalProperties map[string]interface{}
}

type _IncomeTotal IncomeTotal

// NewIncomeTotal instantiates a new IncomeTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeTotal(transactionCount float32, monthlyHistory []MonthlyIncomeEstimate) *IncomeTotal {
	this := IncomeTotal{}
	this.TransactionCount = transactionCount
	this.MonthlyHistory = monthlyHistory
	return &this
}

// NewIncomeTotalWithDefaults instantiates a new IncomeTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeTotalWithDefaults() *IncomeTotal {
	this := IncomeTotal{}
	return &this
}

// GetEstimatedMonthlyIncome returns the EstimatedMonthlyIncome field value if set, zero value otherwise.
func (o *IncomeTotal) GetEstimatedMonthlyIncome() IncomeEstimate {
	if o == nil || IsNil(o.EstimatedMonthlyIncome) {
		var ret IncomeEstimate
		return ret
	}
	return *o.EstimatedMonthlyIncome
}

// GetEstimatedMonthlyIncomeOk returns a tuple with the EstimatedMonthlyIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeTotal) GetEstimatedMonthlyIncomeOk() (*IncomeEstimate, bool) {
	if o == nil || IsNil(o.EstimatedMonthlyIncome) {
		return nil, false
	}
	return o.EstimatedMonthlyIncome, true
}

// HasEstimatedMonthlyIncome returns a boolean if a field has been set.
func (o *IncomeTotal) HasEstimatedMonthlyIncome() bool {
	if o != nil && !IsNil(o.EstimatedMonthlyIncome) {
		return true
	}

	return false
}

// SetEstimatedMonthlyIncome gets a reference to the given IncomeEstimate and assigns it to the EstimatedMonthlyIncome field.
func (o *IncomeTotal) SetEstimatedMonthlyIncome(v IncomeEstimate) {
	o.EstimatedMonthlyIncome = &v
}

// GetTransactionCount returns the TransactionCount field value
func (o *IncomeTotal) GetTransactionCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransactionCount
}

// GetTransactionCountOk returns a tuple with the TransactionCount field value
// and a boolean to check if the value has been set.
func (o *IncomeTotal) GetTransactionCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionCount, true
}

// SetTransactionCount sets field value
func (o *IncomeTotal) SetTransactionCount(v float32) {
	o.TransactionCount = v
}

// GetMonthlyHistory returns the MonthlyHistory field value
func (o *IncomeTotal) GetMonthlyHistory() []MonthlyIncomeEstimate {
	if o == nil {
		var ret []MonthlyIncomeEstimate
		return ret
	}

	return o.MonthlyHistory
}

// GetMonthlyHistoryOk returns a tuple with the MonthlyHistory field value
// and a boolean to check if the value has been set.
func (o *IncomeTotal) GetMonthlyHistoryOk() ([]MonthlyIncomeEstimate, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyHistory, true
}

// SetMonthlyHistory sets field value
func (o *IncomeTotal) SetMonthlyHistory(v []MonthlyIncomeEstimate) {
	o.MonthlyHistory = v
}

func (o IncomeTotal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedMonthlyIncome) {
		toSerialize["estimated_monthly_income"] = o.EstimatedMonthlyIncome
	}
	toSerialize["transaction_count"] = o.TransactionCount
	toSerialize["monthly_history"] = o.MonthlyHistory

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IncomeTotal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_count",
		"monthly_history",
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

	varIncomeTotal := _IncomeTotal{}

	err = json.Unmarshal(data, &varIncomeTotal)

	if err != nil {
		return err
	}

	*o = IncomeTotal(varIncomeTotal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "estimated_monthly_income")
		delete(additionalProperties, "transaction_count")
		delete(additionalProperties, "monthly_history")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeTotal struct {
	value *IncomeTotal
	isSet bool
}

func (v NullableIncomeTotal) Get() *IncomeTotal {
	return v.value
}

func (v *NullableIncomeTotal) Set(val *IncomeTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeTotal(val *IncomeTotal) *NullableIncomeTotal {
	return &NullableIncomeTotal{value: val, isSet: true}
}

func (v NullableIncomeTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
