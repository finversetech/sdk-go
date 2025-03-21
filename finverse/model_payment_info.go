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

// checks if the PaymentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentInfo{}

// PaymentInfo struct for PaymentInfo
type PaymentInfo struct {
	CurrenciesSupported  []string   `json:"currencies_supported,omitempty"`
	PaymentsSupported    []string   `json:"payments_supported"`
	OtherInfo            *OtherInfo `json:"other_info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInfo PaymentInfo

// NewPaymentInfo instantiates a new PaymentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInfo(paymentsSupported []string) *PaymentInfo {
	this := PaymentInfo{}
	this.PaymentsSupported = paymentsSupported
	return &this
}

// NewPaymentInfoWithDefaults instantiates a new PaymentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInfoWithDefaults() *PaymentInfo {
	this := PaymentInfo{}
	return &this
}

// GetCurrenciesSupported returns the CurrenciesSupported field value if set, zero value otherwise.
func (o *PaymentInfo) GetCurrenciesSupported() []string {
	if o == nil || IsNil(o.CurrenciesSupported) {
		var ret []string
		return ret
	}
	return o.CurrenciesSupported
}

// GetCurrenciesSupportedOk returns a tuple with the CurrenciesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInfo) GetCurrenciesSupportedOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrenciesSupported) {
		return nil, false
	}
	return o.CurrenciesSupported, true
}

// HasCurrenciesSupported returns a boolean if a field has been set.
func (o *PaymentInfo) HasCurrenciesSupported() bool {
	if o != nil && !IsNil(o.CurrenciesSupported) {
		return true
	}

	return false
}

// SetCurrenciesSupported gets a reference to the given []string and assigns it to the CurrenciesSupported field.
func (o *PaymentInfo) SetCurrenciesSupported(v []string) {
	o.CurrenciesSupported = v
}

// GetPaymentsSupported returns the PaymentsSupported field value
func (o *PaymentInfo) GetPaymentsSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PaymentsSupported
}

// GetPaymentsSupportedOk returns a tuple with the PaymentsSupported field value
// and a boolean to check if the value has been set.
func (o *PaymentInfo) GetPaymentsSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentsSupported, true
}

// SetPaymentsSupported sets field value
func (o *PaymentInfo) SetPaymentsSupported(v []string) {
	o.PaymentsSupported = v
}

// GetOtherInfo returns the OtherInfo field value if set, zero value otherwise.
func (o *PaymentInfo) GetOtherInfo() OtherInfo {
	if o == nil || IsNil(o.OtherInfo) {
		var ret OtherInfo
		return ret
	}
	return *o.OtherInfo
}

// GetOtherInfoOk returns a tuple with the OtherInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInfo) GetOtherInfoOk() (*OtherInfo, bool) {
	if o == nil || IsNil(o.OtherInfo) {
		return nil, false
	}
	return o.OtherInfo, true
}

// HasOtherInfo returns a boolean if a field has been set.
func (o *PaymentInfo) HasOtherInfo() bool {
	if o != nil && !IsNil(o.OtherInfo) {
		return true
	}

	return false
}

// SetOtherInfo gets a reference to the given OtherInfo and assigns it to the OtherInfo field.
func (o *PaymentInfo) SetOtherInfo(v OtherInfo) {
	o.OtherInfo = &v
}

func (o PaymentInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrenciesSupported) {
		toSerialize["currencies_supported"] = o.CurrenciesSupported
	}
	toSerialize["payments_supported"] = o.PaymentsSupported
	if !IsNil(o.OtherInfo) {
		toSerialize["other_info"] = o.OtherInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payments_supported",
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

	varPaymentInfo := _PaymentInfo{}

	err = json.Unmarshal(data, &varPaymentInfo)

	if err != nil {
		return err
	}

	*o = PaymentInfo(varPaymentInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currencies_supported")
		delete(additionalProperties, "payments_supported")
		delete(additionalProperties, "other_info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInfo struct {
	value *PaymentInfo
	isSet bool
}

func (v NullablePaymentInfo) Get() *PaymentInfo {
	return v.value
}

func (v *NullablePaymentInfo) Set(val *PaymentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInfo(val *PaymentInfo) *NullablePaymentInfo {
	return &NullablePaymentInfo{value: val, isSet: true}
}

func (v NullablePaymentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
