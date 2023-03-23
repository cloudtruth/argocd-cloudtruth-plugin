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

// BackupParameterValue Parameter value data at a point in time.
type BackupParameterValue struct {
	External NullableBackupParameterValueExternal `json:"external"`
	Environment string `json:"environment"`
	Evaluated bool `json:"evaluated"`
	Source NullableString `json:"source,omitempty"`
	Project NullableString `json:"project,omitempty"`
	Value NullableString `json:"value,omitempty"`
	Raw NullableString `json:"raw,omitempty"`
}

// NewBackupParameterValue instantiates a new BackupParameterValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupParameterValue(external NullableBackupParameterValueExternal, environment string, evaluated bool) *BackupParameterValue {
	this := BackupParameterValue{}
	this.External = external
	this.Environment = environment
	this.Evaluated = evaluated
	return &this
}

// NewBackupParameterValueWithDefaults instantiates a new BackupParameterValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupParameterValueWithDefaults() *BackupParameterValue {
	this := BackupParameterValue{}
	return &this
}

// GetExternal returns the External field value
// If the value is explicit nil, the zero value for BackupParameterValueExternal will be returned
func (o *BackupParameterValue) GetExternal() BackupParameterValueExternal {
	if o == nil || o.External.Get() == nil {
		var ret BackupParameterValueExternal
		return ret
	}

	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupParameterValue) GetExternalOk() (*BackupParameterValueExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// SetExternal sets field value
func (o *BackupParameterValue) SetExternal(v BackupParameterValueExternal) {
	o.External.Set(&v)
}

// GetEnvironment returns the Environment field value
func (o *BackupParameterValue) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *BackupParameterValue) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *BackupParameterValue) SetEnvironment(v string) {
	o.Environment = v
}

// GetEvaluated returns the Evaluated field value
func (o *BackupParameterValue) GetEvaluated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Evaluated
}

// GetEvaluatedOk returns a tuple with the Evaluated field value
// and a boolean to check if the value has been set.
func (o *BackupParameterValue) GetEvaluatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Evaluated, true
}

// SetEvaluated sets field value
func (o *BackupParameterValue) SetEvaluated(v bool) {
	o.Evaluated = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupParameterValue) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupParameterValue) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *BackupParameterValue) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *BackupParameterValue) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *BackupParameterValue) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *BackupParameterValue) UnsetSource() {
	o.Source.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupParameterValue) GetProject() string {
	if o == nil || o.Project.Get() == nil {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupParameterValue) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *BackupParameterValue) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *BackupParameterValue) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *BackupParameterValue) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *BackupParameterValue) UnsetProject() {
	o.Project.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupParameterValue) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupParameterValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *BackupParameterValue) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *BackupParameterValue) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *BackupParameterValue) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *BackupParameterValue) UnsetValue() {
	o.Value.Unset()
}

// GetRaw returns the Raw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupParameterValue) GetRaw() string {
	if o == nil || o.Raw.Get() == nil {
		var ret string
		return ret
	}
	return *o.Raw.Get()
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupParameterValue) GetRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raw.Get(), o.Raw.IsSet()
}

// HasRaw returns a boolean if a field has been set.
func (o *BackupParameterValue) HasRaw() bool {
	if o != nil && o.Raw.IsSet() {
		return true
	}

	return false
}

// SetRaw gets a reference to the given NullableString and assigns it to the Raw field.
func (o *BackupParameterValue) SetRaw(v string) {
	o.Raw.Set(&v)
}
// SetRawNil sets the value for Raw to be an explicit nil
func (o *BackupParameterValue) SetRawNil() {
	o.Raw.Set(nil)
}

// UnsetRaw ensures that no value is present for Raw, not even an explicit nil
func (o *BackupParameterValue) UnsetRaw() {
	o.Raw.Unset()
}

func (o BackupParameterValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external"] = o.External.Get()
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["evaluated"] = o.Evaluated
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Raw.IsSet() {
		toSerialize["raw"] = o.Raw.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBackupParameterValue struct {
	value *BackupParameterValue
	isSet bool
}

func (v NullableBackupParameterValue) Get() *BackupParameterValue {
	return v.value
}

func (v *NullableBackupParameterValue) Set(val *BackupParameterValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupParameterValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupParameterValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupParameterValue(val *BackupParameterValue) *NullableBackupParameterValue {
	return &NullableBackupParameterValue{value: val, isSet: true}
}

func (v NullableBackupParameterValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupParameterValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

