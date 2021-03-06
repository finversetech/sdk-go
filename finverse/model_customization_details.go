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

// CustomizationDetails struct for CustomizationDetails
type CustomizationDetails struct {
	LogoId      *string `json:"logo_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

// NewCustomizationDetails instantiates a new CustomizationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomizationDetails() *CustomizationDetails {
	this := CustomizationDetails{}
	return &this
}

// NewCustomizationDetailsWithDefaults instantiates a new CustomizationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomizationDetailsWithDefaults() *CustomizationDetails {
	this := CustomizationDetails{}
	return &this
}

// GetLogoId returns the LogoId field value if set, zero value otherwise.
func (o *CustomizationDetails) GetLogoId() string {
	if o == nil || o.LogoId == nil {
		var ret string
		return ret
	}
	return *o.LogoId
}

// GetLogoIdOk returns a tuple with the LogoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationDetails) GetLogoIdOk() (*string, bool) {
	if o == nil || o.LogoId == nil {
		return nil, false
	}
	return o.LogoId, true
}

// HasLogoId returns a boolean if a field has been set.
func (o *CustomizationDetails) HasLogoId() bool {
	if o != nil && o.LogoId != nil {
		return true
	}

	return false
}

// SetLogoId gets a reference to the given string and assigns it to the LogoId field.
func (o *CustomizationDetails) SetLogoId(v string) {
	o.LogoId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CustomizationDetails) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CustomizationDetails) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CustomizationDetails) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o CustomizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoId != nil {
		toSerialize["logo_id"] = o.LogoId
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableCustomizationDetails struct {
	value *CustomizationDetails
	isSet bool
}

func (v NullableCustomizationDetails) Get() *CustomizationDetails {
	return v.value
}

func (v *NullableCustomizationDetails) Set(val *CustomizationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomizationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomizationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomizationDetails(val *CustomizationDetails) *NullableCustomizationDetails {
	return &NullableCustomizationDetails{value: val, isSet: true}
}

func (v NullableCustomizationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomizationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
