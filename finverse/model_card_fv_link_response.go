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

// CardFvLinkResponse struct for CardFvLinkResponse
type CardFvLinkResponse struct {
	Status      *string            `json:"status,omitempty"`
	CardDetails *CardFvLinkDetails `json:"card_details,omitempty"`
	Recipient   *CardRecipient     `json:"recipient,omitempty"`
}

// NewCardFvLinkResponse instantiates a new CardFvLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardFvLinkResponse() *CardFvLinkResponse {
	this := CardFvLinkResponse{}
	return &this
}

// NewCardFvLinkResponseWithDefaults instantiates a new CardFvLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardFvLinkResponseWithDefaults() *CardFvLinkResponse {
	this := CardFvLinkResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardFvLinkResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFvLinkResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardFvLinkResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CardFvLinkResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCardDetails returns the CardDetails field value if set, zero value otherwise.
func (o *CardFvLinkResponse) GetCardDetails() CardFvLinkDetails {
	if o == nil || o.CardDetails == nil {
		var ret CardFvLinkDetails
		return ret
	}
	return *o.CardDetails
}

// GetCardDetailsOk returns a tuple with the CardDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFvLinkResponse) GetCardDetailsOk() (*CardFvLinkDetails, bool) {
	if o == nil || o.CardDetails == nil {
		return nil, false
	}
	return o.CardDetails, true
}

// HasCardDetails returns a boolean if a field has been set.
func (o *CardFvLinkResponse) HasCardDetails() bool {
	if o != nil && o.CardDetails != nil {
		return true
	}

	return false
}

// SetCardDetails gets a reference to the given CardFvLinkDetails and assigns it to the CardDetails field.
func (o *CardFvLinkResponse) SetCardDetails(v CardFvLinkDetails) {
	o.CardDetails = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *CardFvLinkResponse) GetRecipient() CardRecipient {
	if o == nil || o.Recipient == nil {
		var ret CardRecipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFvLinkResponse) GetRecipientOk() (*CardRecipient, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *CardFvLinkResponse) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given CardRecipient and assigns it to the Recipient field.
func (o *CardFvLinkResponse) SetRecipient(v CardRecipient) {
	o.Recipient = &v
}

func (o CardFvLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CardDetails != nil {
		toSerialize["card_details"] = o.CardDetails
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	return json.Marshal(toSerialize)
}

type NullableCardFvLinkResponse struct {
	value *CardFvLinkResponse
	isSet bool
}

func (v NullableCardFvLinkResponse) Get() *CardFvLinkResponse {
	return v.value
}

func (v *NullableCardFvLinkResponse) Set(val *CardFvLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardFvLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardFvLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardFvLinkResponse(val *CardFvLinkResponse) *NullableCardFvLinkResponse {
	return &NullableCardFvLinkResponse{value: val, isSet: true}
}

func (v NullableCardFvLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardFvLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}