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

// ListPaymentsResponse struct for ListPaymentsResponse
type ListPaymentsResponse struct {
	Payments      []PaymentResponse `json:"payments,omitempty"`
	TotalPayments int32             `json:"total_payments"`
}

// NewListPaymentsResponse instantiates a new ListPaymentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentsResponse(totalPayments int32) *ListPaymentsResponse {
	this := ListPaymentsResponse{}
	this.TotalPayments = totalPayments
	return &this
}

// NewListPaymentsResponseWithDefaults instantiates a new ListPaymentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentsResponseWithDefaults() *ListPaymentsResponse {
	this := ListPaymentsResponse{}
	return &this
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *ListPaymentsResponse) GetPayments() []PaymentResponse {
	if o == nil || o.Payments == nil {
		var ret []PaymentResponse
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentsResponse) GetPaymentsOk() ([]PaymentResponse, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *ListPaymentsResponse) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []PaymentResponse and assigns it to the Payments field.
func (o *ListPaymentsResponse) SetPayments(v []PaymentResponse) {
	o.Payments = v
}

// GetTotalPayments returns the TotalPayments field value
func (o *ListPaymentsResponse) GetTotalPayments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPayments
}

// GetTotalPaymentsOk returns a tuple with the TotalPayments field value
// and a boolean to check if the value has been set.
func (o *ListPaymentsResponse) GetTotalPaymentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPayments, true
}

// SetTotalPayments sets field value
func (o *ListPaymentsResponse) SetTotalPayments(v int32) {
	o.TotalPayments = v
}

func (o ListPaymentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	if true {
		toSerialize["total_payments"] = o.TotalPayments
	}
	return json.Marshal(toSerialize)
}

type NullableListPaymentsResponse struct {
	value *ListPaymentsResponse
	isSet bool
}

func (v NullableListPaymentsResponse) Get() *ListPaymentsResponse {
	return v.value
}

func (v *NullableListPaymentsResponse) Set(val *ListPaymentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentsResponse(val *ListPaymentsResponse) *NullableListPaymentsResponse {
	return &NullableListPaymentsResponse{value: val, isSet: true}
}

func (v NullableListPaymentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
