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

// checks if the CardRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardRecipient{}

// CardRecipient struct for CardRecipient
type CardRecipient struct {
	// Merchant account name
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CardRecipient CardRecipient

// NewCardRecipient instantiates a new CardRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardRecipient() *CardRecipient {
	this := CardRecipient{}
	return &this
}

// NewCardRecipientWithDefaults instantiates a new CardRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardRecipientWithDefaults() *CardRecipient {
	this := CardRecipient{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CardRecipient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardRecipient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CardRecipient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CardRecipient) SetName(v string) {
	o.Name = &v
}

func (o CardRecipient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardRecipient) UnmarshalJSON(data []byte) (err error) {
	varCardRecipient := _CardRecipient{}

	err = json.Unmarshal(data, &varCardRecipient)

	if err != nil {
		return err
	}

	*o = CardRecipient(varCardRecipient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardRecipient struct {
	value *CardRecipient
	isSet bool
}

func (v NullableCardRecipient) Get() *CardRecipient {
	return v.value
}

func (v *NullableCardRecipient) Set(val *CardRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableCardRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableCardRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardRecipient(val *CardRecipient) *NullableCardRecipient {
	return &NullableCardRecipient{value: val, isSet: true}
}

func (v NullableCardRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
