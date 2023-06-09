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

// checks if the Port type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Port{}

// Port Port is a hardware port associated with a reserved or instantiated hardware device.
type Port struct {
	Bond *BondPortData `json:"bond,omitempty"`
	Data *PortData     `json:"data,omitempty"`
	// Indicates whether or not the bond can be broken on the port (when applicable).
	DisbondOperationSupported *bool   `json:"disbond_operation_supported,omitempty"`
	Href                      *string `json:"href,omitempty"`
	Id                        *string `json:"id,omitempty"`
	Name                      *string `json:"name,omitempty"`
	// Type is either \"NetworkBondPort\" for bond ports or \"NetworkPort\" for bondable ethernet ports
	Type *string `json:"type,omitempty"`
	// Composite network type of the bond
	NetworkType          *string         `json:"network_type,omitempty"`
	NativeVirtualNetwork *VirtualNetwork `json:"native_virtual_network,omitempty"`
	VirtualNetworks      []Href          `json:"virtual_networks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Port Port

// NewPort instantiates a new Port object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPort() *Port {
	this := Port{}
	return &this
}

// NewPortWithDefaults instantiates a new Port object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortWithDefaults() *Port {
	this := Port{}
	return &this
}

// GetBond returns the Bond field value if set, zero value otherwise.
func (o *Port) GetBond() BondPortData {
	if o == nil || IsNil(o.Bond) {
		var ret BondPortData
		return ret
	}
	return *o.Bond
}

// GetBondOk returns a tuple with the Bond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetBondOk() (*BondPortData, bool) {
	if o == nil || IsNil(o.Bond) {
		return nil, false
	}
	return o.Bond, true
}

// HasBond returns a boolean if a field has been set.
func (o *Port) HasBond() bool {
	if o != nil && !IsNil(o.Bond) {
		return true
	}

	return false
}

// SetBond gets a reference to the given BondPortData and assigns it to the Bond field.
func (o *Port) SetBond(v BondPortData) {
	o.Bond = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Port) GetData() PortData {
	if o == nil || IsNil(o.Data) {
		var ret PortData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetDataOk() (*PortData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Port) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PortData and assigns it to the Data field.
func (o *Port) SetData(v PortData) {
	o.Data = &v
}

// GetDisbondOperationSupported returns the DisbondOperationSupported field value if set, zero value otherwise.
func (o *Port) GetDisbondOperationSupported() bool {
	if o == nil || IsNil(o.DisbondOperationSupported) {
		var ret bool
		return ret
	}
	return *o.DisbondOperationSupported
}

// GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetDisbondOperationSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisbondOperationSupported) {
		return nil, false
	}
	return o.DisbondOperationSupported, true
}

// HasDisbondOperationSupported returns a boolean if a field has been set.
func (o *Port) HasDisbondOperationSupported() bool {
	if o != nil && !IsNil(o.DisbondOperationSupported) {
		return true
	}

	return false
}

// SetDisbondOperationSupported gets a reference to the given bool and assigns it to the DisbondOperationSupported field.
func (o *Port) SetDisbondOperationSupported(v bool) {
	o.DisbondOperationSupported = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Port) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Port) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Port) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Port) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Port) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Port) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Port) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Port) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Port) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Port) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Port) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Port) SetType(v string) {
	o.Type = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *Port) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *Port) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *Port) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetNativeVirtualNetwork returns the NativeVirtualNetwork field value if set, zero value otherwise.
func (o *Port) GetNativeVirtualNetwork() VirtualNetwork {
	if o == nil || IsNil(o.NativeVirtualNetwork) {
		var ret VirtualNetwork
		return ret
	}
	return *o.NativeVirtualNetwork
}

// GetNativeVirtualNetworkOk returns a tuple with the NativeVirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetNativeVirtualNetworkOk() (*VirtualNetwork, bool) {
	if o == nil || IsNil(o.NativeVirtualNetwork) {
		return nil, false
	}
	return o.NativeVirtualNetwork, true
}

// HasNativeVirtualNetwork returns a boolean if a field has been set.
func (o *Port) HasNativeVirtualNetwork() bool {
	if o != nil && !IsNil(o.NativeVirtualNetwork) {
		return true
	}

	return false
}

// SetNativeVirtualNetwork gets a reference to the given VirtualNetwork and assigns it to the NativeVirtualNetwork field.
func (o *Port) SetNativeVirtualNetwork(v VirtualNetwork) {
	o.NativeVirtualNetwork = &v
}

// GetVirtualNetworks returns the VirtualNetworks field value if set, zero value otherwise.
func (o *Port) GetVirtualNetworks() []Href {
	if o == nil || IsNil(o.VirtualNetworks) {
		var ret []Href
		return ret
	}
	return o.VirtualNetworks
}

// GetVirtualNetworksOk returns a tuple with the VirtualNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetVirtualNetworksOk() ([]Href, bool) {
	if o == nil || IsNil(o.VirtualNetworks) {
		return nil, false
	}
	return o.VirtualNetworks, true
}

// HasVirtualNetworks returns a boolean if a field has been set.
func (o *Port) HasVirtualNetworks() bool {
	if o != nil && !IsNil(o.VirtualNetworks) {
		return true
	}

	return false
}

// SetVirtualNetworks gets a reference to the given []Href and assigns it to the VirtualNetworks field.
func (o *Port) SetVirtualNetworks(v []Href) {
	o.VirtualNetworks = v
}

func (o Port) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Port) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bond) {
		toSerialize["bond"] = o.Bond
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.DisbondOperationSupported) {
		toSerialize["disbond_operation_supported"] = o.DisbondOperationSupported
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}
	if !IsNil(o.NativeVirtualNetwork) {
		toSerialize["native_virtual_network"] = o.NativeVirtualNetwork
	}
	if !IsNil(o.VirtualNetworks) {
		toSerialize["virtual_networks"] = o.VirtualNetworks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Port) UnmarshalJSON(bytes []byte) (err error) {
	varPort := _Port{}

	if err = json.Unmarshal(bytes, &varPort); err == nil {
		*o = Port(varPort)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bond")
		delete(additionalProperties, "data")
		delete(additionalProperties, "disbond_operation_supported")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "network_type")
		delete(additionalProperties, "native_virtual_network")
		delete(additionalProperties, "virtual_networks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePort struct {
	value *Port
	isSet bool
}

func (v NullablePort) Get() *Port {
	return v.value
}

func (v *NullablePort) Set(val *Port) {
	v.value = val
	v.isSet = true
}

func (v NullablePort) IsSet() bool {
	return v.isSet
}

func (v *NullablePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePort(val *Port) *NullablePort {
	return &NullablePort{value: val, isSet: true}
}

func (v NullablePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
