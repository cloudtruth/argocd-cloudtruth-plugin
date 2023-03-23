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

// ParameterCreate A single parameter inside of a project.
type ParameterCreate struct {
	// The parameter name.
	Name string `json:"name"`
	// A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Indicates if this content is secret or not.  External values are inspected on-demand to ensure they align with the parameter's secret setting and if they do not, those external values are not allowed to be used.
	Secret *bool `json:"secret,omitempty"`
	// The type of this Parameter.
	Type *string `json:"type,omitempty"`
}

// NewParameterCreate instantiates a new ParameterCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterCreate(name string) *ParameterCreate {
	this := ParameterCreate{}
	this.Name = name
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// NewParameterCreateWithDefaults instantiates a new ParameterCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterCreateWithDefaults() *ParameterCreate {
	this := ParameterCreate{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetName returns the Name field value
func (o *ParameterCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterCreate) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ParameterCreate) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterCreate) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ParameterCreate) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *ParameterCreate) SetSecret(v bool) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ParameterCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ParameterCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ParameterCreate) SetType(v string) {
	o.Type = &v
}

func (o ParameterCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableParameterCreate struct {
	value *ParameterCreate
	isSet bool
}

func (v NullableParameterCreate) Get() *ParameterCreate {
	return v.value
}

func (v *NullableParameterCreate) Set(val *ParameterCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterCreate(val *ParameterCreate) *NullableParameterCreate {
	return &NullableParameterCreate{value: val, isSet: true}
}

func (v NullableParameterCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


