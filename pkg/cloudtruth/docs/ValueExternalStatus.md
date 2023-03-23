# ValueExternalStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for a task step. | [readonly] 
**Operation** | Pointer to [**NullableOperationEnum**](OperationEnum.md) | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | [optional] 
**Success** | **bool** | Indicates if the operation was successful. | 
**SuccessDetail** | Pointer to **NullableString** | Additional details about the successful operation, if any. | [optional] 
**Fqn** | Pointer to **NullableString** | The fully-qualified name (FQN) this of the value that was changed. | [optional] 
**Environment** | **NullableString** | The environment affected by this step. | [readonly] 
**EnvironmentId** | Pointer to **NullableString** | The environment id involved in the operation. | [optional] 
**EnvironmentName** | Pointer to **NullableString** | The environment name involved in the operation. | [optional] 
**Project** | **NullableString** | The project affected by this step. | [readonly] 
**ProjectId** | Pointer to **NullableString** | The project id involved in the operation. | [optional] 
**ProjectName** | Pointer to **NullableString** | The project name involved in the operation. | [optional] 
**Parameter** | **NullableString** | The parameter affected by this step. | [readonly] 
**ParameterId** | Pointer to **NullableString** | The parameter id involved in the operation. | [optional] 
**ParameterName** | Pointer to **NullableString** | The parameter name involved in the operation. | [optional] 
**VenueId** | Pointer to **NullableString** | The integration-native id for the resource. | [optional] 
**VenueName** | Pointer to **NullableString** | The name of the item or resource as known by the integration. | [optional] 
**Summary** | Pointer to **NullableString** | The summary of this step (what it was trying to do). | [optional] 
**ErrorCode** | Pointer to **NullableString** | An error code, if not successful. | [optional] 
**ErrorDetail** | Pointer to **NullableString** | Details on the error that occurred during processing. | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewValueExternalStatus

`func NewValueExternalStatus(url string, id string, success bool, environment NullableString, project NullableString, parameter NullableString, createdAt time.Time, modifiedAt NullableTime, ) *ValueExternalStatus`

NewValueExternalStatus instantiates a new ValueExternalStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueExternalStatusWithDefaults

`func NewValueExternalStatusWithDefaults() *ValueExternalStatus`

NewValueExternalStatusWithDefaults instantiates a new ValueExternalStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ValueExternalStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ValueExternalStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ValueExternalStatus) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *ValueExternalStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValueExternalStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValueExternalStatus) SetId(v string)`

SetId sets Id field to given value.


### GetOperation

`func (o *ValueExternalStatus) GetOperation() OperationEnum`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ValueExternalStatus) GetOperationOk() (*OperationEnum, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ValueExternalStatus) SetOperation(v OperationEnum)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ValueExternalStatus) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *ValueExternalStatus) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *ValueExternalStatus) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetSuccess

`func (o *ValueExternalStatus) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ValueExternalStatus) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ValueExternalStatus) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetSuccessDetail

`func (o *ValueExternalStatus) GetSuccessDetail() string`

GetSuccessDetail returns the SuccessDetail field if non-nil, zero value otherwise.

### GetSuccessDetailOk

`func (o *ValueExternalStatus) GetSuccessDetailOk() (*string, bool)`

GetSuccessDetailOk returns a tuple with the SuccessDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDetail

`func (o *ValueExternalStatus) SetSuccessDetail(v string)`

SetSuccessDetail sets SuccessDetail field to given value.

### HasSuccessDetail

`func (o *ValueExternalStatus) HasSuccessDetail() bool`

HasSuccessDetail returns a boolean if a field has been set.

### SetSuccessDetailNil

`func (o *ValueExternalStatus) SetSuccessDetailNil(b bool)`

 SetSuccessDetailNil sets the value for SuccessDetail to be an explicit nil

### UnsetSuccessDetail
`func (o *ValueExternalStatus) UnsetSuccessDetail()`

UnsetSuccessDetail ensures that no value is present for SuccessDetail, not even an explicit nil
### GetFqn

`func (o *ValueExternalStatus) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *ValueExternalStatus) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *ValueExternalStatus) SetFqn(v string)`

SetFqn sets Fqn field to given value.

### HasFqn

`func (o *ValueExternalStatus) HasFqn() bool`

HasFqn returns a boolean if a field has been set.

### SetFqnNil

`func (o *ValueExternalStatus) SetFqnNil(b bool)`

 SetFqnNil sets the value for Fqn to be an explicit nil

### UnsetFqn
`func (o *ValueExternalStatus) UnsetFqn()`

UnsetFqn ensures that no value is present for Fqn, not even an explicit nil
### GetEnvironment

`func (o *ValueExternalStatus) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ValueExternalStatus) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ValueExternalStatus) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *ValueExternalStatus) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ValueExternalStatus) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvironmentId

`func (o *ValueExternalStatus) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ValueExternalStatus) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ValueExternalStatus) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ValueExternalStatus) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *ValueExternalStatus) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *ValueExternalStatus) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetEnvironmentName

