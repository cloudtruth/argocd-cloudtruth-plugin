# PatchedEnvironmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The environment name. | [optional] 
**Description** | Pointer to **string** | A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content. | [optional] 
**Parent** | Pointer to **NullableString** | Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed. | [optional] 
**Children** | Pointer to **[]string** | This is the opposite of &#x60;parent&#x60;, see that field for more details. | [optional] [readonly] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | Your role in the environment, if the environment is access-controlled. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedEnvironmentUpdate

`func NewPatchedEnvironmentUpdate() *PatchedEnvironmentUpdate`

NewPatchedEnvironmentUpdate instantiates a new PatchedEnvironmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEnvironmentUpdateWithDefaults

`func NewPatchedEnvironmentUpdateWithDefaults() *PatchedEnvironmentUpdate`

NewPatchedEnvironmentUpdateWithDefaults instantiates a new PatchedEnvironmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedEnvironmentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEnvironmentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEnvironmentUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEnvironmentUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedEnvironmentUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEnvironmentUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEnvironmentUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEnvironmentUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEnvironmentUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEnvironmentUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEnvironmentUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEnvironmentUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedEnvironmentUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedEnvironmentUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedEnvironmentUpdate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedEnvironmentUpdate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedEnvironmentUpdate) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedEnvironmentUpdate) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetChildren

`func (o *PatchedEnvironmentUpdate) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PatchedEnvironmentUpdate) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PatchedEnvironmentUpdate) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PatchedEnvironmentUpdate) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetAccessControlled

`func (o *PatchedEnvironmentUpdate) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *PatchedEnvironmentUpdate) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *PatchedEnvironmentUpdate) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *PatchedEnvironmentUpdate) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *PatchedEnvironmentUpdate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedEnvironmentUpdate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedEnvironmentUpdate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedEnvironmentUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedEnvironmentUpdate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedEnvironmentUpdate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *PatchedEnvironmentUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedEnvironmentUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedEnvironmentUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedEnvironmentUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedEnvironmentUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedEnvironmentUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedEnvironmentUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedEnvironmentUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedEnvironmentUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedEnvironmentUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


