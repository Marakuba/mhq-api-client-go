# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2EgiszProfileList**](EgiszApi.md#ApiV2EgiszProfileList) | **Get** /api/v2/egisz/profile | 
[**ApiV2EgiszRecordList**](EgiszApi.md#ApiV2EgiszRecordList) | **Get** /api/v2/egisz/record | 
[**ApiV2EgiszSendingattemptList**](EgiszApi.md#ApiV2EgiszSendingattemptList) | **Get** /api/v2/egisz/sendingattempt | 
[**ApiV2NsiDirectoriesEntitiesList**](EgiszApi.md#ApiV2NsiDirectoriesEntitiesList) | **Get** /api/v2/nsi/directories/{directory_oid}/entities | 
[**ApiV2NsiDirectoriesEntitiesRetrieve**](EgiszApi.md#ApiV2NsiDirectoriesEntitiesRetrieve) | **Get** /api/v2/nsi/directories/{directory_oid}/entities/{id} | 
[**ApiV2NsiDirectoriesList**](EgiszApi.md#ApiV2NsiDirectoriesList) | **Get** /api/v2/nsi/directories | 
[**ApiV2NsiDirectoriesRetrieve**](EgiszApi.md#ApiV2NsiDirectoriesRetrieve) | **Get** /api/v2/nsi/directories/{oid} | 

# **ApiV2EgiszProfileList**
> PaginatedProfileList ApiV2EgiszProfileList(ctx, optional)


Returns list of EGISZ profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EgiszApiApiV2EgiszProfileListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgiszApiApiV2EgiszProfileListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedProfileList**](PaginatedProfileList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EgiszRecordList**
> PaginatedRecordList ApiV2EgiszRecordList(ctx, optional)


Returns list of EGISZ records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EgiszApiApiV2EgiszRecordListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgiszApiApiV2EgiszRecordListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **patient** | **optional.Int32**|  | 
 **profile** | **optional.Int32**|  | 
 **recordType** | **optional.Int32**| * &#x60;1&#x60; - Информация о пациенте * &#x60;2&#x60; - Добавление завершенного случая * &#x60;3&#x60; - Регистрация случая * &#x60;4&#x60; - Добавление эпизода * &#x60;5&#x60; - Добавление медицинских записей * &#x60;6&#x60; - Закрытие случая * &#x60;7&#x60; - Изменение закрытого случая | 
 **status** | **optional.Int32**| * &#x60;0&#x60; - Новая запись * &#x60;1&#x60; - Успешно отправлена * &#x60;2&#x60; - Ошибка отправки | 

### Return type

[**PaginatedRecordList**](PaginatedRecordList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2EgiszSendingattemptList**
> PaginatedSendingAttemptList ApiV2EgiszSendingattemptList(ctx, optional)


Returns list of EGISZ sending attempts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EgiszApiApiV2EgiszSendingattemptListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgiszApiApiV2EgiszSendingattemptListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **record** | **optional.Int32**|  | 

### Return type

[**PaginatedSendingAttemptList**](PaginatedSendingAttemptList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2NsiDirectoriesEntitiesList**
> PaginatedEntityList ApiV2NsiDirectoriesEntitiesList(ctx, directoryOid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryOid** | **string**|  | 
 **optional** | ***EgiszApiApiV2NsiDirectoriesEntitiesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgiszApiApiV2NsiDirectoriesEntitiesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedEntityList**](PaginatedEntityList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2NsiDirectoriesEntitiesRetrieve**
> Entity ApiV2NsiDirectoriesEntitiesRetrieve(ctx, directoryOid, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryOid** | **string**|  | 
  **id** | **int32**| A unique integer value identifying this entity. | 

### Return type

[**Entity**](Entity.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2NsiDirectoriesList**
> []Directory ApiV2NsiDirectoriesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EgiszApiApiV2NsiDirectoriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgiszApiApiV2NsiDirectoriesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 

### Return type

[**[]Directory**](Directory.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2NsiDirectoriesRetrieve**
> Directory ApiV2NsiDirectoriesRetrieve(ctx, oid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**|  | 

### Return type

[**Directory**](Directory.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

