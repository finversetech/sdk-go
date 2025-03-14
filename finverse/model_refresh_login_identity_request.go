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

// checks if the RefreshLoginIdentityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshLoginIdentityRequest{}

// RefreshLoginIdentityRequest struct for RefreshLoginIdentityRequest
type RefreshLoginIdentityRequest struct {
	// Indicate whether the user is present in this flow. If the user is not present, only institutions that do not require 2fa can be refreshed
	UserPresent          *bool                                   `json:"user_present,omitempty"`
	LinkCustomizations   *RefreshLoginIdentityLinkCustomizations `json:"link_customizations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RefreshLoginIdentityRequest RefreshLoginIdentityRequest

// NewRefreshLoginIdentityRequest instantiates a new RefreshLoginIdentityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshLoginIdentityRequest() *RefreshLoginIdentityRequest {
	this := RefreshLoginIdentityRequest{}
	return &this
}

// NewRefreshLoginIdentityRequestWithDefaults instantiates a new RefreshLoginIdentityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshLoginIdentityRequestWithDefaults() *RefreshLoginIdentityRequest {
	this := RefreshLoginIdentityRequest{}
	return &this
}

// GetUserPresent returns the UserPresent field value if set, zero value otherwise.
func (o *RefreshLoginIdentityRequest) GetUserPresent() bool {
	if o == nil || IsNil(o.UserPresent) {
		var ret bool
		return ret
	}
	return *o.UserPresent
}

// GetUserPresentOk returns a tuple with the UserPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityRequest) GetUserPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.UserPresent) {
		return nil, false
	}
	return o.UserPresent, true
}

// HasUserPresent returns a boolean if a field has been set.
func (o *RefreshLoginIdentityRequest) HasUserPresent() bool {
	if o != nil && !IsNil(o.UserPresent) {
		return true
	}

	return false
}

// SetUserPresent gets a reference to the given bool and assigns it to the UserPresent field.
func (o *RefreshLoginIdentityRequest) SetUserPresent(v bool) {
	o.UserPresent = &v
}

// GetLinkCustomizations returns the LinkCustomizations field value if set, zero value otherwise.
func (o *RefreshLoginIdentityRequest) GetLinkCustomizations() RefreshLoginIdentityLinkCustomizations {
	if o == nil || IsNil(o.LinkCustomizations) {
		var ret RefreshLoginIdentityLinkCustomizations
		return ret
	}
	return *o.LinkCustomizations
}

// GetLinkCustomizationsOk returns a tuple with the LinkCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityRequest) GetLinkCustomizationsOk() (*RefreshLoginIdentityLinkCustomizations, bool) {
	if o == nil || IsNil(o.LinkCustomizations) {
		return nil, false
	}
	return o.LinkCustomizations, true
}

// HasLinkCustomizations returns a boolean if a field has been set.
func (o *RefreshLoginIdentityRequest) HasLinkCustomizations() bool {
	if o != nil && !IsNil(o.LinkCustomizations) {
		return true
	}

	return false
}

// SetLinkCustomizations gets a reference to the given RefreshLoginIdentityLinkCustomizations and assigns it to the LinkCustomizations field.
func (o *RefreshLoginIdentityRequest) SetLinkCustomizations(v RefreshLoginIdentityLinkCustomizations) {
	o.LinkCustomizations = &v
}

func (o RefreshLoginIdentityRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshLoginIdentityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserPresent) {
		toSerialize["user_present"] = o.UserPresent
	}
	if !IsNil(o.LinkCustomizations) {
		toSerialize["link_customizations"] = o.LinkCustomizations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshLoginIdentityRequest) UnmarshalJSON(data []byte) (err error) {
	varRefreshLoginIdentityRequest := _RefreshLoginIdentityRequest{}

	err = json.Unmarshal(data, &varRefreshLoginIdentityRequest)

	if err != nil {
		return err
	}

	*o = RefreshLoginIdentityRequest(varRefreshLoginIdentityRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_present")
		delete(additionalProperties, "link_customizations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshLoginIdentityRequest struct {
	value *RefreshLoginIdentityRequest
	isSet bool
}

func (v NullableRefreshLoginIdentityRequest) Get() *RefreshLoginIdentityRequest {
	return v.value
}

func (v *NullableRefreshLoginIdentityRequest) Set(val *RefreshLoginIdentityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshLoginIdentityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshLoginIdentityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshLoginIdentityRequest(val *RefreshLoginIdentityRequest) *NullableRefreshLoginIdentityRequest {
	return &NullableRefreshLoginIdentityRequest{value: val, isSet: true}
}

func (v NullableRefreshLoginIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshLoginIdentityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
