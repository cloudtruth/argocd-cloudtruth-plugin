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

// EnvironmentUpdate struct for EnvironmentUpdate
type EnvironmentUpdate struct {
	Id string `json:"id"`
	// The environment name.
	Name string `json:"name"`
	// A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed.
	Parent NullableString `json:"parent,omitempty"`
	// This is the opposite of `parent`, see that field for more details.
	Children []string `json:"children"`
	// Indicates if access control is being enforced through grants.
	AccessControlled *bool `json:"access_controlled,omitempty"`
	// Your role in the environment, if the environment is access-controlled.
	Role NullableRoleEnum `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewEnvironmentUpdate instantiates a new EnvironmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentUpdate(id string, name string, children []string, role NullableRoleEnum, createdAt time.Time, modifiedAt NullableTime) *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	this.Id = id
	this.Name = name
	this.Children = children
	this.Role = role
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewEnvironmentUpdateWithDefaults instantiates a new EnvironmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentUpdateWithDefaults() *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentUpdate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EnvironmentUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentUpdate) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *EnvironmentUpdate) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *EnvironmentUpdate) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *EnvironmentUpdate) UnsetParent() {
	o.Parent.Unset()
}

// GetChildren returns the Children field value
func (o *EnvironmentUpdate) GetChildren() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetChildrenOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *EnvironmentUpdate) SetChildren(v []string) {
	o.Children = v
}

// GetAccessControlled returns the AccessControlled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetAccessControlled() bool {
	if o == nil || o.AccessControlled == nil {
		var ret bool
		return ret
	}
	return *o.AccessControlled
}

// GetAccessControlledOk returns a tuple with the AccessControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetAccessControlledOk() (*bool, bool) {
	if o == nil || o.AccessControlled == nil {
		return nil, false
	}
	return o.AccessControlled, true
}

// HasAccessControlled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasAccessControlled() bool {
	if o != nil && o.AccessControlled != nil {
		return true
	}

	return false
}

// SetAccessControlled gets a reference to the given bool and assigns it to the AccessControlled field.
func (o *EnvironmentUpdate) SetAccessControlled(v bool) {
	o.AccessControlled = &v
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for RoleEnum will be returned
func (o *EnvironmentUpdate) GetRole() RoleEnum {
	if o == nil || o.Role.Get() == nil {
		var ret RoleEnum
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentUpdate) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *EnvironmentUpdate) SetRole(v RoleEnum) {
	o.Role.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentUpdate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EnvironmentUpdate) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *EnvironmentUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o EnvironmentUpdate) MarshalJSON() ([]byte, error) {
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
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if true {
		toSerialize["children"] = o.Children
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

type NullableEnvironmentUpdate struct {
	value *EnvironmentUpdate
	isSet bool
}

func (v NullableEnvironmentUpdate) Get() *EnvironmentUpdate {
	return v.value
}

func (v *NullableEnvironmentUpdate) Set(val *EnvironmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentUpdate(val *EnvironmentUpdate) *NullableEnvironmentUpdate {
	return &NullableEnvironmentUpdate{value: val, isSet: true}
}

func (v NullableEnvironmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

