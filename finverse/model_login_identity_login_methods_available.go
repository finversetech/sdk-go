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

// checks if the LoginIdentityLoginMethodsAvailable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginIdentityLoginMethodsAvailable{}

// LoginIdentityLoginMethodsAvailable struct for LoginIdentityLoginMethodsAvailable
type LoginIdentityLoginMethodsAvailable struct {
	HavePassword         *bool `json:"havePassword,omitempty"`
	HaveSecret           *bool `json:"haveSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginIdentityLoginMethodsAvailable LoginIdentityLoginMethodsAvailable

// NewLoginIdentityLoginMethodsAvailable instantiates a new LoginIdentityLoginMethodsAvailable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginIdentityLoginMethodsAvailable() *LoginIdentityLoginMethodsAvailable {
	this := LoginIdentityLoginMethodsAvailable{}
	return &this
}

// NewLoginIdentityLoginMethodsAvailableWithDefaults instantiates a new LoginIdentityLoginMethodsAvailable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginIdentityLoginMethodsAvailableWithDefaults() *LoginIdentityLoginMethodsAvailable {
	this := LoginIdentityLoginMethodsAvailable{}
	return &this
}

// GetHavePassword returns the HavePassword field value if set, zero value otherwise.
func (o *LoginIdentityLoginMethodsAvailable) GetHavePassword() bool {
	if o == nil || IsNil(o.HavePassword) {
		var ret bool
		return ret
	}
	return *o.HavePassword
}

// GetHavePasswordOk returns a tuple with the HavePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityLoginMethodsAvailable) GetHavePasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HavePassword) {
		return nil, false
	}
	return o.HavePassword, true
}

// HasHavePassword returns a boolean if a field has been set.
func (o *LoginIdentityLoginMethodsAvailable) HasHavePassword() bool {
	if o != nil && !IsNil(o.HavePassword) {
		return true
	}

	return false
}

// SetHavePassword gets a reference to the given bool and assigns it to the HavePassword field.
func (o *LoginIdentityLoginMethodsAvailable) SetHavePassword(v bool) {
	o.HavePassword = &v
}

// GetHaveSecret returns the HaveSecret field value if set, zero value otherwise.
func (o *LoginIdentityLoginMethodsAvailable) GetHaveSecret() bool {
	if o == nil || IsNil(o.HaveSecret) {
		var ret bool
		return ret
	}
	return *o.HaveSecret
}

// GetHaveSecretOk returns a tuple with the HaveSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityLoginMethodsAvailable) GetHaveSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.HaveSecret) {
		return nil, false
	}
	return o.HaveSecret, true
}

// HasHaveSecret returns a boolean if a field has been set.
func (o *LoginIdentityLoginMethodsAvailable) HasHaveSecret() bool {
	if o != nil && !IsNil(o.HaveSecret) {
		return true
	}

	return false
}

// SetHaveSecret gets a reference to the given bool and assigns it to the HaveSecret field.
func (o *LoginIdentityLoginMethodsAvailable) SetHaveSecret(v bool) {
	o.HaveSecret = &v
}

func (o LoginIdentityLoginMethodsAvailable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginIdentityLoginMethodsAvailable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HavePassword) {
		toSerialize["havePassword"] = o.HavePassword
	}
	if !IsNil(o.HaveSecret) {
		toSerialize["haveSecret"] = o.HaveSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginIdentityLoginMethodsAvailable) UnmarshalJSON(data []byte) (err error) {
	varLoginIdentityLoginMethodsAvailable := _LoginIdentityLoginMethodsAvailable{}

	err = json.Unmarshal(data, &varLoginIdentityLoginMethodsAvailable)

	if err != nil {
		return err
	}

	*o = LoginIdentityLoginMethodsAvailable(varLoginIdentityLoginMethodsAvailable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "havePassword")
		delete(additionalProperties, "haveSecret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginIdentityLoginMethodsAvailable struct {
	value *LoginIdentityLoginMethodsAvailable
	isSet bool
}

func (v NullableLoginIdentityLoginMethodsAvailable) Get() *LoginIdentityLoginMethodsAvailable {
	return v.value
}

func (v *NullableLoginIdentityLoginMethodsAvailable) Set(val *LoginIdentityLoginMethodsAvailable) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginIdentityLoginMethodsAvailable) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginIdentityLoginMethodsAvailable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginIdentityLoginMethodsAvailable(val *LoginIdentityLoginMethodsAvailable) *NullableLoginIdentityLoginMethodsAvailable {
	return &NullableLoginIdentityLoginMethodsAvailable{value: val, isSet: true}
}

func (v NullableLoginIdentityLoginMethodsAvailable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginIdentityLoginMethodsAvailable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
