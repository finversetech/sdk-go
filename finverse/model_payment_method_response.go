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

// checks if the PaymentMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodResponse{}

// PaymentMethodResponse struct for PaymentMethodResponse
type PaymentMethodResponse struct {
	PaymentMethodId      *string                           `json:"payment_method_id,omitempty"`
	PaymentMethodType    *string                           `json:"payment_method_type,omitempty"`
	Mandate              *GetMandateResponse               `json:"mandate,omitempty"`
	Card                 *FVCard                           `json:"card,omitempty"`
	IntegrationMetadata  *PaymentMethodIntegrationMetadata `json:"integration_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodResponse PaymentMethodResponse

// NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodResponse() *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	return &this
}

// NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	return &this
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *PaymentMethodResponse) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetPaymentMethodType() string {
	if o == nil || IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasPaymentMethodType() bool {
	if o != nil && !IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *PaymentMethodResponse) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetMandate returns the Mandate field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetMandate() GetMandateResponse {
	if o == nil || IsNil(o.Mandate) {
		var ret GetMandateResponse
		return ret
	}
	return *o.Mandate
}

// GetMandateOk returns a tuple with the Mandate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetMandateOk() (*GetMandateResponse, bool) {
	if o == nil || IsNil(o.Mandate) {
		return nil, false
	}
	return o.Mandate, true
}

// HasMandate returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasMandate() bool {
	if o != nil && !IsNil(o.Mandate) {
		return true
	}

	return false
}

// SetMandate gets a reference to the given GetMandateResponse and assigns it to the Mandate field.
func (o *PaymentMethodResponse) SetMandate(v GetMandateResponse) {
	o.Mandate = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetCard() FVCard {
	if o == nil || IsNil(o.Card) {
		var ret FVCard
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetCardOk() (*FVCard, bool) {
	if o == nil || IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasCard() bool {
	if o != nil && !IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given FVCard and assigns it to the Card field.
func (o *PaymentMethodResponse) SetCard(v FVCard) {
	o.Card = &v
}

// GetIntegrationMetadata returns the IntegrationMetadata field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetIntegrationMetadata() PaymentMethodIntegrationMetadata {
	if o == nil || IsNil(o.IntegrationMetadata) {
		var ret PaymentMethodIntegrationMetadata
		return ret
	}
	return *o.IntegrationMetadata
}

// GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetIntegrationMetadataOk() (*PaymentMethodIntegrationMetadata, bool) {
	if o == nil || IsNil(o.IntegrationMetadata) {
		return nil, false
	}
	return o.IntegrationMetadata, true
}

// HasIntegrationMetadata returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasIntegrationMetadata() bool {
	if o != nil && !IsNil(o.IntegrationMetadata) {
		return true
	}

	return false
}

// SetIntegrationMetadata gets a reference to the given PaymentMethodIntegrationMetadata and assigns it to the IntegrationMetadata field.
func (o *PaymentMethodResponse) SetIntegrationMetadata(v PaymentMethodIntegrationMetadata) {
	o.IntegrationMetadata = &v
}

func (o PaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !IsNil(o.PaymentMethodType) {
		toSerialize["payment_method_type"] = o.PaymentMethodType
	}
	if !IsNil(o.Mandate) {
		toSerialize["mandate"] = o.Mandate
	}
	if !IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !IsNil(o.IntegrationMetadata) {
		toSerialize["integration_metadata"] = o.IntegrationMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodResponse) UnmarshalJSON(data []byte) (err error) {
	varPaymentMethodResponse := _PaymentMethodResponse{}

	err = json.Unmarshal(data, &varPaymentMethodResponse)

	if err != nil {
		return err
	}

	*o = PaymentMethodResponse(varPaymentMethodResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_method_id")
		delete(additionalProperties, "payment_method_type")
		delete(additionalProperties, "mandate")
		delete(additionalProperties, "card")
		delete(additionalProperties, "integration_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodResponse struct {
	value *PaymentMethodResponse
	isSet bool
}

func (v NullablePaymentMethodResponse) Get() *PaymentMethodResponse {
	return v.value
}

func (v *NullablePaymentMethodResponse) Set(val *PaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodResponse(val *PaymentMethodResponse) *NullablePaymentMethodResponse {
	return &NullablePaymentMethodResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
