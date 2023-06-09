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
	"time"
)

// checks if the VrfRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfRoute{}

// VrfRoute struct for VrfRoute
type VrfRoute struct {
	// The unique identifier for the newly-created resource
	Id *string `json:"id,omitempty"`
	// The status of the route. Potential values are \"pending\", \"active\", \"deleting\", and \"error\", representing various lifecycle states of the route and whether or not it has been successfully configured on the network
	Status *string `json:"status,omitempty"`
	// The IPv4 prefix for the route, in CIDR-style notation
	Prefix *string `json:"prefix,omitempty"`
	// The next-hop IPv4 address for the route
	NextHop *string `json:"next_hop,omitempty"`
	// VRF route type, like 'bgp', 'connected', and 'static'. Currently, only static routes are supported
	Type                 *string                 `json:"type,omitempty"`
	CreatedAt            *time.Time              `json:"created_at,omitempty"`
	UpdatedAt            *time.Time              `json:"updated_at,omitempty"`
	MetalGateway         *VrfRouteMetalGateway   `json:"metal_gateway,omitempty"`
	VirtualNetwork       *VrfRouteVirtualNetwork `json:"virtual_network,omitempty"`
	Vrf                  *VrfRouteVrf            `json:"vrf,omitempty"`
	Href                 *string                 `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfRoute VrfRoute

// NewVrfRoute instantiates a new VrfRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfRoute() *VrfRoute {
	this := VrfRoute{}
	return &this
}

// NewVrfRouteWithDefaults instantiates a new VrfRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfRouteWithDefaults() *VrfRoute {
	this := VrfRoute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfRoute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfRoute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfRoute) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VrfRoute) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VrfRoute) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VrfRoute) SetStatus(v string) {
	o.Status = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *VrfRoute) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *VrfRoute) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *VrfRoute) SetPrefix(v string) {
	o.Prefix = &v
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *VrfRoute) GetNextHop() string {
	if o == nil || IsNil(o.NextHop) {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetNextHopOk() (*string, bool) {
	if o == nil || IsNil(o.NextHop) {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *VrfRoute) HasNextHop() bool {
	if o != nil && !IsNil(o.NextHop) {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *VrfRoute) SetNextHop(v string) {
	o.NextHop = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VrfRoute) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VrfRoute) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VrfRoute) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfRoute) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfRoute) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfRoute) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VrfRoute) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VrfRoute) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VrfRoute) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *VrfRoute) GetMetalGateway() VrfRouteMetalGateway {
	if o == nil || IsNil(o.MetalGateway) {
		var ret VrfRouteMetalGateway
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetMetalGatewayOk() (*VrfRouteMetalGateway, bool) {
	if o == nil || IsNil(o.MetalGateway) {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *VrfRoute) HasMetalGateway() bool {
	if o != nil && !IsNil(o.MetalGateway) {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given VrfRouteMetalGateway and assigns it to the MetalGateway field.
func (o *VrfRoute) SetMetalGateway(v VrfRouteMetalGateway) {
	o.MetalGateway = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *VrfRoute) GetVirtualNetwork() VrfRouteVirtualNetwork {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret VrfRouteVirtualNetwork
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetVirtualNetworkOk() (*VrfRouteVirtualNetwork, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *VrfRoute) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given VrfRouteVirtualNetwork and assigns it to the VirtualNetwork field.
func (o *VrfRoute) SetVirtualNetwork(v VrfRouteVirtualNetwork) {
	o.VirtualNetwork = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *VrfRoute) GetVrf() VrfRouteVrf {
	if o == nil || IsNil(o.Vrf) {
		var ret VrfRouteVrf
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetVrfOk() (*VrfRouteVrf, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *VrfRoute) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfRouteVrf and assigns it to the Vrf field.
func (o *VrfRoute) SetVrf(v VrfRouteVrf) {
	o.Vrf = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VrfRoute) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRoute) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VrfRoute) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VrfRoute) SetHref(v string) {
	o.Href = &v
}

func (o VrfRoute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: status is readOnly
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.NextHop) {
		toSerialize["next_hop"] = o.NextHop
	}
	// skip: type is readOnly
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	if !IsNil(o.MetalGateway) {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if !IsNil(o.VirtualNetwork) {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}
	// skip: href is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfRoute) UnmarshalJSON(bytes []byte) (err error) {
	varVrfRoute := _VrfRoute{}

	if err = json.Unmarshal(bytes, &varVrfRoute); err == nil {
		*o = VrfRoute(varVrfRoute)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "next_hop")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "metal_gateway")
		delete(additionalProperties, "virtual_network")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfRoute struct {
	value *VrfRoute
	isSet bool
}

func (v NullableVrfRoute) Get() *VrfRoute {
	return v.value
}

func (v *NullableVrfRoute) Set(val *VrfRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfRoute(val *VrfRoute) *NullableVrfRoute {
	return &NullableVrfRoute{value: val, isSet: true}
}

func (v NullableVrfRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
