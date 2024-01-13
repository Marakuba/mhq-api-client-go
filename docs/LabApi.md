# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabFeedRetrieve**](LabApi.md#LabFeedRetrieve) | **Get** /lab/feed | 
[**LabPubBlankDocumentRetrieve**](LabApi.md#LabPubBlankDocumentRetrieve) | **Get** /lab/pub/blank/document/{uid} | 
[**LabPubBlankLabOrderRetrieve**](LabApi.md#LabPubBlankLabOrderRetrieve) | **Get** /lab/pub/blank/lab-order/{uid} | 
[**LabPubResultsRetrieve**](LabApi.md#LabPubResultsRetrieve) | **Get** /lab/pub/results/{barcode_pk} | 

# **LabFeedRetrieve**
> LabFeedRetrieve(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabPubBlankDocumentRetrieve**
> LabPubBlankDocumentRetrieve(ctx, uid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabPubBlankLabOrderRetrieve**
> LabPubBlankLabOrderRetrieve(ctx, uid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabPubResultsRetrieve**
> LabPubResultsRetrieve(ctx, barcodePk)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **barcodePk** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

