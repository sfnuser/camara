/*
QoD for enhanced communication

Service Enabling Network Function API for QoS control

API version: 0.8.0
Contact: project-email@sample.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ErrorInfo struct for ErrorInfo
type ErrorInfo struct {
	// Code given to this error
	Code string `json:"code" bson:"code"`
	// Detailed error description
	Message string `json:"message" bson:"message"`
	AdditionalProperties map[string]interface{}
}

type _ErrorInfo ErrorInfo

// NewErrorInfo instantiates a new ErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorInfo(code string, message string) *ErrorInfo {
	this := ErrorInfo{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorInfoWithDefaults instantiates a new ErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorInfoWithDefaults() *ErrorInfo {
	this := ErrorInfo{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorInfo) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorInfo) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorInfo) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorInfo) SetMessage(v string) {
	o.Message = v
}

func (o ErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorInfo) UnmarshalJSON(bytes []byte) (err error) {
	varErrorInfo := _ErrorInfo{}

	if err = json.Unmarshal(bytes, &varErrorInfo); err == nil {
		*o = ErrorInfo(varErrorInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorInfo struct {
	value *ErrorInfo
	isSet bool
}

func (v NullableErrorInfo) Get() *ErrorInfo {
	return v.value
}

func (v *NullableErrorInfo) Set(val *ErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorInfo(val *ErrorInfo) *NullableErrorInfo {
	return &NullableErrorInfo{value: val, isSet: true}
}

func (v NullableErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

