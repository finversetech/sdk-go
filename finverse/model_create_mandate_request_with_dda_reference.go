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

// checks if the CreateMandateRequestWithDdaReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMandateRequestWithDdaReference{}

// CreateMandateRequestWithDdaReference struct for CreateMandateRequestWithDdaReference
type CreateMandateRequestWithDdaReference struct {
	RecipientAccount MandateRecipientRequest               `json:"recipient_account"`
	SenderAccount    MandateSenderAccountRequest           `json:"sender_account"`
	MandateDetails   MandateDetailsRequestWithDdaReference `json:"mandate_details"`
	Metadata         *map[string]string                    `json:"metadata,omitempty"`
	// The mandate status
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CreateMandateRequestWithDdaReference CreateMandateRequestWithDdaReference

// NewCreateMandateRequestWithDdaReference instantiates a new CreateMandateRequestWithDdaReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMandateRequestWithDdaReference(recipientAccount MandateRecipientRequest, senderAccount MandateSenderAccountRequest, mandateDetails MandateDetailsRequestWithDdaReference, status string) *CreateMandateRequestWithDdaReference {
	this := CreateMandateRequestWithDdaReference{}
	this.RecipientAccount = recipientAccount
	this.SenderAccount = senderAccount
	this.MandateDetails = mandateDetails
	this.Status = status
	return &this
}

// NewCreateMandateRequestWithDdaReferenceWithDefaults instantiates a new CreateMandateRequestWithDdaReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMandateRequestWithDdaReferenceWithDefaults() *CreateMandateRequestWithDdaReference {
	this := CreateMandateRequestWithDdaReference{}
	return &this
}

// GetRecipientAccount returns the RecipientAccount field value
func (o *CreateMandateRequestWithDdaReference) GetRecipientAccount() MandateRecipientRequest {
	if o == nil {
		var ret MandateRecipientRequest
		return ret
	}

	return o.RecipientAccount
}

// GetRecipientAccountOk returns a tuple with the RecipientAccount field value
// and a boolean to check if the value has been set.
func (o *CreateMandateRequestWithDdaReference) GetRecipientAccountOk() (*MandateRecipientRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAccount, true
}

// SetRecipientAccount sets field value
func (o *CreateMandateRequestWithDdaReference) SetRecipientAccount(v MandateRecipientRequest) {
	o.RecipientAccount = v
}

// GetSenderAccount returns the SenderAccount field value
func (o *CreateMandateRequestWithDdaReference) GetSenderAccount() MandateSenderAccountRequest {
	if o == nil {
		var ret MandateSenderAccountRequest
		return ret
	}

	return o.SenderAccount
}

// GetSenderAccountOk returns a tuple with the SenderAccount field value
// and a boolean to check if the value has been set.
func (o *CreateMandateRequestWithDdaReference) GetSenderAccountOk() (*MandateSenderAccountRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAccount, true
}

// SetSenderAccount sets field value
func (o *CreateMandateRequestWithDdaReference) SetSenderAccount(v MandateSenderAccountRequest) {
	o.SenderAccount = v
}

// GetMandateDetails returns the MandateDetails field value
func (o *CreateMandateRequestWithDdaReference) GetMandateDetails() MandateDetailsRequestWithDdaReference {
	if o == nil {
		var ret MandateDetailsRequestWithDdaReference
		return ret
	}

	return o.MandateDetails
}

// GetMandateDetailsOk returns a tuple with the MandateDetails field value
// and a boolean to check if the value has been set.
func (o *CreateMandateRequestWithDdaReference) GetMandateDetailsOk() (*MandateDetailsRequestWithDdaReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateDetails, true
}

// SetMandateDetails sets field value
func (o *CreateMandateRequestWithDdaReference) SetMandateDetails(v MandateDetailsRequestWithDdaReference) {
	o.MandateDetails = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateMandateRequestWithDdaReference) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMandateRequestWithDdaReference) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateMandateRequestWithDdaReference) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateMandateRequestWithDdaReference) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetStatus returns the Status field value
func (o *CreateMandateRequestWithDdaReference) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateMandateRequestWithDdaReference) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateMandateRequestWithDdaReference) SetStatus(v string) {
	o.Status = v
}

func (o CreateMandateRequestWithDdaReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMandateRequestWithDdaReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipient_account"] = o.RecipientAccount
	toSerialize["sender_account"] = o.SenderAccount
	toSerialize["mandate_details"] = o.MandateDetails
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMandateRequestWithDdaReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipient_account",
		"sender_account",
		"mandate_details",
		"status",
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

	varCreateMandateRequestWithDdaReference := _CreateMandateRequestWithDdaReference{}

	err = json.Unmarshal(data, &varCreateMandateRequestWithDdaReference)

	if err != nil {
		return err
	}

	*o = CreateMandateRequestWithDdaReference(varCreateMandateRequestWithDdaReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient_account")
		delete(additionalProperties, "sender_account")
		delete(additionalProperties, "mandate_details")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMandateRequestWithDdaReference struct {
	value *CreateMandateRequestWithDdaReference
	isSet bool
}

func (v NullableCreateMandateRequestWithDdaReference) Get() *CreateMandateRequestWithDdaReference {
	return v.value
}

func (v *NullableCreateMandateRequestWithDdaReference) Set(val *CreateMandateRequestWithDdaReference) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMandateRequestWithDdaReference) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMandateRequestWithDdaReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMandateRequestWithDdaReference(val *CreateMandateRequestWithDdaReference) *NullableCreateMandateRequestWithDdaReference {
	return &NullableCreateMandateRequestWithDdaReference{value: val, isSet: true}
}

func (v NullableCreateMandateRequestWithDdaReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMandateRequestWithDdaReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
