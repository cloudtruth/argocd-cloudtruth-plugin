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

// AwsPushLatestTask The most recent task run for this action.
type AwsPushLatestTask struct {
	Url string `json:"url"`
	// The unique identifier for the task.
	Id string `json:"id"`
	// The reason why this task was triggered.
	Reason NullableString `json:"reason,omitempty"`
	// Indicates task steps were only simulated, not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// The current state of this task.
	State *StateEnum `json:"state,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this code may be helpful in determining the problem.
	ErrorCode NullableString `json:"error_code,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this detail may be helpful in determining the problem.
	ErrorDetail NullableString `json:"error_detail,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewAwsPushLatestTask instantiates a new AwsPushLatestTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPushLatestTask(url string, id string, createdAt time.Time, modifiedAt NullableTime) *AwsPushLatestTask {
	this := AwsPushLatestTask{}
	this.Url = url
	this.Id = id
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAwsPushLatestTaskWithDefaults instantiates a new AwsPushLatestTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPushLatestTaskWithDefaults() *AwsPushLatestTask {
	this := AwsPushLatestTask{}
	return &this
}

// GetUrl returns the Url field value
func (o *AwsPushLatestTask) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AwsPushLatestTask) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AwsPushLatestTask) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AwsPushLatestTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsPushLatestTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsPushLatestTask) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushLatestTask) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushLatestTask) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *AwsPushLatestTask) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *AwsPushLatestTask) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *AwsPushLatestTask) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *AwsPushLatestTask) UnsetReason() {
	o.Reason.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AwsPushLatestTask) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPushLatestTask) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AwsPushLatestTask) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AwsPushLatestTask) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AwsPushLatestTask) GetState() StateEnum {
	if o == nil || o.State == nil {
		var ret StateEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPushLatestTask) GetStateOk() (*StateEnum, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AwsPushLatestTask) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given StateEnum and assigns it to the State field.
func (o *AwsPushLatestTask) SetState(v StateEnum) {
	o.State = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushLatestTask) GetErrorCode() string {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushLatestTask) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AwsPushLatestTask) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AwsPushLatestTask) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AwsPushLatestTask) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AwsPushLatestTask) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushLatestTask) GetErrorDetail() string {
	if o == nil || o.ErrorDetail.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushLatestTask) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *AwsPushLatestTask) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *AwsPushLatestTask) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}
// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *AwsPushLatestTask) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *AwsPushLatestTask) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AwsPushLatestTask) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPushLatestTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AwsPushLatestTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AwsPushLatestTask) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushLatestTask) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AwsPushLatestTask) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o AwsPushLatestTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorDetail.IsSet() {
		toSerialize["error_detail"] = o.ErrorDetail.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsPushLatestTask struct {
	value *AwsPushLatestTask
	isSet bool
}

func (v NullableAwsPushLatestTask) Get() *AwsPushLatestTask {
	return v.value
}

func (v *NullableAwsPushLatestTask) Set(val *AwsPushLatestTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPushLatestTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPushLatestTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPushLatestTask(val *AwsPushLatestTask) *NullableAwsPushLatestTask {
	return &NullableAwsPushLatestTask{value: val, isSet: true}
}

func (v NullableAwsPushLatestTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPushLatestTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

