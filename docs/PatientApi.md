# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2PatientAccountsList**](PatientApi.md#ApiV2PatientAccountsList) | **Get** /api/v2/patient/{patient_uid}/accounts | 
[**ApiV2PatientAccountsRetrieve**](PatientApi.md#ApiV2PatientAccountsRetrieve) | **Get** /api/v2/patient/{patient_uid}/accounts/{id} | 
[**ApiV2PatientCardsList**](PatientApi.md#ApiV2PatientCardsList) | **Get** /api/v2/patient/{patient_uid}/cards | 
[**ApiV2PatientCardsRetrieve**](PatientApi.md#ApiV2PatientCardsRetrieve) | **Get** /api/v2/patient/{patient_uid}/cards/{id} | 
[**ApiV2PatientCertificatesList**](PatientApi.md#ApiV2PatientCertificatesList) | **Get** /api/v2/patient/{patient_uid}/certificates | 
[**ApiV2PatientCertificatesRetrieve**](PatientApi.md#ApiV2PatientCertificatesRetrieve) | **Get** /api/v2/patient/{patient_uid}/certificates/{id} | 
[**ApiV2PatientContractsList**](PatientApi.md#ApiV2PatientContractsList) | **Get** /api/v2/patient/{patient_uid}/contracts | 
[**ApiV2PatientContractsRetrieve**](PatientApi.md#ApiV2PatientContractsRetrieve) | **Get** /api/v2/patient/{patient_uid}/contracts/{id} | 
[**ApiV2PatientLabresultsList**](PatientApi.md#ApiV2PatientLabresultsList) | **Get** /api/v2/patient/{patient_uid}/labresults | 
[**ApiV2PatientLabresultsRetrieve**](PatientApi.md#ApiV2PatientLabresultsRetrieve) | **Get** /api/v2/patient/{patient_uid}/labresults/{uid} | 
[**ApiV2PatientList**](PatientApi.md#ApiV2PatientList) | **Get** /api/v2/patient | 
[**ApiV2PatientMedicalplansCreate**](PatientApi.md#ApiV2PatientMedicalplansCreate) | **Post** /api/v2/patient/{patient_uid}/medicalplans | 
[**ApiV2PatientMedicalplansList**](PatientApi.md#ApiV2PatientMedicalplansList) | **Get** /api/v2/patient/{patient_uid}/medicalplans | 
[**ApiV2PatientMedicalplansPartialUpdate**](PatientApi.md#ApiV2PatientMedicalplansPartialUpdate) | **Patch** /api/v2/patient/{patient_uid}/medicalplans/{id} | 
[**ApiV2PatientMedicalplansRetrieve**](PatientApi.md#ApiV2PatientMedicalplansRetrieve) | **Get** /api/v2/patient/{patient_uid}/medicalplans/{id} | 
[**ApiV2PatientMedicalplansUpdate**](PatientApi.md#ApiV2PatientMedicalplansUpdate) | **Put** /api/v2/patient/{patient_uid}/medicalplans/{id} | 
[**ApiV2PatientPartialUpdate**](PatientApi.md#ApiV2PatientPartialUpdate) | **Patch** /api/v2/patient/{uid} | 
[**ApiV2PatientPaymentsCreate**](PatientApi.md#ApiV2PatientPaymentsCreate) | **Post** /api/v2/patient/{patient_uid}/payments | 
[**ApiV2PatientPaymentsList**](PatientApi.md#ApiV2PatientPaymentsList) | **Get** /api/v2/patient/{patient_uid}/payments | 
[**ApiV2PatientPaymentsPartialUpdate**](PatientApi.md#ApiV2PatientPaymentsPartialUpdate) | **Patch** /api/v2/patient/{patient_uid}/payments/{id} | 
[**ApiV2PatientPaymentsRetrieve**](PatientApi.md#ApiV2PatientPaymentsRetrieve) | **Get** /api/v2/patient/{patient_uid}/payments/{id} | 
[**ApiV2PatientPaymentsUpdate**](PatientApi.md#ApiV2PatientPaymentsUpdate) | **Put** /api/v2/patient/{patient_uid}/payments/{id} | 
[**ApiV2PatientPoliciesList**](PatientApi.md#ApiV2PatientPoliciesList) | **Get** /api/v2/patient/{patient_uid}/policies | 
[**ApiV2PatientPoliciesRetrieve**](PatientApi.md#ApiV2PatientPoliciesRetrieve) | **Get** /api/v2/patient/{patient_uid}/policies/{id} | 
[**ApiV2PatientPreordersCreate**](PatientApi.md#ApiV2PatientPreordersCreate) | **Post** /api/v2/patient/{patient_uid}/preorders | 
[**ApiV2PatientPreordersDestroy**](PatientApi.md#ApiV2PatientPreordersDestroy) | **Delete** /api/v2/patient/{patient_uid}/preorders/{id} | 
[**ApiV2PatientPreordersList**](PatientApi.md#ApiV2PatientPreordersList) | **Get** /api/v2/patient/{patient_uid}/preorders | 
[**ApiV2PatientPreordersPartialUpdate**](PatientApi.md#ApiV2PatientPreordersPartialUpdate) | **Patch** /api/v2/patient/{patient_uid}/preorders/{id} | 
[**ApiV2PatientPreordersRetrieve**](PatientApi.md#ApiV2PatientPreordersRetrieve) | **Get** /api/v2/patient/{patient_uid}/preorders/{id} | 
[**ApiV2PatientPreordersUpdate**](PatientApi.md#ApiV2PatientPreordersUpdate) | **Put** /api/v2/patient/{patient_uid}/preorders/{id} | 
[**ApiV2PatientRetrieve**](PatientApi.md#ApiV2PatientRetrieve) | **Get** /api/v2/patient/{uid} | 
[**ApiV2PatientServicesList**](PatientApi.md#ApiV2PatientServicesList) | **Get** /api/v2/patient/{patient_uid}/services | 
[**ApiV2PatientServicesRetrieve**](PatientApi.md#ApiV2PatientServicesRetrieve) | **Get** /api/v2/patient/{patient_uid}/services/{id} | 
[**ApiV2PatientUpdate**](PatientApi.md#ApiV2PatientUpdate) | **Put** /api/v2/patient/{uid} | 
[**ApiV2PatientUpdatesList**](PatientApi.md#ApiV2PatientUpdatesList) | **Get** /api/v2/patient/updates | 
[**ApiV2PatientVisitsList**](PatientApi.md#ApiV2PatientVisitsList) | **Get** /api/v2/patient/{patient_uid}/visits | 
[**ApiV2PatientVisitsRetrieve**](PatientApi.md#ApiV2PatientVisitsRetrieve) | **Get** /api/v2/patient/{patient_uid}/visits/{id} | 
[**PatientHistoryCreate**](PatientApi.md#PatientHistoryCreate) | **Post** /patient/history/{id}/ | 

# **ApiV2PatientAccountsList**
> PaginatedPatientClientAccountList ApiV2PatientAccountsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientAccountsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientAccountsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientClientAccountList**](PaginatedPatientClientAccountList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientAccountsRetrieve**
> PatientClientAccount ApiV2PatientAccountsRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this client account. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientClientAccount**](PatientClientAccount.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientCardsList**
> PaginatedPatientCardList ApiV2PatientCardsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientCardsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientCardsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barcode** | **optional.Int32**|  | 
 **dateEnd** | **optional.String**|  | 
 **dateStart** | **optional.String**|  | 
 **format** | **optional.String**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientCardList**](PaginatedPatientCardList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/pdf, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientCardsRetrieve**
> PatientCard ApiV2PatientCardsRetrieve(ctx, id, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Карта осмотра. | 
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientCardsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientCardsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.String**|  | 

### Return type

[**PatientCard**](PatientCard.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/pdf, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientCertificatesList**
> PaginatedPatientCertificateList ApiV2PatientCertificatesList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientCertificatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientCertificatesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientCertificateList**](PaginatedPatientCertificateList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientCertificatesRetrieve**
> PatientCertificate ApiV2PatientCertificatesRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Сертификат. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientCertificate**](PatientCertificate.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientContractsList**
> PaginatedPatientContractList ApiV2PatientContractsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientContractsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientContractsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientContractList**](PaginatedPatientContractList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientContractsRetrieve**
> PatientContract ApiV2PatientContractsRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this договор. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientContract**](PatientContract.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientLabresultsList**
> PaginatedPatientLabOrderList ApiV2PatientLabresultsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientLabresultsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientLabresultsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barcode** | **optional.Int32**|  | 
 **dateEnd** | **optional.String**|  | 
 **dateStart** | **optional.String**|  | 
 **format** | **optional.String**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientLabOrderList**](PaginatedPatientLabOrderList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientLabresultsRetrieve**
> PatientLabOrder ApiV2PatientLabresultsRetrieve(ctx, patientUid, uid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
  **uid** | [**string**](.md)|  | 
 **optional** | ***PatientApiApiV2PatientLabresultsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientLabresultsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.String**|  | 

### Return type

[**PatientLabOrder**](PatientLabOrder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientList**
> PaginatedPatientDetailList ApiV2PatientList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatientApiApiV2PatientListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **birthDay** | **optional.String**|  | 
 **contract** | **optional.Int32**|  | 
 **email** | **optional.String**|  | 
 **firstName** | **optional.String**|  | 
 **idCardNumber** | **optional.String**|  | 
 **idCardSeries** | **optional.String**|  | 
 **lastName** | **optional.String**|  | 
 **midName** | **optional.String**|  | 
 **mobilePhone** | **optional.String**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **visit** | **optional.Int32**|  | 

### Return type

[**PaginatedPatientDetailList**](PaginatedPatientDetailList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientMedicalplansCreate**
> PatientMedicalPlan ApiV2PatientMedicalplansCreate(ctx, body, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatientMedicalPlan**](PatientMedicalPlan.md)|  | 
  **patientUid** | **string**|  | 

### Return type

[**PatientMedicalPlan**](PatientMedicalPlan.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientMedicalplansList**
> PaginatedPatientMedicalPlanList ApiV2PatientMedicalplansList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientMedicalplansListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientMedicalplansListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientMedicalPlanList**](PaginatedPatientMedicalPlanList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientMedicalplansPartialUpdate**
> PatientMedicalPlan ApiV2PatientMedicalplansPartialUpdate(ctx, id, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this план обследования. | 
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientMedicalplansPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientMedicalplansPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PatchedPatientMedicalPlan**](PatchedPatientMedicalPlan.md)|  | 

### Return type

[**PatientMedicalPlan**](PatientMedicalPlan.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientMedicalplansRetrieve**
> PatientMedicalPlan ApiV2PatientMedicalplansRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this план обследования. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientMedicalPlan**](PatientMedicalPlan.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientMedicalplansUpdate**
> PatientMedicalPlan ApiV2PatientMedicalplansUpdate(ctx, body, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatientMedicalPlan**](PatientMedicalPlan.md)|  | 
  **id** | **int32**| A unique integer value identifying this план обследования. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientMedicalPlan**](PatientMedicalPlan.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPartialUpdate**
> PatientDetail ApiV2PatientPartialUpdate(ctx, uid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)|  | 
 **optional** | ***PatientApiApiV2PatientPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedPatientDetail**](PatchedPatientDetail.md)|  | 

### Return type

[**PatientDetail**](PatientDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPaymentsCreate**
> PatientPayment ApiV2PatientPaymentsCreate(ctx, body, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatientPayment**](PatientPayment.md)|  | 
  **patientUid** | **string**|  | 

### Return type

[**PatientPayment**](PatientPayment.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPaymentsList**
> PaginatedPatientPaymentList ApiV2PatientPaymentsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPaymentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPaymentsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateEnd** | **optional.String**|  | 
 **dateStart** | **optional.String**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientPaymentList**](PaginatedPatientPaymentList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPaymentsPartialUpdate**
> PatientPayment ApiV2PatientPaymentsPartialUpdate(ctx, id, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this кассовый ордер. | 
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPaymentsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPaymentsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PatchedPatientPayment**](PatchedPatientPayment.md)|  | 

### Return type

[**PatientPayment**](PatientPayment.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPaymentsRetrieve**
> PatientPayment ApiV2PatientPaymentsRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this кассовый ордер. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientPayment**](PatientPayment.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPaymentsUpdate**
> PatientPayment ApiV2PatientPaymentsUpdate(ctx, body, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatientPayment**](PatientPayment.md)|  | 
  **id** | **int32**| A unique integer value identifying this кассовый ордер. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientPayment**](PatientPayment.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPoliciesList**
> PaginatedPatientInsurancePolicyList ApiV2PatientPoliciesList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientInsurancePolicyList**](PaginatedPatientInsurancePolicyList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPoliciesRetrieve**
> PatientInsurancePolicy ApiV2PatientPoliciesRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this страховой полис. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientInsurancePolicy**](PatientInsurancePolicy.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersCreate**
> PatientPreorder ApiV2PatientPreordersCreate(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPreordersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPreordersCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatientPreorder**](PatientPreorder.md)|  | 

### Return type

[**PatientPreorder**](PatientPreorder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersDestroy**
> ApiV2PatientPreordersDestroy(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this предзаказ. | 
  **patientUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersList**
> PaginatedPatientPreorderList ApiV2PatientPreordersList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPreordersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPreordersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientPreorderList**](PaginatedPatientPreorderList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersPartialUpdate**
> PatientPreorder ApiV2PatientPreordersPartialUpdate(ctx, id, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this предзаказ. | 
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPreordersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPreordersPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PatchedPatientPreorder**](PatchedPatientPreorder.md)|  | 

### Return type

[**PatientPreorder**](PatientPreorder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersRetrieve**
> PatientPreorder ApiV2PatientPreordersRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this предзаказ. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientPreorder**](PatientPreorder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientPreordersUpdate**
> PatientPreorder ApiV2PatientPreordersUpdate(ctx, id, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this предзаказ. | 
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientPreordersUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientPreordersUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PatientPreorder**](PatientPreorder.md)|  | 

### Return type

[**PatientPreorder**](PatientPreorder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientRetrieve**
> PatientDetail ApiV2PatientRetrieve(ctx, uid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)|  | 

### Return type

[**PatientDetail**](PatientDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientServicesList**
> PaginatedPatientOrderedServiceList ApiV2PatientServicesList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientServicesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientServicesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientOrderedServiceList**](PaginatedPatientOrderedServiceList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientServicesRetrieve**
> PatientOrderedService ApiV2PatientServicesRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this услуга. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientOrderedService**](PatientOrderedService.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientUpdate**
> PatientDetail ApiV2PatientUpdate(ctx, body, uid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatientDetail**](PatientDetail.md)|  | 
  **uid** | [**string**](.md)|  | 

### Return type

[**PatientDetail**](PatientDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientUpdatesList**
> PaginatedPatientDetailList ApiV2PatientUpdatesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatientApiApiV2PatientUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientUpdatesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **timedelta** | **optional.Int32**| Delta in seconds. | 

### Return type

[**PaginatedPatientDetailList**](PaginatedPatientDetailList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientVisitsList**
> PaginatedPatientVisitList ApiV2PatientVisitsList(ctx, patientUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patientUid** | **string**|  | 
 **optional** | ***PatientApiApiV2PatientVisitsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatientApiApiV2PatientVisitsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barcode** | **optional.Int32**|  | 
 **dateEnd** | **optional.String**|  | 
 **dateStart** | **optional.String**|  | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedPatientVisitList**](PaginatedPatientVisitList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2PatientVisitsRetrieve**
> PatientVisit ApiV2PatientVisitsRetrieve(ctx, id, patientUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this прием. | 
  **patientUid** | **string**|  | 

### Return type

[**PatientVisit**](PatientVisit.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatientHistoryCreate**
> PatientHistoryCreate(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

