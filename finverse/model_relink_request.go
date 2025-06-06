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

// checks if the RelinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelinkRequest{}

// RelinkRequest struct for RelinkRequest
type RelinkRequest struct {
	StoreCredential bool `json:"store_credential"`
	// this is a mandatory field
	Consent              NullableBool `json:"consent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RelinkRequest RelinkRequest

// NewRelinkRequest instantiates a new RelinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelinkRequest(storeCredential bool) *RelinkRequest {
	this := RelinkRequest{}
	this.StoreCredential = storeCredential
	return &this
}

// NewRelinkRequestWithDefaults instantiates a new RelinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelinkRequestWithDefaults() *RelinkRequest {
	this := RelinkRequest{}
	return &this
}

// GetStoreCredential returns the StoreCredential field value
func (o *RelinkRequest) GetStoreCredential() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StoreCredential
}

// GetStoreCredentialOk returns a tuple with the StoreCredential field value
// and a boolean to check if the value has been set.
func (o *RelinkRequest) GetStoreCredentialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreCredential, true
}

// SetStoreCredential sets field value
func (o *RelinkRequest) SetStoreCredential(v bool) {
	o.StoreCredential = v
}

// GetConsent returns the Consent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RelinkRequest) GetConsent() bool {
	if o == nil || IsNil(o.Consent.Get()) {
		var ret bool
		return ret
	}
	return *o.Consent.Get()
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelinkRequest) GetConsentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Consent.Get(), o.Consent.IsSet()
}

// HasConsent returns a boolean if a field has been set.
func (o *RelinkRequest) HasConsent() bool {
	if o != nil && o.Consent.IsSet() {
		return true
	}

	return false
}

// SetConsent gets a reference to the given NullableBool and assigns it to the Consent field.
func (o *RelinkRequest) SetConsent(v bool) {
	o.Consent.Set(&v)
}

// SetConsentNil sets the value for Consent to be an explicit nil
func (o *RelinkRequest) SetConsentNil() {
	o.Consent.Set(nil)
}

// UnsetConsent ensures that no value is present for Consent, not even an explicit nil
func (o *RelinkRequest) UnsetConsent() {
	o.Consent.Unset()
}

func (o RelinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["store_credential"] = o.StoreCredential
	if o.Consent.IsSet() {
		toSerialize["consent"] = o.Consent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RelinkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"store_credential",
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

	varRelinkRequest := _RelinkRequest{}

	err = json.Unmarshal(data, &varRelinkRequest)

	if err != nil {
		return err
	}

	*o = RelinkRequest(varRelinkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "store_credential")
		delete(additionalProperties, "consent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelinkRequest struct {
	value *RelinkRequest
	isSet bool
}

func (v NullableRelinkRequest) Get() *RelinkRequest {
	return v.value
}

func (v *NullableRelinkRequest) Set(val *RelinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRelinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRelinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelinkRequest(val *RelinkRequest) *NullableRelinkRequest {
	return &NullableRelinkRequest{value: val, isSet: true}
}

func (v NullableRelinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
