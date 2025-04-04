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
	"time"
)

// checks if the GetMandateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMandateResponse{}

// GetMandateResponse struct for GetMandateResponse
type GetMandateResponse struct {
	// Timestamp in ISO format (YYYY-MM-DDTHH:MM:SS.SSSZ)
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp in ISO format (YYYY-MM-DDTHH:MM:SS.SSSZ)
	UpdatedAt time.Time `json:"updated_at"`
	// Finverse Mandate ID (ULID)
	MandateId string `json:"mandate_id"`
	// Finverse Payment Method ID (ULID)
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	// Mandate Status
	Status           string                   `json:"status"`
	Recipient        MandateRecipient         `json:"recipient"`
	RecipientAccount *MandateRecipientAccount `json:"recipient_account,omitempty"`
	Sender           GetMandateSender         `json:"sender"`
	SenderAccount    *MandateSenderAccount    `json:"sender_account,omitempty"`
	MandateDetails   MandateDetailsResponse   `json:"mandate_details"`
	Fees             []Fee                    `json:"fees,omitempty"`
	Error            *FvEmbeddedErrorModel    `json:"error,omitempty"`
	// Additional attributes of the mandate in key:value format (e.g. mandate_internal_id: 1234). It supports up to 10 key:value pairs, whereas the key and value supports up to 50 and 1000 characters respectively.
	Metadata             *map[string]string `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMandateResponse GetMandateResponse

// NewGetMandateResponse instantiates a new GetMandateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMandateResponse(updatedAt time.Time, mandateId string, status string, recipient MandateRecipient, sender GetMandateSender, mandateDetails MandateDetailsResponse) *GetMandateResponse {
	this := GetMandateResponse{}
	this.UpdatedAt = updatedAt
	this.MandateId = mandateId
	this.Status = status
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

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetMandateResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetMandateResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetMandateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetMandateResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetMandateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
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

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *GetMandateResponse) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *GetMandateResponse) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *GetMandateResponse) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetStatus returns the Status field value
func (o *GetMandateResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMandateResponse) SetStatus(v string) {
	o.Status = v
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

// GetRecipientAccount returns the RecipientAccount field value if set, zero value otherwise.
func (o *GetMandateResponse) GetRecipientAccount() MandateRecipientAccount {
	if o == nil || IsNil(o.RecipientAccount) {
		var ret MandateRecipientAccount
		return ret
	}
	return *o.RecipientAccount
}

// GetRecipientAccountOk returns a tuple with the RecipientAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetRecipientAccountOk() (*MandateRecipientAccount, bool) {
	if o == nil || IsNil(o.RecipientAccount) {
		return nil, false
	}
	return o.RecipientAccount, true
}

// HasRecipientAccount returns a boolean if a field has been set.
func (o *GetMandateResponse) HasRecipientAccount() bool {
	if o != nil && !IsNil(o.RecipientAccount) {
		return true
	}

	return false
}

// SetRecipientAccount gets a reference to the given MandateRecipientAccount and assigns it to the RecipientAccount field.
func (o *GetMandateResponse) SetRecipientAccount(v MandateRecipientAccount) {
	o.RecipientAccount = &v
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

// GetSenderAccount returns the SenderAccount field value if set, zero value otherwise.
func (o *GetMandateResponse) GetSenderAccount() MandateSenderAccount {
	if o == nil || IsNil(o.SenderAccount) {
		var ret MandateSenderAccount
		return ret
	}
	return *o.SenderAccount
}

// GetSenderAccountOk returns a tuple with the SenderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetSenderAccountOk() (*MandateSenderAccount, bool) {
	if o == nil || IsNil(o.SenderAccount) {
		return nil, false
	}
	return o.SenderAccount, true
}

// HasSenderAccount returns a boolean if a field has been set.
func (o *GetMandateResponse) HasSenderAccount() bool {
	if o != nil && !IsNil(o.SenderAccount) {
		return true
	}

	return false
}

// SetSenderAccount gets a reference to the given MandateSenderAccount and assigns it to the SenderAccount field.
func (o *GetMandateResponse) SetSenderAccount(v MandateSenderAccount) {
	o.SenderAccount = &v
}

// GetMandateDetails returns the MandateDetails field value
func (o *GetMandateResponse) GetMandateDetails() MandateDetailsResponse {
	if o == nil {
		var ret MandateDetailsResponse
		return ret
	}

	return o.MandateDetails
}

// GetMandateDetailsOk returns a tuple with the MandateDetails field value
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetMandateDetailsOk() (*MandateDetailsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateDetails, true
}

// SetMandateDetails sets field value
func (o *GetMandateResponse) SetMandateDetails(v MandateDetailsResponse) {
	o.MandateDetails = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *GetMandateResponse) GetFees() []Fee {
	if o == nil || IsNil(o.Fees) {
		var ret []Fee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetFeesOk() ([]Fee, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *GetMandateResponse) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []Fee and assigns it to the Fees field.
func (o *GetMandateResponse) SetFees(v []Fee) {
	o.Fees = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetMandateResponse) GetError() FvEmbeddedErrorModel {
	if o == nil || IsNil(o.Error) {
		var ret FvEmbeddedErrorModel
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetErrorOk() (*FvEmbeddedErrorModel, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetMandateResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given FvEmbeddedErrorModel and assigns it to the Error field.
func (o *GetMandateResponse) SetError(v FvEmbeddedErrorModel) {
	o.Error = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetMandateResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMandateResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetMandateResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *GetMandateResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o GetMandateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMandateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["mandate_id"] = o.MandateId
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	toSerialize["status"] = o.Status
	toSerialize["recipient"] = o.Recipient
	if !IsNil(o.RecipientAccount) {
		toSerialize["recipient_account"] = o.RecipientAccount
	}
	toSerialize["sender"] = o.Sender
	if !IsNil(o.SenderAccount) {
		toSerialize["sender_account"] = o.SenderAccount
	}
	toSerialize["mandate_details"] = o.MandateDetails
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMandateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"updated_at",
		"mandate_id",
		"status",
		"recipient",
		"sender",
		"mandate_details",
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

	varGetMandateResponse := _GetMandateResponse{}

	err = json.Unmarshal(data, &varGetMandateResponse)

	if err != nil {
		return err
	}

	*o = GetMandateResponse(varGetMandateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "payment_method_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "recipient_account")
		delete(additionalProperties, "sender")
		delete(additionalProperties, "sender_account")
		delete(additionalProperties, "mandate_details")
		delete(additionalProperties, "fees")
		delete(additionalProperties, "error")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
