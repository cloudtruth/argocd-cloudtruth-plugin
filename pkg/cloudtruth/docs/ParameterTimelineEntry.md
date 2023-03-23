# ParameterTimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryType** | [**HistoryTypeEnum**](HistoryTypeEnum.md) |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**HistoryEnvironments** | [**[]ParameterTimelineEntryEnvironment**](ParameterTimelineEntryEnvironment.md) | The affected environment(s). | [readonly] 
**HistoryModel** | [**HistoryModelEnum**](HistoryModelEnum.md) | The component of the parameter that changed. | [readonly] 
**HistoryParameter** | [**ParameterTimelineEntryHistoryParameter**](ParameterTimelineEntryHistoryParameter.md) |  | 

## Methods

### NewParameterTimelineEntry

`func NewParameterTimelineEntry(historyType HistoryTypeEnum, modifiedAt NullableTime, historyEnvironments []ParameterTimelineEntryEnvironment, historyModel HistoryModelEnum, historyParameter ParameterTimelineEntryHistoryParameter, ) *ParameterTimelineEntry`

NewParameterTimelineEntry instantiates a new ParameterTimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTimelineEntryWithDefaults

`func NewParameterTimelineEntryWithDefaults() *ParameterTimelineEntry`

NewParameterTimelineEntryWithDefaults instantiates a new ParameterTimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetModifiedAt

`func (o *ParameterTimelineEntry) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterTimelineEntry) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterTimelineEntry) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ParameterTimelineEntry) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ParameterTimelineEntry) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetModifiedBy

`func (o *ParameterTimelineEntry) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ParameterTimelineEntry) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ParameterTimelineEntry) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ParameterTimelineEntry) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

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

`func (o *ParameterTimelineEntry) GetHistoryParameter() ParameterTimelineEntryHistoryParameter`

GetHistoryParameter returns the HistoryParameter field if non-nil, zero value otherwise.

### GetHistoryParameterOk

`func (o *ParameterTimelineEntry) GetHistoryParameterOk() (*ParameterTimelineEntryHistoryParameter, bool)`

GetHistoryParameterOk returns a tuple with the HistoryParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryParameter

`func (o *ParameterTimelineEntry) SetHistoryParameter(v ParameterTimelineEntryHistoryParameter)`

SetHistoryParameter sets HistoryParameter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


