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

// CreateMandateWithSenderAccountRequest struct for CreateMandateWithSenderAccountRequest
type CreateMandateWithSenderAccountRequest struct {
	RecipientAccount MandateRecipientRequest     `json:"recipient_account"`
	SenderAccount    MandateSenderAccountRequest `json:"sender_account"`
	MandateDetails   MandateDetails              `json:"mandate_details"`
	// Additional attributes of the mandate in key:value format (e.g. mandate_internal_id: 1234). It supports up to 10 key:value pairs, whereas the key and value supports up to 50 and 500 characters respectively.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewCreateMandateWithSenderAccountRequest instantiates a new CreateMandateWithSenderAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMandateWithSenderAccountRequest(recipientAccount MandateRecipientRequest, senderAccount MandateSenderAccountRequest, mandateDetails MandateDetails) *CreateMandateWithSenderAccountRequest {
	this := CreateMandateWithSenderAccountRequest{}
	this.RecipientAccount = recipientAccount
	this.SenderAccount = senderAccount
	this.MandateDetails = mandateDetails
	return &this
}

// NewCreateMandateWithSenderAccountRequestWithDefaults instantiates a new CreateMandateWithSenderAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMandateWithSenderAccountRequestWithDefaults() *CreateMandateWithSenderAccountRequest {
	this := CreateMandateWithSenderAccountRequest{}
	return &this
}

// GetRecipientAccount returns the RecipientAccount field value
func (o *CreateMandateWithSenderAccountRequest) GetRecipientAccount() MandateRecipientRequest {
	if o == nil {
		var ret MandateRecipientRequest
		return ret
	}

	return o.RecipientAccount
}

// GetRecipientAccountOk returns a tuple with the RecipientAccount field value
// and a boolean to check if the value has been set.
func (o *CreateMandateWithSenderAccountRequest) GetRecipientAccountOk() (*MandateRecipientRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAccount, true
}

// SetRecipientAccount sets field value
func (o *CreateMandateWithSenderAccountRequest) SetRecipientAccount(v MandateRecipientRequest) {
	o.RecipientAccount = v
}

// GetSenderAccount returns the SenderAccount field value
func (o *CreateMandateWithSenderAccountRequest) GetSenderAccount() MandateSenderAccountRequest {
	if o == nil {
		var ret MandateSenderAccountRequest
		return ret
	}

	return o.SenderAccount
}

// GetSenderAccountOk returns a tuple with the SenderAccount field value
// and a boolean to check if the value has been set.
func (o *CreateMandateWithSenderAccountRequest) GetSenderAccountOk() (*MandateSenderAccountRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAccount, true
}

// SetSenderAccount sets field value
func (o *CreateMandateWithSenderAccountRequest) SetSenderAccount(v MandateSenderAccountRequest) {
	o.SenderAccount = v
}

// GetMandateDetails returns the MandateDetails field value
func (o *CreateMandateWithSenderAccountRequest) GetMandateDetails() MandateDetails {
	if o == nil {
		var ret MandateDetails
		return ret
	}

	return o.MandateDetails
}

// GetMandateDetailsOk returns a tuple with the MandateDetails field value
// and a boolean to check if the value has been set.
func (o *CreateMandateWithSenderAccountRequest) GetMandateDetailsOk() (*MandateDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateDetails, true
}

// SetMandateDetails sets field value
func (o *CreateMandateWithSenderAccountRequest) SetMandateDetails(v MandateDetails) {
	o.MandateDetails = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateMandateWithSenderAccountRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMandateWithSenderAccountRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateMandateWithSenderAccountRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateMandateWithSenderAccountRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o CreateMandateWithSenderAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_account"] = o.RecipientAccount
	}
	if true {
		toSerialize["sender_account"] = o.SenderAccount
	}
	if true {
		toSerialize["mandate_details"] = o.MandateDetails
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMandateWithSenderAccountRequest struct {
	value *CreateMandateWithSenderAccountRequest
	isSet bool
}

func (v NullableCreateMandateWithSenderAccountRequest) Get() *CreateMandateWithSenderAccountRequest {
	return v.value
}

func (v *NullableCreateMandateWithSenderAccountRequest) Set(val *CreateMandateWithSenderAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMandateWithSenderAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMandateWithSenderAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMandateWithSenderAccountRequest(val *CreateMandateWithSenderAccountRequest) *NullableCreateMandateWithSenderAccountRequest {
	return &NullableCreateMandateWithSenderAccountRequest{value: val, isSet: true}
}

func (v NullableCreateMandateWithSenderAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMandateWithSenderAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}