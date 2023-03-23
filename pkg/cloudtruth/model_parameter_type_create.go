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
)

// ParameterTypeCreate struct for ParameterTypeCreate
type ParameterTypeCreate struct {
	// The parameter type name.
	Name string `json:"name"`
	// A description of the parameter type, provide documentation on how to use this type here.
	Description *string `json:"description,omitempty"`
	// The URL for this parameter type's parent
	Parent string `json:"parent"`
}

// NewParameterTypeCreate instantiates a new ParameterTypeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterTypeCreate(name string, parent string) *ParameterTypeCreate {
	this := ParameterTypeCreate{}
	this.Name = name
	this.Parent = parent
	return &this
}

// NewParameterTypeCreateWithDefaults instantiates a new ParameterTypeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTypeCreateWithDefaults() *ParameterTypeCreate {
	this := ParameterTypeCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ParameterTypeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterTypeCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterTypeCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTypeCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterTypeCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterTypeCreate) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value
func (o *ParameterTypeCreate) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeCreate) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *ParameterTypeCreate) SetParent(v string) {
	o.Parent = v
}

func (o ParameterTypeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["parent"] = o.Parent
	}
	return json.Marshal(toSerialize)
}

type NullableParameterTypeCreate struct {
	value *ParameterTypeCreate
	isSet bool
}

func (v NullableParameterTypeCreate) Get() *ParameterTypeCreate {
	return v.value
}

func (v *NullableParameterTypeCreate) Set(val *ParameterTypeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterTypeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterTypeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterTypeCreate(val *ParameterTypeCreate) *NullableParameterTypeCreate {
	return &NullableParameterTypeCreate{value: val, isSet: true}
}

func (v NullableParameterTypeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterTypeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

