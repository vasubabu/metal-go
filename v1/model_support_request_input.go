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

// SupportRequestInput struct for SupportRequestInput
type SupportRequestInput struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	ProjectId *string `json:"project_id,omitempty"`
	DeviceId *string `json:"device_id,omitempty"`
}

// NewSupportRequestInput instantiates a new SupportRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportRequestInput(subject string, message string) *SupportRequestInput {
	this := SupportRequestInput{}
	this.Subject = subject
	this.Message = message
	return &this
}

// NewSupportRequestInputWithDefaults instantiates a new SupportRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportRequestInputWithDefaults() *SupportRequestInput {
	this := SupportRequestInput{}
	return &this
}

// GetSubject returns the Subject field value
func (o *SupportRequestInput) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *SupportRequestInput) SetSubject(v string) {
	o.Subject = v
}

// GetMessage returns the Message field value
func (o *SupportRequestInput) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SupportRequestInput) SetMessage(v string) {
	o.Message = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *SupportRequestInput) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *SupportRequestInput) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *SupportRequestInput) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SupportRequestInput) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SupportRequestInput) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SupportRequestInput) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o SupportRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableSupportRequestInput struct {
	value *SupportRequestInput
	isSet bool
}

func (v NullableSupportRequestInput) Get() *SupportRequestInput {
	return v.value
}

func (v *NullableSupportRequestInput) Set(val *SupportRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportRequestInput(val *SupportRequestInput) *NullableSupportRequestInput {
	return &NullableSupportRequestInput{value: val, isSet: true}
}

func (v NullableSupportRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


