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

// SenderAccountResponse struct for SenderAccountResponse
type SenderAccountResponse struct {
	// A unique identifier generated after creating sender account
	SenderAccountId *string `json:"sender_account_id,omitempty"`
	// Accountholder name of the sender's account
	AccountholderName *string                 `json:"accountholder_name,omitempty"`
	AccountNumber     *RecipientAccountNumber `json:"account_number,omitempty"`
	// Type of sender account.
	AccountType *string `json:"account_type,omitempty"`
	// Finverse Institution ID for the sender’s institution.
	InstitutionId *string `json:"institution_id,omitempty"`
	// A unique identifier generated after creating sender
	SenderId *string `json:"sender_id,omitempty"`
	// Additional attributes of the sender account in key:value format (e.g. sender_id: 1234). It supports up to 10 key:value pairs, whereas the key and value supports up to 50 and 500 characters respectively.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Timestamp of when the sender was created in ISO format (YYYY-MM-DDTHH:MM:SS.SSSZ)
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of when the sender was last updated in ISO format (YYYY-MM-DDTHH:MM:SS.SSSZ)
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSenderAccountResponse instantiates a new SenderAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSenderAccountResponse() *SenderAccountResponse {
	this := SenderAccountResponse{}
	return &this
}

// NewSenderAccountResponseWithDefaults instantiates a new SenderAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderAccountResponseWithDefaults() *SenderAccountResponse {
	this := SenderAccountResponse{}
	return &this
}

// GetSenderAccountId returns the SenderAccountId field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetSenderAccountId() string {
	if o == nil || o.SenderAccountId == nil {
		var ret string
		return ret
	}
	return *o.SenderAccountId
}

// GetSenderAccountIdOk returns a tuple with the SenderAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetSenderAccountIdOk() (*string, bool) {
	if o == nil || o.SenderAccountId == nil {
		return nil, false
	}
	return o.SenderAccountId, true
}

// HasSenderAccountId returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasSenderAccountId() bool {
	if o != nil && o.SenderAccountId != nil {
		return true
	}

	return false
}

// SetSenderAccountId gets a reference to the given string and assigns it to the SenderAccountId field.
func (o *SenderAccountResponse) SetSenderAccountId(v string) {
	o.SenderAccountId = &v
}

// GetAccountholderName returns the AccountholderName field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetAccountholderName() string {
	if o == nil || o.AccountholderName == nil {
		var ret string
		return ret
	}
	return *o.AccountholderName
}

// GetAccountholderNameOk returns a tuple with the AccountholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetAccountholderNameOk() (*string, bool) {
	if o == nil || o.AccountholderName == nil {
		return nil, false
	}
	return o.AccountholderName, true
}

// HasAccountholderName returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasAccountholderName() bool {
	if o != nil && o.AccountholderName != nil {
		return true
	}

	return false
}

// SetAccountholderName gets a reference to the given string and assigns it to the AccountholderName field.
func (o *SenderAccountResponse) SetAccountholderName(v string) {
	o.AccountholderName = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetAccountNumber() RecipientAccountNumber {
	if o == nil || o.AccountNumber == nil {
		var ret RecipientAccountNumber
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetAccountNumberOk() (*RecipientAccountNumber, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given RecipientAccountNumber and assigns it to the AccountNumber field.
func (o *SenderAccountResponse) SetAccountNumber(v RecipientAccountNumber) {
	o.AccountNumber = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *SenderAccountResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *SenderAccountResponse) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *SenderAccountResponse) SetSenderId(v string) {
	o.SenderId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SenderAccountResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SenderAccountResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SenderAccountResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAccountResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SenderAccountResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SenderAccountResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SenderAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SenderAccountId != nil {
		toSerialize["sender_account_id"] = o.SenderAccountId
	}
	if o.AccountholderName != nil {
		toSerialize["accountholder_name"] = o.AccountholderName
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.SenderId != nil {
		toSerialize["sender_id"] = o.SenderId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSenderAccountResponse struct {
	value *SenderAccountResponse
	isSet bool
}

func (v NullableSenderAccountResponse) Get() *SenderAccountResponse {
	return v.value
}

func (v *NullableSenderAccountResponse) Set(val *SenderAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSenderAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSenderAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSenderAccountResponse(val *SenderAccountResponse) *NullableSenderAccountResponse {
	return &NullableSenderAccountResponse{value: val, isSet: true}
}

func (v NullableSenderAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSenderAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}