# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocsApiV2SchemaRetrieve**](DocsApi.md#DocsApiV2SchemaRetrieve) | **Get** /docs/api/v2/schema | 

# **DocsApiV2SchemaRetrieve**
> map[string]Object DocsApiV2SchemaRetrieve(ctx, optional)


OpenApi3 schema for this API. Format can be selected via content negotiation.  - YAML: application/vnd.oai.openapi - JSON: application/vnd.oai.openapi+json

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocsApiDocsApiV2SchemaRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocsApiDocsApiV2SchemaRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**|  | 
 **lang** | **optional.String**|  | 

### Return type

**map[string]Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

