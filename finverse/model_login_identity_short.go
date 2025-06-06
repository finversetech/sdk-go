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

// checks if the LoginIdentityShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginIdentityShort{}

// LoginIdentityShort struct for LoginIdentityShort
type LoginIdentityShort struct {
	LoginIdentityId      *string `json:"login_identity_id,omitempty"`
	Status               *string `json:"status,omitempty"`
	LastSessionId        *string `json:"last_session_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginIdentityShort LoginIdentityShort

// NewLoginIdentityShort instantiates a new LoginIdentityShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginIdentityShort() *LoginIdentityShort {
	this := LoginIdentityShort{}
	return &this
}

// NewLoginIdentityShortWithDefaults instantiates a new LoginIdentityShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginIdentityShortWithDefaults() *LoginIdentityShort {
	this := LoginIdentityShort{}
	return &this
}

// GetLoginIdentityId returns the LoginIdentityId field value if set, zero value otherwise.
func (o *LoginIdentityShort) GetLoginIdentityId() string {
	if o == nil || IsNil(o.LoginIdentityId) {
		var ret string
		return ret
	}
	return *o.LoginIdentityId
}

// GetLoginIdentityIdOk returns a tuple with the LoginIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityShort) GetLoginIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoginIdentityId) {
		return nil, false
	}
	return o.LoginIdentityId, true
}

// HasLoginIdentityId returns a boolean if a field has been set.
func (o *LoginIdentityShort) HasLoginIdentityId() bool {
	if o != nil && !IsNil(o.LoginIdentityId) {
		return true
	}

	return false
}

// SetLoginIdentityId gets a reference to the given string and assigns it to the LoginIdentityId field.
func (o *LoginIdentityShort) SetLoginIdentityId(v string) {
	o.LoginIdentityId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoginIdentityShort) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityShort) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoginIdentityShort) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LoginIdentityShort) SetStatus(v string) {
	o.Status = &v
}

// GetLastSessionId returns the LastSessionId field value if set, zero value otherwise.
func (o *LoginIdentityShort) GetLastSessionId() string {
	if o == nil || IsNil(o.LastSessionId) {
		var ret string
		return ret
	}
	return *o.LastSessionId
}

// GetLastSessionIdOk returns a tuple with the LastSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginIdentityShort) GetLastSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastSessionId) {
		return nil, false
	}
	return o.LastSessionId, true
}

// HasLastSessionId returns a boolean if a field has been set.
func (o *LoginIdentityShort) HasLastSessionId() bool {
	if o != nil && !IsNil(o.LastSessionId) {
		return true
	}

	return false
}

// SetLastSessionId gets a reference to the given string and assigns it to the LastSessionId field.
func (o *LoginIdentityShort) SetLastSessionId(v string) {
	o.LastSessionId = &v
}

func (o LoginIdentityShort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginIdentityShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoginIdentityId) {
		toSerialize["login_identity_id"] = o.LoginIdentityId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LastSessionId) {
		toSerialize["last_session_id"] = o.LastSessionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginIdentityShort) UnmarshalJSON(data []byte) (err error) {
	varLoginIdentityShort := _LoginIdentityShort{}

	err = json.Unmarshal(data, &varLoginIdentityShort)

	if err != nil {
		return err
	}

	*o = LoginIdentityShort(varLoginIdentityShort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "login_identity_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "last_session_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginIdentityShort struct {
	value *LoginIdentityShort
	isSet bool
}

func (v NullableLoginIdentityShort) Get() *LoginIdentityShort {
	return v.value
}

func (v *NullableLoginIdentityShort) Set(val *LoginIdentityShort) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginIdentityShort) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginIdentityShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginIdentityShort(val *LoginIdentityShort) *NullableLoginIdentityShort {
	return &NullableLoginIdentityShort{value: val, isSet: true}
}

func (v NullableLoginIdentityShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginIdentityShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
