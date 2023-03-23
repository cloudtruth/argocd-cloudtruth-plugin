# PatchedParameterTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The parameter type name. | [optional] 
**Description** | Pointer to **string** | A description of the parameter type, provide documentation on how to use this type here. | [optional] 
**Parent** | Pointer to **string** | The URL for this parameter type&#39;s parent | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterTypeUpdate

`func NewPatchedParameterTypeUpdate() *PatchedParameterTypeUpdate`

NewPatchedParameterTypeUpdate instantiates a new PatchedParameterTypeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterTypeUpdateWithDefaults

`func NewPatchedParameterTypeUpdateWithDefaults() *PatchedParameterTypeUpdate`

NewPatchedParameterTypeUpdateWithDefaults instantiates a new PatchedParameterTypeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedParameterTypeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterTypeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterTypeUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterTypeUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedParameterTypeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedParameterTypeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedParameterTypeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedParameterTypeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedParameterTypeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedParameterTypeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedParameterTypeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedParameterTypeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedParameterTypeUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedParameterTypeUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedParameterTypeUpdate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedParameterTypeUpdate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterTypeUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterTypeUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterTypeUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterTypeUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterTypeUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterTypeUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterTypeUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterTypeUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedParameterTypeUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedParameterTypeUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


