# \IntegrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsAwsCreate**](IntegrationsApi.md#IntegrationsAwsCreate) | **Post** /api/v1/integrations/aws/ | Establishes an AWS Integration.
[**IntegrationsAwsDestroy**](IntegrationsApi.md#IntegrationsAwsDestroy) | **Delete** /api/v1/integrations/aws/{id}/ | Delete an AWS integration.
[**IntegrationsAwsList**](IntegrationsApi.md#IntegrationsAwsList) | **Get** /api/v1/integrations/aws/ | 
[**IntegrationsAwsPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPartialUpdate) | **Patch** /api/v1/integrations/aws/{id}/ | 
[**IntegrationsAwsPullsCreate**](IntegrationsApi.md#IntegrationsAwsPullsCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pulls/ | 
[**IntegrationsAwsPullsDestroy**](IntegrationsApi.md#IntegrationsAwsPullsDestroy) | **Delete** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsList**](IntegrationsApi.md#IntegrationsAwsPullsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/ | 
[**IntegrationsAwsPullsPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPullsPartialUpdate) | **Patch** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsSyncCreate**](IntegrationsApi.md#IntegrationsAwsPullsSyncCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/sync/ | 
[**IntegrationsAwsPullsTasksList**](IntegrationsApi.md#IntegrationsAwsPullsTasksList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/ | 
[**IntegrationsAwsPullsTasksRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsTasksRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{id}/ | 
[**IntegrationsAwsPullsTasksStepsList**](IntegrationsApi.md#IntegrationsAwsPullsTasksStepsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{awspulltask_pk}/steps/ | 
[**IntegrationsAwsPullsTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsTasksStepsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{awspulltask_pk}/steps/{id}/ | 
[**IntegrationsAwsPullsUpdate**](IntegrationsApi.md#IntegrationsAwsPullsUpdate) | **Put** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPushesCreate**](IntegrationsApi.md#IntegrationsAwsPushesCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pushes/ | 
[**IntegrationsAwsPushesDestroy**](IntegrationsApi.md#IntegrationsAwsPushesDestroy) | **Delete** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesList**](IntegrationsApi.md#IntegrationsAwsPushesList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/ | 
[**IntegrationsAwsPushesPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPushesPartialUpdate) | **Patch** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesSyncCreate**](IntegrationsApi.md#IntegrationsAwsPushesSyncCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/sync/ | 
[**IntegrationsAwsPushesTasksList**](IntegrationsApi.md#IntegrationsAwsPushesTasksList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/ | 
[**IntegrationsAwsPushesTasksRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesTasksRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{id}/ | 
[**IntegrationsAwsPushesTasksStepsList**](IntegrationsApi.md#IntegrationsAwsPushesTasksStepsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{awspushtask_pk}/steps/ | 
[**IntegrationsAwsPushesTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesTasksStepsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{awspushtask_pk}/steps/{id}/ | 
[**IntegrationsAwsPushesUpdate**](IntegrationsApi.md#IntegrationsAwsPushesUpdate) | **Put** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsRetrieve**](IntegrationsApi.md#IntegrationsAwsRetrieve) | **Get** /api/v1/integrations/aws/{id}/ | Get details of an AWS Integration.
[**IntegrationsAwsUpdate**](IntegrationsApi.md#IntegrationsAwsUpdate) | **Put** /api/v1/integrations/aws/{id}/ | 
[**IntegrationsExploreList**](IntegrationsApi.md#IntegrationsExploreList) | **Get** /api/v1/integrations/explore/ | Retrieve third-party integration data for the specified FQN.
[**IntegrationsGithubCreate**](IntegrationsApi.md#IntegrationsGithubCreate) | **Post** /api/v1/integrations/github/ | Establishes a GitHub Integration.
[**IntegrationsGithubDestroy**](IntegrationsApi.md#IntegrationsGithubDestroy) | **Delete** /api/v1/integrations/github/{id}/ | Delete a GitHub integration.
[**IntegrationsGithubList**](IntegrationsApi.md#IntegrationsGithubList) | **Get** /api/v1/integrations/github/ | 
[**IntegrationsGithubRetrieve**](IntegrationsApi.md#IntegrationsGithubRetrieve) | **Get** /api/v1/integrations/github/{id}/ | Get details of a GitHub Integration.



## IntegrationsAwsCreate

> AwsIntegration IntegrationsAwsCreate(ctx).AwsIntegrationCreate(awsIntegrationCreate).Execute()

Establishes an AWS Integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsIntegrationCreate := *openapiclient.NewAwsIntegrationCreate("AwsAccountId_example", []openapiclient.AwsRegionEnum{openapiclient.AwsRegionEnum("af-south-1")}, []openapiclient.AwsServiceEnum{openapiclient.AwsServiceEnum("s3")}, "AwsRoleName_example") // AwsIntegrationCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsCreate(context.Background()).AwsIntegrationCreate(awsIntegrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsCreate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsIntegrationCreate** | [**AwsIntegrationCreate**](AwsIntegrationCreate.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsDestroy

> IntegrationsAwsDestroy(ctx, id).InUse(inUse).Execute()

Delete an AWS integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    inUse := "inUse_example" // string | (Optional) Desired behavior if the integration has in-use values.  - `fail` will return HTTP error 409 if there are any values using the integration. - `leave` (default) will leave values in place and future queries may fail; you can control future value query behavior with the `lookup_error` query parameter on those requests. - `remove` will remove the all values using the integration when the integration is removed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsDestroy(context.Background(), id).InUse(inUse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inUse** | **string** | (Optional) Desired behavior if the integration has in-use values.  - &#x60;fail&#x60; will return HTTP error 409 if there are any values using the integration. - &#x60;leave&#x60; (default) will leave values in place and future queries may fail; you can control future value query behavior with the &#x60;lookup_error&#x60; query parameter on those requests. - &#x60;remove&#x60; will remove the all values using the integration when the integration is removed. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsList

> PaginatedAwsIntegrationList IntegrationsAwsList(ctx).AwsAccountId(awsAccountId).AwsRoleName(awsRoleName).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsAccountId := "awsAccountId_example" // string |  (optional)
    awsRoleName := "awsRoleName_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsList(context.Background()).AwsAccountId(awsAccountId).AwsRoleName(awsRoleName).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsList`: PaginatedAwsIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAccountId** | **string** |  | 
 **awsRoleName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAwsIntegrationList**](PaginatedAwsIntegrationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPartialUpdate

> AwsIntegration IntegrationsAwsPartialUpdate(ctx, id).PatchedAwsIntegration(patchedAwsIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    patchedAwsIntegration := *openapiclient.NewPatchedAwsIntegration() // PatchedAwsIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPartialUpdate(context.Background(), id).PatchedAwsIntegration(patchedAwsIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPartialUpdate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAwsIntegration** | [**PatchedAwsIntegration**](PatchedAwsIntegration.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsCreate

> AwsPull IntegrationsAwsPullsCreate(ctx, awsintegrationPk).AwsPull(awsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awsPull := *openapiclient.NewAwsPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), openapiclient.AwsRegionEnum("af-south-1"), openapiclient.AwsServiceEnum("s3"), "Resource_example") // AwsPull | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsCreate(context.Background(), awsintegrationPk).AwsPull(awsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsCreate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsPull** | [**AwsPull**](AwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsDestroy

> IntegrationsAwsPullsDestroy(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsDestroy(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsList

> PaginatedAwsPullList IntegrationsAwsPullsList(ctx, awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsList(context.Background(), awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsList`: PaginatedAwsPullList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAwsPullList**](PaginatedAwsPullList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsPartialUpdate

> AwsPull IntegrationsAwsPullsPartialUpdate(ctx, awsintegrationPk, id).PatchedAwsPull(patchedAwsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    patchedAwsPull := *openapiclient.NewPatchedAwsPull() // PatchedAwsPull |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsPartialUpdate(context.Background(), awsintegrationPk, id).PatchedAwsPull(patchedAwsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsPartialUpdate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAwsPull** | [**PatchedAwsPull**](PatchedAwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsRetrieve

> AwsPull IntegrationsAwsPullsRetrieve(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsRetrieve(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsRetrieve`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsSyncCreate

> IntegrationsAwsPullsSyncCreate(ctx, awsintegrationPk, id).AwsPull(awsPull).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    awsPull := *openapiclient.NewAwsPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), openapiclient.AwsRegionEnum("af-south-1"), openapiclient.AwsServiceEnum("s3"), "Resource_example") // AwsPull | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsSyncCreate(context.Background(), awsintegrationPk, id).AwsPull(awsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPull** | [**AwsPull**](AwsPull.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksList

> PaginatedAwsPushTaskList IntegrationsAwsPullsTasksList(ctx, awsintegrationPk, awspullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspullPk := TODO // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsTasksList(context.Background(), awsintegrationPk, awspullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksList`: PaginatedAwsPushTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspullPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAwsPushTaskList**](PaginatedAwsPushTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksRetrieve

> AwsPushTask IntegrationsAwsPullsTasksRetrieve(ctx, awsintegrationPk, awspullPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspullPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsTasksRetrieve(context.Background(), awsintegrationPk, awspullPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksRetrieve`: AwsPushTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspullPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AwsPushTask**](AwsPushTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksStepsList

> PaginatedAwsPushTaskStepList IntegrationsAwsPullsTasksStepsList(ctx, awsintegrationPk, awspullPk, awspulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspullPk := TODO // string | 
    awspulltaskPk := TODO // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsTasksStepsList(context.Background(), awsintegrationPk, awspullPk, awspulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksStepsList`: PaginatedAwsPushTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspullPk** | [**string**](.md) |  | 
**awspulltaskPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAwsPushTaskStepList**](PaginatedAwsPushTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksStepsRetrieve

> AwsPushTaskStep IntegrationsAwsPullsTasksStepsRetrieve(ctx, awsintegrationPk, awspullPk, awspulltaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspullPk := TODO // string | 
    awspulltaskPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve(context.Background(), awsintegrationPk, awspullPk, awspulltaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksStepsRetrieve`: AwsPushTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspullPk** | [**string**](.md) |  | 
**awspulltaskPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AwsPushTaskStep**](AwsPushTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsUpdate

> AwsPull IntegrationsAwsPullsUpdate(ctx, awsintegrationPk, id).AwsPull(awsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    awsPull := *openapiclient.NewAwsPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), openapiclient.AwsRegionEnum("af-south-1"), openapiclient.AwsServiceEnum("s3"), "Resource_example") // AwsPull | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPullsUpdate(context.Background(), awsintegrationPk, id).AwsPull(awsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsUpdate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPull** | [**AwsPull**](AwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesCreate

> AwsPush IntegrationsAwsPushesCreate(ctx, awsintegrationPk).AwsPush(awsPush).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awsPush := *openapiclient.NewAwsPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, openapiclient.AwsRegionEnum("af-south-1"), openapiclient.AwsServiceEnum("s3"), "Resource_example") // AwsPush | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesCreate(context.Background(), awsintegrationPk).AwsPush(awsPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesCreate`: AwsPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsPush** | [**AwsPush**](AwsPush.md) |  | 

### Return type

[**AwsPush**](AwsPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesDestroy

> IntegrationsAwsPushesDestroy(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesDestroy(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesList

> PaginatedAwsPushList IntegrationsAwsPushesList(ctx, awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesList(context.Background(), awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesList`: PaginatedAwsPushList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAwsPushList**](PaginatedAwsPushList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesPartialUpdate

> AwsPushUpdate IntegrationsAwsPushesPartialUpdate(ctx, awsintegrationPk, id).PatchedAwsPushUpdate(patchedAwsPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    patchedAwsPushUpdate := *openapiclient.NewPatchedAwsPushUpdate() // PatchedAwsPushUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesPartialUpdate(context.Background(), awsintegrationPk, id).PatchedAwsPushUpdate(patchedAwsPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesPartialUpdate`: AwsPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAwsPushUpdate** | [**PatchedAwsPushUpdate**](PatchedAwsPushUpdate.md) |  | 

### Return type

[**AwsPushUpdate**](AwsPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesRetrieve

> AwsPush IntegrationsAwsPushesRetrieve(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesRetrieve(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesRetrieve`: AwsPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AwsPush**](AwsPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesSyncCreate

> IntegrationsAwsPushesSyncCreate(ctx, awsintegrationPk, id).AwsPush(awsPush).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    awsPush := *openapiclient.NewAwsPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, openapiclient.AwsRegionEnum("af-south-1"), openapiclient.AwsServiceEnum("s3"), "Resource_example") // AwsPush | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesSyncCreate(context.Background(), awsintegrationPk, id).AwsPush(awsPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPush** | [**AwsPush**](AwsPush.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksList

> PaginatedAwsPushTaskList IntegrationsAwsPushesTasksList(ctx, awsintegrationPk, awspushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspushPk := TODO // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesTasksList(context.Background(), awsintegrationPk, awspushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksList`: PaginatedAwsPushTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspushPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAwsPushTaskList**](PaginatedAwsPushTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksRetrieve

> AwsPushTask IntegrationsAwsPushesTasksRetrieve(ctx, awsintegrationPk, awspushPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspushPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesTasksRetrieve(context.Background(), awsintegrationPk, awspushPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksRetrieve`: AwsPushTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspushPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AwsPushTask**](AwsPushTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksStepsList

> PaginatedAwsPushTaskStepList IntegrationsAwsPushesTasksStepsList(ctx, awsintegrationPk, awspushPk, awspushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspushPk := TODO // string | 
    awspushtaskPk := TODO // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesTasksStepsList(context.Background(), awsintegrationPk, awspushPk, awspushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksStepsList`: PaginatedAwsPushTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspushPk** | [**string**](.md) |  | 
**awspushtaskPk** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAwsPushTaskStepList**](PaginatedAwsPushTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksStepsRetrieve

> AwsPushTaskStep IntegrationsAwsPushesTasksStepsRetrieve(ctx, awsintegrationPk, awspushPk, awspushtaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    awspushPk := TODO // string | 
    awspushtaskPk := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve(context.Background(), awsintegrationPk, awspushPk, awspushtaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksStepsRetrieve`: AwsPushTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**awspushPk** | [**string**](.md) |  | 
**awspushtaskPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AwsPushTaskStep**](AwsPushTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesUpdate

> AwsPushUpdate IntegrationsAwsPushesUpdate(ctx, awsintegrationPk, id).AwsPushUpdate(awsPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := TODO // string | 
    id := TODO // string | 
    awsPushUpdate := *openapiclient.NewAwsPushUpdate("Name_example", []string{"Projects_example"}, []string{"Tags_example"}, "Resource_example") // AwsPushUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsPushesUpdate(context.Background(), awsintegrationPk, id).AwsPushUpdate(awsPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesUpdate`: AwsPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPushUpdate** | [**AwsPushUpdate**](AwsPushUpdate.md) |  | 

### Return type

[**AwsPushUpdate**](AwsPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsRetrieve

> AwsIntegration IntegrationsAwsRetrieve(ctx, id).RefreshStatus(refreshStatus).Execute()

Get details of an AWS Integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    refreshStatus := true // bool | Trigger a refresh of the integration status before returning the details. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsRetrieve(context.Background(), id).RefreshStatus(refreshStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsRetrieve`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStatus** | **bool** | Trigger a refresh of the integration status before returning the details. | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsUpdate

> AwsIntegration IntegrationsAwsUpdate(ctx, id).AwsIntegration(awsIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    awsIntegration := *openapiclient.NewAwsIntegration("Url_example", "Id_example", "Name_example", openapiclient.StatusEnum("unknown"), "StatusDetail_example", time.Now(), time.Now(), time.Now(), "Fqn_example", "Type_example", "AwsAccountId_example", []openapiclient.AwsRegionEnum{openapiclient.AwsRegionEnum("af-south-1")}, []openapiclient.AwsServiceEnum{openapiclient.AwsServiceEnum("s3")}, "AwsExternalId_example", "AwsRoleName_example") // AwsIntegration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsAwsUpdate(context.Background(), id).AwsIntegration(awsIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsUpdate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsIntegration** | [**AwsIntegration**](AwsIntegration.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsExploreList

> PaginatedIntegrationExplorerList IntegrationsExploreList(ctx).Fqn(fqn).Ordering(ordering).Page(page).PageSize(pageSize).Execute()

Retrieve third-party integration data for the specified FQN.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fqn := "fqn_example" // string | FQN (URL-like) for third-party integration. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsExploreList(context.Background()).Fqn(fqn).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsExploreList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsExploreList`: PaginatedIntegrationExplorerList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsExploreList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsExploreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fqn** | **string** | FQN (URL-like) for third-party integration. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedIntegrationExplorerList**](PaginatedIntegrationExplorerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubCreate

> GitHubIntegration IntegrationsGithubCreate(ctx).GitHubIntegrationCreate(gitHubIntegrationCreate).Execute()

Establishes a GitHub Integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gitHubIntegrationCreate := *openapiclient.NewGitHubIntegrationCreate(int32(123)) // GitHubIntegrationCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsGithubCreate(context.Background()).GitHubIntegrationCreate(gitHubIntegrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubCreate`: GitHubIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitHubIntegrationCreate** | [**GitHubIntegrationCreate**](GitHubIntegrationCreate.md) |  | 

### Return type

[**GitHubIntegration**](GitHubIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubDestroy

> IntegrationsGithubDestroy(ctx, id).InUse(inUse).Execute()

Delete a GitHub integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    inUse := "inUse_example" // string | (Optional) Desired behavior if the integration has in-use values.  - `fail` will return HTTP error 409 if there are any values using the integration. - `leave` (default) will leave values in place and future queries may fail; you can control future value query behavior with the `lookup_error` query parameter on those requests. - `remove` will remove the all values using the integration when the integration is removed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsGithubDestroy(context.Background(), id).InUse(inUse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inUse** | **string** | (Optional) Desired behavior if the integration has in-use values.  - &#x60;fail&#x60; will return HTTP error 409 if there are any values using the integration. - &#x60;leave&#x60; (default) will leave values in place and future queries may fail; you can control future value query behavior with the &#x60;lookup_error&#x60; query parameter on those requests. - &#x60;remove&#x60; will remove the all values using the integration when the integration is removed. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubList

> PaginatedGitHubIntegrationList IntegrationsGithubList(ctx).GhOrganizationSlug(ghOrganizationSlug).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ghOrganizationSlug := "ghOrganizationSlug_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsGithubList(context.Background()).GhOrganizationSlug(ghOrganizationSlug).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubList`: PaginatedGitHubIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ghOrganizationSlug** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedGitHubIntegrationList**](PaginatedGitHubIntegrationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubRetrieve

> GitHubIntegration IntegrationsGithubRetrieve(ctx, id).RefreshStatus(refreshStatus).Execute()

Get details of a GitHub Integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 
    refreshStatus := true // bool | Refresh the integration status before returning the details. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.IntegrationsGithubRetrieve(context.Background(), id).RefreshStatus(refreshStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubRetrieve`: GitHubIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStatus** | **bool** | Refresh the integration status before returning the details. | 

### Return type

[**GitHubIntegration**](GitHubIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

