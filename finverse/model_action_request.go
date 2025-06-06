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

// checks if the ActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionRequest{}

// ActionRequest struct for ActionRequest
type ActionRequest struct {
	EncryptedCredentials EncryptedPayload `json:"encrypted_credentials"`
	// The action id
	ActionId             string `json:"action_id"`
	AdditionalProperties map[string]interface{}
}

type _ActionRequest ActionRequest

// NewActionRequest instantiates a new ActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionRequest(encryptedCredentials EncryptedPayload, actionId string) *ActionRequest {
	this := ActionRequest{}
	this.EncryptedCredentials = encryptedCredentials
	this.ActionId = actionId
	return &this
}

// NewActionRequestWithDefaults instantiates a new ActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionRequestWithDefaults() *ActionRequest {
	this := ActionRequest{}
	return &this
}

// GetEncryptedCredentials returns the EncryptedCredentials field value
func (o *ActionRequest) GetEncryptedCredentials() EncryptedPayload {
	if o == nil {
		var ret EncryptedPayload
		return ret
	}

	return o.EncryptedCredentials
}

// GetEncryptedCredentialsOk returns a tuple with the EncryptedCredentials field value
// and a boolean to check if the value has been set.
func (o *ActionRequest) GetEncryptedCredentialsOk() (*EncryptedPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedCredentials, true
}

// SetEncryptedCredentials sets field value
func (o *ActionRequest) SetEncryptedCredentials(v EncryptedPayload) {
	o.EncryptedCredentials = v
}

// GetActionId returns the ActionId field value
func (o *ActionRequest) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *ActionRequest) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *ActionRequest) SetActionId(v string) {
	o.ActionId = v
}

func (o ActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encrypted_credentials"] = o.EncryptedCredentials
	toSerialize["action_id"] = o.ActionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ActionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encrypted_credentials",
		"action_id",
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

	varActionRequest := _ActionRequest{}

	err = json.Unmarshal(data, &varActionRequest)

	if err != nil {
		return err
	}

	*o = ActionRequest(varActionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "encrypted_credentials")
		delete(additionalProperties, "action_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActionRequest struct {
	value *ActionRequest
	isSet bool
}

func (v NullableActionRequest) Get() *ActionRequest {
	return v.value
}

func (v *NullableActionRequest) Set(val *ActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRequest(val *ActionRequest) *NullableActionRequest {
	return &NullableActionRequest{value: val, isSet: true}
}

func (v NullableActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
