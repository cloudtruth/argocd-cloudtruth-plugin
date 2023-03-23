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

// ParameterDualityT1 struct for ParameterDualityT1
type ParameterDualityT1 struct {
	// The parameter url.
	Url string `json:"url"`
	Id string `json:"id"`
	LedgerId string `json:"ledger_id"`
	// The parameter name.
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Secret *bool `json:"secret,omitempty"`
	// The type of this Parameter.
	Type *string `json:"type,omitempty"`
	// Rules applied to this parameter.
	Rules []ParameterRule `json:"rules"`
	// The project url.
	Project string `json:"project"`
	// The project name that the parameter is within.
	ProjectName string `json:"project_name"`
	// Templates that reference this Parameter.  This field is not valid for history requests.
	ReferencingTemplates []string `json:"referencing_templates"`
	// Dynamic values that reference this Parameter.  This field is not valid for history requests.
	ReferencingValues []string `json:"referencing_values"`
	//              This dictionary has keys that correspond to environment urls, and values             that correspond to the effective value for this parameter in that environment.             Each parameter has an effective value in every environment based on             project dependencies and environment inheritance.              The effective value is found by looking (within the keyed environment) up             the project dependencies by parameter name.  If a value is not found, the             parent environment is consulted with the same logic to locate a value.  It             is possible for there to be a `null` value record for an environment, which             means there is no value set; it is also possible for there to be a value record             with a `value` of `null`, which means the value was explicitly set to `null`.              If the value's parameter does not match the enclosing parameter (holding the             values array) then that value is flowing in through project dependencies.             Clients must recognize this in case the user asks to modify the value; in this             case the client must POST a new Value to the current parameter to override the             value coming in from the project dependency.              If the Value.environment matches the key, then it is an explicit value set for             that environment.  If they differ, the value was obtained from a parent             environment (directly or indirectly).  If the value is None then no value has             ever been set in any environment for this parameter within all the project             dependencies.         
	Values map[string]ParameterValuesValue `json:"values"`
	//          Identical to values, except the dictionary is flattened to a list.         Note that in this case, the environment in the Value is the environment         asked for, not the environment where it was obtained.         
	ValuesFlat []Value `json:"values_flat"`
	// If this parameter's project depends on another project which provides a parameter of the same name, this parameter overrides the one provided by the dependee.  You can use this field to determine if there will be side-effects the user should know about when deleting a parameter.  Deleting a parameter that overrides another one due to an identical name will uncover the one from the dependee project.
	Overrides NullableString `json:"overrides"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewParameterDualityT1 instantiates a new ParameterDualityT1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterDualityT1(url string, id string, ledgerId string, name string, rules []ParameterRule, project string, projectName string, referencingTemplates []string, referencingValues []string, values map[string]ParameterValuesValue, valuesFlat []Value, overrides NullableString, createdAt time.Time, modifiedAt NullableTime) *ParameterDualityT1 {
	this := ParameterDualityT1{}
	this.Url = url
	this.Id = id
	this.LedgerId = ledgerId
	this.Name = name
	var type_ string = "string"
	this.Type = &type_
	this.Rules = rules
	this.Project = project
	this.ProjectName = projectName
	this.ReferencingTemplates = referencingTemplates
	this.ReferencingValues = referencingValues
	this.Values = values
	this.ValuesFlat = valuesFlat
	this.Overrides = overrides
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewParameterDualityT1WithDefaults instantiates a new ParameterDualityT1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterDualityT1WithDefaults() *ParameterDualityT1 {
	this := ParameterDualityT1{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetUrl returns the Url field value
func (o *ParameterDualityT1) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ParameterDualityT1) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ParameterDualityT1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterDualityT1) SetId(v string) {
	o.Id = v
}

// GetLedgerId returns the LedgerId field value
func (o *ParameterDualityT1) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *ParameterDualityT1) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetName returns the Name field value
func (o *ParameterDualityT1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterDualityT1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterDualityT1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterDualityT1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterDualityT1) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ParameterDualityT1) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ParameterDualityT1) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *ParameterDualityT1) SetSecret(v bool) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ParameterDualityT1) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ParameterDualityT1) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ParameterDualityT1) SetType(v string) {
	o.Type = &v
}

// GetRules returns the Rules field value
func (o *ParameterDualityT1) GetRules() []ParameterRule {
	if o == nil {
		var ret []ParameterRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetRulesOk() ([]ParameterRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ParameterDualityT1) SetRules(v []ParameterRule) {
	o.Rules = v
}

// GetProject returns the Project field value
func (o *ParameterDualityT1) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ParameterDualityT1) SetProject(v string) {
	o.Project = v
}

// GetProjectName returns the ProjectName field value
func (o *ParameterDualityT1) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ParameterDualityT1) SetProjectName(v string) {
	o.ProjectName = v
}

// GetReferencingTemplates returns the ReferencingTemplates field value
func (o *ParameterDualityT1) GetReferencingTemplates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferencingTemplates
}

// GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetReferencingTemplatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferencingTemplates, true
}

// SetReferencingTemplates sets field value
func (o *ParameterDualityT1) SetReferencingTemplates(v []string) {
	o.ReferencingTemplates = v
}

// GetReferencingValues returns the ReferencingValues field value
func (o *ParameterDualityT1) GetReferencingValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferencingValues
}

// GetReferencingValuesOk returns a tuple with the ReferencingValues field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetReferencingValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferencingValues, true
}

// SetReferencingValues sets field value
func (o *ParameterDualityT1) SetReferencingValues(v []string) {
	o.ReferencingValues = v
}

// GetValues returns the Values field value
func (o *ParameterDualityT1) GetValues() map[string]ParameterValuesValue {
	if o == nil {
		var ret map[string]ParameterValuesValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetValuesOk() (*map[string]ParameterValuesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *ParameterDualityT1) SetValues(v map[string]ParameterValuesValue) {
	o.Values = v
}

// GetValuesFlat returns the ValuesFlat field value
func (o *ParameterDualityT1) GetValuesFlat() []Value {
	if o == nil {
		var ret []Value
		return ret
	}

	return o.ValuesFlat
}

// GetValuesFlatOk returns a tuple with the ValuesFlat field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetValuesFlatOk() ([]Value, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValuesFlat, true
}

// SetValuesFlat sets field value
func (o *ParameterDualityT1) SetValuesFlat(v []Value) {
	o.ValuesFlat = v
}

// GetOverrides returns the Overrides field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ParameterDualityT1) GetOverrides() string {
	if o == nil || o.Overrides.Get() == nil {
		var ret string
		return ret
	}

	return *o.Overrides.Get()
}

// GetOverridesOk returns a tuple with the Overrides field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterDualityT1) GetOverridesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Overrides.Get(), o.Overrides.IsSet()
}

// SetOverrides sets field value
func (o *ParameterDualityT1) SetOverrides(v string) {
	o.Overrides.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ParameterDualityT1) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterDualityT1) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ParameterDualityT1) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ParameterDualityT1) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterDualityT1) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ParameterDualityT1) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o ParameterDualityT1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ledger_id"] = o.LedgerId
	}
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
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if true {
		toSerialize["project_name"] = o.ProjectName
	}
	if true {
		toSerialize["referencing_templates"] = o.ReferencingTemplates
	}
	if true {
		toSerialize["referencing_values"] = o.ReferencingValues
	}
	if true {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["values_flat"] = o.ValuesFlat
	}
	if true {
		toSerialize["overrides"] = o.Overrides.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableParameterDualityT1 struct {
	value *ParameterDualityT1
	isSet bool
}

func (v NullableParameterDualityT1) Get() *ParameterDualityT1 {
	return v.value
}

func (v *NullableParameterDualityT1) Set(val *ParameterDualityT1) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterDualityT1) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterDualityT1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterDualityT1(val *ParameterDualityT1) *NullableParameterDualityT1 {
	return &NullableParameterDualityT1{value: val, isSet: true}
}

func (v NullableParameterDualityT1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterDualityT1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


