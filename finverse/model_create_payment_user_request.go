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

// CreatePaymentUserRequest struct for CreatePaymentUserRequest
type CreatePaymentUserRequest struct {
	Name           string             `json:"name"`
	ExternalUserId string             `json:"external_user_id"`
	UserType       *string            `json:"user_type,omitempty"`
	Email          *string            `json:"email,omitempty"`
	UserDetails    []SenderDetail     `json:"user_details,omitempty"`
	Metadata       *map[string]string `json:"metadata,omitempty"`
}

// NewCreatePaymentUserRequest instantiates a new CreatePaymentUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentUserRequest(name string, externalUserId string) *CreatePaymentUserRequest {
	this := CreatePaymentUserRequest{}
	this.Name = name
	this.ExternalUserId = externalUserId
	return &this
}

// NewCreatePaymentUserRequestWithDefaults instantiates a new CreatePaymentUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentUserRequestWithDefaults() *CreatePaymentUserRequest {
	this := CreatePaymentUserRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePaymentUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePaymentUserRequest) SetName(v string) {
	o.Name = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *CreatePaymentUserRequest) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *CreatePaymentUserRequest) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *CreatePaymentUserRequest) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *CreatePaymentUserRequest) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *CreatePaymentUserRequest) SetUserType(v string) {
	o.UserType = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreatePaymentUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreatePaymentUserRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreatePaymentUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetUserDetails returns the UserDetails field value if set, zero value otherwise.
func (o *CreatePaymentUserRequest) GetUserDetails() []SenderDetail {
	if o == nil || o.UserDetails == nil {
		var ret []SenderDetail
		return ret
	}
	return o.UserDetails
}

// GetUserDetailsOk returns a tuple with the UserDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetUserDetailsOk() ([]SenderDetail, bool) {
	if o == nil || o.UserDetails == nil {
		return nil, false
	}
	return o.UserDetails, true
}

// HasUserDetails returns a boolean if a field has been set.
func (o *CreatePaymentUserRequest) HasUserDetails() bool {
	if o != nil && o.UserDetails != nil {
		return true
	}

	return false
}

// SetUserDetails gets a reference to the given []SenderDetail and assigns it to the UserDetails field.
func (o *CreatePaymentUserRequest) SetUserDetails(v []SenderDetail) {
	o.UserDetails = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreatePaymentUserRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentUserRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreatePaymentUserRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreatePaymentUserRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o CreatePaymentUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["external_user_id"] = o.ExternalUserId
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.UserDetails != nil {
		toSerialize["user_details"] = o.UserDetails
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePaymentUserRequest struct {
	value *CreatePaymentUserRequest
	isSet bool
}

func (v NullableCreatePaymentUserRequest) Get() *CreatePaymentUserRequest {
	return v.value
}

func (v *NullableCreatePaymentUserRequest) Set(val *CreatePaymentUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentUserRequest(val *CreatePaymentUserRequest) *NullableCreatePaymentUserRequest {
	return &NullableCreatePaymentUserRequest{value: val, isSet: true}
}

func (v NullableCreatePaymentUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}