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

// checks if the ProjectUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectUsage{}

// ProjectUsage struct for ProjectUsage
type ProjectUsage struct {
	Facility    *string `json:"facility,omitempty"`
	Name        *string `json:"name,omitempty"`
	Plan        *string `json:"plan,omitempty"`
	PlanVersion *string `json:"plan_version,omitempty"`
	Price       *string `json:"price,omitempty"`
	Quantity    *string `json:"quantity,omitempty"`
	Total       *string `json:"total,omitempty"`
	Type        *string `json:"type,omitempty"`
	Unit        *string `json:"unit,omitempty"`
}

// NewProjectUsage instantiates a new ProjectUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsage() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// NewProjectUsageWithDefaults instantiates a new ProjectUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsageWithDefaults() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *ProjectUsage) GetFacility() string {
	if o == nil || isNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetFacilityOk() (*string, bool) {
	if o == nil || isNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *ProjectUsage) HasFacility() bool {
	if o != nil && !isNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *ProjectUsage) SetFacility(v string) {
	o.Facility = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectUsage) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectUsage) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectUsage) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *ProjectUsage) GetPlan() string {
	if o == nil || isNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPlanOk() (*string, bool) {
	if o == nil || isNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *ProjectUsage) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *ProjectUsage) SetPlan(v string) {
	o.Plan = &v
}

// GetPlanVersion returns the PlanVersion field value if set, zero value otherwise.
func (o *ProjectUsage) GetPlanVersion() string {
	if o == nil || isNil(o.PlanVersion) {
		var ret string
		return ret
	}
	return *o.PlanVersion
}

// GetPlanVersionOk returns a tuple with the PlanVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPlanVersionOk() (*string, bool) {
	if o == nil || isNil(o.PlanVersion) {
		return nil, false
	}
	return o.PlanVersion, true
}

// HasPlanVersion returns a boolean if a field has been set.
func (o *ProjectUsage) HasPlanVersion() bool {
	if o != nil && !isNil(o.PlanVersion) {
		return true
	}

	return false
}

// SetPlanVersion gets a reference to the given string and assigns it to the PlanVersion field.
func (o *ProjectUsage) SetPlanVersion(v string) {
	o.PlanVersion = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProjectUsage) GetPrice() string {
	if o == nil || isNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPriceOk() (*string, bool) {
	if o == nil || isNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProjectUsage) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ProjectUsage) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProjectUsage) GetQuantity() string {
	if o == nil || isNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetQuantityOk() (*string, bool) {
	if o == nil || isNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProjectUsage) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *ProjectUsage) SetQuantity(v string) {
	o.Quantity = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectUsage) GetTotal() string {
	if o == nil || isNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetTotalOk() (*string, bool) {
	if o == nil || isNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *ProjectUsage) SetTotal(v string) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectUsage) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectUsage) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectUsage) SetType(v string) {
	o.Type = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProjectUsage) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProjectUsage) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ProjectUsage) SetUnit(v string) {
	o.Unit = &v
}

func (o ProjectUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.PlanVersion) {
		toSerialize["plan_version"] = o.PlanVersion
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableProjectUsage struct {
	value *ProjectUsage
	isSet bool
}

func (v NullableProjectUsage) Get() *ProjectUsage {
	return v.value
}

func (v *NullableProjectUsage) Set(val *ProjectUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsage(val *ProjectUsage) *NullableProjectUsage {
	return &NullableProjectUsage{value: val, isSet: true}
}

func (v NullableProjectUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}