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

// Message struct for Message
type Message struct {
	// Message severity
	Severity string `json:"severity" bson:"severity"`
	// Detailed message text
	Description string `json:"description" bson:"description"`
	AdditionalProperties map[string]interface{}
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(severity string, description string) *Message {
	this := Message{}
	this.Severity = severity
	this.Description = description
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetSeverity returns the Severity field value
func (o *Message) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Message) GetSeverityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Message) SetSeverity(v string) {
	o.Severity = v
}

// GetDescription returns the Description field value
func (o *Message) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Message) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Message) SetDescription(v string) {
	o.Description = v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Message) UnmarshalJSON(bytes []byte) (err error) {
	varMessage := _Message{}

	if err = json.Unmarshal(bytes, &varMessage); err == nil {
		*o = Message(varMessage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "severity")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

