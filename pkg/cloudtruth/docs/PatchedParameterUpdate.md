# PatchedParameterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The parameter name. | [optional] 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Secret** | Pointer to **bool** | Indicates if this content is secret or not.  External values are inspected on-demand to ensure they align with the parameter&#39;s secret setting and if they do not, those external values are not allowed to be used. | [optional] 
**Type** | Pointer to **string** | The type of this Parameter. | [optional] [default to "string"]
**Project** | Pointer to **string** | The project url. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterUpdate

`func NewPatchedParameterUpdate() *PatchedParameterUpdate`

NewPatchedParameterUpdate instantiates a new PatchedParameterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterUpdateWithDefaults

`func NewPatchedParameterUpdateWithDefaults() *PatchedParameterUpdate`

NewPatchedParameterUpdateWithDefaults instantiates a new PatchedParameterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedParameterUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedParameterUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedParameterUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedParameterUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedParameterUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedParameterUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedParameterUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedParameterUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedParameterUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedParameterUpdate) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedParameterUpdate) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedParameterUpdate) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedParameterUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameterUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameterUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameterUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameterUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *PatchedParameterUpdate) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedParameterUpdate) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedParameterUpdate) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedParameterUpdate) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedParameterUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedParameterUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


