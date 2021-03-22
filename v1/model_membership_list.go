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

// MembershipList struct for MembershipList
type MembershipList struct {
	Memberships *[]Membership `json:"memberships,omitempty"`
}

// NewMembershipList instantiates a new MembershipList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipList() *MembershipList {
	this := MembershipList{}
	return &this
}

// NewMembershipListWithDefaults instantiates a new MembershipList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipListWithDefaults() *MembershipList {
	this := MembershipList{}
	return &this
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *MembershipList) GetMemberships() []Membership {
	if o == nil || o.Memberships == nil {
		var ret []Membership
		return ret
	}
	return *o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipList) GetMembershipsOk() (*[]Membership, bool) {
	if o == nil || o.Memberships == nil {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *MembershipList) HasMemberships() bool {
	if o != nil && o.Memberships != nil {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []Membership and assigns it to the Memberships field.
func (o *MembershipList) SetMemberships(v []Membership) {
	o.Memberships = &v
}

func (o MembershipList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Memberships != nil {
		toSerialize["memberships"] = o.Memberships
	}
	return json.Marshal(toSerialize)
}

type NullableMembershipList struct {
	value *MembershipList
	isSet bool
}

func (v NullableMembershipList) Get() *MembershipList {
	return v.value
}

func (v *NullableMembershipList) Set(val *MembershipList) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipList) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipList(val *MembershipList) *NullableMembershipList {
	return &NullableMembershipList{value: val, isSet: true}
}

func (v NullableMembershipList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


