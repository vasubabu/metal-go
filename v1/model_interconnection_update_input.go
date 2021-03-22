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

// InterconnectionUpdateInput struct for InterconnectionUpdateInput
type InterconnectionUpdateInput struct {
	Name *string `json:"name,omitempty"`
	// Updating from 'redundant' to 'primary' will remove a secondary port, while updating from 'primary' to 'redundant' will add one.
	Redundancy *string `json:"redundancy,omitempty"`
	Description *string `json:"description,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
}

// NewInterconnectionUpdateInput instantiates a new InterconnectionUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterconnectionUpdateInput() *InterconnectionUpdateInput {
	this := InterconnectionUpdateInput{}
	return &this
}

// NewInterconnectionUpdateInputWithDefaults instantiates a new InterconnectionUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterconnectionUpdateInputWithDefaults() *InterconnectionUpdateInput {
	this := InterconnectionUpdateInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterconnectionUpdateInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterconnectionUpdateInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterconnectionUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *InterconnectionUpdateInput) GetRedundancy() string {
	if o == nil || o.Redundancy == nil {
		var ret string
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionUpdateInput) GetRedundancyOk() (*string, bool) {
	if o == nil || o.Redundancy == nil {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *InterconnectionUpdateInput) HasRedundancy() bool {
	if o != nil && o.Redundancy != nil {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given string and assigns it to the Redundancy field.
func (o *InterconnectionUpdateInput) SetRedundancy(v string) {
	o.Redundancy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterconnectionUpdateInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterconnectionUpdateInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterconnectionUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *InterconnectionUpdateInput) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionUpdateInput) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *InterconnectionUpdateInput) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *InterconnectionUpdateInput) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InterconnectionUpdateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionUpdateInput) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InterconnectionUpdateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InterconnectionUpdateInput) SetTags(v []string) {
	o.Tags = &v
}

func (o InterconnectionUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Redundancy != nil {
		toSerialize["redundancy"] = o.Redundancy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInterconnectionUpdateInput struct {
	value *InterconnectionUpdateInput
	isSet bool
}

func (v NullableInterconnectionUpdateInput) Get() *InterconnectionUpdateInput {
	return v.value
}

func (v *NullableInterconnectionUpdateInput) Set(val *InterconnectionUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnectionUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnectionUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnectionUpdateInput(val *InterconnectionUpdateInput) *NullableInterconnectionUpdateInput {
	return &NullableInterconnectionUpdateInput{value: val, isSet: true}
}

func (v NullableInterconnectionUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnectionUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


