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

// PaymentLinkCustomizations struct for PaymentLinkCustomizations
type PaymentLinkCustomizations struct {
	// ISO639-1 language code. Language to display when user open the link, default to English (en) if not specified
	Language *string `json:"language,omitempty"`
	// The UI mode link is intended to be used in - \"iframe\", \"auto_redirect\", \"redirect\" or \"standalone\"
	UiMode *string `json:"ui_mode,omitempty"`
	// URI to redirect to. Only needed if ui_mode = redirect
	RedirectUri *string `json:"redirect_uri,omitempty"`
}

// NewPaymentLinkCustomizations instantiates a new PaymentLinkCustomizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinkCustomizations() *PaymentLinkCustomizations {
	this := PaymentLinkCustomizations{}
	return &this
}

// NewPaymentLinkCustomizationsWithDefaults instantiates a new PaymentLinkCustomizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinkCustomizationsWithDefaults() *PaymentLinkCustomizations {
	this := PaymentLinkCustomizations{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PaymentLinkCustomizations) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkCustomizations) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PaymentLinkCustomizations) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PaymentLinkCustomizations) SetLanguage(v string) {
	o.Language = &v
}

// GetUiMode returns the UiMode field value if set, zero value otherwise.
func (o *PaymentLinkCustomizations) GetUiMode() string {
	if o == nil || o.UiMode == nil {
		var ret string
		return ret
	}
	return *o.UiMode
}

// GetUiModeOk returns a tuple with the UiMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkCustomizations) GetUiModeOk() (*string, bool) {
	if o == nil || o.UiMode == nil {
		return nil, false
	}
	return o.UiMode, true
}

// HasUiMode returns a boolean if a field has been set.
func (o *PaymentLinkCustomizations) HasUiMode() bool {
	if o != nil && o.UiMode != nil {
		return true
	}

	return false
}

// SetUiMode gets a reference to the given string and assigns it to the UiMode field.
func (o *PaymentLinkCustomizations) SetUiMode(v string) {
	o.UiMode = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *PaymentLinkCustomizations) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkCustomizations) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *PaymentLinkCustomizations) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *PaymentLinkCustomizations) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

func (o PaymentLinkCustomizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.UiMode != nil {
		toSerialize["ui_mode"] = o.UiMode
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentLinkCustomizations struct {
	value *PaymentLinkCustomizations
	isSet bool
}

func (v NullablePaymentLinkCustomizations) Get() *PaymentLinkCustomizations {
	return v.value
}

func (v *NullablePaymentLinkCustomizations) Set(val *PaymentLinkCustomizations) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinkCustomizations) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinkCustomizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinkCustomizations(val *PaymentLinkCustomizations) *NullablePaymentLinkCustomizations {
	return &NullablePaymentLinkCustomizations{value: val, isSet: true}
}

func (v NullablePaymentLinkCustomizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinkCustomizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
