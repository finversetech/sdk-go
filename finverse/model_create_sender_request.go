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

// CreateSenderRequest struct for CreateSenderRequest
type CreateSenderRequest struct {
	// Sender’s name/nickname (note: this does not need to match the actual accountholder name of the sender’s account)
	Name string `json:"name"`
	// Customer App's internal ID for the sender
	UserId string `json:"user_id"`
	// Sender details which will be used for fraud checking.
	SenderDetails []SenderDetail `json:"sender_details,omitempty"`
	// Type of account held by the Sender at the Institution. Possible values are INDIVIDUAL, BUSINESS
	SenderType string `json:"sender_type"`
	// Additional attributes of the sender in key:value format (e.g. employer_name: Apple Inc for a payroll case where sender is an employee)
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewCreateSenderRequest instantiates a new CreateSenderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSenderRequest(name string, userId string, senderType string) *CreateSenderRequest {
	this := CreateSenderRequest{}
	this.Name = name
	this.UserId = userId
	this.SenderType = senderType
	return &this
}

// NewCreateSenderRequestWithDefaults instantiates a new CreateSenderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSenderRequestWithDefaults() *CreateSenderRequest {
	this := CreateSenderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSenderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSenderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSenderRequest) SetName(v string) {
	o.Name = v
}

// GetUserId returns the UserId field value
func (o *CreateSenderRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateSenderRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateSenderRequest) SetUserId(v string) {
	o.UserId = v
}

// GetSenderDetails returns the SenderDetails field value if set, zero value otherwise.
func (o *CreateSenderRequest) GetSenderDetails() []SenderDetail {
	if o == nil || o.SenderDetails == nil {
		var ret []SenderDetail
		return ret
	}
	return o.SenderDetails
}

// GetSenderDetailsOk returns a tuple with the SenderDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSenderRequest) GetSenderDetailsOk() ([]SenderDetail, bool) {
	if o == nil || o.SenderDetails == nil {
		return nil, false
	}
	return o.SenderDetails, true
}

// HasSenderDetails returns a boolean if a field has been set.
func (o *CreateSenderRequest) HasSenderDetails() bool {
	if o != nil && o.SenderDetails != nil {
		return true
	}

	return false
}

// SetSenderDetails gets a reference to the given []SenderDetail and assigns it to the SenderDetails field.
func (o *CreateSenderRequest) SetSenderDetails(v []SenderDetail) {
	o.SenderDetails = v
}

// GetSenderType returns the SenderType field value
func (o *CreateSenderRequest) GetSenderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderType
}

// GetSenderTypeOk returns a tuple with the SenderType field value
// and a boolean to check if the value has been set.
func (o *CreateSenderRequest) GetSenderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderType, true
}

// SetSenderType sets field value
func (o *CreateSenderRequest) SetSenderType(v string) {
	o.SenderType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateSenderRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSenderRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateSenderRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateSenderRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o CreateSenderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.SenderDetails != nil {
		toSerialize["sender_details"] = o.SenderDetails
	}
	if true {
		toSerialize["sender_type"] = o.SenderType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSenderRequest struct {
	value *CreateSenderRequest
	isSet bool
}

func (v NullableCreateSenderRequest) Get() *CreateSenderRequest {
	return v.value
}

func (v *NullableCreateSenderRequest) Set(val *CreateSenderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSenderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSenderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSenderRequest(val *CreateSenderRequest) *NullableCreateSenderRequest {
	return &NullableCreateSenderRequest{value: val, isSet: true}
}

func (v NullableCreateSenderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSenderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}