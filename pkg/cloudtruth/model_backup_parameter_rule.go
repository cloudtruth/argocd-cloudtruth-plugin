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

// BackupParameterRule Rule that is applied to a parameter or parameter-type at a point in time.
type BackupParameterRule struct {
	RuleType string `json:"rule_type"`
	Constraint string `json:"constraint"`
}

// NewBackupParameterRule instantiates a new BackupParameterRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupParameterRule(ruleType string, constraint string) *BackupParameterRule {
	this := BackupParameterRule{}
	this.RuleType = ruleType
	this.Constraint = constraint
	return &this
}

// NewBackupParameterRuleWithDefaults instantiates a new BackupParameterRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupParameterRuleWithDefaults() *BackupParameterRule {
	this := BackupParameterRule{}
	return &this
}

// GetRuleType returns the RuleType field value
func (o *BackupParameterRule) GetRuleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *BackupParameterRule) GetRuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *BackupParameterRule) SetRuleType(v string) {
	o.RuleType = v
}

// GetConstraint returns the Constraint field value
func (o *BackupParameterRule) GetConstraint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value
// and a boolean to check if the value has been set.
func (o *BackupParameterRule) GetConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraint, true
}

// SetConstraint sets field value
func (o *BackupParameterRule) SetConstraint(v string) {
	o.Constraint = v
}

func (o BackupParameterRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rule_type"] = o.RuleType
	}
	if true {
		toSerialize["constraint"] = o.Constraint
	}
	return json.Marshal(toSerialize)
}

type NullableBackupParameterRule struct {
	value *BackupParameterRule
	isSet bool
}

func (v NullableBackupParameterRule) Get() *BackupParameterRule {
	return v.value
}

func (v *NullableBackupParameterRule) Set(val *BackupParameterRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupParameterRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupParameterRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupParameterRule(val *BackupParameterRule) *NullableBackupParameterRule {
	return &NullableBackupParameterRule{value: val, isSet: true}
}

func (v NullableBackupParameterRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupParameterRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


