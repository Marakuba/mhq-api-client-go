# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BookingAppointmentCreate**](BookingApi.md#ApiV2BookingAppointmentCreate) | **Post** /api/v2/booking/appointment | 
[**ApiV2BookingAppointmentDestroy**](BookingApi.md#ApiV2BookingAppointmentDestroy) | **Delete** /api/v2/booking/appointment/{uid} | 
[**ApiV2BookingBranchList**](BookingApi.md#ApiV2BookingBranchList) | **Get** /api/v2/booking/branch | 
[**ApiV2BookingServiceList**](BookingApi.md#ApiV2BookingServiceList) | **Get** /api/v2/booking/service | 
[**ApiV2BookingSpecializationList**](BookingApi.md#ApiV2BookingSpecializationList) | **Get** /api/v2/booking/specialization | 
[**ApiV2BookingStaffList**](BookingApi.md#ApiV2BookingStaffList) | **Get** /api/v2/booking/staff | 
[**ApiV2BookingStaffRetrieve**](BookingApi.md#ApiV2BookingStaffRetrieve) | **Get** /api/v2/booking/staff/{id} | 
[**ApiV2BookingStaffScheduleList**](BookingApi.md#ApiV2BookingStaffScheduleList) | **Get** /api/v2/booking/staff/{id}/schedule | 
[**ApiV2BookingTimeslotList**](BookingApi.md#ApiV2BookingTimeslotList) | **Get** /api/v2/booking/timeslot | 
[**ApiV2BookingTimeslotMonthList**](BookingApi.md#ApiV2BookingTimeslotMonthList) | **Get** /api/v2/booking/timeslot/month | 
[**ApiV2BookingTimeslotRecentList**](BookingApi.md#ApiV2BookingTimeslotRecentList) | **Get** /api/v2/booking/timeslot/recent | 

# **ApiV2BookingAppointmentCreate**
> CreatedAppointmentUuids ApiV2BookingAppointmentCreate(ctx, body)


Creates doctor appointments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Appointment**](Appointment.md)|  | 

### Return type

[**CreatedAppointmentUuids**](CreatedAppointmentUUIDs.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingAppointmentDestroy**
> ApiV2BookingAppointmentDestroy(ctx, uid)


Removes doctor appointment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingBranchList**
> PaginatedBranchList ApiV2BookingBranchList(ctx, optional)


Returns list of clinic branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingApiApiV2BookingBranchListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingApiApiV2BookingBranchListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **spec** | [**optional.Interface of []int32**](int32.md)| List of specialization IDs. | 
 **staff** | [**optional.Interface of []int32**](int32.md)| List of staff IDs. | 

### Return type

[**PaginatedBranchList**](PaginatedBranchList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingServiceList**
> PaginatedServiceList ApiV2BookingServiceList(ctx, optional)


Returns list of available services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingApiApiV2BookingServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingApiApiV2BookingServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **branch** | [**optional.Interface of []int32**](int32.md)| List of branch IDs. | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **spec** | [**optional.Interface of []int32**](int32.md)| List of specialization IDs. | 
 **staff** | [**optional.Interface of []int32**](int32.md)| List of staff IDs. | 

### Return type

[**PaginatedServiceList**](PaginatedServiceList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingSpecializationList**
> PaginatedStaffGroupList ApiV2BookingSpecializationList(ctx, optional)


Returns list of actual specializations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingApiApiV2BookingSpecializationListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingApiApiV2BookingSpecializationListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **optional.Bool**| Set to &#x27;true&#x27; to get list of all available specializations in the clinic | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**PaginatedStaffGroupList**](PaginatedStaffGroupList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingStaffList**
> PaginatedStaffList ApiV2BookingStaffList(ctx, optional)


Returns list of staff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingApiApiV2BookingStaffListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingApiApiV2BookingStaffListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **spec** | [**optional.Interface of []int32**](int32.md)| List of specialization IDs. | 

### Return type

[**PaginatedStaffList**](PaginatedStaffList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingStaffRetrieve**
> Staff ApiV2BookingStaffRetrieve(ctx, id)


Returns staff details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Staff. | 

### Return type

[**Staff**](Staff.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingStaffScheduleList**
> []StaffSchedule ApiV2BookingStaffScheduleList(ctx, id)


Returns staff schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this Staff. | 

### Return type

[**[]StaffSchedule**](StaffSchedule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingTimeslotList**
> []Timeslot ApiV2BookingTimeslotList(ctx, date, staff, optional)


Returns list of staff timeslots for a date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**|  | 
  **staff** | **int32**| Staff ID. | 
 **optional** | ***BookingApiApiV2BookingTimeslotListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingApiApiV2BookingTimeslotListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | [**optional.Interface of []int32**](int32.md)| List of branch IDs. | 

### Return type

[**[]Timeslot**](Timeslot.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingTimeslotMonthList**
> []TimeslotYearMonth ApiV2BookingTimeslotMonthList(ctx, month, staff, year)


Returns list of staff working dates for a month

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **month** | **float64**|  | 
  **staff** | **int32**| Staff ID. | 
  **year** | **float64**|  | 

### Return type

[**[]TimeslotYearMonth**](TimeslotYearMonth.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BookingTimeslotRecentList**
> []Timeslot ApiV2BookingTimeslotRecentList(ctx, staff)


Returns list of last ten staff timeslots

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **staff** | **int32**| Staff ID. | 

### Return type

[**[]Timeslot**](Timeslot.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

