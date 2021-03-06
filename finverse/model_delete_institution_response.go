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

// DeleteInstitutionResponse struct for DeleteInstitutionResponse
type DeleteInstitutionResponse struct {
	Success *bool `json:"success,omitempty"`
}

// NewDeleteInstitutionResponse instantiates a new DeleteInstitutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInstitutionResponse() *DeleteInstitutionResponse {
	this := DeleteInstitutionResponse{}
	return &this
}

// NewDeleteInstitutionResponseWithDefaults instantiates a new DeleteInstitutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInstitutionResponseWithDefaults() *DeleteInstitutionResponse {
	this := DeleteInstitutionResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteInstitutionResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInstitutionResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteInstitutionResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteInstitutionResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o DeleteInstitutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteInstitutionResponse struct {
	value *DeleteInstitutionResponse
	isSet bool
}

func (v NullableDeleteInstitutionResponse) Get() *DeleteInstitutionResponse {
	return v.value
}

func (v *NullableDeleteInstitutionResponse) Set(val *DeleteInstitutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInstitutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInstitutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInstitutionResponse(val *DeleteInstitutionResponse) *NullableDeleteInstitutionResponse {
	return &NullableDeleteInstitutionResponse{value: val, isSet: true}
}

func (v NullableDeleteInstitutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInstitutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
