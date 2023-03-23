# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The value url. | [readonly] 
**Id** | **string** |  | [readonly] 
**LedgerId** | **string** |  | [readonly] 
**Environment** | **string** | The environment this value is set in. | [readonly] 
**EnvironmentId** | **string** | The environment id for this value. | [readonly] 
**EnvironmentName** | **string** | The environment name for this value.  This is a convenience to avoid another query against the server to resolve the environment url into a name. | [readonly] 
**Parameter** | **string** | The parameter this value is for. | [readonly] 
**ParameterId** | **string** | The parameter id for this value. | [readonly] 
**External** | Pointer to **bool** |  | [optional] 
**ExternalFqn** | Pointer to **string** |  | [optional] 
**ExternalFilter** | Pointer to **NullableString** |  | [optional] 
**ExternalError** | **NullableString** | This field is deprecated and unused. | [readonly] 
**ExternalStatus** | [**NullableValueExternalStatus**](ValueExternalStatus.md) |  | 
**InternalValue** | Pointer to **NullableString** |  | [optional] 
**Interpolated** | Pointer to **bool** |  | [optional] 
**Value** | **NullableString** | This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  Clients applying this value to a shell environment should set &#x60;&lt;parameter_name&gt;&#x3D;&lt;value&gt;&#x60; even if &#x60;value&#x60; is the empty string.  If &#x60;value&#x60; is &#x60;null&#x60;, the client should unset that shell environment variable. | [readonly] 
**Evaluated** | **bool** | If true, the &#x60;value&#x60; field has undergone template evaluation. | [readonly] 
**Secret** | **NullableBool** | Indicates the value content is a secret.  Normally this is &#x60;true&#x60; when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret. | [readonly] 
**ReferencedProjects** | **[]string** | The projects this value references, if dynamic.  This field is not valid for history requests. | [readonly] 
**ReferencedParameters** | **[]string** | The parameters this value references, if dynamic.  this field is not valid for history requests. | [readonly] 
**ReferencedTemplates** | **[]string** | The templates this value references, if dynamic.  This field is not valid for history requests. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewValue

`func NewValue(url string, id string, ledgerId string, environment string, environmentId string, environmentName string, parameter string, parameterId string, externalError NullableString, externalStatus NullableValueExternalStatus, value NullableString, evaluated bool, secret NullableBool, referencedProjects []string, referencedParameters []string, referencedTemplates []string, createdAt time.Time, modifiedAt NullableTime, ) *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Value) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Value) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Value) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Value) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Value) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Value) SetId(v string)`

SetId sets Id field to given value.


### GetLedgerId

`func (o *Value) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *Value) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *Value) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetEnvironment

`func (o *Value) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Value) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Value) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEnvironmentId

`func (o *Value) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Value) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Value) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEnvironmentName

`func (o *Value) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *Value) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *Value) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetParameter

`func (o *Value) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *Value) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *Value) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### GetParameterId

`func (o *Value) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *Value) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *Value) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.


### GetExternal

