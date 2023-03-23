# PatchedParameterTypeRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ParameterType** | Pointer to **string** | The type this rule is for. | [optional] [readonly] 
**Type** | Pointer to [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterTypeRuleUpdate

`func NewPatchedParameterTypeRuleUpdate() *PatchedParameterTypeRuleUpdate`

NewPatchedParameterTypeRuleUpdate instantiates a new PatchedParameterTypeRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterTypeRuleUpdateWithDefaults

`func NewPatchedParameterTypeRuleUpdateWithDefaults() *PatchedParameterTypeRuleUpdate`

NewPatchedParameterTypeRuleUpdateWithDefaults instantiates a new PatchedParameterTypeRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedParameterTypeRuleUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterTypeRuleUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterTypeRuleUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterTypeRuleUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameterType

`func (o *PatchedParameterTypeRuleUpdate) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *PatchedParameterTypeRuleUpdate) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *PatchedParameterTypeRuleUpdate) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *PatchedParameterTypeRuleUpdate) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameterTypeRuleUpdate) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameterTypeRuleUpdate) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameterTypeRuleUpdate) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameterTypeRuleUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConstraint

`func (o *PatchedParameterTypeRuleUpdate) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *PatchedParameterTypeRuleUpdate) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *PatchedParameterTypeRuleUpdate) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *PatchedParameterTypeRuleUpdate) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterTypeRuleUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterTypeRuleUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterTypeRuleUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterTypeRuleUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterTypeRuleUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterTypeRuleUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterTypeRuleUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterTypeRuleUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedParameterTypeRuleUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedParameterTypeRuleUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


