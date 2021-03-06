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

// CreateCustomerRequest struct for CreateCustomerRequest
type CreateCustomerRequest struct {
	// The email of the customer. This has to be unique.
	Name string `json:"name"`
	// The email of the customer. This has to be unique.
	Email string `json:"email"`
	// Primary key to identity the customer
	CustomerAppId string `json:"customer_app_id"`
	// The redirect callback
	RedirectUris []string `json:"redirect_uris"`
	// The webhook uris
	WebhookUris []string `json:"webhook_uris,omitempty"`
}

// NewCreateCustomerRequest instantiates a new CreateCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerRequest(name string, email string, customerAppId string, redirectUris []string) *CreateCustomerRequest {
	this := CreateCustomerRequest{}
	this.Name = name
	this.Email = email
	this.CustomerAppId = customerAppId
	this.RedirectUris = redirectUris
	return &this
}

// NewCreateCustomerRequestWithDefaults instantiates a new CreateCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerRequestWithDefaults() *CreateCustomerRequest {
	this := CreateCustomerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustomerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomerRequest) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CreateCustomerRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateCustomerRequest) SetEmail(v string) {
	o.Email = v
}

// GetCustomerAppId returns the CustomerAppId field value
func (o *CreateCustomerRequest) GetCustomerAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerAppId
}

// GetCustomerAppIdOk returns a tuple with the CustomerAppId field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetCustomerAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAppId, true
}

// SetCustomerAppId sets field value
func (o *CreateCustomerRequest) SetCustomerAppId(v string) {
	o.CustomerAppId = v
}

// GetRedirectUris returns the RedirectUris field value
func (o *CreateCustomerRequest) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetRedirectUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *CreateCustomerRequest) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetWebhookUris returns the WebhookUris field value if set, zero value otherwise.
func (o *CreateCustomerRequest) GetWebhookUris() []string {
	if o == nil || o.WebhookUris == nil {
		var ret []string
		return ret
	}
	return o.WebhookUris
}

// GetWebhookUrisOk returns a tuple with the WebhookUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetWebhookUrisOk() ([]string, bool) {
	if o == nil || o.WebhookUris == nil {
		return nil, false
	}
	return o.WebhookUris, true
}

// HasWebhookUris returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasWebhookUris() bool {
	if o != nil && o.WebhookUris != nil {
		return true
	}

	return false
}

// SetWebhookUris gets a reference to the given []string and assigns it to the WebhookUris field.
func (o *CreateCustomerRequest) SetWebhookUris(v []string) {
	o.WebhookUris = v
}

func (o CreateCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["customer_app_id"] = o.CustomerAppId
	}
	if true {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.WebhookUris != nil {
		toSerialize["webhook_uris"] = o.WebhookUris
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCustomerRequest struct {
	value *CreateCustomerRequest
	isSet bool
}

func (v NullableCreateCustomerRequest) Get() *CreateCustomerRequest {
	return v.value
}

func (v *NullableCreateCustomerRequest) Set(val *CreateCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerRequest(val *CreateCustomerRequest) *NullableCreateCustomerRequest {
	return &NullableCreateCustomerRequest{value: val, isSet: true}
}

func (v NullableCreateCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
