# PatchedAwsPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the action. | [optional] [readonly] 
**Name** | Pointer to **string** | The action name. | [optional] 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | Pointer to [**NullableAwsPullTask**](AwsPullTask.md) | The most recent task run for this action. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**Region** | Pointer to [**AwsRegionEnum**](AwsRegionEnum.md) | The AWS region this pull uses.  This region must be enabled in the integration. | [optional] 
**Service** | Pointer to [**AwsServiceEnum**](AwsServiceEnum.md) | The AWS service this pull uses.  This service must be enabled in the integration. | [optional] 
**Resource** | Pointer to **string** | Defines a path through the integration to the location where values will be pulled.  The following mustache-style substitutions must be used in the resource locator string:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name | [optional] 

## Methods

### NewPatchedAwsPull

`func NewPatchedAwsPull() *PatchedAwsPull`

NewPatchedAwsPull instantiates a new PatchedAwsPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAwsPullWithDefaults

`func NewPatchedAwsPullWithDefaults() *PatchedAwsPull`

NewPatchedAwsPullWithDefaults instantiates a new PatchedAwsPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedAwsPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedAwsPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedAwsPull) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedAwsPull) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedAwsPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAwsPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAwsPull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAwsPull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedAwsPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAwsPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAwsPull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAwsPull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedAwsPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAwsPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAwsPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAwsPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *PatchedAwsPull) GetLatestTask() AwsPullTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *PatchedAwsPull) GetLatestTaskOk() (*AwsPullTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *PatchedAwsPull) SetLatestTask(v AwsPullTask)`

SetLatestTask sets LatestTask field to given value.

### HasLatestTask

`func (o *PatchedAwsPull) HasLatestTask() bool`

HasLatestTask returns a boolean if a field has been set.

### SetLatestTaskNil

`func (o *PatchedAwsPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *PatchedAwsPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *PatchedAwsPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAwsPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAwsPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAwsPull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedAwsPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedAwsPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedAwsPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedAwsPull) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetDryRun

`func (o *PatchedAwsPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedAwsPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedAwsPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedAwsPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetRegion

`func (o *PatchedAwsPull) GetRegion() AwsRegionEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchedAwsPull) GetRegionOk() (*AwsRegionEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchedAwsPull) SetRegion(v AwsRegionEnum)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatchedAwsPull) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *PatchedAwsPull) GetService() AwsServiceEnum`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchedAwsPull) GetServiceOk() (*AwsServiceEnum, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchedAwsPull) SetService(v AwsServiceEnum)`

SetService sets Service field to given value.

### HasService

`func (o *PatchedAwsPull) HasService() bool`

HasService returns a boolean if a field has been set.

### GetResource

`func (o *PatchedAwsPull) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PatchedAwsPull) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PatchedAwsPull) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PatchedAwsPull) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


