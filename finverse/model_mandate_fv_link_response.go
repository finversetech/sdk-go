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

// checks if the MandateFvLinkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MandateFvLinkResponse{}

// MandateFvLinkResponse struct for MandateFvLinkResponse
type MandateFvLinkResponse struct {
	MandateId            *string                      `json:"mandate_id,omitempty"`
	InstitutionId        *string                      `json:"institution_id,omitempty"`
	MandateStatus        *string                      `json:"mandate_status,omitempty"`
	Recipient            *MandateRecipient            `json:"recipient,omitempty"`
	SenderAccount        *SenderAccountFvLinkResponse `json:"sender_account,omitempty"`
	Error                *FvEmbeddedErrorModel        `json:"error,omitempty"`
	MandateDetails       *MandateFvLinkDetails        `json:"mandate_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MandateFvLinkResponse MandateFvLinkResponse

// NewMandateFvLinkResponse instantiates a new MandateFvLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandateFvLinkResponse() *MandateFvLinkResponse {
	this := MandateFvLinkResponse{}
	return &this
}

// NewMandateFvLinkResponseWithDefaults instantiates a new MandateFvLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateFvLinkResponseWithDefaults() *MandateFvLinkResponse {
	this := MandateFvLinkResponse{}
	return &this
}

// GetMandateId returns the MandateId field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetMandateId() string {
	if o == nil || IsNil(o.MandateId) {
		var ret string
		return ret
	}
	return *o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetMandateIdOk() (*string, bool) {
	if o == nil || IsNil(o.MandateId) {
		return nil, false
	}
	return o.MandateId, true
}

// HasMandateId returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasMandateId() bool {
	if o != nil && !IsNil(o.MandateId) {
		return true
	}

	return false
}

// SetMandateId gets a reference to the given string and assigns it to the MandateId field.
func (o *MandateFvLinkResponse) SetMandateId(v string) {
	o.MandateId = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetInstitutionId() string {
	if o == nil || IsNil(o.InstitutionId) {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetInstitutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionId) {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasInstitutionId() bool {
	if o != nil && !IsNil(o.InstitutionId) {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *MandateFvLinkResponse) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetMandateStatus returns the MandateStatus field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetMandateStatus() string {
	if o == nil || IsNil(o.MandateStatus) {
		var ret string
		return ret
	}
	return *o.MandateStatus
}

// GetMandateStatusOk returns a tuple with the MandateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetMandateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MandateStatus) {
		return nil, false
	}
	return o.MandateStatus, true
}

// HasMandateStatus returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasMandateStatus() bool {
	if o != nil && !IsNil(o.MandateStatus) {
		return true
	}

	return false
}

// SetMandateStatus gets a reference to the given string and assigns it to the MandateStatus field.
func (o *MandateFvLinkResponse) SetMandateStatus(v string) {
	o.MandateStatus = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetRecipient() MandateRecipient {
	if o == nil || IsNil(o.Recipient) {
		var ret MandateRecipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetRecipientOk() (*MandateRecipient, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given MandateRecipient and assigns it to the Recipient field.
func (o *MandateFvLinkResponse) SetRecipient(v MandateRecipient) {
	o.Recipient = &v
}

// GetSenderAccount returns the SenderAccount field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetSenderAccount() SenderAccountFvLinkResponse {
	if o == nil || IsNil(o.SenderAccount) {
		var ret SenderAccountFvLinkResponse
		return ret
	}
	return *o.SenderAccount
}

// GetSenderAccountOk returns a tuple with the SenderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetSenderAccountOk() (*SenderAccountFvLinkResponse, bool) {
	if o == nil || IsNil(o.SenderAccount) {
		return nil, false
	}
	return o.SenderAccount, true
}

// HasSenderAccount returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasSenderAccount() bool {
	if o != nil && !IsNil(o.SenderAccount) {
		return true
	}

	return false
}

// SetSenderAccount gets a reference to the given SenderAccountFvLinkResponse and assigns it to the SenderAccount field.
func (o *MandateFvLinkResponse) SetSenderAccount(v SenderAccountFvLinkResponse) {
	o.SenderAccount = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetError() FvEmbeddedErrorModel {
	if o == nil || IsNil(o.Error) {
		var ret FvEmbeddedErrorModel
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetErrorOk() (*FvEmbeddedErrorModel, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given FvEmbeddedErrorModel and assigns it to the Error field.
func (o *MandateFvLinkResponse) SetError(v FvEmbeddedErrorModel) {
	o.Error = &v
}

// GetMandateDetails returns the MandateDetails field value if set, zero value otherwise.
func (o *MandateFvLinkResponse) GetMandateDetails() MandateFvLinkDetails {
	if o == nil || IsNil(o.MandateDetails) {
		var ret MandateFvLinkDetails
		return ret
	}
	return *o.MandateDetails
}

// GetMandateDetailsOk returns a tuple with the MandateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateFvLinkResponse) GetMandateDetailsOk() (*MandateFvLinkDetails, bool) {
	if o == nil || IsNil(o.MandateDetails) {
		return nil, false
	}
	return o.MandateDetails, true
}

// HasMandateDetails returns a boolean if a field has been set.
func (o *MandateFvLinkResponse) HasMandateDetails() bool {
	if o != nil && !IsNil(o.MandateDetails) {
		return true
	}

	return false
}

// SetMandateDetails gets a reference to the given MandateFvLinkDetails and assigns it to the MandateDetails field.
func (o *MandateFvLinkResponse) SetMandateDetails(v MandateFvLinkDetails) {
	o.MandateDetails = &v
}

func (o MandateFvLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MandateFvLinkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MandateId) {
		toSerialize["mandate_id"] = o.MandateId
	}
	if !IsNil(o.InstitutionId) {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if !IsNil(o.MandateStatus) {
		toSerialize["mandate_status"] = o.MandateStatus
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.SenderAccount) {
		toSerialize["sender_account"] = o.SenderAccount
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.MandateDetails) {
		toSerialize["mandate_details"] = o.MandateDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MandateFvLinkResponse) UnmarshalJSON(data []byte) (err error) {
	varMandateFvLinkResponse := _MandateFvLinkResponse{}

	err = json.Unmarshal(data, &varMandateFvLinkResponse)

	if err != nil {
		return err
	}

	*o = MandateFvLinkResponse(varMandateFvLinkResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "mandate_status")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "sender_account")
		delete(additionalProperties, "error")
		delete(additionalProperties, "mandate_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMandateFvLinkResponse struct {
	value *MandateFvLinkResponse
	isSet bool
}

func (v NullableMandateFvLinkResponse) Get() *MandateFvLinkResponse {
	return v.value
}

func (v *NullableMandateFvLinkResponse) Set(val *MandateFvLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMandateFvLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMandateFvLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandateFvLinkResponse(val *MandateFvLinkResponse) *NullableMandateFvLinkResponse {
	return &NullableMandateFvLinkResponse{value: val, isSet: true}
}

func (v NullableMandateFvLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandateFvLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
