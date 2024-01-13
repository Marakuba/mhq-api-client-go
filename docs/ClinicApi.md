# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ClinicBranchesList**](ClinicApi.md#ApiV2ClinicBranchesList) | **Get** /api/v2/clinic/branches | 
[**ApiV2ClinicBranchesRetrieve**](ClinicApi.md#ApiV2ClinicBranchesRetrieve) | **Get** /api/v2/clinic/branches/{id} | 
[**ApiV2ClinicCertificatesCreate**](ClinicApi.md#ApiV2ClinicCertificatesCreate) | **Post** /api/v2/clinic/certificates | 
[**ApiV2ClinicCertificatesList**](ClinicApi.md#ApiV2ClinicCertificatesList) | **Get** /api/v2/clinic/certificates | 
[**ApiV2ClinicDiscountsList**](ClinicApi.md#ApiV2ClinicDiscountsList) | **Get** /api/v2/clinic/discounts | 
[**ApiV2ClinicDiscountsRetrieve**](ClinicApi.md#ApiV2ClinicDiscountsRetrieve) | **Get** /api/v2/clinic/discounts/{id} | 
[**ApiV2ClinicLegalentitiesList**](ClinicApi.md#ApiV2ClinicLegalentitiesList) | **Get** /api/v2/clinic/legalentities | 
[**ApiV2ClinicLegalentitiesRetrieve**](ClinicApi.md#ApiV2ClinicLegalentitiesRetrieve) | **Get** /api/v2/clinic/legalentities/{id} | 
[**ApiV2ClinicPromotionsList**](ClinicApi.md#ApiV2ClinicPromotionsList) | **Get** /api/v2/clinic/promotions | 
[**ApiV2ClinicPromotionsRetrieve**](ClinicApi.md#ApiV2ClinicPromotionsRetrieve) | **Get** /api/v2/clinic/promotions/{id} | 
[**ApiV2ClinicServicegroupsList**](ClinicApi.md#ApiV2ClinicServicegroupsList) | **Get** /api/v2/clinic/servicegroups | 
[**ApiV2ClinicServicegroupsRetrieve**](ClinicApi.md#ApiV2ClinicServicegroupsRetrieve) | **Get** /api/v2/clinic/servicegroups/{id} | 
[**ApiV2ClinicServicesList**](ClinicApi.md#ApiV2ClinicServicesList) | **Get** /api/v2/clinic/services | 
[**ApiV2ClinicServicesRetrieve**](ClinicApi.md#ApiV2ClinicServicesRetrieve) | **Get** /api/v2/clinic/services/{id} | 
[**ApiV2ClinicStaffList**](ClinicApi.md#ApiV2ClinicStaffList) | **Get** /api/v2/clinic/staff | 
[**ApiV2ClinicStaffRetrieve**](ClinicApi.md#ApiV2ClinicStaffRetrieve) | **Get** /api/v2/clinic/staff/{id} | 
[**ApiV2ClinicStaffgroupsList**](ClinicApi.md#ApiV2ClinicStaffgroupsList) | **Get** /api/v2/clinic/staffgroups | 
[**ApiV2ClinicStaffgroupsRetrieve**](ClinicApi.md#ApiV2ClinicStaffgroupsRetrieve) | **Get** /api/v2/clinic/staffgroups/{id} | 
[**ApiV2ClinicStatesList**](ClinicApi.md#ApiV2ClinicStatesList) | **Get** /api/v2/clinic/states | 
[**ApiV2ClinicStatesRetrieve**](ClinicApi.md#ApiV2ClinicStatesRetrieve) | **Get** /api/v2/clinic/states/{id} | 

# **ApiV2ClinicBranchesList**
> PaginatedBranchList ApiV2ClinicBranchesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicBranchesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicBranchesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActive** | **optional.Bool**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedBranchList**](PaginatedBranchList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicBranchesRetrieve**
> Branch ApiV2ClinicBranchesRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this филиал. | 

### Return type

[**Branch**](Branch.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicCertificatesCreate**
> Certificate ApiV2ClinicCertificatesCreate(ctx, body)


Returns list of actual specializations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Certificate**](Certificate.md)|  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicCertificatesList**
> PaginatedCertificateList ApiV2ClinicCertificatesList(ctx, numberText, optional)


Returns list of actual specializations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberText** | **string**|  | 
 **optional** | ***ClinicApiApiV2ClinicCertificatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicCertificatesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedCertificateList**](PaginatedCertificateList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicDiscountsList**
> PaginatedDiscountList ApiV2ClinicDiscountsList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicDiscountsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicDiscountsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedDiscountList**](PaginatedDiscountList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicDiscountsRetrieve**
> Discount ApiV2ClinicDiscountsRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this скидка. | 

### Return type

[**Discount**](Discount.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicLegalentitiesList**
> PaginatedLegalEntityList ApiV2ClinicLegalentitiesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicLegalentitiesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicLegalentitiesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedLegalEntityList**](PaginatedLegalEntityList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicLegalentitiesRetrieve**
> LegalEntity ApiV2ClinicLegalentitiesRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this юридическое лицо. | 

### Return type

[**LegalEntity**](LegalEntity.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicPromotionsList**
> PaginatedPromotionList ApiV2ClinicPromotionsList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicPromotionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicPromotionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedPromotionList**](PaginatedPromotionList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicPromotionsRetrieve**
> Promotion ApiV2ClinicPromotionsRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this комплексная услуга / акция. | 

### Return type

[**Promotion**](Promotion.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicServicegroupsList**
> PaginatedBaseServiceGroupList ApiV2ClinicServicegroupsList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicServicegroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicServicegroupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedBaseServiceGroupList**](PaginatedBaseServiceGroupList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicServicegroupsRetrieve**
> BaseServiceGroup ApiV2ClinicServicegroupsRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this группа услуг. | 

### Return type

[**BaseServiceGroup**](BaseServiceGroup.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicServicesList**
> PaginatedBaseServiceList ApiV2ClinicServicesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicServicesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicServicesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseGroup** | **optional.Int32**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **patient** | [**optional.Interface of string**](.md)|  | 
 **search** | **optional.String**| A search term. | 
 **type_** | **optional.String**| * &#x60;cons&#x60; - консультативная услуга * &#x60;man&#x60; - манипуляция * &#x60;exam&#x60; - обследование * &#x60;lab&#x60; - лабораторная услуга * &#x60;diag&#x60; - диагностические исследования * &#x60;article&#x60; - товар * &#x60;material&#x60; - материал * &#x60;group&#x60; - группа * &#x60;addval&#x60; - дополнительная услуга * &#x60;archive&#x60; - архив | 

### Return type

[**PaginatedBaseServiceList**](PaginatedBaseServiceList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicServicesRetrieve**
> BaseService ApiV2ClinicServicesRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this услуга клиники. | 

### Return type

[**BaseService**](BaseService.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStaffList**
> PaginatedStaffDetailList ApiV2ClinicStaffList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicStaffListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicStaffListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gender** | **optional.String**| * &#x60;М&#x60; - Мужской * &#x60;Ж&#x60; - Женский | 
 **groups** | [**optional.Interface of []int32**](int32.md)|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedStaffDetailList**](PaginatedStaffDetailList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStaffRetrieve**
> StaffDetail ApiV2ClinicStaffRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Staff. | 

### Return type

[**StaffDetail**](StaffDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStaffgroupsList**
> PaginatedStaffGroupList ApiV2ClinicStaffgroupsList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicStaffgroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicStaffgroupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 

### Return type

[**PaginatedStaffGroupList**](PaginatedStaffGroupList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStaffgroupsRetrieve**
> StaffGroup ApiV2ClinicStaffgroupsRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Группа сотрудников. | 

### Return type

[**StaffGroup**](StaffGroup.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStatesList**
> PaginatedStateList ApiV2ClinicStatesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClinicApiApiV2ClinicStatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClinicApiApiV2ClinicStatesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **branch** | **optional.Int32**|  | 
 **legalEntity** | **optional.Int32**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **search** | **optional.String**| A search term. | 
 **type_** | **optional.String**| * &#x60;m&#x60; - Медицинское учреждение * &#x60;b&#x60; - Собственный филиал * &#x60;i&#x60; - Страховая компания * &#x60;j&#x60; - Прочее юридическое лицо * &#x60;p&#x60; - Медицинский пункт * &#x60;o&#x60; - Плательщик ОМС | 

### Return type

[**PaginatedStateList**](PaginatedStateList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2ClinicStatesRetrieve**
> State ApiV2ClinicStatesRetrieve(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this организация. | 

### Return type

[**State**](State.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

