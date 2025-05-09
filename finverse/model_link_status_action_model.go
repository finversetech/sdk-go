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

// checks if the LinkStatusActionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkStatusActionModel{}

// LinkStatusActionModel struct for LinkStatusActionModel
type LinkStatusActionModel struct {
	// Unique identifier
	ActionId string `json:"action_id"`
	// The type of user screen the UI is to render
	Type string `json:"type"`
	// The name of the user screen the UI is to render
	Name                 string        `json:"name"`
	Messages             []UserMessage `json:"messages"`
	Fields               []UserField   `json:"fields"`
	Buttons              []UserButton  `json:"buttons,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkStatusActionModel LinkStatusActionModel

// NewLinkStatusActionModel instantiates a new LinkStatusActionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkStatusActionModel(actionId string, type_ string, name string, messages []UserMessage, fields []UserField) *LinkStatusActionModel {
	this := LinkStatusActionModel{}
	this.ActionId = actionId
	this.Type = type_
	this.Name = name
	this.Messages = messages
	this.Fields = fields
	return &this
}

// NewLinkStatusActionModelWithDefaults instantiates a new LinkStatusActionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkStatusActionModelWithDefaults() *LinkStatusActionModel {
	this := LinkStatusActionModel{}
	return &this
}

// GetActionId returns the ActionId field value
func (o *LinkStatusActionModel) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *LinkStatusActionModel) SetActionId(v string) {
	o.ActionId = v
}

// GetType returns the Type field value
func (o *LinkStatusActionModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkStatusActionModel) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *LinkStatusActionModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkStatusActionModel) SetName(v string) {
	o.Name = v
}

// GetMessages returns the Messages field value
func (o *LinkStatusActionModel) GetMessages() []UserMessage {
	if o == nil {
		var ret []UserMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetMessagesOk() ([]UserMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *LinkStatusActionModel) SetMessages(v []UserMessage) {
	o.Messages = v
}

// GetFields returns the Fields field value
func (o *LinkStatusActionModel) GetFields() []UserField {
	if o == nil {
		var ret []UserField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetFieldsOk() ([]UserField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *LinkStatusActionModel) SetFields(v []UserField) {
	o.Fields = v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *LinkStatusActionModel) GetButtons() []UserButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []UserButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkStatusActionModel) GetButtonsOk() ([]UserButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *LinkStatusActionModel) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []UserButton and assigns it to the Buttons field.
func (o *LinkStatusActionModel) SetButtons(v []UserButton) {
	o.Buttons = v
}

func (o LinkStatusActionModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkStatusActionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_id"] = o.ActionId
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["messages"] = o.Messages
	toSerialize["fields"] = o.Fields
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkStatusActionModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_id",
		"type",
		"name",
		"messages",
		"fields",
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

	varLinkStatusActionModel := _LinkStatusActionModel{}

	err = json.Unmarshal(data, &varLinkStatusActionModel)

	if err != nil {
		return err
	}

	*o = LinkStatusActionModel(varLinkStatusActionModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "buttons")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkStatusActionModel struct {
	value *LinkStatusActionModel
	isSet bool
}

func (v NullableLinkStatusActionModel) Get() *LinkStatusActionModel {
	return v.value
}

func (v *NullableLinkStatusActionModel) Set(val *LinkStatusActionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkStatusActionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkStatusActionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkStatusActionModel(val *LinkStatusActionModel) *NullableLinkStatusActionModel {
	return &NullableLinkStatusActionModel{value: val, isSet: true}
}

func (v NullableLinkStatusActionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkStatusActionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
