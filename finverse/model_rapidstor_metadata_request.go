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

// checks if the RapidstorMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RapidstorMetadataRequest{}

// RapidstorMetadataRequest struct for RapidstorMetadataRequest
type RapidstorMetadataRequest struct {
	CorpCode             string   `json:"corp_code"`
	SLocationCode        string   `json:"s_location_code"`
	TenantId             string   `json:"tenant_id"`
	IAnnivDays           *float32 `json:"i_anniv_days,omitempty"`
	AccountToken         string   `json:"account_token"`
	AdditionalProperties map[string]interface{}
}

type _RapidstorMetadataRequest RapidstorMetadataRequest

// NewRapidstorMetadataRequest instantiates a new RapidstorMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRapidstorMetadataRequest(corpCode string, sLocationCode string, tenantId string, accountToken string) *RapidstorMetadataRequest {
	this := RapidstorMetadataRequest{}
	this.CorpCode = corpCode
	this.SLocationCode = sLocationCode
	this.TenantId = tenantId
	this.AccountToken = accountToken
	return &this
}

// NewRapidstorMetadataRequestWithDefaults instantiates a new RapidstorMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRapidstorMetadataRequestWithDefaults() *RapidstorMetadataRequest {
	this := RapidstorMetadataRequest{}
	return &this
}

// GetCorpCode returns the CorpCode field value
func (o *RapidstorMetadataRequest) GetCorpCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorpCode
}

// GetCorpCodeOk returns a tuple with the CorpCode field value
// and a boolean to check if the value has been set.
func (o *RapidstorMetadataRequest) GetCorpCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorpCode, true
}

// SetCorpCode sets field value
func (o *RapidstorMetadataRequest) SetCorpCode(v string) {
	o.CorpCode = v
}

// GetSLocationCode returns the SLocationCode field value
func (o *RapidstorMetadataRequest) GetSLocationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLocationCode
}

// GetSLocationCodeOk returns a tuple with the SLocationCode field value
// and a boolean to check if the value has been set.
func (o *RapidstorMetadataRequest) GetSLocationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLocationCode, true
}

// SetSLocationCode sets field value
func (o *RapidstorMetadataRequest) SetSLocationCode(v string) {
	o.SLocationCode = v
}

// GetTenantId returns the TenantId field value
func (o *RapidstorMetadataRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RapidstorMetadataRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RapidstorMetadataRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetIAnnivDays returns the IAnnivDays field value if set, zero value otherwise.
func (o *RapidstorMetadataRequest) GetIAnnivDays() float32 {
	if o == nil || IsNil(o.IAnnivDays) {
		var ret float32
		return ret
	}
	return *o.IAnnivDays
}

// GetIAnnivDaysOk returns a tuple with the IAnnivDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RapidstorMetadataRequest) GetIAnnivDaysOk() (*float32, bool) {
	if o == nil || IsNil(o.IAnnivDays) {
		return nil, false
	}
	return o.IAnnivDays, true
}

// HasIAnnivDays returns a boolean if a field has been set.
func (o *RapidstorMetadataRequest) HasIAnnivDays() bool {
	if o != nil && !IsNil(o.IAnnivDays) {
		return true
	}

	return false
}

// SetIAnnivDays gets a reference to the given float32 and assigns it to the IAnnivDays field.
func (o *RapidstorMetadataRequest) SetIAnnivDays(v float32) {
	o.IAnnivDays = &v
}

// GetAccountToken returns the AccountToken field value
func (o *RapidstorMetadataRequest) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *RapidstorMetadataRequest) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *RapidstorMetadataRequest) SetAccountToken(v string) {
	o.AccountToken = v
}

func (o RapidstorMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RapidstorMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["corp_code"] = o.CorpCode
	toSerialize["s_location_code"] = o.SLocationCode
	toSerialize["tenant_id"] = o.TenantId
	if !IsNil(o.IAnnivDays) {
		toSerialize["i_anniv_days"] = o.IAnnivDays
	}
	toSerialize["account_token"] = o.AccountToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RapidstorMetadataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"corp_code",
		"s_location_code",
		"tenant_id",
		"account_token",
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

	varRapidstorMetadataRequest := _RapidstorMetadataRequest{}

	err = json.Unmarshal(data, &varRapidstorMetadataRequest)

	if err != nil {
		return err
	}

	*o = RapidstorMetadataRequest(varRapidstorMetadataRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "corp_code")
		delete(additionalProperties, "s_location_code")
		delete(additionalProperties, "tenant_id")
		delete(additionalProperties, "i_anniv_days")
		delete(additionalProperties, "account_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRapidstorMetadataRequest struct {
	value *RapidstorMetadataRequest
	isSet bool
}

func (v NullableRapidstorMetadataRequest) Get() *RapidstorMetadataRequest {
	return v.value
}

func (v *NullableRapidstorMetadataRequest) Set(val *RapidstorMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRapidstorMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRapidstorMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRapidstorMetadataRequest(val *RapidstorMetadataRequest) *NullableRapidstorMetadataRequest {
	return &NullableRapidstorMetadataRequest{value: val, isSet: true}
}

func (v NullableRapidstorMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRapidstorMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
