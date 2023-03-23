# PatchedValueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**External** | Pointer to **bool** | An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is &#x60;false&#x60; the value is stored by CloudTruth and considered to be _internal_.  When this is &#x60;true&#x60;, the &#x60;external_fqn&#x60; field must be set. | [optional] 
**ExternalFqn** | Pointer to **string** | The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is &#x60;external&#x60;. | [optional] 
**ExternalFilter** | Pointer to **string** | If the value is &#x60;external&#x60;, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on &#x60;json&#x60;, &#x60;yaml&#x60;, and &#x60;dotenv&#x60; content. | [optional] 
**InternalValue** | Pointer to **NullableString** | This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  This field cannot be specified when creating or updating an &#x60;external&#x60; value. | [optional] 
**Interpolated** | Pointer to **bool** | If &#x60;true&#x60;, apply template substitution rules to this value.  If &#x60;false&#x60;, this value is a literal value.  Note: secrets cannot be interpolated. | [optional] 
**Secret** | Pointer to **NullableBool** | Indicates the value content is a secret.  Normally this is &#x60;true&#x60; when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  Clients applying this value to a shell environment should set &#x60;&lt;parameter_name&gt;&#x3D;&lt;value&gt;&#x60; even if &#x60;value&#x60; is the empty string.  If &#x60;value&#x60; is &#x60;null&#x60;, the client should unset that shell environment variable. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedValueUpdate

`func NewPatchedValueUpdate() *PatchedValueUpdate`

NewPatchedValueUpdate instantiates a new PatchedValueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedValueUpdateWithDefaults

`func NewPatchedValueUpdateWithDefaults() *PatchedValueUpdate`

NewPatchedValueUpdateWithDefaults instantiates a new PatchedValueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedValueUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedValueUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedValueUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedValueUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternal

`func (o *PatchedValueUpdate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *PatchedValueUpdate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *PatchedValueUpdate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *PatchedValueUpdate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalFqn

`func (o *PatchedValueUpdate) GetExternalFqn() string`

GetExternalFqn returns the ExternalFqn field if non-nil, zero value otherwise.

### GetExternalFqnOk

`func (o *PatchedValueUpdate) GetExternalFqnOk() (*string, bool)`

GetExternalFqnOk returns a tuple with the ExternalFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqn

`func (o *PatchedValueUpdate) SetExternalFqn(v string)`

SetExternalFqn sets ExternalFqn field to given value.

### HasExternalFqn

`func (o *PatchedValueUpdate) HasExternalFqn() bool`

HasExternalFqn returns a boolean if a field has been set.

### GetExternalFilter

`func (o *PatchedValueUpdate) GetExternalFilter() string`

GetExternalFilter returns the ExternalFilter field if non-nil, zero value otherwise.

### GetExternalFilterOk

`func (o *PatchedValueUpdate) GetExternalFilterOk() (*string, bool)`

GetExternalFilterOk returns a tuple with the ExternalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilter

`func (o *PatchedValueUpdate) SetExternalFilter(v string)`

SetExternalFilter sets ExternalFilter field to given value.

### HasExternalFilter

`func (o *PatchedValueUpdate) HasExternalFilter() bool`

HasExternalFilter returns a boolean if a field has been set.

### GetInternalValue

`func (o *PatchedValueUpdate) GetInternalValue() string`

GetInternalValue returns the InternalValue field if non-nil, zero value otherwise.

### GetInternalValueOk

`func (o *PatchedValueUpdate) GetInternalValueOk() (*string, bool)`

GetInternalValueOk returns a tuple with the InternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalValue

`func (o *PatchedValueUpdate) SetInternalValue(v string)`

SetInternalValue sets InternalValue field to given value.

### HasInternalValue

`func (o *PatchedValueUpdate) HasInternalValue() bool`

HasInternalValue returns a boolean if a field has been set.

### SetInternalValueNil

`func (o *PatchedValueUpdate) SetInternalValueNil(b bool)`

 SetInternalValueNil sets the value for InternalValue to be an explicit nil

### UnsetInternalValue
`func (o *PatchedValueUpdate) UnsetInternalValue()`

UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
### GetInterpolated

`func (o *PatchedValueUpdate) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *PatchedValueUpdate) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *PatchedValueUpdate) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *PatchedValueUpdate) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedValueUpdate) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedValueUpdate) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedValueUpdate) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedValueUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *PatchedValueUpdate) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *PatchedValueUpdate) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetValue

`func (o *PatchedValueUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedValueUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedValueUpdate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedValueUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchedValueUpdate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchedValueUpdate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCreatedAt

`func (o *PatchedValueUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedValueUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedValueUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedValueUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedValueUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedValueUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedValueUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedValueUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedValueUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedValueUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


