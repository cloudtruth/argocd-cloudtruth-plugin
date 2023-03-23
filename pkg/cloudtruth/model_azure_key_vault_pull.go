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

// AzureKeyVaultPull Pull actions can be configured to get configuration and secrets from integrations on demand.
type AzureKeyVaultPull struct {
	Url string `json:"url"`
	// Unique identifier for the action.
	Id string `json:"id"`
	// The action name.
	Name string `json:"name"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	LatestTask NullableAzureKeyVaultPullLatestTask `json:"latest_task"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
	// Allow the pull to create environments.  Any automatically created environments will be children of the `default` environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateEnvironments *bool `json:"create_environments,omitempty"`
	// Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateProjects *bool `json:"create_projects,omitempty"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// Values being managed by a mapped pull.
	MappedValues []ValueCreate `json:"mapped_values"`
	// The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations.
	Mode ModeEnum `json:"mode"`
	// Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - ``{{ environment }}`` to identify the environment name   - ``{{ parameter }}`` to identify the parameter name   - ``{{ project }}`` to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the `environment`, `project`, and `parameter`.
	Resource NullableString `json:"resource"`
}

// NewAzureKeyVaultPull instantiates a new AzureKeyVaultPull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPull(url string, id string, name string, latestTask NullableAzureKeyVaultPullLatestTask, createdAt time.Time, modifiedAt NullableTime, mappedValues []ValueCreate, mode ModeEnum, resource NullableString) *AzureKeyVaultPull {
	this := AzureKeyVaultPull{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.LatestTask = latestTask
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.MappedValues = mappedValues
	this.Mode = mode
	this.Resource = resource
	return &this
}

// NewAzureKeyVaultPullWithDefaults instantiates a new AzureKeyVaultPull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPullWithDefaults() *AzureKeyVaultPull {
	this := AzureKeyVaultPull{}
	return &this
}

// GetUrl returns the Url field value
func (o *AzureKeyVaultPull) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AzureKeyVaultPull) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AzureKeyVaultPull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultPull) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AzureKeyVaultPull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureKeyVaultPull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultPull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultPull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultPull) SetDescription(v string) {
	o.Description = &v
}

// GetLatestTask returns the LatestTask field value
// If the value is explicit nil, the zero value for AzureKeyVaultPullLatestTask will be returned
func (o *AzureKeyVaultPull) GetLatestTask() AzureKeyVaultPullLatestTask {
	if o == nil || o.LatestTask.Get() == nil {
		var ret AzureKeyVaultPullLatestTask
		return ret
	}

	return *o.LatestTask.Get()
}

// GetLatestTaskOk returns a tuple with the LatestTask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPull) GetLatestTaskOk() (*AzureKeyVaultPullLatestTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestTask.Get(), o.LatestTask.IsSet()
}

// SetLatestTask sets field value
func (o *AzureKeyVaultPull) SetLatestTask(v AzureKeyVaultPullLatestTask) {
	o.LatestTask.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AzureKeyVaultPull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AzureKeyVaultPull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AzureKeyVaultPull) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPull) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AzureKeyVaultPull) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetCreateEnvironments returns the CreateEnvironments field value if set, zero value otherwise.
func (o *AzureKeyVaultPull) GetCreateEnvironments() bool {
	if o == nil || o.CreateEnvironments == nil {
		var ret bool
		return ret
	}
	return *o.CreateEnvironments
}

// GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetCreateEnvironmentsOk() (*bool, bool) {
	if o == nil || o.CreateEnvironments == nil {
		return nil, false
	}
	return o.CreateEnvironments, true
}

// HasCreateEnvironments returns a boolean if a field has been set.
func (o *AzureKeyVaultPull) HasCreateEnvironments() bool {
	if o != nil && o.CreateEnvironments != nil {
		return true
	}

	return false
}

// SetCreateEnvironments gets a reference to the given bool and assigns it to the CreateEnvironments field.
func (o *AzureKeyVaultPull) SetCreateEnvironments(v bool) {
	o.CreateEnvironments = &v
}

// GetCreateProjects returns the CreateProjects field value if set, zero value otherwise.
func (o *AzureKeyVaultPull) GetCreateProjects() bool {
	if o == nil || o.CreateProjects == nil {
		var ret bool
		return ret
	}
	return *o.CreateProjects
}

// GetCreateProjectsOk returns a tuple with the CreateProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetCreateProjectsOk() (*bool, bool) {
	if o == nil || o.CreateProjects == nil {
		return nil, false
	}
	return o.CreateProjects, true
}

// HasCreateProjects returns a boolean if a field has been set.
func (o *AzureKeyVaultPull) HasCreateProjects() bool {
	if o != nil && o.CreateProjects != nil {
		return true
	}

	return false
}

// SetCreateProjects gets a reference to the given bool and assigns it to the CreateProjects field.
func (o *AzureKeyVaultPull) SetCreateProjects(v bool) {
	o.CreateProjects = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AzureKeyVaultPull) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AzureKeyVaultPull) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AzureKeyVaultPull) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMappedValues returns the MappedValues field value
func (o *AzureKeyVaultPull) GetMappedValues() []ValueCreate {
	if o == nil {
		var ret []ValueCreate
		return ret
	}

	return o.MappedValues
}

// GetMappedValuesOk returns a tuple with the MappedValues field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetMappedValuesOk() ([]ValueCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.MappedValues, true
}

// SetMappedValues sets field value
func (o *AzureKeyVaultPull) SetMappedValues(v []ValueCreate) {
	o.MappedValues = v
}

// GetMode returns the Mode field value
func (o *AzureKeyVaultPull) GetMode() ModeEnum {
	if o == nil {
		var ret ModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPull) GetModeOk() (*ModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *AzureKeyVaultPull) SetMode(v ModeEnum) {
	o.Mode = v
}

// GetResource returns the Resource field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureKeyVaultPull) GetResource() string {
	if o == nil || o.Resource.Get() == nil {
		var ret string
		return ret
	}

	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPull) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// SetResource sets field value
func (o *AzureKeyVaultPull) SetResource(v string) {
	o.Resource.Set(&v)
}

func (o AzureKeyVaultPull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["latest_task"] = o.LatestTask.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if o.CreateEnvironments != nil {
		toSerialize["create_environments"] = o.CreateEnvironments
	}
	if o.CreateProjects != nil {
		toSerialize["create_projects"] = o.CreateProjects
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if true {
		toSerialize["mapped_values"] = o.MappedValues
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["resource"] = o.Resource.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureKeyVaultPull struct {
	value *AzureKeyVaultPull
	isSet bool
}

func (v NullableAzureKeyVaultPull) Get() *AzureKeyVaultPull {
	return v.value
}

func (v *NullableAzureKeyVaultPull) Set(val *AzureKeyVaultPull) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPull) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPull(val *AzureKeyVaultPull) *NullableAzureKeyVaultPull {
	return &NullableAzureKeyVaultPull{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

