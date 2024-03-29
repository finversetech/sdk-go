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

// MandateAuthLinkCustomizations struct for MandateAuthLinkCustomizations
type MandateAuthLinkCustomizations struct {
	// institution's country filter
	Countries []string `json:"countries,omitempty"`
	// Institution to preselect
	InstitutionId *string `json:"institution_id,omitempty"`
	// institution's status filter
	InstitutionStatus *string `json:"institution_status,omitempty"`
	// ISO639-1 language code. Language to display when user open the link, default to English (en) if not specified
	Language *string `json:"language,omitempty"`
	// Space separated list of the tags of the institutions to view.
	LinkMode *string `json:"link_mode,omitempty"`
	// institution's supported product filter. For mandate authorization, this field should contain [\"MANDATE\"]
	ProductsSupported []string `json:"products_supported,omitempty"`
	// The UI mode link is intended to be used in - \"iframe\", \"auto_redirect\", \"redirect\" or \"standalone\"
	UiMode *string `json:"ui_mode,omitempty"`
	// The URI to redirect to. Required if ui_mode is \"redirect\" or \"auto_redirect\"
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// institution's supported user_type filter
	UserType []string `json:"user_type,omitempty"`
}

// NewMandateAuthLinkCustomizations instantiates a new MandateAuthLinkCustomizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandateAuthLinkCustomizations() *MandateAuthLinkCustomizations {
	this := MandateAuthLinkCustomizations{}
	return &this
}

// NewMandateAuthLinkCustomizationsWithDefaults instantiates a new MandateAuthLinkCustomizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateAuthLinkCustomizationsWithDefaults() *MandateAuthLinkCustomizations {
	this := MandateAuthLinkCustomizations{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetCountries() []string {
	if o == nil || o.Countries == nil {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetCountriesOk() ([]string, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *MandateAuthLinkCustomizations) SetCountries(v []string) {
	o.Countries = v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *MandateAuthLinkCustomizations) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetInstitutionStatus returns the InstitutionStatus field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetInstitutionStatus() string {
	if o == nil || o.InstitutionStatus == nil {
		var ret string
		return ret
	}
	return *o.InstitutionStatus
}

// GetInstitutionStatusOk returns a tuple with the InstitutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetInstitutionStatusOk() (*string, bool) {
	if o == nil || o.InstitutionStatus == nil {
		return nil, false
	}
	return o.InstitutionStatus, true
}

// HasInstitutionStatus returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasInstitutionStatus() bool {
	if o != nil && o.InstitutionStatus != nil {
		return true
	}

	return false
}

// SetInstitutionStatus gets a reference to the given string and assigns it to the InstitutionStatus field.
func (o *MandateAuthLinkCustomizations) SetInstitutionStatus(v string) {
	o.InstitutionStatus = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MandateAuthLinkCustomizations) SetLanguage(v string) {
	o.Language = &v
}

// GetLinkMode returns the LinkMode field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetLinkMode() string {
	if o == nil || o.LinkMode == nil {
		var ret string
		return ret
	}
	return *o.LinkMode
}

// GetLinkModeOk returns a tuple with the LinkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetLinkModeOk() (*string, bool) {
	if o == nil || o.LinkMode == nil {
		return nil, false
	}
	return o.LinkMode, true
}

// HasLinkMode returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasLinkMode() bool {
	if o != nil && o.LinkMode != nil {
		return true
	}

	return false
}

// SetLinkMode gets a reference to the given string and assigns it to the LinkMode field.
func (o *MandateAuthLinkCustomizations) SetLinkMode(v string) {
	o.LinkMode = &v
}

// GetProductsSupported returns the ProductsSupported field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetProductsSupported() []string {
	if o == nil || o.ProductsSupported == nil {
		var ret []string
		return ret
	}
	return o.ProductsSupported
}

// GetProductsSupportedOk returns a tuple with the ProductsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetProductsSupportedOk() ([]string, bool) {
	if o == nil || o.ProductsSupported == nil {
		return nil, false
	}
	return o.ProductsSupported, true
}

// HasProductsSupported returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasProductsSupported() bool {
	if o != nil && o.ProductsSupported != nil {
		return true
	}

	return false
}

// SetProductsSupported gets a reference to the given []string and assigns it to the ProductsSupported field.
func (o *MandateAuthLinkCustomizations) SetProductsSupported(v []string) {
	o.ProductsSupported = v
}

// GetUiMode returns the UiMode field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetUiMode() string {
	if o == nil || o.UiMode == nil {
		var ret string
		return ret
	}
	return *o.UiMode
}

// GetUiModeOk returns a tuple with the UiMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetUiModeOk() (*string, bool) {
	if o == nil || o.UiMode == nil {
		return nil, false
	}
	return o.UiMode, true
}

// HasUiMode returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasUiMode() bool {
	if o != nil && o.UiMode != nil {
		return true
	}

	return false
}

// SetUiMode gets a reference to the given string and assigns it to the UiMode field.
func (o *MandateAuthLinkCustomizations) SetUiMode(v string) {
	o.UiMode = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *MandateAuthLinkCustomizations) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *MandateAuthLinkCustomizations) GetUserType() []string {
	if o == nil || o.UserType == nil {
		var ret []string
		return ret
	}
	return o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MandateAuthLinkCustomizations) GetUserTypeOk() ([]string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *MandateAuthLinkCustomizations) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given []string and assigns it to the UserType field.
func (o *MandateAuthLinkCustomizations) SetUserType(v []string) {
	o.UserType = v
}

func (o MandateAuthLinkCustomizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.InstitutionStatus != nil {
		toSerialize["institution_status"] = o.InstitutionStatus
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.LinkMode != nil {
		toSerialize["link_mode"] = o.LinkMode
	}
	if o.ProductsSupported != nil {
		toSerialize["products_supported"] = o.ProductsSupported
	}
	if o.UiMode != nil {
		toSerialize["ui_mode"] = o.UiMode
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	return json.Marshal(toSerialize)
}

type NullableMandateAuthLinkCustomizations struct {
	value *MandateAuthLinkCustomizations
	isSet bool
}

func (v NullableMandateAuthLinkCustomizations) Get() *MandateAuthLinkCustomizations {
	return v.value
}

func (v *NullableMandateAuthLinkCustomizations) Set(val *MandateAuthLinkCustomizations) {
	v.value = val
	v.isSet = true
}

func (v NullableMandateAuthLinkCustomizations) IsSet() bool {
	return v.isSet
}

func (v *NullableMandateAuthLinkCustomizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandateAuthLinkCustomizations(val *MandateAuthLinkCustomizations) *NullableMandateAuthLinkCustomizations {
	return &NullableMandateAuthLinkCustomizations{value: val, isSet: true}
}

func (v NullableMandateAuthLinkCustomizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandateAuthLinkCustomizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
