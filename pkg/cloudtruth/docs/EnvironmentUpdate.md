# EnvironmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | The environment name. | 
**Description** | Pointer to **string** | A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content. | [optional] 
**Parent** | Pointer to **NullableString** | Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed. | [optional] 
**Children** | **[]string** | This is the opposite of &#x60;parent&#x60;, see that field for more details. | [readonly] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | Your role in the environment, if the environment is access-controlled. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewEnvironmentUpdate

`func NewEnvironmentUpdate(id string, name string, children []string, role NullableRoleEnum, createdAt time.Time, modifiedAt NullableTime, ) *EnvironmentUpdate`

NewEnvironmentUpdate instantiates a new EnvironmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUpdateWithDefaults

`func NewEnvironmentUpdateWithDefaults() *EnvironmentUpdate`

NewEnvironmentUpdateWithDefaults instantiates a new EnvironmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnvironmentUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnvironmentUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *EnvironmentUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EnvironmentUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EnvironmentUpdate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EnvironmentUpdate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *EnvironmentUpdate) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *EnvironmentUpdate) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetChildren

`func (o *EnvironmentUpdate) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *EnvironmentUpdate) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *EnvironmentUpdate) SetChildren(v []string)`

SetChildren sets Children field to given value.


### GetAccessControlled

`func (o *EnvironmentUpdate) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *EnvironmentUpdate) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *EnvironmentUpdate) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *EnvironmentUpdate) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *EnvironmentUpdate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EnvironmentUpdate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EnvironmentUpdate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *EnvironmentUpdate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *EnvironmentUpdate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *EnvironmentUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EnvironmentUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EnvironmentUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *EnvironmentUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *EnvironmentUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


