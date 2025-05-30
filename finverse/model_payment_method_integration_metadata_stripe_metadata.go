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

// checks if the PaymentMethodIntegrationMetadataStripeMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodIntegrationMetadataStripeMetadata{}

// PaymentMethodIntegrationMetadataStripeMetadata struct for PaymentMethodIntegrationMetadataStripeMetadata
type PaymentMethodIntegrationMetadataStripeMetadata struct {
	Customer             PaymentMethodIntegrationMetadataStripeMetadataCustomer `json:"customer"`
	PaymentMethod        PaymentMethodIntegrationMetadataStripeMetadataCustomer `json:"payment_method"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodIntegrationMetadataStripeMetadata PaymentMethodIntegrationMetadataStripeMetadata

// NewPaymentMethodIntegrationMetadataStripeMetadata instantiates a new PaymentMethodIntegrationMetadataStripeMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodIntegrationMetadataStripeMetadata(customer PaymentMethodIntegrationMetadataStripeMetadataCustomer, paymentMethod PaymentMethodIntegrationMetadataStripeMetadataCustomer) *PaymentMethodIntegrationMetadataStripeMetadata {
	this := PaymentMethodIntegrationMetadataStripeMetadata{}
	this.Customer = customer
	this.PaymentMethod = paymentMethod
	return &this
}

// NewPaymentMethodIntegrationMetadataStripeMetadataWithDefaults instantiates a new PaymentMethodIntegrationMetadataStripeMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodIntegrationMetadataStripeMetadataWithDefaults() *PaymentMethodIntegrationMetadataStripeMetadata {
	this := PaymentMethodIntegrationMetadataStripeMetadata{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *PaymentMethodIntegrationMetadataStripeMetadata) GetCustomer() PaymentMethodIntegrationMetadataStripeMetadataCustomer {
	if o == nil {
		var ret PaymentMethodIntegrationMetadataStripeMetadataCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodIntegrationMetadataStripeMetadata) GetCustomerOk() (*PaymentMethodIntegrationMetadataStripeMetadataCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *PaymentMethodIntegrationMetadataStripeMetadata) SetCustomer(v PaymentMethodIntegrationMetadataStripeMetadataCustomer) {
	o.Customer = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PaymentMethodIntegrationMetadataStripeMetadata) GetPaymentMethod() PaymentMethodIntegrationMetadataStripeMetadataCustomer {
	if o == nil {
		var ret PaymentMethodIntegrationMetadataStripeMetadataCustomer
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodIntegrationMetadataStripeMetadata) GetPaymentMethodOk() (*PaymentMethodIntegrationMetadataStripeMetadataCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PaymentMethodIntegrationMetadataStripeMetadata) SetPaymentMethod(v PaymentMethodIntegrationMetadataStripeMetadataCustomer) {
	o.PaymentMethod = v
}

func (o PaymentMethodIntegrationMetadataStripeMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodIntegrationMetadataStripeMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer"] = o.Customer
	toSerialize["payment_method"] = o.PaymentMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodIntegrationMetadataStripeMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer",
		"payment_method",
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

	varPaymentMethodIntegrationMetadataStripeMetadata := _PaymentMethodIntegrationMetadataStripeMetadata{}

	err = json.Unmarshal(data, &varPaymentMethodIntegrationMetadataStripeMetadata)

	if err != nil {
		return err
	}

	*o = PaymentMethodIntegrationMetadataStripeMetadata(varPaymentMethodIntegrationMetadataStripeMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		delete(additionalProperties, "payment_method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodIntegrationMetadataStripeMetadata struct {
	value *PaymentMethodIntegrationMetadataStripeMetadata
	isSet bool
}

func (v NullablePaymentMethodIntegrationMetadataStripeMetadata) Get() *PaymentMethodIntegrationMetadataStripeMetadata {
	return v.value
}

func (v *NullablePaymentMethodIntegrationMetadataStripeMetadata) Set(val *PaymentMethodIntegrationMetadataStripeMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodIntegrationMetadataStripeMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodIntegrationMetadataStripeMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodIntegrationMetadataStripeMetadata(val *PaymentMethodIntegrationMetadataStripeMetadata) *NullablePaymentMethodIntegrationMetadataStripeMetadata {
	return &NullablePaymentMethodIntegrationMetadataStripeMetadata{value: val, isSet: true}
}

func (v NullablePaymentMethodIntegrationMetadataStripeMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodIntegrationMetadataStripeMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
