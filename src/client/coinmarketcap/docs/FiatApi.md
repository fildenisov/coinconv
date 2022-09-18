# \FiatApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1FiatMap**](FiatApi.md#GetV1FiatMap) | **Get** /v1/fiat/map | CoinMarketCap ID Map


# **GetV1FiatMap**
> FiatMapResponseModel GetV1FiatMap(ctx, optional)
CoinMarketCap ID Map

Returns a mapping of all supported fiat currencies to unique CoinMarketCap ids. Per our Best Practices we recommend utilizing CMC ID instead of currency symbols to securely identify assets with our other endpoints and in your own application logic.     **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Mapping data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request no matter query size.   **CMC equivalent pages:** No equivalent, this data is only available via API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FiatApiGetV1FiatMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiatApiGetV1FiatMapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | 
 **sort** | **optional.String**| What field to sort the list by. | [default to id]
 **includeMetals** | **optional.Bool**| Pass &#x60;true&#x60; to include precious metals. | [default to false]

### Return type

[**FiatMapResponseModel**](Fiat Map - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

