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

// PortsSpecRangesInner struct for PortsSpecRangesInner
type PortsSpecRangesInner struct {
	From int32 `json:"from" bson:"from"`
	To int32 `json:"to" bson:"to"`
	AdditionalProperties map[string]interface{}
}

type _PortsSpecRangesInner PortsSpecRangesInner

// NewPortsSpecRangesInner instantiates a new PortsSpecRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortsSpecRangesInner(from int32, to int32) *PortsSpecRangesInner {
	this := PortsSpecRangesInner{}
	this.From = from
	this.To = to
	return &this
}

// NewPortsSpecRangesInnerWithDefaults instantiates a new PortsSpecRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortsSpecRangesInnerWithDefaults() *PortsSpecRangesInner {
	this := PortsSpecRangesInner{}
	return &this
}

// GetFrom returns the From field value
func (o *PortsSpecRangesInner) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PortsSpecRangesInner) GetFromOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PortsSpecRangesInner) SetFrom(v int32) {
	o.From = v
}

// GetTo returns the To field value
func (o *PortsSpecRangesInner) GetTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PortsSpecRangesInner) GetToOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PortsSpecRangesInner) SetTo(v int32) {
	o.To = v
}

func (o PortsSpecRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortsSpecRangesInner) UnmarshalJSON(bytes []byte) (err error) {
	varPortsSpecRangesInner := _PortsSpecRangesInner{}

	if err = json.Unmarshal(bytes, &varPortsSpecRangesInner); err == nil {
		*o = PortsSpecRangesInner(varPortsSpecRangesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortsSpecRangesInner struct {
	value *PortsSpecRangesInner
	isSet bool
}

func (v NullablePortsSpecRangesInner) Get() *PortsSpecRangesInner {
	return v.value
}

func (v *NullablePortsSpecRangesInner) Set(val *PortsSpecRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortsSpecRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortsSpecRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortsSpecRangesInner(val *PortsSpecRangesInner) *NullablePortsSpecRangesInner {
	return &NullablePortsSpecRangesInner{value: val, isSet: true}
}

func (v NullablePortsSpecRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortsSpecRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


