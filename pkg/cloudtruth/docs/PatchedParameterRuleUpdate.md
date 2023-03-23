# PatchedParameterRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Parameter** | Pointer to **string** | The parameter this rule is for. | [optional] [readonly] 
**Type** | Pointer to [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterRuleUpdate

`func NewPatchedParameterRuleUpdate() *PatchedParameterRuleUpdate`

NewPatchedParameterRuleUpdate instantiates a new PatchedParameterRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterRuleUpdateWithDefaults

`func NewPatchedParameterRuleUpdateWithDefaults() *PatchedParameterRuleUpdate`

NewPatchedParameterRuleUpdateWithDefaults instantiates a new PatchedParameterRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedParameterRuleUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterRuleUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterRuleUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterRuleUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameter

`func (o *PatchedParameterRuleUpdate) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *PatchedParameterRuleUpdate) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *PatchedParameterRuleUpdate) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *PatchedParameterRuleUpdate) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameterRuleUpdate) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameterRuleUpdate) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameterRuleUpdate) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameterRuleUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConstraint

`func (o *PatchedParameterRuleUpdate) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *PatchedParameterRuleUpdate) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *PatchedParameterRuleUpdate) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *PatchedParameterRuleUpdate) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterRuleUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterRuleUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterRuleUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterRuleUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterRuleUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterRuleUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterRuleUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterRuleUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedParameterRuleUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedParameterRuleUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


