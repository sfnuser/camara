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

// Ipv6Addr IPv6 address, following IETF 5952 format, may be specified in form <address/mask> as:   - address - The /128 subnet is optional for single addresses:     - 2001:db8:85a3:8d3:1319:8a2e:370:7344     - 2001:db8:85a3:8d3:1319:8a2e:370:7344/128   - address/mask - an IP v6 number with a mask:     - 2001:db8:85a3:8d3::0/64     - 2001:db8:85a3:8d3::/64 
type Ipv6Addr struct {
}

// NewIpv6Addr instantiates a new Ipv6Addr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Addr() *Ipv6Addr {
	this := Ipv6Addr{}
	return &this
}

// NewIpv6AddrWithDefaults instantiates a new Ipv6Addr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6AddrWithDefaults() *Ipv6Addr {
	this := Ipv6Addr{}
	return &this
}

func (o Ipv6Addr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableIpv6Addr struct {
	value *Ipv6Addr
	isSet bool
}

func (v NullableIpv6Addr) Get() *Ipv6Addr {
	return v.value
}

func (v *NullableIpv6Addr) Set(val *Ipv6Addr) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Addr) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Addr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Addr(val *Ipv6Addr) *NullableIpv6Addr {
	return &NullableIpv6Addr{value: val, isSet: true}
}

func (v NullableIpv6Addr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Addr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


