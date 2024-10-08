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

// RefreshLoginIdentityLinkCustomizations struct for RefreshLoginIdentityLinkCustomizations
type RefreshLoginIdentityLinkCustomizations struct {
	// ISO639-1 language code. Language to display when user open the link, default to English (en) if not specified
	Language *string `json:"language,omitempty"`
	UiMode   *string `json:"ui_mode,omitempty"`
	// Required if ui_mode is redirect or auto_redirect
	RedirectUri *string `json:"redirect_uri,omitempty"`
	State       *string `json:"state,omitempty"`
}

// NewRefreshLoginIdentityLinkCustomizations instantiates a new RefreshLoginIdentityLinkCustomizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshLoginIdentityLinkCustomizations() *RefreshLoginIdentityLinkCustomizations {
	this := RefreshLoginIdentityLinkCustomizations{}
	return &this
}

// NewRefreshLoginIdentityLinkCustomizationsWithDefaults instantiates a new RefreshLoginIdentityLinkCustomizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshLoginIdentityLinkCustomizationsWithDefaults() *RefreshLoginIdentityLinkCustomizations {
	this := RefreshLoginIdentityLinkCustomizations{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *RefreshLoginIdentityLinkCustomizations) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityLinkCustomizations) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *RefreshLoginIdentityLinkCustomizations) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *RefreshLoginIdentityLinkCustomizations) SetLanguage(v string) {
	o.Language = &v
}

// GetUiMode returns the UiMode field value if set, zero value otherwise.
func (o *RefreshLoginIdentityLinkCustomizations) GetUiMode() string {
	if o == nil || o.UiMode == nil {
		var ret string
		return ret
	}
	return *o.UiMode
}

// GetUiModeOk returns a tuple with the UiMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityLinkCustomizations) GetUiModeOk() (*string, bool) {
	if o == nil || o.UiMode == nil {
		return nil, false
	}
	return o.UiMode, true
}

// HasUiMode returns a boolean if a field has been set.
func (o *RefreshLoginIdentityLinkCustomizations) HasUiMode() bool {
	if o != nil && o.UiMode != nil {
		return true
	}

	return false
}

// SetUiMode gets a reference to the given string and assigns it to the UiMode field.
func (o *RefreshLoginIdentityLinkCustomizations) SetUiMode(v string) {
	o.UiMode = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *RefreshLoginIdentityLinkCustomizations) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityLinkCustomizations) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *RefreshLoginIdentityLinkCustomizations) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *RefreshLoginIdentityLinkCustomizations) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RefreshLoginIdentityLinkCustomizations) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshLoginIdentityLinkCustomizations) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RefreshLoginIdentityLinkCustomizations) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RefreshLoginIdentityLinkCustomizations) SetState(v string) {
	o.State = &v
}

func (o RefreshLoginIdentityLinkCustomizations) MarshalJSON() ([]byte, error) {
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
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshLoginIdentityLinkCustomizations struct {
	value *RefreshLoginIdentityLinkCustomizations
	isSet bool
}

func (v NullableRefreshLoginIdentityLinkCustomizations) Get() *RefreshLoginIdentityLinkCustomizations {
	return v.value
}

func (v *NullableRefreshLoginIdentityLinkCustomizations) Set(val *RefreshLoginIdentityLinkCustomizations) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshLoginIdentityLinkCustomizations) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshLoginIdentityLinkCustomizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshLoginIdentityLinkCustomizations(val *RefreshLoginIdentityLinkCustomizations) *NullableRefreshLoginIdentityLinkCustomizations {
	return &NullableRefreshLoginIdentityLinkCustomizations{value: val, isSet: true}
}

func (v NullableRefreshLoginIdentityLinkCustomizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshLoginIdentityLinkCustomizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
