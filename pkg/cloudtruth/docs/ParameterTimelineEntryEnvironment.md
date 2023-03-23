# ParameterTimelineEntryEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **NullableString** | A unique identifier for the environment. | [readonly] 
**Name** | **string** | The environment name. | 
**Override** | **bool** | Indicates if the value change was direct or if it flowed into the environment. If &#x60;true&#x60; then the value was actually set directly into this environment. If &#x60;false&#x60; then the environment has no value set directly so it inherited the value from its parent. | [readonly] 

## Methods

### NewParameterTimelineEntryEnvironment

`func NewParameterTimelineEntryEnvironment(environmentId NullableString, name string, override bool, ) *ParameterTimelineEntryEnvironment`

NewParameterTimelineEntryEnvironment instantiates a new ParameterTimelineEntryEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTimelineEntryEnvironmentWithDefaults

`func NewParameterTimelineEntryEnvironmentWithDefaults() *ParameterTimelineEntryEnvironment`

NewParameterTimelineEntryEnvironmentWithDefaults instantiates a new ParameterTimelineEntryEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ParameterTimelineEntryEnvironment) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ParameterTimelineEntryEnvironment) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ParameterTimelineEntryEnvironment) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### SetEnvironmentIdNil

`func (o *ParameterTimelineEntryEnvironment) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *ParameterTimelineEntryEnvironment) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetName

`func (o *ParameterTimelineEntryEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterTimelineEntryEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterTimelineEntryEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetOverride

`func (o *ParameterTimelineEntryEnvironment) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *ParameterTimelineEntryEnvironment) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *ParameterTimelineEntryEnvironment) SetOverride(v bool)`

SetOverride sets Override field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


