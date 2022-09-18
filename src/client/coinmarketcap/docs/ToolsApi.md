# \ToolsApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ToolsPostman**](ToolsApi.md#GetV1ToolsPostman) | **Get** /v1/tools/postman | Postman Conversion v1
[**GetV2ToolsPriceconversion**](ToolsApi.md#GetV2ToolsPriceconversion) | **Get** /v2/tools/price-conversion | Price Conversion v2


# **GetV1ToolsPostman**
> GetV1ToolsPostman(ctx, )
Postman Conversion v1

Convert APIs into postman format. You can reference the operation from <a href=\"https://coinmarketcap.com/alexandria/article/register-for-coinmarketcap-api\" target=\"_blank\"><b>this article</b></a>.         **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Free   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise    **Technical Notes**   - Set the env variables in the postman: {{baseUrl}}, {{API_KEY}} 

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

# **GetV2ToolsPriceconversion**
> ToolsPriceConversionResponseModel GetV2ToolsPriceconversion(ctx, amount, optional)
Price Conversion v2

Convert an amount of one cryptocurrency or fiat currency into one or more different currencies utilizing the latest market rate for each currency. You may optionally pass a historical timestamp as `time` to convert values based on historical rates (as your API plan supports).       **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br> **Technical Notes** - Latest market rate conversions are accurate to 1 minute of specificity. Historical conversions are accurate to 1 minute of specificity outside of non-USD fiat conversions which have 5 minute specificity.  - You may reference a current list of all supported cryptocurrencies via the <a href=\"/api/v1/#section/Standards-and-Conventions\" target=\"_blank\">cryptocurrency/map</a> endpoint. This endpoint also returns the supported date ranges for historical conversions via the `first_historical_data` and `last_historical_data` properties.    - Conversions are supported in 93 different fiat currencies and 4 precious metals <a href=\"/api/v1/#section/Standards-and-Conventions\" target=\"_blank\">as outlined here</a>. Historical fiat conversions are supported as far back as 2013-04-28. - A `last_updated` timestamp is included for both your source currency and each conversion currency. This is the timestamp of the closest market rate record referenced for each currency during the conversion.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic (Latest market price conversions) - Hobbyist (Latest market price conversions + 1 month historical) - Startup (Latest market price conversions + 1 month historical) - Standard (Latest market price conversions + 3 months historical) - Professional (Latest market price conversions + 12 months historical) - Enterprise (Latest market price conversions + up to 6 years historical)  **Cache / Update frequency:** Every 60 seconds for the lastest cryptocurrency and fiat currency rates.     **Plan credit use:** 1 call credit per call and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our cryptocurrency conversion page at [coinmarketcap.com/converter/](https://coinmarketcap.com/converter/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| An amount of currency to convert. Example: 10.43 | 
 **optional** | ***ToolsApiGetV2ToolsPriceconversionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ToolsApiGetV2ToolsPriceconversionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| The CoinMarketCap currency ID of the base cryptocurrency or fiat to convert from. Example: \&quot;1\&quot; | 
 **symbol** | **optional.String**| Alternatively the currency symbol of the base cryptocurrency or fiat to convert from. Example: \&quot;BTC\&quot;. One \&quot;id\&quot; *or* \&quot;symbol\&quot; is required. | 
 **time** | **optional.String**| Optional timestamp (Unix or ISO 8601) to reference historical pricing during conversion. If not passed, the current time will be used. If passed, we&#39;ll reference the closest historic values available for this conversion. | 
 **convert** | **optional.String**| Pass up to 120 comma-separated fiat or cryptocurrency symbols to convert the source amount to. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**ToolsPriceConversionResponseModel**](Tools Price Conversion - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

