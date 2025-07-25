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

// checks if the FVCardProcessorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FVCardProcessorDetails{}

// FVCardProcessorDetails struct for FVCardProcessorDetails
type FVCardProcessorDetails struct {
	AuthCode             *string `json:"auth_code,omitempty"`
	ProcessorId          *string `json:"processor_id,omitempty"`
	ProcessorReference   *string `json:"processor_reference,omitempty"`
	TokenId              *string `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FVCardProcessorDetails FVCardProcessorDetails

// NewFVCardProcessorDetails instantiates a new FVCardProcessorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFVCardProcessorDetails() *FVCardProcessorDetails {
	this := FVCardProcessorDetails{}
	return &this
}

// NewFVCardProcessorDetailsWithDefaults instantiates a new FVCardProcessorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFVCardProcessorDetailsWithDefaults() *FVCardProcessorDetails {
	this := FVCardProcessorDetails{}
	return &this
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *FVCardProcessorDetails) GetAuthCode() string {
	if o == nil || IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FVCardProcessorDetails) GetAuthCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *FVCardProcessorDetails) HasAuthCode() bool {
	if o != nil && !IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *FVCardProcessorDetails) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetProcessorId returns the ProcessorId field value if set, zero value otherwise.
func (o *FVCardProcessorDetails) GetProcessorId() string {
	if o == nil || IsNil(o.ProcessorId) {
		var ret string
		return ret
	}
	return *o.ProcessorId
}

// GetProcessorIdOk returns a tuple with the ProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FVCardProcessorDetails) GetProcessorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorId) {
		return nil, false
	}
	return o.ProcessorId, true
}

// HasProcessorId returns a boolean if a field has been set.
func (o *FVCardProcessorDetails) HasProcessorId() bool {
	if o != nil && !IsNil(o.ProcessorId) {
		return true
	}

	return false
}

// SetProcessorId gets a reference to the given string and assigns it to the ProcessorId field.
func (o *FVCardProcessorDetails) SetProcessorId(v string) {
	o.ProcessorId = &v
}

// GetProcessorReference returns the ProcessorReference field value if set, zero value otherwise.
func (o *FVCardProcessorDetails) GetProcessorReference() string {
	if o == nil || IsNil(o.ProcessorReference) {
		var ret string
		return ret
	}
	return *o.ProcessorReference
}

// GetProcessorReferenceOk returns a tuple with the ProcessorReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FVCardProcessorDetails) GetProcessorReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorReference) {
		return nil, false
	}
	return o.ProcessorReference, true
}

// HasProcessorReference returns a boolean if a field has been set.
func (o *FVCardProcessorDetails) HasProcessorReference() bool {
	if o != nil && !IsNil(o.ProcessorReference) {
		return true
	}

	return false
}

// SetProcessorReference gets a reference to the given string and assigns it to the ProcessorReference field.
func (o *FVCardProcessorDetails) SetProcessorReference(v string) {
	o.ProcessorReference = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *FVCardProcessorDetails) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FVCardProcessorDetails) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *FVCardProcessorDetails) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *FVCardProcessorDetails) SetTokenId(v string) {
	o.TokenId = &v
}

func (o FVCardProcessorDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FVCardProcessorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthCode) {
		toSerialize["auth_code"] = o.AuthCode
	}
	if !IsNil(o.ProcessorId) {
		toSerialize["processor_id"] = o.ProcessorId
	}
	if !IsNil(o.ProcessorReference) {
		toSerialize["processor_reference"] = o.ProcessorReference
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FVCardProcessorDetails) UnmarshalJSON(data []byte) (err error) {
	varFVCardProcessorDetails := _FVCardProcessorDetails{}

	err = json.Unmarshal(data, &varFVCardProcessorDetails)

	if err != nil {
		return err
	}

	*o = FVCardProcessorDetails(varFVCardProcessorDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_code")
		delete(additionalProperties, "processor_id")
		delete(additionalProperties, "processor_reference")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFVCardProcessorDetails struct {
	value *FVCardProcessorDetails
	isSet bool
}

func (v NullableFVCardProcessorDetails) Get() *FVCardProcessorDetails {
	return v.value
}

func (v *NullableFVCardProcessorDetails) Set(val *FVCardProcessorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFVCardProcessorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFVCardProcessorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFVCardProcessorDetails(val *FVCardProcessorDetails) *NullableFVCardProcessorDetails {
	return &NullableFVCardProcessorDetails{value: val, isSet: true}
}

func (v NullableFVCardProcessorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFVCardProcessorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
