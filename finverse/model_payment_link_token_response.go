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
)

// checks if the PaymentLinkTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinkTokenResponse{}

// PaymentLinkTokenResponse struct for PaymentLinkTokenResponse
type PaymentLinkTokenResponse struct {
	// Short-lived access-token to interact with Finverse Link
	AccessToken string `json:"access_token"`
	// Access token validity duration (in seconds)
	ExpiresIn int64 `json:"expires_in"`
	// URL to launch Finverse Link to authorize the mandate
	LinkUrl              string `json:"link_url"`
	TokenType            string `json:"token_type"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinkTokenResponse PaymentLinkTokenResponse

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinkTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["link_url"] = o.LinkUrl
	toSerialize["token_type"] = o.TokenType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinkTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"expires_in",
		"link_url",
		"token_type",
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

	varPaymentLinkTokenResponse := _PaymentLinkTokenResponse{}

	err = json.Unmarshal(data, &varPaymentLinkTokenResponse)

	if err != nil {
		return err
	}

	*o = PaymentLinkTokenResponse(varPaymentLinkTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "link_url")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
