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

// CreateFpsTokenResponse struct for CreateFpsTokenResponse
type CreateFpsTokenResponse struct {
	FpsToken PaymentLinkTokenResponse `json:"fps_token"`
}

// NewCreateFpsTokenResponse instantiates a new CreateFpsTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFpsTokenResponse(fpsToken PaymentLinkTokenResponse) *CreateFpsTokenResponse {
	this := CreateFpsTokenResponse{}
	this.FpsToken = fpsToken
	return &this
}

// NewCreateFpsTokenResponseWithDefaults instantiates a new CreateFpsTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFpsTokenResponseWithDefaults() *CreateFpsTokenResponse {
	this := CreateFpsTokenResponse{}
	return &this
}

// GetFpsToken returns the FpsToken field value
func (o *CreateFpsTokenResponse) GetFpsToken() PaymentLinkTokenResponse {
	if o == nil {
		var ret PaymentLinkTokenResponse
		return ret
	}

	return o.FpsToken
}

// GetFpsTokenOk returns a tuple with the FpsToken field value
// and a boolean to check if the value has been set.
func (o *CreateFpsTokenResponse) GetFpsTokenOk() (*PaymentLinkTokenResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FpsToken, true
}

// SetFpsToken sets field value
func (o *CreateFpsTokenResponse) SetFpsToken(v PaymentLinkTokenResponse) {
	o.FpsToken = v
}

func (o CreateFpsTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fps_token"] = o.FpsToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFpsTokenResponse struct {
	value *CreateFpsTokenResponse
	isSet bool
}

func (v NullableCreateFpsTokenResponse) Get() *CreateFpsTokenResponse {
	return v.value
}

func (v *NullableCreateFpsTokenResponse) Set(val *CreateFpsTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFpsTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFpsTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFpsTokenResponse(val *CreateFpsTokenResponse) *NullableCreateFpsTokenResponse {
	return &NullableCreateFpsTokenResponse{value: val, isSet: true}
}

func (v NullableCreateFpsTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFpsTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}