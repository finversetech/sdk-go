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

// ApiRelinkRequest struct for ApiRelinkRequest
type ApiRelinkRequest struct {
	StoreCredential      *bool            `json:"store_credential,omitempty"`
	Consent              bool             `json:"consent"`
	EncryptedCredentials EncryptedPayload `json:"encrypted_credentials"`
}

// NewApiRelinkRequest instantiates a new ApiRelinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRelinkRequest(consent bool, encryptedCredentials EncryptedPayload) *ApiRelinkRequest {
	this := ApiRelinkRequest{}
	this.Consent = consent
	this.EncryptedCredentials = encryptedCredentials
	return &this
}

// NewApiRelinkRequestWithDefaults instantiates a new ApiRelinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRelinkRequestWithDefaults() *ApiRelinkRequest {
	this := ApiRelinkRequest{}
	return &this
}

// GetStoreCredential returns the StoreCredential field value if set, zero value otherwise.
func (o *ApiRelinkRequest) GetStoreCredential() bool {
	if o == nil || o.StoreCredential == nil {
		var ret bool
		return ret
	}
	return *o.StoreCredential
}

// GetStoreCredentialOk returns a tuple with the StoreCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRelinkRequest) GetStoreCredentialOk() (*bool, bool) {
	if o == nil || o.StoreCredential == nil {
		return nil, false
	}
	return o.StoreCredential, true
}

// HasStoreCredential returns a boolean if a field has been set.
func (o *ApiRelinkRequest) HasStoreCredential() bool {
	if o != nil && o.StoreCredential != nil {
		return true
	}

	return false
}

// SetStoreCredential gets a reference to the given bool and assigns it to the StoreCredential field.
func (o *ApiRelinkRequest) SetStoreCredential(v bool) {
	o.StoreCredential = &v
}

// GetConsent returns the Consent field value
func (o *ApiRelinkRequest) GetConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *ApiRelinkRequest) GetConsentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *ApiRelinkRequest) SetConsent(v bool) {
	o.Consent = v
}

// GetEncryptedCredentials returns the EncryptedCredentials field value
func (o *ApiRelinkRequest) GetEncryptedCredentials() EncryptedPayload {
	if o == nil {
		var ret EncryptedPayload
		return ret
	}

	return o.EncryptedCredentials
}

// GetEncryptedCredentialsOk returns a tuple with the EncryptedCredentials field value
// and a boolean to check if the value has been set.
func (o *ApiRelinkRequest) GetEncryptedCredentialsOk() (*EncryptedPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedCredentials, true
}

// SetEncryptedCredentials sets field value
func (o *ApiRelinkRequest) SetEncryptedCredentials(v EncryptedPayload) {
	o.EncryptedCredentials = v
}

func (o ApiRelinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoreCredential != nil {
		toSerialize["store_credential"] = o.StoreCredential
	}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if true {
		toSerialize["encrypted_credentials"] = o.EncryptedCredentials
	}
	return json.Marshal(toSerialize)
}

type NullableApiRelinkRequest struct {
	value *ApiRelinkRequest
	isSet bool
}

func (v NullableApiRelinkRequest) Get() *ApiRelinkRequest {
	return v.value
}

func (v *NullableApiRelinkRequest) Set(val *ApiRelinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRelinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRelinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRelinkRequest(val *ApiRelinkRequest) *NullableApiRelinkRequest {
	return &NullableApiRelinkRequest{value: val, isSet: true}
}

func (v NullableApiRelinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRelinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
