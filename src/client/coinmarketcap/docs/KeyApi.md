# \KeyApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1KeyInfo**](KeyApi.md#GetV1KeyInfo) | **Get** /v1/key/info | Key Info


# **GetV1KeyInfo**
> AccountInfoResponseModel GetV1KeyInfo(ctx, )
Key Info

Returns API key details and usage stats. This endpoint can be used to programmatically monitor your key usage compared to the rate limit and daily/monthly credit limits available to your API plan. You may use the Developer Portal's account dashboard as an alternative to this endpoint.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise    **Cache / Update frequency:** No cache, this endpoint updates as requests are made with your key.     **Plan credit use:** No API credit cost. Requests to this endpoint do contribute to your minute based rate limit however.     **CMC equivalent pages:** Our Developer Portal dashboard for your API Key at [pro.coinmarketcap.com/account](https://pro.coinmarketcap.com/account).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountInfoResponseModel**](Account Info - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

