# \DefaultApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ContentLatest**](DefaultApi.md#GetV1ContentLatest) | **Get** /v1/content/latest | Content Latest
[**GetV1ContentPostsComments**](DefaultApi.md#GetV1ContentPostsComments) | **Get** /v1/content/posts/comments | Content Post Comments
[**GetV1ContentPostsLatest**](DefaultApi.md#GetV1ContentPostsLatest) | **Get** /v1/content/posts/latest | Content Latest Posts
[**GetV1ContentPostsTop**](DefaultApi.md#GetV1ContentPostsTop) | **Get** /v1/content/posts/top | Content Top Posts


# **GetV1ContentLatest**
> ContentLatestResponseModel GetV1ContentLatest(ctx, optional)
Content Latest

Returns a paginated list of content pulled from CMC News/Headlines and Alexandria articles.       **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Professional   - Enterprise  **Cache / Update frequency:** Five Minutes **Plan credit use:** 1 call credit per 100 items returned (rounded up).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetV1ContentLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetV1ContentLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **id** | **optional.String**| Optional cryptocurrency CoinMarketCap IDs. Example: 1027 | 
 **slug** | **optional.String**| Alternatively pass one cryptocurrency slug. Example: \&quot;ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one cryptocurrency symbols. Example: \&quot;ETH\&quot; | 
 **newsType** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields: &#x60;news&#x60;, &#x60;gravity&#x60;, or &#x60;alexandria&#x60; to filter news sources. Pass &#x60;all&#x60; or leave it blank to include all news types. | [default to all]

### Return type

[**ContentLatestResponseModel**](Content Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ContentPostsComments**
> ContentPostCommentsResponseModel GetV1ContentPostsComments(ctx, postId)
Content Post Comments

Returns comments of the CMC Community post.       **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Professional   - Enterprise  **Cache / Update frequency:** Five Minutes **Plan credit use:** 0 credit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postId** | **string**| Required post ID. Example: 325670123 | 

### Return type

[**ContentPostCommentsResponseModel**](Content Post Comments - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ContentPostsLatest**
> ContentLatestPostsResponseModel GetV1ContentPostsLatest(ctx, optional)
Content Latest Posts

Returns the latest crypto-related posts from the CMC Community.       **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Professional   - Enterprise  **Cache / Update frequency:** Five Minutes **Plan credit use:** 0 credit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetV1ContentPostsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetV1ContentPostsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Optional one cryptocurrency CoinMarketCap ID. Example: 1027 | 
 **slug** | **optional.String**| Alternatively pass one cryptocurrency slug. Example: \&quot;ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one cryptocurrency symbols. Example: \&quot;ETH\&quot; | 
 **lastScore** | **optional.String**| Optional. The score is given in the response for finding next batch posts. Example: 1662903634322 | 

### Return type

[**ContentLatestPostsResponseModel**](Content Latest Posts - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ContentPostsTop**
> ContentTopPostsResponseModel GetV1ContentPostsTop(ctx, optional)
Content Top Posts

Returns the top crypto-related posts from the CMC Community.       **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Professional   - Enterprise  **Cache / Update frequency:** Five Minutes **Plan credit use:** 0 credit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetV1ContentPostsTopOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetV1ContentPostsTopOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Optional one cryptocurrency CoinMarketCap ID. Example: 1027 | 
 **slug** | **optional.String**| Alternatively pass one cryptocurrency slug. Example: \&quot;ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one cryptocurrency symbols. Example: \&quot;ETH\&quot; | 
 **lastScore** | **optional.String**| Optional. The score is given in the response for finding next batch of related posts. Example: 38507.8865 | 

### Return type

[**ContentTopPostsResponseModel**](Content Top Posts - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

