/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the SelfServiceReservationItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceReservationItemResponse{}

// SelfServiceReservationItemResponse struct for SelfServiceReservationItemResponse
type SelfServiceReservationItemResponse struct {
	Amount               *float32 `json:"amount,omitempty"`
	Id                   *string  `json:"id,omitempty"`
	MetroCode            *string  `json:"metro_code,omitempty"`
	MetroId              *string  `json:"metro_id,omitempty"`
	MetroName            *string  `json:"metro_name,omitempty"`
	PlanId               *string  `json:"plan_id,omitempty"`
	PlanName             *string  `json:"plan_name,omitempty"`
	PlanSlug             *string  `json:"plan_slug,omitempty"`
	Quantity             *int32   `json:"quantity,omitempty"`
	Term                 *string  `json:"term,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceReservationItemResponse SelfServiceReservationItemResponse

// NewSelfServiceReservationItemResponse instantiates a new SelfServiceReservationItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationItemResponse() *SelfServiceReservationItemResponse {
	this := SelfServiceReservationItemResponse{}
	return &this
}

// NewSelfServiceReservationItemResponseWithDefaults instantiates a new SelfServiceReservationItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationItemResponseWithDefaults() *SelfServiceReservationItemResponse {
	this := SelfServiceReservationItemResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SelfServiceReservationItemResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SelfServiceReservationItemResponse) SetId(v string) {
	o.Id = &v
}

// GetMetroCode returns the MetroCode field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroCode() string {
	if o == nil || IsNil(o.MetroCode) {
		var ret string
		return ret
	}
	return *o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MetroCode) {
		return nil, false
	}
	return o.MetroCode, true
}

// HasMetroCode returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroCode() bool {
	if o != nil && !IsNil(o.MetroCode) {
		return true
	}

	return false
}

// SetMetroCode gets a reference to the given string and assigns it to the MetroCode field.
func (o *SelfServiceReservationItemResponse) SetMetroCode(v string) {
	o.MetroCode = &v
}

// GetMetroId returns the MetroId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroId() string {
	if o == nil || IsNil(o.MetroId) {
		var ret string
		return ret
	}
	return *o.MetroId
}

// GetMetroIdOk returns a tuple with the MetroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetroId) {
		return nil, false
	}
	return o.MetroId, true
}

// HasMetroId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroId() bool {
	if o != nil && !IsNil(o.MetroId) {
		return true
	}

	return false
}

// SetMetroId gets a reference to the given string and assigns it to the MetroId field.
func (o *SelfServiceReservationItemResponse) SetMetroId(v string) {
	o.MetroId = &v
}

// GetMetroName returns the MetroName field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroName() string {
	if o == nil || IsNil(o.MetroName) {
		var ret string
		return ret
	}
	return *o.MetroName
}

// GetMetroNameOk returns a tuple with the MetroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetroName) {
		return nil, false
	}
	return o.MetroName, true
}

// HasMetroName returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroName() bool {
	if o != nil && !IsNil(o.MetroName) {
		return true
	}

	return false
}

// SetMetroName gets a reference to the given string and assigns it to the MetroName field.
func (o *SelfServiceReservationItemResponse) SetMetroName(v string) {
	o.MetroName = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SelfServiceReservationItemResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *SelfServiceReservationItemResponse) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPlanSlug returns the PlanSlug field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanSlug() string {
	if o == nil || IsNil(o.PlanSlug) {
		var ret string
		return ret
	}
	return *o.PlanSlug
}

// GetPlanSlugOk returns a tuple with the PlanSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanSlugOk() (*string, bool) {
	if o == nil || IsNil(o.PlanSlug) {
		return nil, false
	}
	return o.PlanSlug, true
}

// HasPlanSlug returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanSlug() bool {
	if o != nil && !IsNil(o.PlanSlug) {
		return true
	}

	return false
}

// SetPlanSlug gets a reference to the given string and assigns it to the PlanSlug field.
func (o *SelfServiceReservationItemResponse) SetPlanSlug(v string) {
	o.PlanSlug = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SelfServiceReservationItemResponse) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SelfServiceReservationItemResponse) SetTerm(v string) {
	o.Term = &v
}

func (o SelfServiceReservationItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceReservationItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MetroCode) {
		toSerialize["metro_code"] = o.MetroCode
	}
	if !IsNil(o.MetroId) {
		toSerialize["metro_id"] = o.MetroId
	}
	if !IsNil(o.MetroName) {
		toSerialize["metro_name"] = o.MetroName
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.PlanName) {
		toSerialize["plan_name"] = o.PlanName
	}
	if !IsNil(o.PlanSlug) {
		toSerialize["plan_slug"] = o.PlanSlug
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceReservationItemResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSelfServiceReservationItemResponse := _SelfServiceReservationItemResponse{}

	if err = json.Unmarshal(bytes, &varSelfServiceReservationItemResponse); err == nil {
		*o = SelfServiceReservationItemResponse(varSelfServiceReservationItemResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metro_code")
		delete(additionalProperties, "metro_id")
		delete(additionalProperties, "metro_name")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "plan_name")
		delete(additionalProperties, "plan_slug")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "term")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceReservationItemResponse struct {
	value *SelfServiceReservationItemResponse
	isSet bool
}

func (v NullableSelfServiceReservationItemResponse) Get() *SelfServiceReservationItemResponse {
	return v.value
}

func (v *NullableSelfServiceReservationItemResponse) Set(val *SelfServiceReservationItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationItemResponse(val *SelfServiceReservationItemResponse) *NullableSelfServiceReservationItemResponse {
	return &NullableSelfServiceReservationItemResponse{value: val, isSet: true}
}

func (v NullableSelfServiceReservationItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
