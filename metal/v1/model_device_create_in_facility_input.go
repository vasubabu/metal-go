/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:    ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// DeviceCreateInFacilityInput struct for DeviceCreateInFacilityInput
type DeviceCreateInFacilityInput struct {
	Facility DeviceCreateInFacilityInputAllOfFacility `json:"facility"`
	// When true, devices with a `custom_ipxe` OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot.
	AlwaysPxe *bool `json:"always_pxe,omitempty"`
	// The billing cycle of the device.
	BillingCycle *string `json:"billing_cycle,omitempty"`
	// Customdata is an arbitrary JSON value that can be accessed via the metadata service.
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	// Any description of the device or how it will be used. This may be used to inform other API consumers with project access.
	Description *string `json:"description,omitempty"`
	// The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are `required`:  ``` { \"features\": [\"tpm\"] } ```  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a `preferred` value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  ``` { \"features\": { \"tpm\": \"required\", \"raid\": \"preferred\" } } ```  The request will only fail if there are no available servers matching the required `tpm` criteria.
	Features []string `json:"features,omitempty"`
	// The Hardware Reservation UUID to provision. Alternatively, `next-available` can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.  See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details.
	HardwareReservationId *string `json:"hardware_reservation_id,omitempty"`
	// The hostname to use within the operating system. The same hostname may be used on multiple devices within a project.
	Hostname *string `json:"hostname,omitempty"`
	// The `ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying `ip_addresses`, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send `{ \"ip_addresses\": [{ \"address_family\": 4, \"public\": false }] }`.  It is possible to request a subnet size larger than a `/30` by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, `{ \"ip_addresses\": [..., {\"address_family\": 4, \"public\": true, \"ip_reservations\": [\"uuid1\", \"uuid2\"]}] }`  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled.
	IpAddresses []DeviceCreateInputIpAddressesInner `json:"ip_addresses,omitempty"`
	// When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.  See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details.
	IpxeScriptUrl *string `json:"ipxe_script_url,omitempty"`
	// Whether the device should be locked, preventing accidental deletion.
	Locked *bool `json:"locked,omitempty"`
	// Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified
	NoSshKeys *bool `json:"no_ssh_keys,omitempty"`
	// The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans.
	OperatingSystem string `json:"operating_system"`
	// The slug of the device plan to provision.
	Plan string `json:"plan"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device.
	PrivateIpv4SubnetSize *float32 `json:"private_ipv4_subnet_size,omitempty"`
	// A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added.
	ProjectSshKeys []string `json:"project_ssh_keys,omitempty"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request.
	PublicIpv4SubnetSize *float32 `json:"public_ipv4_subnet_size,omitempty"`
	// Create a spot instance. Spot instances are created with a maximum bid price. If the bid price is not met, the spot instance will be terminated as indicated by the `termination_time` field.
	SpotInstance *bool `json:"spot_instance,omitempty"`
	// The maximum amount to bid for a spot instance.
	SpotPriceMax *float32 `json:"spot_price_max,omitempty"`
	// A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   `project_ssh_keys` and `user_ssh_keys`.
	SshKeys         []SSHKeyInput `json:"ssh_keys,omitempty"`
	Tags            []string      `json:"tags,omitempty"`
	TerminationTime *time.Time    `json:"termination_time,omitempty"`
	// A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added.
	UserSshKeys []string `json:"user_ssh_keys,omitempty"`
	// The userdata presented in the metadata service for this device.  Userdata is fetched and interpreted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.  See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details.
	Userdata *string `json:"userdata,omitempty"`
}

// NewDeviceCreateInFacilityInput instantiates a new DeviceCreateInFacilityInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateInFacilityInput(facility DeviceCreateInFacilityInputAllOfFacility, operatingSystem string, plan string) *DeviceCreateInFacilityInput {
	this := DeviceCreateInFacilityInput{}
	this.Facility = facility
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var locked bool = false
	this.Locked = &locked
	var noSshKeys bool = false
	this.NoSshKeys = &noSshKeys
	this.OperatingSystem = operatingSystem
	this.Plan = plan
	var privateIpv4SubnetSize float32 = 28
	this.PrivateIpv4SubnetSize = &privateIpv4SubnetSize
	var publicIpv4SubnetSize float32 = 31
	this.PublicIpv4SubnetSize = &publicIpv4SubnetSize
	return &this
}

