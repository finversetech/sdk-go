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

// CreateRecipientRequest struct for CreateRecipientRequest
type CreateRecipientRequest struct {
	// Recipient's name/nickname (note: this does not need to match the actual accountholder name of the recipient's account)
	Name             string                 `json:"name"`
	RecipientAccount CreateRecipientAccount `json:"recipient_account"`
	// Customer App's internal ID for the recipient
	UserId string `json:"user_id"`
	// Additional attributes of the recipient in key:value format (e.g. employer_name: Apple Inc for a payroll case where recipient is an employee)
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewCreateRecipientRequest instantiates a new CreateRecipientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecipientRequest(name string, recipientAccount CreateRecipientAccount, userId string) *CreateRecipientRequest {
	this := CreateRecipientRequest{}
	this.Name = name
	this.RecipientAccount = recipientAccount
	this.UserId = userId
	return &this
}

// NewCreateRecipientRequestWithDefaults instantiates a new CreateRecipientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecipientRequestWithDefaults() *CreateRecipientRequest {
	this := CreateRecipientRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRecipientRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRecipientRequest) SetName(v string) {
	o.Name = v
}

// GetRecipientAccount returns the RecipientAccount field value
func (o *CreateRecipientRequest) GetRecipientAccount() CreateRecipientAccount {
	if o == nil {
		var ret CreateRecipientAccount
		return ret
	}

	return o.RecipientAccount
}

// GetRecipientAccountOk returns a tuple with the RecipientAccount field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientRequest) GetRecipientAccountOk() (*CreateRecipientAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAccount, true
}

// SetRecipientAccount sets field value
func (o *CreateRecipientRequest) SetRecipientAccount(v CreateRecipientAccount) {
	o.RecipientAccount = v
}

// GetUserId returns the UserId field value
func (o *CreateRecipientRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateRecipientRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateRecipientRequest) SetUserId(v string) {
	o.UserId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateRecipientRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecipientRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateRecipientRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateRecipientRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o CreateRecipientRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["recipient_account"] = o.RecipientAccount
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRecipientRequest struct {
	value *CreateRecipientRequest
	isSet bool
}

func (v NullableCreateRecipientRequest) Get() *CreateRecipientRequest {
	return v.value
}

func (v *NullableCreateRecipientRequest) Set(val *CreateRecipientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecipientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecipientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecipientRequest(val *CreateRecipientRequest) *NullableCreateRecipientRequest {
	return &NullableCreateRecipientRequest{value: val, isSet: true}
}

func (v NullableCreateRecipientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecipientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
