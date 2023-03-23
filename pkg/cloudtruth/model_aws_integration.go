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

// AwsIntegration struct for AwsIntegration
type AwsIntegration struct {
	Url string `json:"url"`
	// The unique identifier for the integration.
	Id string `json:"id"`
	Name string `json:"name"`
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	// The status of the integration connection with the third-party provider as of the `status_last_checked_at` field.  The status is updated automatically by the server when the integration is modified.
	Status StatusEnum `json:"status"`
	// If an error occurs, more details will be available in this field.
	StatusDetail string `json:"status_detail"`
	// The last time the status was evaluated.
	StatusLastCheckedAt NullableTime `json:"status_last_checked_at"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
	Fqn string `json:"fqn"`
	// The type of integration.
	Type string `json:"type"`
	// Allow actions to write to the integration.
	Writable *bool `json:"writable,omitempty"`
	// The AWS Account ID.
	AwsAccountId string `json:"aws_account_id"`
	// The AWS regions to integrate with.
	AwsEnabledRegions []AwsRegionEnum `json:"aws_enabled_regions"`
	// The AWS services to integrate with.
	AwsEnabledServices []AwsServiceEnum `json:"aws_enabled_services"`
	// This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  CloudTruth will generate a random value for you to give to your AWS Administrator in order to create the necessary IAM role for proper access.
	AwsExternalId string `json:"aws_external_id"`
	// If present, this is the KMS Key Id that is used to push values.  This key must be accessible in the AWS account (it cannot be an ARN to a key in another AWS account). 
	AwsKmsKeyId NullableString `json:"aws_kms_key_id,omitempty"`
	// The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator.
	AwsRoleName string `json:"aws_role_name"`
}

// NewAwsIntegration instantiates a new AwsIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsIntegration(url string, id string, name string, status StatusEnum, statusDetail string, statusLastCheckedAt NullableTime, createdAt time.Time, modifiedAt NullableTime, fqn string, type_ string, awsAccountId string, awsEnabledRegions []AwsRegionEnum, awsEnabledServices []AwsServiceEnum, awsExternalId string, awsRoleName string) *AwsIntegration {
	this := AwsIntegration{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusDetail = statusDetail
	this.StatusLastCheckedAt = statusLastCheckedAt
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Fqn = fqn
	this.Type = type_
	this.AwsAccountId = awsAccountId
	this.AwsEnabledRegions = awsEnabledRegions
	this.AwsEnabledServices = awsEnabledServices
	this.AwsExternalId = awsExternalId
	this.AwsRoleName = awsRoleName
	return &this
}

// NewAwsIntegrationWithDefaults instantiates a new AwsIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsIntegrationWithDefaults() *AwsIntegration {
	this := AwsIntegration{}
	return &this
}

// GetUrl returns the Url field value
func (o *AwsIntegration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AwsIntegration) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AwsIntegration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsIntegration) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AwsIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsIntegration) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsIntegration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsIntegration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsIntegration) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *AwsIntegration) GetStatus() StatusEnum {
	if o == nil {
		var ret StatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AwsIntegration) SetStatus(v StatusEnum) {
	o.Status = v
}

// GetStatusDetail returns the StatusDetail field value
func (o *AwsIntegration) GetStatusDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetStatusDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDetail, true
}

// SetStatusDetail sets field value
func (o *AwsIntegration) SetStatusDetail(v string) {
	o.StatusDetail = v
}

// GetStatusLastCheckedAt returns the StatusLastCheckedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AwsIntegration) GetStatusLastCheckedAt() time.Time {
	if o == nil || o.StatusLastCheckedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StatusLastCheckedAt.Get()
}

// GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusLastCheckedAt.Get(), o.StatusLastCheckedAt.IsSet()
}

// SetStatusLastCheckedAt sets field value
func (o *AwsIntegration) SetStatusLastCheckedAt(v time.Time) {
	o.StatusLastCheckedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AwsIntegration) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AwsIntegration) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AwsIntegration) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegration) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AwsIntegration) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetFqn returns the Fqn field value
func (o *AwsIntegration) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value
func (o *AwsIntegration) SetFqn(v string) {
	o.Fqn = v
}

// GetType returns the Type field value
func (o *AwsIntegration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AwsIntegration) SetType(v string) {
	o.Type = v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *AwsIntegration) GetWritable() bool {
	if o == nil || o.Writable == nil {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetWritableOk() (*bool, bool) {
	if o == nil || o.Writable == nil {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *AwsIntegration) HasWritable() bool {
	if o != nil && o.Writable != nil {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *AwsIntegration) SetWritable(v bool) {
	o.Writable = &v
}

// GetAwsAccountId returns the AwsAccountId field value
func (o *AwsIntegration) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value
func (o *AwsIntegration) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsEnabledRegions returns the AwsEnabledRegions field value
func (o *AwsIntegration) GetAwsEnabledRegions() []AwsRegionEnum {
	if o == nil {
		var ret []AwsRegionEnum
		return ret
	}

	return o.AwsEnabledRegions
}

// GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetAwsEnabledRegionsOk() ([]AwsRegionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsEnabledRegions, true
}

// SetAwsEnabledRegions sets field value
func (o *AwsIntegration) SetAwsEnabledRegions(v []AwsRegionEnum) {
	o.AwsEnabledRegions = v
}

// GetAwsEnabledServices returns the AwsEnabledServices field value
func (o *AwsIntegration) GetAwsEnabledServices() []AwsServiceEnum {
	if o == nil {
		var ret []AwsServiceEnum
		return ret
	}

	return o.AwsEnabledServices
}

// GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetAwsEnabledServicesOk() ([]AwsServiceEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsEnabledServices, true
}

// SetAwsEnabledServices sets field value
func (o *AwsIntegration) SetAwsEnabledServices(v []AwsServiceEnum) {
	o.AwsEnabledServices = v
}

// GetAwsExternalId returns the AwsExternalId field value
func (o *AwsIntegration) GetAwsExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsExternalId
}

// GetAwsExternalIdOk returns a tuple with the AwsExternalId field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetAwsExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsExternalId, true
}

// SetAwsExternalId sets field value
func (o *AwsIntegration) SetAwsExternalId(v string) {
	o.AwsExternalId = v
}

// GetAwsKmsKeyId returns the AwsKmsKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsIntegration) GetAwsKmsKeyId() string {
	if o == nil || o.AwsKmsKeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AwsKmsKeyId.Get()
}

// GetAwsKmsKeyIdOk returns a tuple with the AwsKmsKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegration) GetAwsKmsKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsKmsKeyId.Get(), o.AwsKmsKeyId.IsSet()
}

// HasAwsKmsKeyId returns a boolean if a field has been set.
func (o *AwsIntegration) HasAwsKmsKeyId() bool {
	if o != nil && o.AwsKmsKeyId.IsSet() {
		return true
	}

	return false
}

// SetAwsKmsKeyId gets a reference to the given NullableString and assigns it to the AwsKmsKeyId field.
func (o *AwsIntegration) SetAwsKmsKeyId(v string) {
	o.AwsKmsKeyId.Set(&v)
}
// SetAwsKmsKeyIdNil sets the value for AwsKmsKeyId to be an explicit nil
func (o *AwsIntegration) SetAwsKmsKeyIdNil() {
	o.AwsKmsKeyId.Set(nil)
}

// UnsetAwsKmsKeyId ensures that no value is present for AwsKmsKeyId, not even an explicit nil
func (o *AwsIntegration) UnsetAwsKmsKeyId() {
	o.AwsKmsKeyId.Unset()
}

// GetAwsRoleName returns the AwsRoleName field value
func (o *AwsIntegration) GetAwsRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsRoleName
}

// GetAwsRoleNameOk returns a tuple with the AwsRoleName field value
// and a boolean to check if the value has been set.
func (o *AwsIntegration) GetAwsRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRoleName, true
}

// SetAwsRoleName sets field value
func (o *AwsIntegration) SetAwsRoleName(v string) {
	o.AwsRoleName = v
}

func (o AwsIntegration) MarshalJSON() ([]byte, error) {
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
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["status_detail"] = o.StatusDetail
	}
	if true {
		toSerialize["status_last_checked_at"] = o.StatusLastCheckedAt.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if true {
		toSerialize["fqn"] = o.Fqn
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Writable != nil {
		toSerialize["writable"] = o.Writable
	}
	if true {
		toSerialize["aws_account_id"] = o.AwsAccountId
	}
	if true {
		toSerialize["aws_enabled_regions"] = o.AwsEnabledRegions
	}
	if true {
		toSerialize["aws_enabled_services"] = o.AwsEnabledServices
	}
	if true {
		toSerialize["aws_external_id"] = o.AwsExternalId
	}
	if o.AwsKmsKeyId.IsSet() {
		toSerialize["aws_kms_key_id"] = o.AwsKmsKeyId.Get()
	}
	if true {
		toSerialize["aws_role_name"] = o.AwsRoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAwsIntegration struct {
	value *AwsIntegration
	isSet bool
}

func (v NullableAwsIntegration) Get() *AwsIntegration {
	return v.value
}

func (v *NullableAwsIntegration) Set(val *AwsIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsIntegration(val *AwsIntegration) *NullableAwsIntegration {
	return &NullableAwsIntegration{value: val, isSet: true}
}

func (v NullableAwsIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