`func (o *ValueExternalStatus) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ValueExternalStatus) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ValueExternalStatus) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *ValueExternalStatus) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### SetEnvironmentNameNil

`func (o *ValueExternalStatus) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *ValueExternalStatus) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil
### GetProject

`func (o *ValueExternalStatus) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ValueExternalStatus) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ValueExternalStatus) SetProject(v string)`

SetProject sets Project field to given value.


### SetProjectNil

`func (o *ValueExternalStatus) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ValueExternalStatus) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProjectId

`func (o *ValueExternalStatus) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ValueExternalStatus) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ValueExternalStatus) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ValueExternalStatus) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ValueExternalStatus) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ValueExternalStatus) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *ValueExternalStatus) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ValueExternalStatus) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ValueExternalStatus) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ValueExternalStatus) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ValueExternalStatus) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ValueExternalStatus) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetParameter

`func (o *ValueExternalStatus) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ValueExternalStatus) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ValueExternalStatus) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### SetParameterNil

`func (o *ValueExternalStatus) SetParameterNil(b bool)`

 SetParameterNil sets the value for Parameter to be an explicit nil

### UnsetParameter
`func (o *ValueExternalStatus) UnsetParameter()`

UnsetParameter ensures that no value is present for Parameter, not even an explicit nil
### GetParameterId

`func (o *ValueExternalStatus) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *ValueExternalStatus) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *ValueExternalStatus) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.

### HasParameterId

`func (o *ValueExternalStatus) HasParameterId() bool`

HasParameterId returns a boolean if a field has been set.

### SetParameterIdNil

`func (o *ValueExternalStatus) SetParameterIdNil(b bool)`

 SetParameterIdNil sets the value for ParameterId to be an explicit nil

### UnsetParameterId
`func (o *ValueExternalStatus) UnsetParameterId()`

UnsetParameterId ensures that no value is present for ParameterId, not even an explicit nil
### GetParameterName

`func (o *ValueExternalStatus) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *ValueExternalStatus) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *ValueExternalStatus) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *ValueExternalStatus) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### SetParameterNameNil

`func (o *ValueExternalStatus) SetParameterNameNil(b bool)`

 SetParameterNameNil sets the value for ParameterName to be an explicit nil

### UnsetParameterName
`func (o *ValueExternalStatus) UnsetParameterName()`

UnsetParameterName ensures that no value is present for ParameterName, not even an explicit nil
### GetVenueId

`func (o *ValueExternalStatus) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *ValueExternalStatus) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *ValueExternalStatus) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.

### HasVenueId

`func (o *ValueExternalStatus) HasVenueId() bool`

HasVenueId returns a boolean if a field has been set.

### SetVenueIdNil

`func (o *ValueExternalStatus) SetVenueIdNil(b bool)`

 SetVenueIdNil sets the value for VenueId to be an explicit nil

### UnsetVenueId
`func (o *ValueExternalStatus) UnsetVenueId()`

UnsetVenueId ensures that no value is present for VenueId, not even an explicit nil
### GetVenueName

`func (o *ValueExternalStatus) GetVenueName() string`

GetVenueName returns the VenueName field if non-nil, zero value otherwise.

### GetVenueNameOk

`func (o *ValueExternalStatus) GetVenueNameOk() (*string, bool)`

GetVenueNameOk returns a tuple with the VenueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueName

`func (o *ValueExternalStatus) SetVenueName(v string)`

SetVenueName sets VenueName field to given value.

### HasVenueName

`func (o *ValueExternalStatus) HasVenueName() bool`

HasVenueName returns a boolean if a field has been set.

### SetVenueNameNil

`func (o *ValueExternalStatus) SetVenueNameNil(b bool)`

 SetVenueNameNil sets the value for VenueName to be an explicit nil

### UnsetVenueName
`func (o *ValueExternalStatus) UnsetVenueName()`

UnsetVenueName ensures that no value is present for VenueName, not even an explicit nil
### GetSummary

`func (o *ValueExternalStatus) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ValueExternalStatus) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ValueExternalStatus) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ValueExternalStatus) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *ValueExternalStatus) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *ValueExternalStatus) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetErrorCode

`func (o *ValueExternalStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ValueExternalStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ValueExternalStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ValueExternalStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ValueExternalStatus) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ValueExternalStatus) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *ValueExternalStatus) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *ValueExternalStatus) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *ValueExternalStatus) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *ValueExternalStatus) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *ValueExternalStatus) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *ValueExternalStatus) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetCreatedAt

`func (o *ValueExternalStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ValueExternalStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ValueExternalStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ValueExternalStatus) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ValueExternalStatus) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ValueExternalStatus) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *ValueExternalStatus) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ValueExternalStatus) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


