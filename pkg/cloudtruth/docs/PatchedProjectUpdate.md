# PatchedProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The project name. | [optional] 
**Description** | Pointer to **string** | A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content. | [optional] 
**ParameterNamePattern** | Pointer to **string** | A regular expression parameter names must match | [optional] 
**DependsOn** | Pointer to **NullableString** | Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project. | [optional] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | Your role in the project, if the project is access-controlled. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedProjectUpdate

`func NewPatchedProjectUpdate() *PatchedProjectUpdate`

NewPatchedProjectUpdate instantiates a new PatchedProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProjectUpdateWithDefaults

`func NewPatchedProjectUpdateWithDefaults() *PatchedProjectUpdate`

NewPatchedProjectUpdateWithDefaults instantiates a new PatchedProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedProjectUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedProjectUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedProjectUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedProjectUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedProjectUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProjectUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProjectUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProjectUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedProjectUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedProjectUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedProjectUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedProjectUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameterNamePattern

`func (o *PatchedProjectUpdate) GetParameterNamePattern() string`

GetParameterNamePattern returns the ParameterNamePattern field if non-nil, zero value otherwise.

### GetParameterNamePatternOk

`func (o *PatchedProjectUpdate) GetParameterNamePatternOk() (*string, bool)`

GetParameterNamePatternOk returns a tuple with the ParameterNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterNamePattern

`func (o *PatchedProjectUpdate) SetParameterNamePattern(v string)`

SetParameterNamePattern sets ParameterNamePattern field to given value.

### HasParameterNamePattern

`func (o *PatchedProjectUpdate) HasParameterNamePattern() bool`

HasParameterNamePattern returns a boolean if a field has been set.

### GetDependsOn

`func (o *PatchedProjectUpdate) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *PatchedProjectUpdate) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *PatchedProjectUpdate) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *PatchedProjectUpdate) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *PatchedProjectUpdate) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *PatchedProjectUpdate) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetAccessControlled

`func (o *PatchedProjectUpdate) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *PatchedProjectUpdate) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *PatchedProjectUpdate) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *PatchedProjectUpdate) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *PatchedProjectUpdate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedProjectUpdate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedProjectUpdate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedProjectUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedProjectUpdate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedProjectUpdate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *PatchedProjectUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedProjectUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedProjectUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedProjectUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedProjectUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedProjectUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedProjectUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedProjectUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedProjectUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedProjectUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


