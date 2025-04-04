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

// checks if the FvErrorModelV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FvErrorModelV2{}

// FvErrorModelV2 struct for FvErrorModelV2
type FvErrorModelV2 struct {
	// The error type
	Type      string `json:"type"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Details   string `json:"details"`
	// The request_id provided in the request header
	RequestId            string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _FvErrorModelV2 FvErrorModelV2

// NewFvErrorModelV2 instantiates a new FvErrorModelV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFvErrorModelV2(type_ string, errorCode string, message string, details string, requestId string) *FvErrorModelV2 {
	this := FvErrorModelV2{}
	this.Type = type_
	this.ErrorCode = errorCode
	this.Message = message
	this.Details = details
	this.RequestId = requestId
	return &this
}

// NewFvErrorModelV2WithDefaults instantiates a new FvErrorModelV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFvErrorModelV2WithDefaults() *FvErrorModelV2 {
	this := FvErrorModelV2{}
	return &this
}

// GetType returns the Type field value
func (o *FvErrorModelV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FvErrorModelV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FvErrorModelV2) SetType(v string) {
	o.Type = v
}

// GetErrorCode returns the ErrorCode field value
func (o *FvErrorModelV2) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *FvErrorModelV2) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *FvErrorModelV2) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *FvErrorModelV2) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FvErrorModelV2) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FvErrorModelV2) SetMessage(v string) {
	o.Message = v
}

// GetDetails returns the Details field value
func (o *FvErrorModelV2) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *FvErrorModelV2) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *FvErrorModelV2) SetDetails(v string) {
	o.Details = v
}

// GetRequestId returns the RequestId field value
func (o *FvErrorModelV2) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *FvErrorModelV2) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *FvErrorModelV2) SetRequestId(v string) {
	o.RequestId = v
}

func (o FvErrorModelV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FvErrorModelV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	toSerialize["details"] = o.Details
	toSerialize["request_id"] = o.RequestId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FvErrorModelV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"error_code",
		"message",
		"details",
		"request_id",
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

	varFvErrorModelV2 := _FvErrorModelV2{}

	err = json.Unmarshal(data, &varFvErrorModelV2)

	if err != nil {
		return err
	}

	*o = FvErrorModelV2(varFvErrorModelV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "details")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFvErrorModelV2 struct {
	value *FvErrorModelV2
	isSet bool
}

func (v NullableFvErrorModelV2) Get() *FvErrorModelV2 {
	return v.value
}

func (v *NullableFvErrorModelV2) Set(val *FvErrorModelV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFvErrorModelV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFvErrorModelV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFvErrorModelV2(val *FvErrorModelV2) *NullableFvErrorModelV2 {
	return &NullableFvErrorModelV2{value: val, isSet: true}
}

func (v NullableFvErrorModelV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFvErrorModelV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
