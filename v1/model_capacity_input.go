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

// CapacityInput struct for CapacityInput
type CapacityInput struct {
	Servers *[]ServerInfo `json:"servers,omitempty"`
}

// NewCapacityInput instantiates a new CapacityInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityInput() *CapacityInput {
	this := CapacityInput{}
	return &this
}

// NewCapacityInputWithDefaults instantiates a new CapacityInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityInputWithDefaults() *CapacityInput {
	this := CapacityInput{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *CapacityInput) GetServers() []ServerInfo {
	if o == nil || o.Servers == nil {
		var ret []ServerInfo
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityInput) GetServersOk() (*[]ServerInfo, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *CapacityInput) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerInfo and assigns it to the Servers field.
func (o *CapacityInput) SetServers(v []ServerInfo) {
	o.Servers = &v
}

func (o CapacityInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityInput struct {
	value *CapacityInput
	isSet bool
}

func (v NullableCapacityInput) Get() *CapacityInput {
	return v.value
}

func (v *NullableCapacityInput) Set(val *CapacityInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityInput(val *CapacityInput) *NullableCapacityInput {
	return &NullableCapacityInput{value: val, isSet: true}
}

func (v NullableCapacityInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


