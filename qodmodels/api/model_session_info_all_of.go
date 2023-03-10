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

// SessionInfoAllOf struct for SessionInfoAllOf
type SessionInfoAllOf struct {
	// Session ID in UUID format
	Id string `json:"id" bson:"id"`
	// Timestamp of session start in seconds since unix epoch
	StartedAt int64 `json:"startedAt" bson:"startedAt"`
	// Timestamp of session expiration if the session was not deleted, in seconds since unix epoch
	ExpiresAt int64 `json:"expiresAt" bson:"expiresAt"`
	Messages []Message `json:"messages,omitempty" bson:"messages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionInfoAllOf SessionInfoAllOf

// NewSessionInfoAllOf instantiates a new SessionInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoAllOf(id string, startedAt int64, expiresAt int64) *SessionInfoAllOf {
	this := SessionInfoAllOf{}
	this.Id = id
	this.StartedAt = startedAt
	this.ExpiresAt = expiresAt
	return &this
}

// NewSessionInfoAllOfWithDefaults instantiates a new SessionInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoAllOfWithDefaults() *SessionInfoAllOf {
	this := SessionInfoAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *SessionInfoAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SessionInfoAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SessionInfoAllOf) SetId(v string) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value
func (o *SessionInfoAllOf) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *SessionInfoAllOf) GetStartedAtOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *SessionInfoAllOf) SetStartedAt(v int64) {
	o.StartedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *SessionInfoAllOf) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *SessionInfoAllOf) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *SessionInfoAllOf) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SessionInfoAllOf) GetMessages() []Message {
	if o == nil || isNil(o.Messages) {
		var ret []Message
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoAllOf) GetMessagesOk() ([]Message, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SessionInfoAllOf) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []Message and assigns it to the Messages field.
func (o *SessionInfoAllOf) SetMessages(v []Message) {
	o.Messages = v
}

func (o SessionInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSessionInfoAllOf := _SessionInfoAllOf{}

	if err = json.Unmarshal(bytes, &varSessionInfoAllOf); err == nil {
		*o = SessionInfoAllOf(varSessionInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "messages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionInfoAllOf struct {
	value *SessionInfoAllOf
	isSet bool
}

func (v NullableSessionInfoAllOf) Get() *SessionInfoAllOf {
	return v.value
}

func (v *NullableSessionInfoAllOf) Set(val *SessionInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoAllOf(val *SessionInfoAllOf) *NullableSessionInfoAllOf {
	return &NullableSessionInfoAllOf{value: val, isSet: true}
}

func (v NullableSessionInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


