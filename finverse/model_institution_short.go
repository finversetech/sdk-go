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

// InstitutionShort struct for InstitutionShort
type InstitutionShort struct {
	InstitutionId   *string  `json:"institution_id,omitempty"`
	Countries       []string `json:"countries,omitempty"`
	InstitutionName *string  `json:"institution_name,omitempty"`
	PortalName      *string  `json:"portal_name,omitempty"`
}

// NewInstitutionShort instantiates a new InstitutionShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionShort() *InstitutionShort {
	this := InstitutionShort{}
	return &this
}

// NewInstitutionShortWithDefaults instantiates a new InstitutionShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionShortWithDefaults() *InstitutionShort {
	this := InstitutionShort{}
	return &this
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *InstitutionShort) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionShort) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *InstitutionShort) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *InstitutionShort) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *InstitutionShort) GetCountries() []string {
	if o == nil || o.Countries == nil {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionShort) GetCountriesOk() ([]string, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *InstitutionShort) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *InstitutionShort) SetCountries(v []string) {
	o.Countries = v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *InstitutionShort) GetInstitutionName() string {
	if o == nil || o.InstitutionName == nil {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionShort) GetInstitutionNameOk() (*string, bool) {
	if o == nil || o.InstitutionName == nil {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *InstitutionShort) HasInstitutionName() bool {
	if o != nil && o.InstitutionName != nil {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *InstitutionShort) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

// GetPortalName returns the PortalName field value if set, zero value otherwise.
func (o *InstitutionShort) GetPortalName() string {
	if o == nil || o.PortalName == nil {
		var ret string
		return ret
	}
	return *o.PortalName
}

// GetPortalNameOk returns a tuple with the PortalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionShort) GetPortalNameOk() (*string, bool) {
	if o == nil || o.PortalName == nil {
		return nil, false
	}
	return o.PortalName, true
}

// HasPortalName returns a boolean if a field has been set.
func (o *InstitutionShort) HasPortalName() bool {
	if o != nil && o.PortalName != nil {
		return true
	}

	return false
}

// SetPortalName gets a reference to the given string and assigns it to the PortalName field.
func (o *InstitutionShort) SetPortalName(v string) {
	o.PortalName = &v
}

func (o InstitutionShort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	if o.InstitutionName != nil {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if o.PortalName != nil {
		toSerialize["portal_name"] = o.PortalName
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionShort struct {
	value *InstitutionShort
	isSet bool
}

func (v NullableInstitutionShort) Get() *InstitutionShort {
	return v.value
}

func (v *NullableInstitutionShort) Set(val *InstitutionShort) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionShort) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionShort(val *InstitutionShort) *NullableInstitutionShort {
	return &NullableInstitutionShort{value: val, isSet: true}
}

func (v NullableInstitutionShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
