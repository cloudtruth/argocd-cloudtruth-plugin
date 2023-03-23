# ParameterTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | The parameter type name. | 
**Description** | Pointer to **string** | A description of the parameter type, provide documentation on how to use this type here. | [optional] 
**Parent** | **string** | The URL for this parameter type&#39;s parent | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewParameterTypeUpdate

`func NewParameterTypeUpdate(id string, name string, parent string, createdAt time.Time, modifiedAt NullableTime, ) *ParameterTypeUpdate`

NewParameterTypeUpdate instantiates a new ParameterTypeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeUpdateWithDefaults

`func NewParameterTypeUpdateWithDefaults() *ParameterTypeUpdate`

NewParameterTypeUpdateWithDefaults instantiates a new ParameterTypeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterTypeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterTypeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterTypeUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ParameterTypeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterTypeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterTypeUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParameterTypeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterTypeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterTypeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterTypeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *ParameterTypeUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ParameterTypeUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ParameterTypeUpdate) SetParent(v string)`

SetParent sets Parent field to given value.


### GetCreatedAt

`func (o *ParameterTypeUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterTypeUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterTypeUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterTypeUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterTypeUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterTypeUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ParameterTypeUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ParameterTypeUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


