# \GlobalMetricsApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GlobalmetricsQuotesHistorical**](GlobalMetricsApi.md#GetV1GlobalmetricsQuotesHistorical) | **Get** /v1/global-metrics/quotes/historical | Quotes Historical
[**GetV1GlobalmetricsQuotesLatest**](GlobalMetricsApi.md#GetV1GlobalmetricsQuotesLatest) | **Get** /v1/global-metrics/quotes/latest | Quotes Latest


# **GetV1GlobalmetricsQuotesHistorical**
> GlobalMetricsQuotesHistoricResponseModel GetV1GlobalmetricsQuotesHistorical(ctx, optional)
Quotes Historical

Returns an interval of historical global cryptocurrency market metrics based on time and interval parameters.  **Technical Notes** - A historic quote for every \"interval\" period between your \"time_start\" and \"time_end\" will be returned. - If a \"time_start\" is not supplied, the \"interval\" will be applied in reverse from \"time_end\". - If \"time_end\" is not supplied, it defaults to the current time. - At each \"interval\" period, the historic quote that is closest in time to the requested time will be returned. - If no historic quotes are available in a given \"interval\" period up until the next interval period, it will be skipped.  **Interval Options**   There are 2 types of time interval formats that may be used for \"interval\".    The first are calendar year and time constants in UTC time:   **\"hourly\"** - Get the first quote available at the beginning of each calendar hour.   **\"daily\"** - Get the first quote available at the beginning of each calendar day.   **\"weekly\"** - Get the first quote available at the beginning of each calendar week.   **\"monthly\"** - Get the first quote available at the beginning of each calendar month.   **\"yearly\"** - Get the first quote available at the beginning of each calendar year.      The second are relative time intervals.   **\"m\"**: Get the first quote available every \"m\" minutes (60 second intervals). Supported minutes are: \"5m\", \"10m\", \"15m\", \"30m\", \"45m\".   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Get the first quote available every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - ~~Basic~~ - ~~Hobbyist~~ - ~~Startup~~ - Standard (3 months) - Professional (12 months) - Enterprise (Up to 6 years)  **Cache / Update frequency:** Every 5 minutes.   **Plan credit use:** 1 call credit per 100 historical data points returned (rounded up).   **CMC equivalent pages:** Our Total Market Capitalization global chart [coinmarketcap.com/charts/](https://coinmarketcap.com/charts/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalMetricsApiGetV1GlobalmetricsQuotesHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalMetricsApiGetV1GlobalmetricsQuotesHistoricalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeStart** | **optional.String**| Timestamp (Unix or ISO 8601) to start returning quotes for. Optional, if not passed, we&#39;ll return quotes calculated in reverse from \&quot;time_end\&quot;. | 
 **timeEnd** | **optional.String**| Timestamp (Unix or ISO 8601) to stop returning quotes for (inclusive). Optional, if not passed, we&#39;ll default to the current time. If no \&quot;time_start\&quot; is passed, we return quotes in reverse order starting from this time. | 
 **count** | **optional.Float32**| The number of interval periods to return results for. Optional, required if both \&quot;time_start\&quot; and \&quot;time_end\&quot; aren&#39;t supplied. The default is 10 items. The current query limit is 10000. | [default to 10]
 **interval** | **optional.String**| Interval of time to return data points for. See details in endpoint description. | [default to 1d]
 **convert** | **optional.String**| By default market quotes are returned in USD. Optionally calculate market quotes in up to 3 other fiat currencies or cryptocurrencies. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;btc_dominance,active_cryptocurrencies,active_exchanges,active_market_pairs,total_volume_24h,total_volume_24h_reported,altcoin_market_cap,altcoin_volume_24h,altcoin_volume_24h_reported,search_interval&#x60; to include all auxiliary fields. | [default to btc_dominance,active_cryptocurrencies,active_exchanges,active_market_pairs,total_volume_24h,total_volume_24h_reported,altcoin_market_cap,altcoin_volume_24h,altcoin_volume_24h_reported]

### Return type

[**GlobalMetricsQuotesHistoricResponseModel**](Global Metrics Quotes Historic - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1GlobalmetricsQuotesLatest**
> GlobalMetricsQuotesLatestResponseModel GetV1GlobalmetricsQuotesLatest(ctx, optional)
Quotes Latest

Returns the latest global cryptocurrency market metrics. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.  **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Hobbyist - Startup - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 5 minute.    **Plan credit use:** 1 call credit per call and 1 call credit per `convert` option beyond the first.    **CMC equivalent pages:** The latest aggregate global market stats ticker across all CMC pages like [coinmarketcap.com](https://coinmarketcap.com/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalMetricsApiGetV1GlobalmetricsQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalMetricsApiGetV1GlobalmetricsQuotesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**GlobalMetricsQuotesLatestResponseModel**](Global Metrics Quotes Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

