# ParameterRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Parameter** | **string** | The parameter this rule is for. | [readonly] 
**Type** | [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | 
**Constraint** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewParameterRuleUpdate

`func NewParameterRuleUpdate(id string, parameter string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt NullableTime, ) *ParameterRuleUpdate`

NewParameterRuleUpdate instantiates a new ParameterRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterRuleUpdateWithDefaults

`func NewParameterRuleUpdateWithDefaults() *ParameterRuleUpdate`

NewParameterRuleUpdateWithDefaults instantiates a new ParameterRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterRuleUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterRuleUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterRuleUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetParameter

`func (o *ParameterRuleUpdate) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ParameterRuleUpdate) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ParameterRuleUpdate) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### GetType

`func (o *ParameterRuleUpdate) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterRuleUpdate) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterRuleUpdate) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.


### GetConstraint

`func (o *ParameterRuleUpdate) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *ParameterRuleUpdate) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *ParameterRuleUpdate) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.


### GetCreatedAt

`func (o *ParameterRuleUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterRuleUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterRuleUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterRuleUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterRuleUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterRuleUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ParameterRuleUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ParameterRuleUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


