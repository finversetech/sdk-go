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

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	Err                  *string  `json:"err,omitempty"`
	HttpStatusCode       *float32 `json:"http_status_code,omitempty"`
	StatusText           *string  `json:"status_text,omitempty"`
	AppCode              *float32 `json:"app_code,omitempty"`
	ErrorCategory        *string  `json:"error_category,omitempty"`
	ErrorText            *string  `json:"error_text,omitempty"`
	RequestId            *string  `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetErr returns the Err field value if set, zero value otherwise.
func (o *ErrorResponse) GetErr() string {
	if o == nil || IsNil(o.Err) {
		var ret string
		return ret
	}
	return *o.Err
}

// GetErrOk returns a tuple with the Err field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrOk() (*string, bool) {
	if o == nil || IsNil(o.Err) {
		return nil, false
	}
	return o.Err, true
}

// HasErr returns a boolean if a field has been set.
func (o *ErrorResponse) HasErr() bool {
	if o != nil && !IsNil(o.Err) {
		return true
	}

	return false
}

// SetErr gets a reference to the given string and assigns it to the Err field.
func (o *ErrorResponse) SetErr(v string) {
	o.Err = &v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *ErrorResponse) GetHttpStatusCode() float32 {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret float32
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetHttpStatusCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *ErrorResponse) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given float32 and assigns it to the HttpStatusCode field.
func (o *ErrorResponse) SetHttpStatusCode(v float32) {
	o.HttpStatusCode = &v
}

// GetStatusText returns the StatusText field value if set, zero value otherwise.
func (o *ErrorResponse) GetStatusText() string {
	if o == nil || IsNil(o.StatusText) {
		var ret string
		return ret
	}
	return *o.StatusText
}

// GetStatusTextOk returns a tuple with the StatusText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetStatusTextOk() (*string, bool) {
	if o == nil || IsNil(o.StatusText) {
		return nil, false
	}
	return o.StatusText, true
}

// HasStatusText returns a boolean if a field has been set.
func (o *ErrorResponse) HasStatusText() bool {
	if o != nil && !IsNil(o.StatusText) {
		return true
	}

	return false
}

// SetStatusText gets a reference to the given string and assigns it to the StatusText field.
func (o *ErrorResponse) SetStatusText(v string) {
	o.StatusText = &v
}

// GetAppCode returns the AppCode field value if set, zero value otherwise.
func (o *ErrorResponse) GetAppCode() float32 {
	if o == nil || IsNil(o.AppCode) {
		var ret float32
		return ret
	}
	return *o.AppCode
}

// GetAppCodeOk returns a tuple with the AppCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetAppCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.AppCode) {
		return nil, false
	}
	return o.AppCode, true
}

// HasAppCode returns a boolean if a field has been set.
func (o *ErrorResponse) HasAppCode() bool {
	if o != nil && !IsNil(o.AppCode) {
		return true
	}

	return false
}

// SetAppCode gets a reference to the given float32 and assigns it to the AppCode field.
func (o *ErrorResponse) SetAppCode(v float32) {
	o.AppCode = &v
}

// GetErrorCategory returns the ErrorCategory field value if set, zero value otherwise.
func (o *ErrorResponse) GetErrorCategory() string {
	if o == nil || IsNil(o.ErrorCategory) {
		var ret string
		return ret
	}
	return *o.ErrorCategory
}

// GetErrorCategoryOk returns a tuple with the ErrorCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCategory) {
		return nil, false
	}
	return o.ErrorCategory, true
}

// HasErrorCategory returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrorCategory() bool {
	if o != nil && !IsNil(o.ErrorCategory) {
		return true
	}

	return false
}

// SetErrorCategory gets a reference to the given string and assigns it to the ErrorCategory field.
func (o *ErrorResponse) SetErrorCategory(v string) {
	o.ErrorCategory = &v
}

// GetErrorText returns the ErrorText field value if set, zero value otherwise.
func (o *ErrorResponse) GetErrorText() string {
	if o == nil || IsNil(o.ErrorText) {
		var ret string
		return ret
	}
	return *o.ErrorText
}

// GetErrorTextOk returns a tuple with the ErrorText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorTextOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorText) {
		return nil, false
	}
	return o.ErrorText, true
}

// HasErrorText returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrorText() bool {
	if o != nil && !IsNil(o.ErrorText) {
		return true
	}

	return false
}

// SetErrorText gets a reference to the given string and assigns it to the ErrorText field.
func (o *ErrorResponse) SetErrorText(v string) {
	o.ErrorText = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ErrorResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ErrorResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ErrorResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Err) {
		toSerialize["err"] = o.Err
	}
	if !IsNil(o.HttpStatusCode) {
		toSerialize["http_status_code"] = o.HttpStatusCode
	}
	if !IsNil(o.StatusText) {
		toSerialize["status_text"] = o.StatusText
	}
	if !IsNil(o.AppCode) {
		toSerialize["app_code"] = o.AppCode
	}
	if !IsNil(o.ErrorCategory) {
		toSerialize["error_category"] = o.ErrorCategory
	}
	if !IsNil(o.ErrorText) {
		toSerialize["error_text"] = o.ErrorText
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorResponse) UnmarshalJSON(data []byte) (err error) {
	varErrorResponse := _ErrorResponse{}

	err = json.Unmarshal(data, &varErrorResponse)

	if err != nil {
		return err
	}

	*o = ErrorResponse(varErrorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "err")
		delete(additionalProperties, "http_status_code")
		delete(additionalProperties, "status_text")
		delete(additionalProperties, "app_code")
		delete(additionalProperties, "error_category")
		delete(additionalProperties, "error_text")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
