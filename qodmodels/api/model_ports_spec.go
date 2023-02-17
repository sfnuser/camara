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

// PortsSpec Ports may be specified as a list of ranges or single ports. 
type PortsSpec struct {
	Ranges []PortsSpecRangesInner `json:"ranges,omitempty" bson:"ranges,omitempty"`
	Ports []int32 `json:"ports,omitempty" bson:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortsSpec PortsSpec

// NewPortsSpec instantiates a new PortsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortsSpec() *PortsSpec {
	this := PortsSpec{}
	return &this
}

// NewPortsSpecWithDefaults instantiates a new PortsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortsSpecWithDefaults() *PortsSpec {
	this := PortsSpec{}
	return &this
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *PortsSpec) GetRanges() []PortsSpecRangesInner {
	if o == nil || isNil(o.Ranges) {
		var ret []PortsSpecRangesInner
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortsSpec) GetRangesOk() ([]PortsSpecRangesInner, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *PortsSpec) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []PortsSpecRangesInner and assigns it to the Ranges field.
func (o *PortsSpec) SetRanges(v []PortsSpecRangesInner) {
	o.Ranges = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *PortsSpec) GetPorts() []int32 {
	if o == nil || isNil(o.Ports) {
		var ret []int32
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortsSpec) GetPortsOk() ([]int32, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *PortsSpec) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int32 and assigns it to the Ports field.
func (o *PortsSpec) SetPorts(v []int32) {
	o.Ports = v
}

func (o PortsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortsSpec) UnmarshalJSON(bytes []byte) (err error) {
	varPortsSpec := _PortsSpec{}

	if err = json.Unmarshal(bytes, &varPortsSpec); err == nil {
		*o = PortsSpec(varPortsSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ranges")
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortsSpec struct {
	value *PortsSpec
	isSet bool
}

func (v NullablePortsSpec) Get() *PortsSpec {
	return v.value
}

func (v *NullablePortsSpec) Set(val *PortsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePortsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePortsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortsSpec(val *PortsSpec) *NullablePortsSpec {
	return &NullablePortsSpec{value: val, isSet: true}
}

func (v NullablePortsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


