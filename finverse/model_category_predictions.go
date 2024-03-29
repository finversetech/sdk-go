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

// CategoryPredictions struct for CategoryPredictions
type CategoryPredictions struct {
	Categories []string `json:"categories,omitempty"`
	Source     *string  `json:"source,omitempty"`
	SourceId   *string  `json:"source_id,omitempty"`
}

// NewCategoryPredictions instantiates a new CategoryPredictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryPredictions() *CategoryPredictions {
	this := CategoryPredictions{}
	return &this
}

// NewCategoryPredictionsWithDefaults instantiates a new CategoryPredictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryPredictionsWithDefaults() *CategoryPredictions {
	this := CategoryPredictions{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CategoryPredictions) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryPredictions) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CategoryPredictions) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *CategoryPredictions) SetCategories(v []string) {
	o.Categories = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CategoryPredictions) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryPredictions) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CategoryPredictions) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CategoryPredictions) SetSource(v string) {
	o.Source = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *CategoryPredictions) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryPredictions) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *CategoryPredictions) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *CategoryPredictions) SetSourceId(v string) {
	o.SourceId = &v
}

func (o CategoryPredictions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryPredictions struct {
	value *CategoryPredictions
	isSet bool
}

func (v NullableCategoryPredictions) Get() *CategoryPredictions {
	return v.value
}

func (v *NullableCategoryPredictions) Set(val *CategoryPredictions) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryPredictions) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryPredictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryPredictions(val *CategoryPredictions) *NullableCategoryPredictions {
	return &NullableCategoryPredictions{value: val, isSet: true}
}

func (v NullableCategoryPredictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryPredictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
