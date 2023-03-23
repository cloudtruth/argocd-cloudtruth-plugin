# ServiceAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional description of the process or system using the service account. | [optional] 
**Role** | **string** | The role for the service acount | [readonly] 

## Methods

### NewServiceAccountUpdateRequest

`func NewServiceAccountUpdateRequest(role string, ) *ServiceAccountUpdateRequest`

NewServiceAccountUpdateRequest instantiates a new ServiceAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountUpdateRequestWithDefaults

`func NewServiceAccountUpdateRequestWithDefaults() *ServiceAccountUpdateRequest`

NewServiceAccountUpdateRequestWithDefaults instantiates a new ServiceAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceAccountUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *ServiceAccountUpdateRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceAccountUpdateRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceAccountUpdateRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