// NewDeviceCreateInFacilityInputWithDefaults instantiates a new DeviceCreateInFacilityInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateInFacilityInputWithDefaults() *DeviceCreateInFacilityInput {
	this := DeviceCreateInFacilityInput{}
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var locked bool = false
	this.Locked = &locked
	var noSshKeys bool = false
	this.NoSshKeys = &noSshKeys
	var privateIpv4SubnetSize float32 = 28
	this.PrivateIpv4SubnetSize = &privateIpv4SubnetSize
	var publicIpv4SubnetSize float32 = 31
	this.PublicIpv4SubnetSize = &publicIpv4SubnetSize
	return &this
}

// GetFacility returns the Facility field value
func (o *DeviceCreateInFacilityInput) GetFacility() DeviceCreateInFacilityInputAllOfFacility {
	if o == nil {
		var ret DeviceCreateInFacilityInputAllOfFacility
		return ret
	}

	return o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetFacilityOk() (*DeviceCreateInFacilityInputAllOfFacility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facility, true
}

// SetFacility sets field value
func (o *DeviceCreateInFacilityInput) SetFacility(v DeviceCreateInFacilityInputAllOfFacility) {
	o.Facility = v
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetAlwaysPxe() bool {
	if o == nil || isNil(o.AlwaysPxe) {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || isNil(o.AlwaysPxe) {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasAlwaysPxe() bool {
	if o != nil && !isNil(o.AlwaysPxe) {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *DeviceCreateInFacilityInput) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetBillingCycle() string {
	if o == nil || isNil(o.BillingCycle) {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetBillingCycleOk() (*string, bool) {
	if o == nil || isNil(o.BillingCycle) {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasBillingCycle() bool {
	if o != nil && !isNil(o.BillingCycle) {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *DeviceCreateInFacilityInput) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetCustomdata() map[string]interface{} {
	if o == nil || isNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasCustomdata() bool {
	if o != nil && !isNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *DeviceCreateInFacilityInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceCreateInFacilityInput) SetDescription(v string) {
	o.Description = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetFeatures() []string {
	if o == nil || isNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetFeaturesOk() ([]string, bool) {
	if o == nil || isNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *DeviceCreateInFacilityInput) SetFeatures(v []string) {
	o.Features = v
}

// GetHardwareReservationId returns the HardwareReservationId field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetHardwareReservationId() string {
	if o == nil || isNil(o.HardwareReservationId) {
		var ret string
		return ret
	}
	return *o.HardwareReservationId
}

// GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetHardwareReservationIdOk() (*string, bool) {
	if o == nil || isNil(o.HardwareReservationId) {
		return nil, false
	}
	return o.HardwareReservationId, true
}

// HasHardwareReservationId returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasHardwareReservationId() bool {
	if o != nil && !isNil(o.HardwareReservationId) {
		return true
	}

	return false
}

// SetHardwareReservationId gets a reference to the given string and assigns it to the HardwareReservationId field.
func (o *DeviceCreateInFacilityInput) SetHardwareReservationId(v string) {
	o.HardwareReservationId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeviceCreateInFacilityInput) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetIpAddresses() []DeviceCreateInputIpAddressesInner {
	if o == nil || isNil(o.IpAddresses) {
		var ret []DeviceCreateInputIpAddressesInner
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetIpAddressesOk() ([]DeviceCreateInputIpAddressesInner, bool) {
	if o == nil || isNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasIpAddresses() bool {
	if o != nil && !isNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []DeviceCreateInputIpAddressesInner and assigns it to the IpAddresses field.
func (o *DeviceCreateInFacilityInput) SetIpAddresses(v []DeviceCreateInputIpAddressesInner) {
	o.IpAddresses = v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetIpxeScriptUrl() string {
	if o == nil || isNil(o.IpxeScriptUrl) {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || isNil(o.IpxeScriptUrl) {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasIpxeScriptUrl() bool {
	if o != nil && !isNil(o.IpxeScriptUrl) {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *DeviceCreateInFacilityInput) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *DeviceCreateInFacilityInput) SetLocked(v bool) {
	o.Locked = &v
}

// GetNoSshKeys returns the NoSshKeys field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetNoSshKeys() bool {
	if o == nil || isNil(o.NoSshKeys) {
		var ret bool
		return ret
	}
	return *o.NoSshKeys
}

// GetNoSshKeysOk returns a tuple with the NoSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetNoSshKeysOk() (*bool, bool) {
	if o == nil || isNil(o.NoSshKeys) {
		return nil, false
	}
	return o.NoSshKeys, true
}

// HasNoSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasNoSshKeys() bool {
	if o != nil && !isNil(o.NoSshKeys) {
		return true
	}

	return false
}

// SetNoSshKeys gets a reference to the given bool and assigns it to the NoSshKeys field.
func (o *DeviceCreateInFacilityInput) SetNoSshKeys(v bool) {
	o.NoSshKeys = &v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *DeviceCreateInFacilityInput) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *DeviceCreateInFacilityInput) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetPlan returns the Plan field value
func (o *DeviceCreateInFacilityInput) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *DeviceCreateInFacilityInput) SetPlan(v string) {
	o.Plan = v
}

// GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetPrivateIpv4SubnetSize() float32 {
	if o == nil || isNil(o.PrivateIpv4SubnetSize) {
		var ret float32
		return ret
	}
	return *o.PrivateIpv4SubnetSize
}

// GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetPrivateIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || isNil(o.PrivateIpv4SubnetSize) {
		return nil, false
	}
	return o.PrivateIpv4SubnetSize, true
}

// HasPrivateIpv4SubnetSize returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasPrivateIpv4SubnetSize() bool {
	if o != nil && !isNil(o.PrivateIpv4SubnetSize) {
		return true
	}

	return false
}

// SetPrivateIpv4SubnetSize gets a reference to the given float32 and assigns it to the PrivateIpv4SubnetSize field.
func (o *DeviceCreateInFacilityInput) SetPrivateIpv4SubnetSize(v float32) {
	o.PrivateIpv4SubnetSize = &v
}

// GetProjectSshKeys returns the ProjectSshKeys field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetProjectSshKeys() []string {
	if o == nil || isNil(o.ProjectSshKeys) {
		var ret []string
		return ret
	}
	return o.ProjectSshKeys
}

// GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetProjectSshKeysOk() ([]string, bool) {
	if o == nil || isNil(o.ProjectSshKeys) {
		return nil, false
	}
	return o.ProjectSshKeys, true
}

// HasProjectSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasProjectSshKeys() bool {
	if o != nil && !isNil(o.ProjectSshKeys) {
		return true
	}

	return false
}

// SetProjectSshKeys gets a reference to the given []string and assigns it to the ProjectSshKeys field.
func (o *DeviceCreateInFacilityInput) SetProjectSshKeys(v []string) {
	o.ProjectSshKeys = v
}

// GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetPublicIpv4SubnetSize() float32 {
	if o == nil || isNil(o.PublicIpv4SubnetSize) {
		var ret float32
		return ret
	}
	return *o.PublicIpv4SubnetSize
}

// GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetPublicIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || isNil(o.PublicIpv4SubnetSize) {
		return nil, false
	}
	return o.PublicIpv4SubnetSize, true
}

// HasPublicIpv4SubnetSize returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasPublicIpv4SubnetSize() bool {
	if o != nil && !isNil(o.PublicIpv4SubnetSize) {
		return true
	}

	return false
}

// SetPublicIpv4SubnetSize gets a reference to the given float32 and assigns it to the PublicIpv4SubnetSize field.
func (o *DeviceCreateInFacilityInput) SetPublicIpv4SubnetSize(v float32) {
	o.PublicIpv4SubnetSize = &v
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetSpotInstance() bool {
	if o == nil || isNil(o.SpotInstance) {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || isNil(o.SpotInstance) {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasSpotInstance() bool {
	if o != nil && !isNil(o.SpotInstance) {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *DeviceCreateInFacilityInput) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetSpotPriceMax returns the SpotPriceMax field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetSpotPriceMax() float32 {
	if o == nil || isNil(o.SpotPriceMax) {
		var ret float32
		return ret
	}
	return *o.SpotPriceMax
}

// GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetSpotPriceMaxOk() (*float32, bool) {
	if o == nil || isNil(o.SpotPriceMax) {
		return nil, false
	}
	return o.SpotPriceMax, true
}

// HasSpotPriceMax returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasSpotPriceMax() bool {
	if o != nil && !isNil(o.SpotPriceMax) {
		return true
	}

	return false
}

// SetSpotPriceMax gets a reference to the given float32 and assigns it to the SpotPriceMax field.
func (o *DeviceCreateInFacilityInput) SetSpotPriceMax(v float32) {
	o.SpotPriceMax = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetSshKeys() []SSHKeyInput {
	if o == nil || isNil(o.SshKeys) {
		var ret []SSHKeyInput
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetSshKeysOk() ([]SSHKeyInput, bool) {
	if o == nil || isNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []SSHKeyInput and assigns it to the SshKeys field.
func (o *DeviceCreateInFacilityInput) SetSshKeys(v []SSHKeyInput) {
	o.SshKeys = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceCreateInFacilityInput) SetTags(v []string) {
	o.Tags = v
}

// GetTerminationTime returns the TerminationTime field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetTerminationTime() time.Time {
	if o == nil || isNil(o.TerminationTime) {
		var ret time.Time
		return ret
	}
	return *o.TerminationTime
}

// GetTerminationTimeOk returns a tuple with the TerminationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetTerminationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.TerminationTime) {
		return nil, false
	}
	return o.TerminationTime, true
}

// HasTerminationTime returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasTerminationTime() bool {
	if o != nil && !isNil(o.TerminationTime) {
		return true
	}

	return false
}

// SetTerminationTime gets a reference to the given time.Time and assigns it to the TerminationTime field.
func (o *DeviceCreateInFacilityInput) SetTerminationTime(v time.Time) {
	o.TerminationTime = &v
}

// GetUserSshKeys returns the UserSshKeys field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetUserSshKeys() []string {
	if o == nil || isNil(o.UserSshKeys) {
		var ret []string
		return ret
	}
	return o.UserSshKeys
}

// GetUserSshKeysOk returns a tuple with the UserSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetUserSshKeysOk() ([]string, bool) {
	if o == nil || isNil(o.UserSshKeys) {
		return nil, false
	}
	return o.UserSshKeys, true
}

// HasUserSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasUserSshKeys() bool {
	if o != nil && !isNil(o.UserSshKeys) {
		return true
	}

	return false
}

// SetUserSshKeys gets a reference to the given []string and assigns it to the UserSshKeys field.
func (o *DeviceCreateInFacilityInput) SetUserSshKeys(v []string) {
	o.UserSshKeys = v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *DeviceCreateInFacilityInput) GetUserdata() string {
	if o == nil || isNil(o.Userdata) {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInFacilityInput) GetUserdataOk() (*string, bool) {
	if o == nil || isNil(o.Userdata) {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *DeviceCreateInFacilityInput) HasUserdata() bool {
	if o != nil && !isNil(o.Userdata) {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *DeviceCreateInFacilityInput) SetUserdata(v string) {
	o.Userdata = &v
}

func (o DeviceCreateInFacilityInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["facility"] = o.Facility
	}
	if !isNil(o.AlwaysPxe) {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if !isNil(o.BillingCycle) {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if !isNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !isNil(o.HardwareReservationId) {
		toSerialize["hardware_reservation_id"] = o.HardwareReservationId
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.IpAddresses) {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if !isNil(o.IpxeScriptUrl) {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.NoSshKeys) {
		toSerialize["no_ssh_keys"] = o.NoSshKeys
	}
	if true {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.PrivateIpv4SubnetSize) {
		toSerialize["private_ipv4_subnet_size"] = o.PrivateIpv4SubnetSize
	}
	if !isNil(o.ProjectSshKeys) {
		toSerialize["project_ssh_keys"] = o.ProjectSshKeys
	}
	if !isNil(o.PublicIpv4SubnetSize) {
		toSerialize["public_ipv4_subnet_size"] = o.PublicIpv4SubnetSize
	}
	if !isNil(o.SpotInstance) {
		toSerialize["spot_instance"] = o.SpotInstance
	}
	if !isNil(o.SpotPriceMax) {
		toSerialize["spot_price_max"] = o.SpotPriceMax
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.TerminationTime) {
		toSerialize["termination_time"] = o.TerminationTime
	}
	if !isNil(o.UserSshKeys) {
		toSerialize["user_ssh_keys"] = o.UserSshKeys
	}
	if !isNil(o.Userdata) {
		toSerialize["userdata"] = o.Userdata
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCreateInFacilityInput struct {
	value *DeviceCreateInFacilityInput
	isSet bool
}

func (v NullableDeviceCreateInFacilityInput) Get() *DeviceCreateInFacilityInput {
	return v.value
}

func (v *NullableDeviceCreateInFacilityInput) Set(val *DeviceCreateInFacilityInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateInFacilityInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateInFacilityInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateInFacilityInput(val *DeviceCreateInFacilityInput) *NullableDeviceCreateInFacilityInput {
	return &NullableDeviceCreateInFacilityInput{value: val, isSet: true}
}

func (v NullableDeviceCreateInFacilityInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateInFacilityInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
