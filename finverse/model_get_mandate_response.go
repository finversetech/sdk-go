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
	"time"
)

// GetMandateResponse struct for GetMandateResponse
type GetMandateResponse struct {
	// Timestamp in ISO format (YYYY-MM-DDTHH:MM:SS.SSSZ)
	LastUpdate time.Time `json:"last_update"`
	// Finverse Mandate ID (ULID)
	MandateId string `json:"mandate_id"`
	// Mandate status
	MandateStatus  string           `json:"mandate_status"`
	Recipient      MandateRecipient `json:"recipient"`
	Sender         GetMandateSender `json:"sender"`
	MandateDetails MandateDetails   `json:"mandate_details"`
}

// NewGetMandateResponse instantiates a new GetMandateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMandateResponse(lastUpdate time.Time, mandateId string, mandateStatus string, recipient MandateRecipient, sender GetMandateSender, mandateDetails MandateDetails) *GetMandateResponse {
	this := GetMandateResponse{}
	this.LastUpdate = lastUpdate
	this.MandateId = mandateId
	this.MandateStatus = mandateStatus
	this.Recipient = recipient
	this.Sender = sender
	this.MandateDetails = mandateDetails
	return &this
}

// NewGetMandateResponseWithDefaults instantiates a new GetMandateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMandateResponseWithDefaults() *GetMandateResponse {
	this := GetMandateResponse{}
	return &this
}

// GetLastUpdate returns the LastUpdate field value
func (o *GetMandateResponse) GetLastUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *GetMandateResponse) SetLastUpdate(v time.Time) {
	o.LastUpdate = v
}

// GetMandateId returns the MandateId field value
func (o *GetMandateResponse) GetMandateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateId, true
}

// SetMandateId sets field value
func (o *GetMandateResponse) SetMandateId(v string) {
	o.MandateId = v
}

// GetMandateStatus returns the MandateStatus field value
func (o *GetMandateResponse) GetMandateStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateStatus
}

// GetMandateStatusOk returns a tuple with the MandateStatus field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetMandateStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateStatus, true
}

// SetMandateStatus sets field value
func (o *GetMandateResponse) SetMandateStatus(v string) {
	o.MandateStatus = v
}

// GetRecipient returns the Recipient field value
func (o *GetMandateResponse) GetRecipient() MandateRecipient {
	if o == nil {
		var ret MandateRecipient
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetRecipientOk() (*MandateRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *GetMandateResponse) SetRecipient(v MandateRecipient) {
	o.Recipient = v
}

// GetSender returns the Sender field value
func (o *GetMandateResponse) GetSender() GetMandateSender {
	if o == nil {
		var ret GetMandateSender
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetSenderOk() (*GetMandateSender, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *GetMandateResponse) SetSender(v GetMandateSender) {
	o.Sender = v
}

// GetMandateDetails returns the MandateDetails field value
func (o *GetMandateResponse) GetMandateDetails() MandateDetails {
	if o == nil {
		var ret MandateDetails
		return ret
	}

	return o.MandateDetails
}

// GetMandateDetailsOk returns a tuple with the MandateDetails field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetMandateDetailsOk() (*MandateDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateDetails, true
}

// SetMandateDetails sets field value
func (o *GetMandateResponse) SetMandateDetails(v MandateDetails) {
	o.MandateDetails = v
}

func (o GetMandateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["last_update"] = o.LastUpdate
	}
	if true {
		toSerialize["mandate_id"] = o.MandateId
	}
	if true {
		toSerialize["mandate_status"] = o.MandateStatus
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["mandate_details"] = o.MandateDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGetMandateResponse struct {
	value *GetMandateResponse
	isSet bool
}

func (v NullableGetMandateResponse) Get() *GetMandateResponse {
	return v.value
}

func (v *NullableGetMandateResponse) Set(val *GetMandateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMandateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMandateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMandateResponse(val *GetMandateResponse) *NullableGetMandateResponse {
	return &NullableGetMandateResponse{value: val, isSet: true}
}

func (v NullableGetMandateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMandateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
