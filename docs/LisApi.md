# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2LisOrderCreate**](LisApi.md#ApiV2LisOrderCreate) | **Post** /api/v2/lis/order | 
[**ApiV2LisResultsList**](LisApi.md#ApiV2LisResultsList) | **Get** /api/v2/lis/results | 
[**ApiV2LisResultsRetrieve**](LisApi.md#ApiV2LisResultsRetrieve) | **Get** /api/v2/lis/results/{id} | 
[**ApiV2LisSamplingCreate**](LisApi.md#ApiV2LisSamplingCreate) | **Post** /api/v2/lis/sampling | 
[**ApiV2LisServicesList**](LisApi.md#ApiV2LisServicesList) | **Get** /api/v2/lis/services | 

# **ApiV2LisOrderCreate**
> Order ApiV2LisOrderCreate(ctx, body)


Creates visit and lab orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Order**](Order.md)|  | 

### Return type

[**Order**](Order.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2LisResultsList**
> []LabOrder ApiV2LisResultsList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LisApiApiV2LisResultsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LisApiApiV2LisResultsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultId** | **optional.Bool**|  | 
 **specimen** | **optional.Bool**|  | 
 **taskId** | **optional.Bool**|  | 

### Return type

[**[]LabOrder**](LabOrder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2LisResultsRetrieve**
> LabOrder ApiV2LisResultsRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this заказ. | 

### Return type

[**LabOrder**](LabOrder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2LisSamplingCreate**
> Sampling ApiV2LisSamplingCreate(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LisApiApiV2LisSamplingCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LisApiApiV2LisSamplingCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Sampling**](Sampling.md)|  | 

### Return type

[**Sampling**](Sampling.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2LisServicesList**
> []ServiceInformation ApiV2LisServicesList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ServiceInformation**](ServiceInformation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

