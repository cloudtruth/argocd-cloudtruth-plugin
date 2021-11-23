# ParameterTimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryDate** | **time.Time** |  | 
**HistoryType** | [**HistoryTypeEnum**](HistoryTypeEnum.md) |  | [readonly] 
**HistoryUser** | Pointer to **NullableString** | The unique identifier of a user. | [optional] 
**HistoryEnvironments** | [**[]ParameterTimelineEntryEnvironment**](ParameterTimelineEntryEnvironment.md) | The affected environment(s). | [readonly] 
**HistoryModel** | [**HistoryModelEnum**](HistoryModelEnum.md) | The component of the parameter that changed. | [readonly] 
**HistoryParameter** | [**ParameterTimelineEntryParameter**](ParameterTimelineEntryParameter.md) | The affected parameter. | [readonly] 

## Methods

### NewParameterTimelineEntry

`func NewParameterTimelineEntry(historyDate time.Time, historyType HistoryTypeEnum, historyEnvironments []ParameterTimelineEntryEnvironment, historyModel HistoryModelEnum, historyParameter ParameterTimelineEntryParameter, ) *ParameterTimelineEntry`

NewParameterTimelineEntry instantiates a new ParameterTimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTimelineEntryWithDefaults

`func NewParameterTimelineEntryWithDefaults() *ParameterTimelineEntry`

NewParameterTimelineEntryWithDefaults instantiates a new ParameterTimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryDate

`func (o *ParameterTimelineEntry) GetHistoryDate() time.Time`

GetHistoryDate returns the HistoryDate field if non-nil, zero value otherwise.

### GetHistoryDateOk

`func (o *ParameterTimelineEntry) GetHistoryDateOk() (*time.Time, bool)`

GetHistoryDateOk returns a tuple with the HistoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryDate

`func (o *ParameterTimelineEntry) SetHistoryDate(v time.Time)`

SetHistoryDate sets HistoryDate field to given value.


### GetHistoryType

`func (o *ParameterTimelineEntry) GetHistoryType() HistoryTypeEnum`

GetHistoryType returns the HistoryType field if non-nil, zero value otherwise.

### GetHistoryTypeOk

`func (o *ParameterTimelineEntry) GetHistoryTypeOk() (*HistoryTypeEnum, bool)`

GetHistoryTypeOk returns a tuple with the HistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryType

`func (o *ParameterTimelineEntry) SetHistoryType(v HistoryTypeEnum)`

SetHistoryType sets HistoryType field to given value.


### GetHistoryUser

`func (o *ParameterTimelineEntry) GetHistoryUser() string`

GetHistoryUser returns the HistoryUser field if non-nil, zero value otherwise.

### GetHistoryUserOk

`func (o *ParameterTimelineEntry) GetHistoryUserOk() (*string, bool)`

GetHistoryUserOk returns a tuple with the HistoryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryUser

`func (o *ParameterTimelineEntry) SetHistoryUser(v string)`

SetHistoryUser sets HistoryUser field to given value.

### HasHistoryUser

`func (o *ParameterTimelineEntry) HasHistoryUser() bool`

HasHistoryUser returns a boolean if a field has been set.

### SetHistoryUserNil

`func (o *ParameterTimelineEntry) SetHistoryUserNil(b bool)`

 SetHistoryUserNil sets the value for HistoryUser to be an explicit nil

### UnsetHistoryUser
`func (o *ParameterTimelineEntry) UnsetHistoryUser()`

UnsetHistoryUser ensures that no value is present for HistoryUser, not even an explicit nil
### GetHistoryEnvironments

`func (o *ParameterTimelineEntry) GetHistoryEnvironments() []ParameterTimelineEntryEnvironment`

GetHistoryEnvironments returns the HistoryEnvironments field if non-nil, zero value otherwise.

### GetHistoryEnvironmentsOk

`func (o *ParameterTimelineEntry) GetHistoryEnvironmentsOk() (*[]ParameterTimelineEntryEnvironment, bool)`

GetHistoryEnvironmentsOk returns a tuple with the HistoryEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryEnvironments

`func (o *ParameterTimelineEntry) SetHistoryEnvironments(v []ParameterTimelineEntryEnvironment)`

SetHistoryEnvironments sets HistoryEnvironments field to given value.


### GetHistoryModel

`func (o *ParameterTimelineEntry) GetHistoryModel() HistoryModelEnum`

GetHistoryModel returns the HistoryModel field if non-nil, zero value otherwise.

### GetHistoryModelOk

`func (o *ParameterTimelineEntry) GetHistoryModelOk() (*HistoryModelEnum, bool)`

GetHistoryModelOk returns a tuple with the HistoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryModel

`func (o *ParameterTimelineEntry) SetHistoryModel(v HistoryModelEnum)`

SetHistoryModel sets HistoryModel field to given value.


### GetHistoryParameter

`func (o *ParameterTimelineEntry) GetHistoryParameter() ParameterTimelineEntryParameter`

GetHistoryParameter returns the HistoryParameter field if non-nil, zero value otherwise.

### GetHistoryParameterOk

`func (o *ParameterTimelineEntry) GetHistoryParameterOk() (*ParameterTimelineEntryParameter, bool)`

GetHistoryParameterOk returns a tuple with the HistoryParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryParameter

`func (o *ParameterTimelineEntry) SetHistoryParameter(v ParameterTimelineEntryParameter)`

SetHistoryParameter sets HistoryParameter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


