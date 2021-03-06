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

// BadRequestModel struct for BadRequestModel
type BadRequestModel struct {
	Error *BadRequestModelError `json:"error,omitempty"`
}

// NewBadRequestModel instantiates a new BadRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestModel() *BadRequestModel {
	this := BadRequestModel{}
	return &this
}

// NewBadRequestModelWithDefaults instantiates a new BadRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestModelWithDefaults() *BadRequestModel {
	this := BadRequestModel{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BadRequestModel) GetError() BadRequestModelError {
	if o == nil || o.Error == nil {
		var ret BadRequestModelError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestModel) GetErrorOk() (*BadRequestModelError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BadRequestModel) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given BadRequestModelError and assigns it to the Error field.
func (o *BadRequestModel) SetError(v BadRequestModelError) {
	o.Error = &v
}

func (o BadRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequestModel struct {
	value *BadRequestModel
	isSet bool
}

func (v NullableBadRequestModel) Get() *BadRequestModel {
	return v.value
}

func (v *NullableBadRequestModel) Set(val *BadRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestModel(val *BadRequestModel) *NullableBadRequestModel {
	return &NullableBadRequestModel{value: val, isSet: true}
}

func (v NullableBadRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
