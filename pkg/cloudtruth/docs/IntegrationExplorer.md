# IntegrationExplorer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqn** | **string** |  | 
**NodeType** | [**NodeTypeEnum**](NodeTypeEnum.md) |  | 
**Secret** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ContentData** | Pointer to **NullableString** |  | [optional] 
**ContentSize** | Pointer to **NullableInt32** |  | [optional] 
**ContentKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIntegrationExplorer

`func NewIntegrationExplorer(fqn string, nodeType NodeTypeEnum, ) *IntegrationExplorer`

NewIntegrationExplorer instantiates a new IntegrationExplorer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationExplorerWithDefaults

`func NewIntegrationExplorerWithDefaults() *IntegrationExplorer`

NewIntegrationExplorerWithDefaults instantiates a new IntegrationExplorer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqn

`func (o *IntegrationExplorer) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *IntegrationExplorer) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *IntegrationExplorer) SetFqn(v string)`

SetFqn sets Fqn field to given value.


### GetNodeType

`func (o *IntegrationExplorer) GetNodeType() NodeTypeEnum`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *IntegrationExplorer) GetNodeTypeOk() (*NodeTypeEnum, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *IntegrationExplorer) SetNodeType(v NodeTypeEnum)`

SetNodeType sets NodeType field to given value.


### GetSecret

`func (o *IntegrationExplorer) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IntegrationExplorer) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IntegrationExplorer) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IntegrationExplorer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetName

`func (o *IntegrationExplorer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationExplorer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationExplorer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationExplorer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *IntegrationExplorer) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *IntegrationExplorer) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *IntegrationExplorer) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *IntegrationExplorer) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *IntegrationExplorer) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *IntegrationExplorer) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentData

`func (o *IntegrationExplorer) GetContentData() string`

GetContentData returns the ContentData field if non-nil, zero value otherwise.

### GetContentDataOk

`func (o *IntegrationExplorer) GetContentDataOk() (*string, bool)`

GetContentDataOk returns a tuple with the ContentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentData

`func (o *IntegrationExplorer) SetContentData(v string)`

SetContentData sets ContentData field to given value.

### HasContentData

`func (o *IntegrationExplorer) HasContentData() bool`

HasContentData returns a boolean if a field has been set.

### SetContentDataNil

`func (o *IntegrationExplorer) SetContentDataNil(b bool)`

 SetContentDataNil sets the value for ContentData to be an explicit nil

### UnsetContentData
`func (o *IntegrationExplorer) UnsetContentData()`

UnsetContentData ensures that no value is present for ContentData, not even an explicit nil
### GetContentSize

`func (o *IntegrationExplorer) GetContentSize() int32`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *IntegrationExplorer) GetContentSizeOk() (*int32, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *IntegrationExplorer) SetContentSize(v int32)`

SetContentSize sets ContentSize field to given value.

### HasContentSize

`func (o *IntegrationExplorer) HasContentSize() bool`

HasContentSize returns a boolean if a field has been set.

### SetContentSizeNil

`func (o *IntegrationExplorer) SetContentSizeNil(b bool)`

 SetContentSizeNil sets the value for ContentSize to be an explicit nil

### UnsetContentSize
`func (o *IntegrationExplorer) UnsetContentSize()`

UnsetContentSize ensures that no value is present for ContentSize, not even an explicit nil
### GetContentKeys

`func (o *IntegrationExplorer) GetContentKeys() []string`

GetContentKeys returns the ContentKeys field if non-nil, zero value otherwise.

### GetContentKeysOk

`func (o *IntegrationExplorer) GetContentKeysOk() (*[]string, bool)`

GetContentKeysOk returns a tuple with the ContentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentKeys

`func (o *IntegrationExplorer) SetContentKeys(v []string)`

SetContentKeys sets ContentKeys field to given value.

### HasContentKeys

`func (o *IntegrationExplorer) HasContentKeys() bool`

HasContentKeys returns a boolean if a field has been set.

### SetContentKeysNil

`func (o *IntegrationExplorer) SetContentKeysNil(b bool)`

 SetContentKeysNil sets the value for ContentKeys to be an explicit nil

### UnsetContentKeys
`func (o *IntegrationExplorer) UnsetContentKeys()`

UnsetContentKeys ensures that no value is present for ContentKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


