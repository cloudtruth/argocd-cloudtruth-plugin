/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruth

import (
	"encoding/json"
	"time"
)

// ProjectUpdate struct for ProjectUpdate
type ProjectUpdate struct {
	Id string `json:"id"`
	// The project name.
	Name string `json:"name"`
	// A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// A regular expression parameter names must match
	ParameterNamePattern *string `json:"parameter_name_pattern,omitempty"`
	// Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project.
	DependsOn NullableString `json:"depends_on,omitempty"`
	// Indicates if access control is being enforced through grants.
	AccessControlled *bool `json:"access_controlled,omitempty"`
	// Your role in the project, if the project is access-controlled.
	Role NullableRoleEnum `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewProjectUpdate instantiates a new ProjectUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUpdate(id string, name string, role NullableRoleEnum, createdAt time.Time, modifiedAt NullableTime) *ProjectUpdate {
	this := ProjectUpdate{}
	this.Id = id
	this.Name = name
	this.Role = role
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUpdateWithDefaults() *ProjectUpdate {
	this := ProjectUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectUpdate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetParameterNamePattern returns the ParameterNamePattern field value if set, zero value otherwise.
func (o *ProjectUpdate) GetParameterNamePattern() string {
	if o == nil || o.ParameterNamePattern == nil {
		var ret string
		return ret
	}
	return *o.ParameterNamePattern
}

// GetParameterNamePatternOk returns a tuple with the ParameterNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetParameterNamePatternOk() (*string, bool) {
	if o == nil || o.ParameterNamePattern == nil {
		return nil, false
	}
	return o.ParameterNamePattern, true
}

// HasParameterNamePattern returns a boolean if a field has been set.
func (o *ProjectUpdate) HasParameterNamePattern() bool {
	if o != nil && o.ParameterNamePattern != nil {
		return true
	}

	return false
}

// SetParameterNamePattern gets a reference to the given string and assigns it to the ParameterNamePattern field.
func (o *ProjectUpdate) SetParameterNamePattern(v string) {
	o.ParameterNamePattern = &v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectUpdate) GetDependsOn() string {
	if o == nil || o.DependsOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUpdate) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *ProjectUpdate) HasDependsOn() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *ProjectUpdate) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}
// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *ProjectUpdate) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *ProjectUpdate) UnsetDependsOn() {
	o.DependsOn.Unset()
}

// GetAccessControlled returns the AccessControlled field value if set, zero value otherwise.
func (o *ProjectUpdate) GetAccessControlled() bool {
	if o == nil || o.AccessControlled == nil {
		var ret bool
		return ret
	}
	return *o.AccessControlled
}

// GetAccessControlledOk returns a tuple with the AccessControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetAccessControlledOk() (*bool, bool) {
	if o == nil || o.AccessControlled == nil {
		return nil, false
	}
	return o.AccessControlled, true
}

// HasAccessControlled returns a boolean if a field has been set.
func (o *ProjectUpdate) HasAccessControlled() bool {
	if o != nil && o.AccessControlled != nil {
		return true
	}

	return false
}

// SetAccessControlled gets a reference to the given bool and assigns it to the AccessControlled field.
func (o *ProjectUpdate) SetAccessControlled(v bool) {
	o.AccessControlled = &v
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for RoleEnum will be returned
func (o *ProjectUpdate) GetRole() RoleEnum {
	if o == nil || o.Role.Get() == nil {
		var ret RoleEnum
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUpdate) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *ProjectUpdate) SetRole(v RoleEnum) {
	o.Role.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectUpdate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ProjectUpdate) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ProjectUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o ProjectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ParameterNamePattern != nil {
		toSerialize["parameter_name_pattern"] = o.ParameterNamePattern
	}
	if o.DependsOn.IsSet() {
		toSerialize["depends_on"] = o.DependsOn.Get()
	}
	if o.AccessControlled != nil {
		toSerialize["access_controlled"] = o.AccessControlled
	}
	if true {
		toSerialize["role"] = o.Role.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProjectUpdate struct {
	value *ProjectUpdate
	isSet bool
}

func (v NullableProjectUpdate) Get() *ProjectUpdate {
	return v.value
}

func (v *NullableProjectUpdate) Set(val *ProjectUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUpdate(val *ProjectUpdate) *NullableProjectUpdate {
	return &NullableProjectUpdate{value: val, isSet: true}
}

func (v NullableProjectUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


