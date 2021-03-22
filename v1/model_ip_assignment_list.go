/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// IPAssignmentList struct for IPAssignmentList
type IPAssignmentList struct {
	IpAddresses *[]IPAssignment `json:"ip_addresses,omitempty"`
}

// NewIPAssignmentList instantiates a new IPAssignmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAssignmentList() *IPAssignmentList {
	this := IPAssignmentList{}
	return &this
}

// NewIPAssignmentListWithDefaults instantiates a new IPAssignmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAssignmentListWithDefaults() *IPAssignmentList {
	this := IPAssignmentList{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *IPAssignmentList) GetIpAddresses() []IPAssignment {
	if o == nil || o.IpAddresses == nil {
		var ret []IPAssignment
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAssignmentList) GetIpAddressesOk() (*[]IPAssignment, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *IPAssignmentList) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IPAssignment and assigns it to the IpAddresses field.
func (o *IPAssignmentList) SetIpAddresses(v []IPAssignment) {
	o.IpAddresses = &v
}

func (o IPAssignmentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableIPAssignmentList struct {
	value *IPAssignmentList
	isSet bool
}

func (v NullableIPAssignmentList) Get() *IPAssignmentList {
	return v.value
}

func (v *NullableIPAssignmentList) Set(val *IPAssignmentList) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAssignmentList) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAssignmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAssignmentList(val *IPAssignmentList) *NullableIPAssignmentList {
	return &NullableIPAssignmentList{value: val, isSet: true}
}

func (v NullableIPAssignmentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAssignmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


