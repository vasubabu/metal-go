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

// ProjectUpdateInput struct for ProjectUpdateInput
type ProjectUpdateInput struct {
	Name *string `json:"name,omitempty"`
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	BackendTransferEnabled *bool `json:"backend_transfer_enabled,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
}

// NewProjectUpdateInput instantiates a new ProjectUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUpdateInput() *ProjectUpdateInput {
	this := ProjectUpdateInput{}
	return &this
}

// NewProjectUpdateInputWithDefaults instantiates a new ProjectUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUpdateInputWithDefaults() *ProjectUpdateInput {
	this := ProjectUpdateInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectUpdateInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectUpdateInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *ProjectUpdateInput) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateInput) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *ProjectUpdateInput) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *ProjectUpdateInput) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetBackendTransferEnabled returns the BackendTransferEnabled field value if set, zero value otherwise.
func (o *ProjectUpdateInput) GetBackendTransferEnabled() bool {
	if o == nil || o.BackendTransferEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BackendTransferEnabled
}

// GetBackendTransferEnabledOk returns a tuple with the BackendTransferEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateInput) GetBackendTransferEnabledOk() (*bool, bool) {
	if o == nil || o.BackendTransferEnabled == nil {
		return nil, false
	}
	return o.BackendTransferEnabled, true
}

// HasBackendTransferEnabled returns a boolean if a field has been set.
func (o *ProjectUpdateInput) HasBackendTransferEnabled() bool {
	if o != nil && o.BackendTransferEnabled != nil {
		return true
	}

	return false
}

// SetBackendTransferEnabled gets a reference to the given bool and assigns it to the BackendTransferEnabled field.
func (o *ProjectUpdateInput) SetBackendTransferEnabled(v bool) {
	o.BackendTransferEnabled = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *ProjectUpdateInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *ProjectUpdateInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *ProjectUpdateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

func (o ProjectUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if o.BackendTransferEnabled != nil {
		toSerialize["backend_transfer_enabled"] = o.BackendTransferEnabled
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	return json.Marshal(toSerialize)
}

type NullableProjectUpdateInput struct {
	value *ProjectUpdateInput
	isSet bool
}

func (v NullableProjectUpdateInput) Get() *ProjectUpdateInput {
	return v.value
}

func (v *NullableProjectUpdateInput) Set(val *ProjectUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUpdateInput(val *ProjectUpdateInput) *NullableProjectUpdateInput {
	return &NullableProjectUpdateInput{value: val, isSet: true}
}

func (v NullableProjectUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


