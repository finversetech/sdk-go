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

// PaymentLinkTokenResponse struct for PaymentLinkTokenResponse
type PaymentLinkTokenResponse struct {
	// Short-lived access-token to interact with Finverse Link
	AccessToken string `json:"access_token"`
	// Access token validity duration (in seconds)
	ExpiresIn int64 `json:"expires_in"`
	// URL to launch Finverse Link to authorize the mandate
	LinkUrl string `json:"link_url"`
	// The FPS QR code in base64
	QrCode    *string `json:"qr_code,omitempty"`
	TokenType string  `json:"token_type"`
}

// NewPaymentLinkTokenResponse instantiates a new PaymentLinkTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinkTokenResponse(accessToken string, expiresIn int64, linkUrl string, tokenType string) *PaymentLinkTokenResponse {
	this := PaymentLinkTokenResponse{}
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.LinkUrl = linkUrl
	this.TokenType = tokenType
	return &this
}

// NewPaymentLinkTokenResponseWithDefaults instantiates a new PaymentLinkTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinkTokenResponseWithDefaults() *PaymentLinkTokenResponse {
	this := PaymentLinkTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *PaymentLinkTokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *PaymentLinkTokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *PaymentLinkTokenResponse) GetExpiresIn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkTokenResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *PaymentLinkTokenResponse) SetExpiresIn(v int64) {
	o.ExpiresIn = v
}

// GetLinkUrl returns the LinkUrl field value
func (o *PaymentLinkTokenResponse) GetLinkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkTokenResponse) GetLinkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkUrl, true
}

// SetLinkUrl sets field value
func (o *PaymentLinkTokenResponse) SetLinkUrl(v string) {
	o.LinkUrl = v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *PaymentLinkTokenResponse) GetQrCode() string {
	if o == nil || o.QrCode == nil {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkTokenResponse) GetQrCodeOk() (*string, bool) {
	if o == nil || o.QrCode == nil {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *PaymentLinkTokenResponse) HasQrCode() bool {
	if o != nil && o.QrCode != nil {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *PaymentLinkTokenResponse) SetQrCode(v string) {
	o.QrCode = &v
}

// GetTokenType returns the TokenType field value
func (o *PaymentLinkTokenResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *PaymentLinkTokenResponse) SetTokenType(v string) {
	o.TokenType = v
}

func (o PaymentLinkTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if true {
		toSerialize["link_url"] = o.LinkUrl
	}
	if o.QrCode != nil {
		toSerialize["qr_code"] = o.QrCode
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentLinkTokenResponse struct {
	value *PaymentLinkTokenResponse
	isSet bool
}

func (v NullablePaymentLinkTokenResponse) Get() *PaymentLinkTokenResponse {
	return v.value
}

func (v *NullablePaymentLinkTokenResponse) Set(val *PaymentLinkTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinkTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinkTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinkTokenResponse(val *PaymentLinkTokenResponse) *NullablePaymentLinkTokenResponse {
	return &NullablePaymentLinkTokenResponse{value: val, isSet: true}
}

func (v NullablePaymentLinkTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinkTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
