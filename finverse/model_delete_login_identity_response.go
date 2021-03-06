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

// DeleteLoginIdentityResponse struct for DeleteLoginIdentityResponse
type DeleteLoginIdentityResponse struct {
	Success *bool `json:"success,omitempty"`
}

// NewDeleteLoginIdentityResponse instantiates a new DeleteLoginIdentityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLoginIdentityResponse() *DeleteLoginIdentityResponse {
	this := DeleteLoginIdentityResponse{}
	return &this
}

// NewDeleteLoginIdentityResponseWithDefaults instantiates a new DeleteLoginIdentityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLoginIdentityResponseWithDefaults() *DeleteLoginIdentityResponse {
	this := DeleteLoginIdentityResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteLoginIdentityResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoginIdentityResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteLoginIdentityResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteLoginIdentityResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o DeleteLoginIdentityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLoginIdentityResponse struct {
	value *DeleteLoginIdentityResponse
	isSet bool
}

func (v NullableDeleteLoginIdentityResponse) Get() *DeleteLoginIdentityResponse {
	return v.value
}

func (v *NullableDeleteLoginIdentityResponse) Set(val *DeleteLoginIdentityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLoginIdentityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLoginIdentityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLoginIdentityResponse(val *DeleteLoginIdentityResponse) *NullableDeleteLoginIdentityResponse {
	return &NullableDeleteLoginIdentityResponse{value: val, isSet: true}
}

func (v NullableDeleteLoginIdentityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLoginIdentityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
