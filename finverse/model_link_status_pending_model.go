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

// LinkStatusPendingModel struct for LinkStatusPendingModel
type LinkStatusPendingModel struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *string `json:"details,omitempty"`
}

// NewLinkStatusPendingModel instantiates a new LinkStatusPendingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkStatusPendingModel() *LinkStatusPendingModel {
	this := LinkStatusPendingModel{}
	return &this
}

// NewLinkStatusPendingModelWithDefaults instantiates a new LinkStatusPendingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkStatusPendingModelWithDefaults() *LinkStatusPendingModel {
	this := LinkStatusPendingModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *LinkStatusPendingModel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkStatusPendingModel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *LinkStatusPendingModel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *LinkStatusPendingModel) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LinkStatusPendingModel) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkStatusPendingModel) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LinkStatusPendingModel) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LinkStatusPendingModel) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *LinkStatusPendingModel) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkStatusPendingModel) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *LinkStatusPendingModel) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *LinkStatusPendingModel) SetDetails(v string) {
	o.Details = &v
}

func (o LinkStatusPendingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableLinkStatusPendingModel struct {
	value *LinkStatusPendingModel
	isSet bool
}

func (v NullableLinkStatusPendingModel) Get() *LinkStatusPendingModel {
	return v.value
}

func (v *NullableLinkStatusPendingModel) Set(val *LinkStatusPendingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkStatusPendingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkStatusPendingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkStatusPendingModel(val *LinkStatusPendingModel) *NullableLinkStatusPendingModel {
	return &NullableLinkStatusPendingModel{value: val, isSet: true}
}

func (v NullableLinkStatusPendingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkStatusPendingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}