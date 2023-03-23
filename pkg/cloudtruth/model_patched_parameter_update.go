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

// PatchedParameterUpdate A single parameter inside of a project.
type PatchedParameterUpdate struct {
	Id *string `json:"id,omitempty"`
	// The parameter name.
	Name *string `json:"name,omitempty"`
	// A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Indicates if this content is secret or not.  External values are inspected on-demand to ensure they align with the parameter's secret setting and if they do not, those external values are not allowed to be used.
	Secret *bool `json:"secret,omitempty"`
	// The type of this Parameter.
	Type *string `json:"type,omitempty"`
	// The project url.
	Project *string `json:"project,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt NullableTime `json:"modified_at,omitempty"`
}

// NewPatchedParameterUpdate instantiates a new PatchedParameterUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedParameterUpdate() *PatchedParameterUpdate {
	this := PatchedParameterUpdate{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// NewPatchedParameterUpdateWithDefaults instantiates a new PatchedParameterUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedParameterUpdateWithDefaults() *PatchedParameterUpdate {
	this := PatchedParameterUpdate{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedParameterUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedParameterUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedParameterUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *PatchedParameterUpdate) SetSecret(v bool) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedParameterUpdate) SetType(v string) {
	o.Type = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *PatchedParameterUpdate) SetProject(v string) {
	o.Project = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedParameterUpdate) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedParameterUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedParameterUpdate) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedParameterUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedParameterUpdate) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *PatchedParameterUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}
// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *PatchedParameterUpdate) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *PatchedParameterUpdate) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

func (o PatchedParameterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedParameterUpdate struct {
	value *PatchedParameterUpdate
	isSet bool
}

func (v NullablePatchedParameterUpdate) Get() *PatchedParameterUpdate {
	return v.value
}

func (v *NullablePatchedParameterUpdate) Set(val *PatchedParameterUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedParameterUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedParameterUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedParameterUpdate(val *PatchedParameterUpdate) *NullablePatchedParameterUpdate {
	return &NullablePatchedParameterUpdate{value: val, isSet: true}
}

func (v NullablePatchedParameterUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedParameterUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


