# ProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | The project name. | 
**Description** | Pointer to **string** | A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content. | [optional] 
**ParameterNamePattern** | Pointer to **string** | A regular expression parameter names must match | [optional] 
**DependsOn** | Pointer to **NullableString** | Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project. | [optional] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | Your role in the project, if the project is access-controlled. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewProjectUpdate

`func NewProjectUpdate(id string, name string, role NullableRoleEnum, createdAt time.Time, modifiedAt NullableTime, ) *ProjectUpdate`

NewProjectUpdate instantiates a new ProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateWithDefaults

`func NewProjectUpdateWithDefaults() *ProjectUpdate`

NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameterNamePattern

`func (o *ProjectUpdate) GetParameterNamePattern() string`

GetParameterNamePattern returns the ParameterNamePattern field if non-nil, zero value otherwise.

### GetParameterNamePatternOk

`func (o *ProjectUpdate) GetParameterNamePatternOk() (*string, bool)`

GetParameterNamePatternOk returns a tuple with the ParameterNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterNamePattern

`func (o *ProjectUpdate) SetParameterNamePattern(v string)`

SetParameterNamePattern sets ParameterNamePattern field to given value.

### HasParameterNamePattern

`func (o *ProjectUpdate) HasParameterNamePattern() bool`

HasParameterNamePattern returns a boolean if a field has been set.

### GetDependsOn

`func (o *ProjectUpdate) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ProjectUpdate) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ProjectUpdate) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ProjectUpdate) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *ProjectUpdate) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *ProjectUpdate) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetAccessControlled

`func (o *ProjectUpdate) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *ProjectUpdate) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *ProjectUpdate) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *ProjectUpdate) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *ProjectUpdate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectUpdate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectUpdate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *ProjectUpdate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ProjectUpdate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *ProjectUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ProjectUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ProjectUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ProjectUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ProjectUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ProjectUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