`func (o *Value) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Value) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Value) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *Value) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalFqn

`func (o *Value) GetExternalFqn() string`

GetExternalFqn returns the ExternalFqn field if non-nil, zero value otherwise.

### GetExternalFqnOk

`func (o *Value) GetExternalFqnOk() (*string, bool)`

GetExternalFqnOk returns a tuple with the ExternalFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqn

`func (o *Value) SetExternalFqn(v string)`

SetExternalFqn sets ExternalFqn field to given value.

### HasExternalFqn

`func (o *Value) HasExternalFqn() bool`

HasExternalFqn returns a boolean if a field has been set.

### GetExternalFilter

`func (o *Value) GetExternalFilter() string`

GetExternalFilter returns the ExternalFilter field if non-nil, zero value otherwise.

### GetExternalFilterOk

`func (o *Value) GetExternalFilterOk() (*string, bool)`

GetExternalFilterOk returns a tuple with the ExternalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilter

`func (o *Value) SetExternalFilter(v string)`

SetExternalFilter sets ExternalFilter field to given value.

### HasExternalFilter

`func (o *Value) HasExternalFilter() bool`

HasExternalFilter returns a boolean if a field has been set.

### SetExternalFilterNil

`func (o *Value) SetExternalFilterNil(b bool)`

 SetExternalFilterNil sets the value for ExternalFilter to be an explicit nil

### UnsetExternalFilter
`func (o *Value) UnsetExternalFilter()`

UnsetExternalFilter ensures that no value is present for ExternalFilter, not even an explicit nil
### GetExternalError

`func (o *Value) GetExternalError() string`

GetExternalError returns the ExternalError field if non-nil, zero value otherwise.

### GetExternalErrorOk

`func (o *Value) GetExternalErrorOk() (*string, bool)`

GetExternalErrorOk returns a tuple with the ExternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalError

`func (o *Value) SetExternalError(v string)`

SetExternalError sets ExternalError field to given value.


### SetExternalErrorNil

`func (o *Value) SetExternalErrorNil(b bool)`

 SetExternalErrorNil sets the value for ExternalError to be an explicit nil

### UnsetExternalError
`func (o *Value) UnsetExternalError()`

UnsetExternalError ensures that no value is present for ExternalError, not even an explicit nil
### GetExternalStatus

`func (o *Value) GetExternalStatus() ValueExternalStatus`

GetExternalStatus returns the ExternalStatus field if non-nil, zero value otherwise.

### GetExternalStatusOk

`func (o *Value) GetExternalStatusOk() (*ValueExternalStatus, bool)`

GetExternalStatusOk returns a tuple with the ExternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatus

`func (o *Value) SetExternalStatus(v ValueExternalStatus)`

SetExternalStatus sets ExternalStatus field to given value.


### SetExternalStatusNil

`func (o *Value) SetExternalStatusNil(b bool)`

 SetExternalStatusNil sets the value for ExternalStatus to be an explicit nil

### UnsetExternalStatus
`func (o *Value) UnsetExternalStatus()`

UnsetExternalStatus ensures that no value is present for ExternalStatus, not even an explicit nil
### GetInternalValue

`func (o *Value) GetInternalValue() string`

GetInternalValue returns the InternalValue field if non-nil, zero value otherwise.

### GetInternalValueOk

`func (o *Value) GetInternalValueOk() (*string, bool)`

GetInternalValueOk returns a tuple with the InternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalValue

`func (o *Value) SetInternalValue(v string)`

SetInternalValue sets InternalValue field to given value.

### HasInternalValue

`func (o *Value) HasInternalValue() bool`

HasInternalValue returns a boolean if a field has been set.

### SetInternalValueNil

`func (o *Value) SetInternalValueNil(b bool)`

 SetInternalValueNil sets the value for InternalValue to be an explicit nil

### UnsetInternalValue
`func (o *Value) UnsetInternalValue()`

UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
### GetInterpolated

`func (o *Value) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *Value) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *Value) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *Value) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### GetValue

`func (o *Value) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Value) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Value) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Value) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Value) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEvaluated

`func (o *Value) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *Value) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *Value) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetSecret

`func (o *Value) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Value) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Value) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### SetSecretNil

`func (o *Value) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *Value) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetReferencedProjects

`func (o *Value) GetReferencedProjects() []string`

GetReferencedProjects returns the ReferencedProjects field if non-nil, zero value otherwise.

### GetReferencedProjectsOk

`func (o *Value) GetReferencedProjectsOk() (*[]string, bool)`

GetReferencedProjectsOk returns a tuple with the ReferencedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedProjects

`func (o *Value) SetReferencedProjects(v []string)`

SetReferencedProjects sets ReferencedProjects field to given value.


### GetReferencedParameters

`func (o *Value) GetReferencedParameters() []string`

GetReferencedParameters returns the ReferencedParameters field if non-nil, zero value otherwise.

### GetReferencedParametersOk

`func (o *Value) GetReferencedParametersOk() (*[]string, bool)`

GetReferencedParametersOk returns a tuple with the ReferencedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedParameters

`func (o *Value) SetReferencedParameters(v []string)`

SetReferencedParameters sets ReferencedParameters field to given value.


### GetReferencedTemplates

`func (o *Value) GetReferencedTemplates() []string`

GetReferencedTemplates returns the ReferencedTemplates field if non-nil, zero value otherwise.

### GetReferencedTemplatesOk

`func (o *Value) GetReferencedTemplatesOk() (*[]string, bool)`

GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTemplates

`func (o *Value) SetReferencedTemplates(v []string)`

SetReferencedTemplates sets ReferencedTemplates field to given value.


### GetCreatedAt

`func (o *Value) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Value) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Value) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Value) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Value) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Value) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *Value) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Value) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


