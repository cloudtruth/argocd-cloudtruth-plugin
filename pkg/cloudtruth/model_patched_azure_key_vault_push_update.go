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

// PatchedAzureKeyVaultPushUpdate Update a push.
type PatchedAzureKeyVaultPushUpdate struct {
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// Projects that are included in the push.
	Projects []string `json:"projects,omitempty"`
	// Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push.
	Tags []string `json:"tags,omitempty"`
	// Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - ``{{ environment }}`` to insert the environment name   - ``{{ parameter }}`` to insert the parameter name   - ``{{ project }}`` to insert the project name   - ``{{ push }}`` to insert the push name   - ``{{ tag }}`` to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the `project` substitution is required.  If you include multiple tags from different environments in the push, the `environment` substitution is required.  If you include multiple tags from the same environment in the push, the `tag` substitution is required.  In all cases, the `parameter` substitution is always required.
	Resource NullableString `json:"resource,omitempty"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// Normally, push will check to see if it originated the values in the destination before making changes to them.  Forcing a push disables the ownership check.
	Force *bool `json:"force,omitempty"`
	// Normally, push will process all parameters including those that flow in from project dependencies.  Declaring a push as `local` will cause it to only process the parameters defined in the selected projects.
	Local *bool `json:"local,omitempty"`
	// This setting allows parameters (non-secrets) to be pushed to a destination that only supports storing secrets.  This may increase your overall cost from the cloud provider as some cloud providers charge a premium for secrets-only storage.
	CoerceParameters *bool `json:"coerce_parameters,omitempty"`
	// Include parameters (non-secrets) in the values being pushed.  This setting requires the destination to support parameters or for the `coerce_parameters` flag to be enabled, otherwise the push will fail.
	IncludeParameters *bool `json:"include_parameters,omitempty"`
	// Include secrets in the values being pushed.  This setting requires the destination to support secrets, otherwise the push will fail.
	IncludeSecrets *bool `json:"include_secrets,omitempty"`
}

// NewPatchedAzureKeyVaultPushUpdate instantiates a new PatchedAzureKeyVaultPushUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAzureKeyVaultPushUpdate() *PatchedAzureKeyVaultPushUpdate {
	this := PatchedAzureKeyVaultPushUpdate{}
	return &this
}

// NewPatchedAzureKeyVaultPushUpdateWithDefaults instantiates a new PatchedAzureKeyVaultPushUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAzureKeyVaultPushUpdateWithDefaults() *PatchedAzureKeyVaultPushUpdate {
	this := PatchedAzureKeyVaultPushUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAzureKeyVaultPushUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAzureKeyVaultPushUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetProjects() []string {
	if o == nil || o.Projects == nil {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetProjectsOk() ([]string, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *PatchedAzureKeyVaultPushUpdate) SetProjects(v []string) {
	o.Projects = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchedAzureKeyVaultPushUpdate) SetTags(v []string) {
	o.Tags = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAzureKeyVaultPushUpdate) GetResource() string {
	if o == nil || o.Resource.Get() == nil {
		var ret string
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAzureKeyVaultPushUpdate) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableString and assigns it to the Resource field.
func (o *PatchedAzureKeyVaultPushUpdate) SetResource(v string) {
	o.Resource.Set(&v)
}
// SetResourceNil sets the value for Resource to be an explicit nil
func (o *PatchedAzureKeyVaultPushUpdate) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *PatchedAzureKeyVaultPushUpdate) UnsetResource() {
	o.Resource.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PatchedAzureKeyVaultPushUpdate) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *PatchedAzureKeyVaultPushUpdate) SetForce(v bool) {
	o.Force = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *PatchedAzureKeyVaultPushUpdate) SetLocal(v bool) {
	o.Local = &v
}

// GetCoerceParameters returns the CoerceParameters field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetCoerceParameters() bool {
	if o == nil || o.CoerceParameters == nil {
		var ret bool
		return ret
	}
	return *o.CoerceParameters
}

// GetCoerceParametersOk returns a tuple with the CoerceParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetCoerceParametersOk() (*bool, bool) {
	if o == nil || o.CoerceParameters == nil {
		return nil, false
	}
	return o.CoerceParameters, true
}

// HasCoerceParameters returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasCoerceParameters() bool {
	if o != nil && o.CoerceParameters != nil {
		return true
	}

	return false
}

// SetCoerceParameters gets a reference to the given bool and assigns it to the CoerceParameters field.
func (o *PatchedAzureKeyVaultPushUpdate) SetCoerceParameters(v bool) {
	o.CoerceParameters = &v
}

// GetIncludeParameters returns the IncludeParameters field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetIncludeParameters() bool {
	if o == nil || o.IncludeParameters == nil {
		var ret bool
		return ret
	}
	return *o.IncludeParameters
}

// GetIncludeParametersOk returns a tuple with the IncludeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetIncludeParametersOk() (*bool, bool) {
	if o == nil || o.IncludeParameters == nil {
		return nil, false
	}
	return o.IncludeParameters, true
}

// HasIncludeParameters returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasIncludeParameters() bool {
	if o != nil && o.IncludeParameters != nil {
		return true
	}

	return false
}

// SetIncludeParameters gets a reference to the given bool and assigns it to the IncludeParameters field.
func (o *PatchedAzureKeyVaultPushUpdate) SetIncludeParameters(v bool) {
	o.IncludeParameters = &v
}

// GetIncludeSecrets returns the IncludeSecrets field value if set, zero value otherwise.
func (o *PatchedAzureKeyVaultPushUpdate) GetIncludeSecrets() bool {
	if o == nil || o.IncludeSecrets == nil {
		var ret bool
		return ret
	}
	return *o.IncludeSecrets
}

// GetIncludeSecretsOk returns a tuple with the IncludeSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAzureKeyVaultPushUpdate) GetIncludeSecretsOk() (*bool, bool) {
	if o == nil || o.IncludeSecrets == nil {
		return nil, false
	}
	return o.IncludeSecrets, true
}

// HasIncludeSecrets returns a boolean if a field has been set.
func (o *PatchedAzureKeyVaultPushUpdate) HasIncludeSecrets() bool {
	if o != nil && o.IncludeSecrets != nil {
		return true
	}

	return false
}

// SetIncludeSecrets gets a reference to the given bool and assigns it to the IncludeSecrets field.
func (o *PatchedAzureKeyVaultPushUpdate) SetIncludeSecrets(v bool) {
	o.IncludeSecrets = &v
}

func (o PatchedAzureKeyVaultPushUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Resource.IsSet() {
		toSerialize["resource"] = o.Resource.Get()
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.CoerceParameters != nil {
		toSerialize["coerce_parameters"] = o.CoerceParameters
	}
	if o.IncludeParameters != nil {
		toSerialize["include_parameters"] = o.IncludeParameters
	}
	if o.IncludeSecrets != nil {
		toSerialize["include_secrets"] = o.IncludeSecrets
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAzureKeyVaultPushUpdate struct {
	value *PatchedAzureKeyVaultPushUpdate
	isSet bool
}

func (v NullablePatchedAzureKeyVaultPushUpdate) Get() *PatchedAzureKeyVaultPushUpdate {
	return v.value
}

func (v *NullablePatchedAzureKeyVaultPushUpdate) Set(val *PatchedAzureKeyVaultPushUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAzureKeyVaultPushUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAzureKeyVaultPushUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAzureKeyVaultPushUpdate(val *PatchedAzureKeyVaultPushUpdate) *NullablePatchedAzureKeyVaultPushUpdate {
	return &NullablePatchedAzureKeyVaultPushUpdate{value: val, isSet: true}
}

func (v NullablePatchedAzureKeyVaultPushUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAzureKeyVaultPushUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


