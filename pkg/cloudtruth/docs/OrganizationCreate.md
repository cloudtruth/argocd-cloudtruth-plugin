# OrganizationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The organization name. | 
**MfaEnabled** | Pointer to **bool** | Multi-factor authentication for the organization | [optional] 

## Methods

### NewOrganizationCreate

`func NewOrganizationCreate(name string, ) *OrganizationCreate`

NewOrganizationCreate instantiates a new OrganizationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateWithDefaults

`func NewOrganizationCreateWithDefaults() *OrganizationCreate`

NewOrganizationCreateWithDefaults instantiates a new OrganizationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetMfaEnabled

`func (o *OrganizationCreate) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *OrganizationCreate) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *OrganizationCreate) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *OrganizationCreate) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


