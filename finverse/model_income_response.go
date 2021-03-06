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

// IncomeResponse struct for IncomeResponse
type IncomeResponse struct {
	Income        []SingleSourceIncome `json:"income"`
	LoginIdentity LoginIdentityShort   `json:"login_identity"`
	Institution   InstitutionShort     `json:"institution"`
}

// NewIncomeResponse instantiates a new IncomeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeResponse(income []SingleSourceIncome, loginIdentity LoginIdentityShort, institution InstitutionShort) *IncomeResponse {
	this := IncomeResponse{}
	this.Income = income
	this.LoginIdentity = loginIdentity
	this.Institution = institution
	return &this
}

// NewIncomeResponseWithDefaults instantiates a new IncomeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeResponseWithDefaults() *IncomeResponse {
	this := IncomeResponse{}
	return &this
}

// GetIncome returns the Income field value
func (o *IncomeResponse) GetIncome() []SingleSourceIncome {
	if o == nil {
		var ret []SingleSourceIncome
		return ret
	}

	return o.Income
}

// GetIncomeOk returns a tuple with the Income field value
// and a boolean to check if the value has been set.
func (o *IncomeResponse) GetIncomeOk() ([]SingleSourceIncome, bool) {
	if o == nil {
		return nil, false
	}
	return o.Income, true
}

// SetIncome sets field value
func (o *IncomeResponse) SetIncome(v []SingleSourceIncome) {
	o.Income = v
}

// GetLoginIdentity returns the LoginIdentity field value
func (o *IncomeResponse) GetLoginIdentity() LoginIdentityShort {
	if o == nil {
		var ret LoginIdentityShort
		return ret
	}

	return o.LoginIdentity
}

// GetLoginIdentityOk returns a tuple with the LoginIdentity field value
// and a boolean to check if the value has been set.
func (o *IncomeResponse) GetLoginIdentityOk() (*LoginIdentityShort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginIdentity, true
}

// SetLoginIdentity sets field value
func (o *IncomeResponse) SetLoginIdentity(v LoginIdentityShort) {
	o.LoginIdentity = v
}

// GetInstitution returns the Institution field value
func (o *IncomeResponse) GetInstitution() InstitutionShort {
	if o == nil {
		var ret InstitutionShort
		return ret
	}

	return o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value
// and a boolean to check if the value has been set.
func (o *IncomeResponse) GetInstitutionOk() (*InstitutionShort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Institution, true
}

// SetInstitution sets field value
func (o *IncomeResponse) SetInstitution(v InstitutionShort) {
	o.Institution = v
}

func (o IncomeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["income"] = o.Income
	}
	if true {
		toSerialize["login_identity"] = o.LoginIdentity
	}
	if true {
		toSerialize["institution"] = o.Institution
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeResponse struct {
	value *IncomeResponse
	isSet bool
}

func (v NullableIncomeResponse) Get() *IncomeResponse {
	return v.value
}

func (v *NullableIncomeResponse) Set(val *IncomeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeResponse(val *IncomeResponse) *NullableIncomeResponse {
	return &NullableIncomeResponse{value: val, isSet: true}
}

func (v NullableIncomeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
