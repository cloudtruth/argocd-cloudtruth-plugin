# ParameterTypeRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ParameterType** | **string** | The type this rule is for. | [readonly] 
**Type** | [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | 
**Constraint** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewParameterTypeRuleUpdate

`func NewParameterTypeRuleUpdate(id string, parameterType string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt NullableTime, ) *ParameterTypeRuleUpdate`

NewParameterTypeRuleUpdate instantiates a new ParameterTypeRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeRuleUpdateWithDefaults

`func NewParameterTypeRuleUpdateWithDefaults() *ParameterTypeRuleUpdate`

NewParameterTypeRuleUpdateWithDefaults instantiates a new ParameterTypeRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterTypeRuleUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterTypeRuleUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterTypeRuleUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetParameterType

`func (o *ParameterTypeRuleUpdate) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ParameterTypeRuleUpdate) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ParameterTypeRuleUpdate) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.


### GetType

`func (o *ParameterTypeRuleUpdate) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterTypeRuleUpdate) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterTypeRuleUpdate) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.


### GetConstraint

`func (o *ParameterTypeRuleUpdate) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *ParameterTypeRuleUpdate) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *ParameterTypeRuleUpdate) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.


### GetCreatedAt

`func (o *ParameterTypeRuleUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterTypeRuleUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterTypeRuleUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterTypeRuleUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterTypeRuleUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterTypeRuleUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ParameterTypeRuleUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ParameterTypeRuleUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


