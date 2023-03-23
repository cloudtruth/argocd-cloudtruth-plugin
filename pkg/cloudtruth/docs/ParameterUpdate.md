# ParameterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | The parameter name. | 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Secret** | Pointer to **bool** | Indicates if this content is secret or not.  External values are inspected on-demand to ensure they align with the parameter&#39;s secret setting and if they do not, those external values are not allowed to be used. | [optional] 
**Type** | Pointer to **string** | The type of this Parameter. | [optional] [default to "string"]
**Project** | **string** | The project url. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewParameterUpdate

`func NewParameterUpdate(id string, name string, project string, createdAt time.Time, modifiedAt NullableTime, ) *ParameterUpdate`

NewParameterUpdate instantiates a new ParameterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterUpdateWithDefaults

`func NewParameterUpdateWithDefaults() *ParameterUpdate`

NewParameterUpdateWithDefaults instantiates a new ParameterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ParameterUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParameterUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *ParameterUpdate) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ParameterUpdate) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ParameterUpdate) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ParameterUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *ParameterUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ParameterUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *ParameterUpdate) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ParameterUpdate) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ParameterUpdate) SetProject(v string)`

SetProject sets Project field to given value.


### GetCreatedAt

`func (o *ParameterUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ParameterUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ParameterUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


