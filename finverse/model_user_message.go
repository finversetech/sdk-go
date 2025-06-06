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

// checks if the UserMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMessage{}

// UserMessage struct for UserMessage
type UserMessage struct {
	// The name of the message
	Name string `json:"name"`
	// The type of the message. This will help how the UI renders this text.
	Type string `json:"type"`
	// The actual text value.
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _UserMessage UserMessage

// NewUserMessage instantiates a new UserMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMessage(name string, type_ string, value string) *UserMessage {
	this := UserMessage{}
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewUserMessageWithDefaults instantiates a new UserMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMessageWithDefaults() *UserMessage {
	this := UserMessage{}
	return &this
}

// GetName returns the Name field value
func (o *UserMessage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserMessage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserMessage) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *UserMessage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserMessage) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *UserMessage) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserMessage) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserMessage) SetValue(v string) {
	o.Value = v
}

func (o UserMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"value",
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

	varUserMessage := _UserMessage{}

	err = json.Unmarshal(data, &varUserMessage)

	if err != nil {
		return err
	}

	*o = UserMessage(varUserMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserMessage struct {
	value *UserMessage
	isSet bool
}

func (v NullableUserMessage) Get() *UserMessage {
	return v.value
}

func (v *NullableUserMessage) Set(val *UserMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMessage(val *UserMessage) *NullableUserMessage {
	return &NullableUserMessage{value: val, isSet: true}
}

func (v NullableUserMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
