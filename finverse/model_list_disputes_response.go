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

// checks if the ListDisputesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDisputesResponse{}

// ListDisputesResponse struct for ListDisputesResponse
type ListDisputesResponse struct {
	Disputes             []DisputeResponse `json:"disputes"`
	TotalDisputes        int32             `json:"total_disputes"`
	AdditionalProperties map[string]interface{}
}

type _ListDisputesResponse ListDisputesResponse

// NewListDisputesResponse instantiates a new ListDisputesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDisputesResponse(disputes []DisputeResponse, totalDisputes int32) *ListDisputesResponse {
	this := ListDisputesResponse{}
	this.Disputes = disputes
	this.TotalDisputes = totalDisputes
	return &this
}

// NewListDisputesResponseWithDefaults instantiates a new ListDisputesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDisputesResponseWithDefaults() *ListDisputesResponse {
	this := ListDisputesResponse{}
	return &this
}

// GetDisputes returns the Disputes field value
func (o *ListDisputesResponse) GetDisputes() []DisputeResponse {
	if o == nil {
		var ret []DisputeResponse
		return ret
	}

	return o.Disputes
}

// GetDisputesOk returns a tuple with the Disputes field value
// and a boolean to check if the value has been set.
func (o *ListDisputesResponse) GetDisputesOk() ([]DisputeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disputes, true
}

// SetDisputes sets field value
func (o *ListDisputesResponse) SetDisputes(v []DisputeResponse) {
	o.Disputes = v
}

// GetTotalDisputes returns the TotalDisputes field value
func (o *ListDisputesResponse) GetTotalDisputes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalDisputes
}

// GetTotalDisputesOk returns a tuple with the TotalDisputes field value
// and a boolean to check if the value has been set.
func (o *ListDisputesResponse) GetTotalDisputesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDisputes, true
}

// SetTotalDisputes sets field value
func (o *ListDisputesResponse) SetTotalDisputes(v int32) {
	o.TotalDisputes = v
}

func (o ListDisputesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDisputesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputes"] = o.Disputes
	toSerialize["total_disputes"] = o.TotalDisputes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDisputesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disputes",
		"total_disputes",
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

	varListDisputesResponse := _ListDisputesResponse{}

	err = json.Unmarshal(data, &varListDisputesResponse)

	if err != nil {
		return err
	}

	*o = ListDisputesResponse(varListDisputesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disputes")
		delete(additionalProperties, "total_disputes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDisputesResponse struct {
	value *ListDisputesResponse
	isSet bool
}

func (v NullableListDisputesResponse) Get() *ListDisputesResponse {
	return v.value
}

func (v *NullableListDisputesResponse) Set(val *ListDisputesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDisputesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDisputesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDisputesResponse(val *ListDisputesResponse) *NullableListDisputesResponse {
	return &NullableListDisputesResponse{value: val, isSet: true}
}

func (v NullableListDisputesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDisputesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
