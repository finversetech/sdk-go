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

// checks if the AuthorizeMandateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeMandateRequest{}

// AuthorizeMandateRequest struct for AuthorizeMandateRequest
type AuthorizeMandateRequest struct {
	// Whether a consent was provided by the enduser to authorize a mandate
	EnduserConsent       bool `json:"enduser_consent"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizeMandateRequest AuthorizeMandateRequest

// NewAuthorizeMandateRequest instantiates a new AuthorizeMandateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeMandateRequest(enduserConsent bool) *AuthorizeMandateRequest {
	this := AuthorizeMandateRequest{}
	this.EnduserConsent = enduserConsent
	return &this
}

// NewAuthorizeMandateRequestWithDefaults instantiates a new AuthorizeMandateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeMandateRequestWithDefaults() *AuthorizeMandateRequest {
	this := AuthorizeMandateRequest{}
	return &this
}

// GetEnduserConsent returns the EnduserConsent field value
func (o *AuthorizeMandateRequest) GetEnduserConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnduserConsent
}

// GetEnduserConsentOk returns a tuple with the EnduserConsent field value
// and a boolean to check if the value has been set.
func (o *AuthorizeMandateRequest) GetEnduserConsentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnduserConsent, true
}

// SetEnduserConsent sets field value
func (o *AuthorizeMandateRequest) SetEnduserConsent(v bool) {
	o.EnduserConsent = v
}

func (o AuthorizeMandateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeMandateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enduser_consent"] = o.EnduserConsent

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizeMandateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enduser_consent",
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

	varAuthorizeMandateRequest := _AuthorizeMandateRequest{}

	err = json.Unmarshal(data, &varAuthorizeMandateRequest)

	if err != nil {
		return err
	}

	*o = AuthorizeMandateRequest(varAuthorizeMandateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enduser_consent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizeMandateRequest struct {
	value *AuthorizeMandateRequest
	isSet bool
}

func (v NullableAuthorizeMandateRequest) Get() *AuthorizeMandateRequest {
	return v.value
}

func (v *NullableAuthorizeMandateRequest) Set(val *AuthorizeMandateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeMandateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeMandateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeMandateRequest(val *AuthorizeMandateRequest) *NullableAuthorizeMandateRequest {
	return &NullableAuthorizeMandateRequest{value: val, isSet: true}
}

func (v NullableAuthorizeMandateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeMandateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
