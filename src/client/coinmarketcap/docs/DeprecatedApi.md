# \DeprecatedApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CryptocurrencyInfo**](DeprecatedApi.md#GetV1CryptocurrencyInfo) | **Get** /v1/cryptocurrency/info | Metadata v1 (deprecated)
[**GetV1CryptocurrencyMarketpairsLatest**](DeprecatedApi.md#GetV1CryptocurrencyMarketpairsLatest) | **Get** /v1/cryptocurrency/market-pairs/latest | Market Pairs Latest v1 (deprecated)
[**GetV1CryptocurrencyOhlcvHistorical**](DeprecatedApi.md#GetV1CryptocurrencyOhlcvHistorical) | **Get** /v1/cryptocurrency/ohlcv/historical | OHLCV Historical v1 (deprecated)
[**GetV1CryptocurrencyOhlcvLatest**](DeprecatedApi.md#GetV1CryptocurrencyOhlcvLatest) | **Get** /v1/cryptocurrency/ohlcv/latest | OHLCV Latest v1 (deprecated)
[**GetV1CryptocurrencyPriceperformancestatsLatest**](DeprecatedApi.md#GetV1CryptocurrencyPriceperformancestatsLatest) | **Get** /v1/cryptocurrency/price-performance-stats/latest | Price Performance Stats v1 (deprecated)
[**GetV1CryptocurrencyQuotesHistorical**](DeprecatedApi.md#GetV1CryptocurrencyQuotesHistorical) | **Get** /v1/cryptocurrency/quotes/historical | Quotes Historical v1 (deprecated)
[**GetV1CryptocurrencyQuotesLatest**](DeprecatedApi.md#GetV1CryptocurrencyQuotesLatest) | **Get** /v1/cryptocurrency/quotes/latest | Quotes Latest v1 (deprecated)
[**GetV1PartnersFlipsidecryptoFcasListingsLatest**](DeprecatedApi.md#GetV1PartnersFlipsidecryptoFcasListingsLatest) | **Get** /v1/partners/flipside-crypto/fcas/listings/latest | FCAS Listings Latest (deprecated)
[**GetV1PartnersFlipsidecryptoFcasQuotesLatest**](DeprecatedApi.md#GetV1PartnersFlipsidecryptoFcasQuotesLatest) | **Get** /v1/partners/flipside-crypto/fcas/quotes/latest | FCAS Quotes Latest (deprecated)
[**GetV1ToolsPriceconversion**](DeprecatedApi.md#GetV1ToolsPriceconversion) | **Get** /v1/tools/price-conversion | Price Conversion v1 (deprecated)


# **GetV1CryptocurrencyInfo**
> CryptocurrenciesInfoResponseModel GetV1CryptocurrencyInfo(ctx, optional)
Metadata v1 (deprecated)

Returns all static metadata available for one or more cryptocurrencies. This information includes details like logo, description, official website URL, social links, and links to a cryptocurrency's technical documentation.   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Static data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up).   **CMC equivalent pages:** Cryptocurrency detail page metadata like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated CoinMarketCap cryptocurrency IDs. Example: \&quot;1,2\&quot; | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **address** | **optional.String**| Alternatively pass in a contract address. Example: \&quot;0xc40af1e4fecfa05ce6bab79dcd8b373d2e436c4e\&quot; | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;urls,logo,description,tags,platform,date_added,notice,status&#x60; to include all auxiliary fields. | [default to urls,logo,description,tags,platform,date_added,notice]

### Return type

[**CryptocurrenciesInfoResponseModel**](Cryptocurrencies Info - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyMarketpairsLatest**
> CryptocurrencyMarketPairsLatestResponseModel GetV1CryptocurrencyMarketpairsLatest(ctx, optional)
Market Pairs Latest v1 (deprecated)

Lists all active market pairs that CoinMarketCap tracks for a given cryptocurrency or fiat currency. All markets with this currency as the pair base *or* pair quote will be returned. The latest price and volume information is returned for each market. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 1 minute.   **Plan credit use:** 1 call credit per 100 market pairs returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our active cryptocurrency markets pages like [coinmarketcap.com/currencies/bitcoin/#markets](https://coinmarketcap.com/currencies/bitcoin/#markets).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyMarketpairsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyMarketpairsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| A cryptocurrency or fiat currency by CoinMarketCap ID to list market pairs for. Example: \&quot;1\&quot; | 
 **slug** | **optional.String**| Alternatively pass a cryptocurrency by slug. Example: \&quot;bitcoin\&quot; | 
 **symbol** | **optional.String**| Alternatively pass a cryptocurrency by symbol. Fiat currencies are not supported by this field. Example: \&quot;BTC\&quot;. A single cryptocurrency \&quot;id\&quot;, \&quot;slug\&quot;, *or* \&quot;symbol\&quot; is required. | 
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **sortDir** | **optional.String**| Optionally specify the sort direction of markets returned. | [default to desc]
 **sort** | **optional.String**| Optionally specify the sort order of markets returned. By default we return a strict sort on 24 hour reported volume. Pass &#x60;cmc_rank&#x60; to return a CMC methodology based sort where markets with excluded volumes are returned last. | [default to volume_24h_strict]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,category,fee_type,market_url,currency_name,currency_slug,price_quote,notice,cmc_rank,effective_liquidity,market_score,market_reputation&#x60; to include all auxiliary fields. | [default to num_market_pairs,category,fee_type]
 **matchedId** | **optional.String**| Optionally include one or more fiat or cryptocurrency IDs to filter market pairs by. For example &#x60;?id&#x3D;1&amp;matched_id&#x3D;2781&#x60; would only return BTC markets that matched: \&quot;BTC/USD\&quot; or \&quot;USD/BTC\&quot;. This parameter cannot be used when &#x60;matched_symbol&#x60; is used. | 
 **matchedSymbol** | **optional.String**| Optionally include one or more fiat or cryptocurrency symbols to filter market pairs by. For example &#x60;?symbol&#x3D;BTC&amp;matched_symbol&#x3D;USD&#x60; would only return BTC markets that matched: \&quot;BTC/USD\&quot; or \&quot;USD/BTC\&quot;. This parameter cannot be used when &#x60;matched_id&#x60; is used. | 
 **category** | **optional.String**| The category of trading this market falls under. Spot markets are the most common but options include derivatives and OTC. | [default to all]
 **feeType** | **optional.String**| The fee type the exchange enforces for this market. | [default to all]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**CryptocurrencyMarketPairsLatestResponseModel**](Cryptocurrency Market Pairs Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyOhlcvHistorical**
> CryptocurrencyOhlcvHistoricalResponseModel GetV1CryptocurrencyOhlcvHistorical(ctx, optional)
OHLCV Historical v1 (deprecated)

Returns historical OHLCV (Open, High, Low, Close, Volume) data along with market cap for any cryptocurrency using time interval parameters. Currently daily and hourly OHLCV periods are supported. Volume is not currently supported for hourly OHLCV intervals before 2020-09-22.     **Technical Notes** - Only the date portion of the timestamp is used for daily OHLCV so it's recommended to send an ISO date format like \"2018-09-19\" without time for this \"time_period\".  - One OHLCV quote will be returned for every \"time_period\" between your \"time_start\" (exclusive) and \"time_end\" (inclusive).   - If a \"time_start\" is not supplied, the \"time_period\" will be calculated in reverse from \"time_end\" using the \"count\" parameter which defaults to 10 results.   - If \"time_end\" is not supplied, it defaults to the current time.    - If you don't need every \"time_period\" between your dates you may adjust the frequency that \"time_period\" is sampled using the \"interval\" parameter. For example with \"time_period\" set to \"daily\" you may set \"interval\" to \"2d\" to get the daily OHLCV for every other day. You could set \"interval\" to \"monthly\" to get the first daily OHLCV for each month, or set it to \"yearly\" to get the daily OHLCV value against the same date every year.    **Implementation Tips** - If querying for a specific OHLCV date your \"time_start\" should specify a timestamp of 1 interval prior as \"time_start\" is an exclusive time parameter (as opposed to \"time_end\" which is inclusive to the search). This means that when you pass a \"time_start\" results will be returned for the *next* complete \"time_period\". For example, if you are querying for a daily OHLCV datapoint for 2018-11-30 your \"time_start\" should be \"2018-11-29\".    - If only specifying a \"count\" parameter to return latest OHLCV periods, your \"count\" should be 1 number higher than the number of results you expect to receive. \"Count\" defines the number of \"time_period\" intervals queried, *not* the number of results to return, and this includes the currently active time period which is incomplete when working backwards from current time. For example, if you want the last daily OHLCV value available simply pass \"count=2\" to skip the incomplete active time period. - This endpoint supports requesting multiple cryptocurrencies in the same call. Please note the API response will be wrapped in an additional object in this case.      **Interval Options**      There are 2 types of time interval formats that may be used for \"time_period\" and \"interval\" parameters. For \"time_period\" these return aggregate OHLCV data from the beginning to end of each interval period. Apply these time intervals to \"interval\" to adjust how frequently \"time_period\" is sampled.      The first are calendar year and time constants in UTC time:   **\"hourly\"** - Hour intervals in UTC.   **\"daily\"** - Calendar day intervals for each UTC day.   **\"weekly\"** - Calendar week intervals for each calendar week.   **\"monthly\"** - Calendar month intervals for each calendar month.     **\"yearly\"** - Calendar year intervals for each calendar year.      The second are relative time intervals.   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Time periods that repeat every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".      Please note that \"time_period\" currently supports the \"daily\" and \"hourly\" options. \"interval\" supports all interval options.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~ - ~~Hobbyist~~ - Startup (1 month) - Standard (3 months) - Professional (12 months) - Enterprise (Up to 6 years)  **Cache / Update frequency:** Latest Daily OHLCV record is available ~5 to ~10 minutes after each midnight UTC. The latest hourly OHLCV record is available 5 minutes after each UTC hour.   **Plan credit use:** 1 call credit per 100 OHLCV data points returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our historical cryptocurrency data pages like [coinmarketcap.com/currencies/bitcoin/historical-data/](https://coinmarketcap.com/currencies/bitcoin/historical-data/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyOhlcvHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyOhlcvHistoricalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated CoinMarketCap cryptocurrency IDs. Example: \&quot;1,1027\&quot; | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **timePeriod** | **optional.String**| Time period to return OHLCV data for. The default is \&quot;daily\&quot;. See the main endpoint description for details. | [default to daily]
 **timeStart** | **optional.String**| Timestamp (Unix or ISO 8601) to start returning OHLCV time periods for. Only the date portion of the timestamp is used for daily OHLCV so it&#39;s recommended to send an ISO date format like \&quot;2018-09-19\&quot; without time. | 
 **timeEnd** | **optional.String**| Timestamp (Unix or ISO 8601) to stop returning OHLCV time periods for (inclusive). Optional, if not passed we&#39;ll default to the current time. Only the date portion of the timestamp is used for daily OHLCV so it&#39;s recommended to send an ISO date format like \&quot;2018-09-19\&quot; without time. | 
 **count** | **optional.Float32**| Optionally limit the number of time periods to return results for. The default is 10 items. The current query limit is 10000 items. | [default to 10]
 **interval** | **optional.String**| Optionally adjust the interval that \&quot;time_period\&quot; is sampled. See main endpoint description for available options. | [default to daily]
 **convert** | **optional.String**| By default market quotes are returned in USD. Optionally calculate market quotes in up to 3 fiat currencies or cryptocurrencies. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if any invalid cryptocurrencies are requested or a cryptocurrency does not have matching records in the requested timeframe. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | [default to true]

### Return type

[**CryptocurrencyOhlcvHistoricalResponseModel**](Cryptocurrency OHLCV Historical - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyOhlcvLatest**
> CryptocurrencyOhlcvLatestResponseModel GetV1CryptocurrencyOhlcvLatest(ctx, optional)
OHLCV Latest v1 (deprecated)

Returns the latest OHLCV (Open, High, Low, Close, Volume) market values for one or more cryptocurrencies for the current UTC day. Since the current UTC day is still active these values are updated frequently. You can find the final calculated OHLCV values for the last completed UTC day along with all historic days using /cryptocurrency/ohlcv/historical.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - Startup   - Standard   - Professional   - Enterprise    **Cache / Update frequency:** Every 10 minutes. Additional OHLCV intervals and 1 minute updates will be available in the future.     **Plan credit use:** 1 call credit per 100 OHLCV values returned (rounded up) and 1 call credit per `convert` option beyond the first.     **CMC equivalent pages:** No equivalent, this data is only available via API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyOhlcvLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyOhlcvLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;symbol\&quot; is required. | 
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if any invalid cryptocurrencies are requested or a cryptocurrency does not have matching records in the requested timeframe. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | [default to true]

### Return type

[**CryptocurrencyOhlcvLatestResponseModel**](Cryptocurrency OHLCV Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyPriceperformancestatsLatest**
> CryptocurrencyPricePerformanceStatsLatestResponseModel GetV1CryptocurrencyPriceperformancestatsLatest(ctx, optional)
Price Performance Stats v1 (deprecated)

Returns price performance statistics for one or more cryptocurrencies including launch price ROI and all-time high / all-time low. Stats are returned for an `all_time` period by default. UTC `yesterday` and a number of *rolling time periods* may be requested using the `time_period` parameter. Utilize the `convert` parameter to translate values into multiple fiats or cryptocurrencies using historical rates.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** The statistics module displayed on cryptocurrency pages like [Bitcoin](https://coinmarketcap.com/currencies/bitcoin/).         ***NOTE:** You may also use [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) for traditional OHLCV data at historical daily and hourly intervals. You may also use [/v1/cryptocurrency/ohlcv/latest](#operation/getV1CryptocurrencyOhlcvLatest) for OHLCV data for the current UTC day.* 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyPriceperformancestatsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyPriceperformancestatsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **timePeriod** | **optional.String**| Specify one or more comma-delimited time periods to return stats for. &#x60;all_time&#x60; is the default. Pass &#x60;all_time,yesterday,24h,7d,30d,90d,365d&#x60; to return all supported time periods. All rolling periods have a rolling close time of the current request time. For example &#x60;24h&#x60; would have a close time of now and an open time of 24 hours before now. *Please note: &#x60;yesterday&#x60; is a UTC period and currently does not currently support &#x60;high&#x60; and &#x60;low&#x60; timestamps.* | [default to all_time]
 **convert** | **optional.String**| Optionally calculate quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if no match is found for 1 or more requested cryptocurrencies. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | [default to true]

### Return type

[**CryptocurrencyPricePerformanceStatsLatestResponseModel**](Cryptocurrency Price Performance Stats Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyQuotesHistorical**
> CryptocurrencyQuotesHistoricalResponseModel GetV1CryptocurrencyQuotesHistorical(ctx, optional)
Quotes Historical v1 (deprecated)

Returns an interval of historic market quotes for any cryptocurrency based on time and interval parameters.   **Technical Notes**   - A historic quote for every \"interval\" period between your \"time_start\" and \"time_end\" will be returned.   - If a \"time_start\" is not supplied, the \"interval\" will be applied in reverse from \"time_end\".   - If \"time_end\" is not supplied, it defaults to the current time.   - At each \"interval\" period, the historic quote that is closest in time to the requested time will be returned.   - If no historic quotes are available in a given \"interval\" period up until the next interval period, it will be skipped.    **Implementation Tips** - Want to get the last quote of each UTC day? Don't use \"interval=daily\" as that returns the first quote. Instead use \"interval=24h\" to repeat a specific timestamp search every 24 hours and pass ex. \"time_start=2019-01-04T23:59:00.000Z\" to query for the last record of each UTC day. - This endpoint supports requesting multiple cryptocurrencies in the same call. Please note the API response will be wrapped in an additional object in this case.      **Interval Options**   There are 2 types of time interval formats that may be used for \"interval\".  The first are calendar year and time constants in UTC time:   **\"hourly\"** - Get the first quote available at the beginning of each calendar hour.   **\"daily\"** - Get the first quote available at the beginning of each calendar day.   **\"weekly\"** - Get the first quote available at the beginning of each calendar week.   **\"monthly\"** - Get the first quote available at the beginning of each calendar month.   **\"yearly\"** - Get the first quote available at the beginning of each calendar year.    The second are relative time intervals.   **\"m\"**: Get the first quote available every \"m\" minutes (60 second intervals). Supported minutes are: \"5m\", \"10m\", \"15m\", \"30m\", \"45m\".   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Get the first quote available every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - ~~Basic~~ - ~~Hobbyist~~ - ~~Startup~~ - Standard (3 month) - Professional (12 months) - Enterprise (Up to 6 years)  **Cache / Update frequency:** Every 5 minutes.     **Plan credit use:** 1 call credit per 100 historical data points returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our historical cryptocurrency charts like [coinmarketcap.com/currencies/bitcoin/#charts](https://coinmarketcap.com/currencies/bitcoin/#charts).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyQuotesHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyQuotesHistoricalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated CoinMarketCap cryptocurrency IDs. Example: \&quot;1,2\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **timeStart** | **optional.String**| Timestamp (Unix or ISO 8601) to start returning quotes for. Optional, if not passed, we&#39;ll return quotes calculated in reverse from \&quot;time_end\&quot;. | 
 **timeEnd** | **optional.String**| Timestamp (Unix or ISO 8601) to stop returning quotes for (inclusive). Optional, if not passed, we&#39;ll default to the current time. If no \&quot;time_start\&quot; is passed, we return quotes in reverse order starting from this time. | 
 **count** | **optional.Float32**| The number of interval periods to return results for. Optional, required if both \&quot;time_start\&quot; and \&quot;time_end\&quot; aren&#39;t supplied. The default is 10 items. The current query limit is 10000. | [default to 10]
 **interval** | **optional.String**| Interval of time to return data points for. See details in endpoint description. | [default to 5m]
 **convert** | **optional.String**| By default market quotes are returned in USD. Optionally calculate market quotes in up to 3 other fiat currencies or cryptocurrencies. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;price,volume,market_cap,circulating_supply,total_supply,quote_timestamp,is_active,is_fiat,search_interval&#x60; to include all auxiliary fields. | [default to price,volume,market_cap,circulating_supply,total_supply,quote_timestamp,is_active,is_fiat]
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if no match is found for 1 or more requested cryptocurrencies. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | [default to true]

### Return type

[**CryptocurrencyQuotesHistoricalResponseModel**](Cryptocurrency Quotes Historical - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyQuotesLatest**
> CryptocurrencyQuotesLatestResponseModel GetV1CryptocurrencyQuotesLatest(ctx, optional)
Quotes Latest v1 (deprecated)

Returns the latest market quote for 1 or more cryptocurrencies. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Latest market data pages for specific cryptocurrencies like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/).       ***NOTE:** Use this endpoint to request the latest quote for specific cryptocurrencies. If you need to request all cryptocurrencies use [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1CryptocurrencyQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1CryptocurrencyQuotesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_active,is_fiat&#x60; to include all auxiliary fields. | [default to num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,is_active,is_fiat]
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if no match is found for 1 or more requested cryptocurrencies. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | [default to true]

### Return type

[**CryptocurrencyQuotesLatestResponseModel**](Cryptocurrency Quotes Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1PartnersFlipsidecryptoFcasListingsLatest**
> FcasListingsLatestResponseModel GetV1PartnersFlipsidecryptoFcasListingsLatest(ctx, optional)
FCAS Listings Latest (deprecated)

Returns a paginated list of FCAS scores for all cryptocurrencies currently supported by FCAS. FCAS ratings are on a 0-1000 point scale with a corresponding letter grade and is updated once a day at UTC midnight.           FCAS stands for Fundamental Crypto Asset Score, a single, consistently comparable value for measuring cryptocurrency project health. FCAS measures User Activity, Developer Behavior and Market Maturity and is provided by <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/\" target=\"_blank\">FlipSide Crypto</a>. Find out more about <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/fcas-explained\" target=\"_blank\">FCAS methodology</a>. Users interested in FCAS historical data including sub-component scoring may inquire through our <a rel=\"noopener noreferrer\" href=\"https://pro.coinmarketcap.com/contact-data/\" target=\"_blank\">CSV Data Delivery</a> request form.      *Disclaimer: Ratings that are calculated by third party organizations and are not influenced or endorsed by CoinMarketCap in any way.*        **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Once a day at UTC midnight.   **Plan credit use:** 1 call credit per 100 FCAS scores returned (rounded up).   **CMC equivalent pages:** The FCAS ratings available under our cryptocurrency ratings tab like [coinmarketcap.com/currencies/bitcoin/#ratings](https://coinmarketcap.com/currencies/bitcoin/#ratings).         ***NOTE:** Use this endpoint to request the latest FCAS score for all supported cryptocurrencies at the same time. If you require FCAS for only specific cryptocurrencies use [/v1/partners/flipside-crypto/fcas/quotes/latest](#operation/getV1PartnersFlipsidecryptoFcasQuotesLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.* 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1PartnersFlipsidecryptoFcasListingsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1PartnersFlipsidecryptoFcasListingsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;point_change_24h,percent_change_24h&#x60; to include all auxiliary fields. | [default to point_change_24h,percent_change_24h]

### Return type

[**FcasListingsLatestResponseModel**](FCAS Listings Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1PartnersFlipsidecryptoFcasQuotesLatest**
> FcasQuoteLatestResponseModel GetV1PartnersFlipsidecryptoFcasQuotesLatest(ctx, optional)
FCAS Quotes Latest (deprecated)

Returns the latest FCAS score for 1 or more cryptocurrencies. FCAS ratings are on a 0-1000 point scale with a corresponding letter grade and is updated once a day at UTC midnight.           FCAS stands for Fundamental Crypto Asset Score, a single, consistently comparable value for measuring cryptocurrency project health. FCAS measures User Activity, Developer Behavior and Market Maturity and is provided by <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/\" target=\"_blank\">FlipSide Crypto</a>. Find out more about <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/fcas-explained\" target=\"_blank\">FCAS methodology</a>. Users interested in FCAS historical data including sub-component scoring may inquire through our <a rel=\"noopener noreferrer\" href=\"https://pro.coinmarketcap.com/contact-data/\" target=\"_blank\">CSV Data Delivery</a> request form.    *Disclaimer: Ratings that are calculated by third party organizations and are not influenced or endorsed by CoinMarketCap in any way.*        **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Once a day at UTC midnight.   **Plan credit use:** 1 call credit per 100 FCAS scores returned (rounded up).   **CMC equivalent pages:** The FCAS ratings available under our cryptocurrency ratings tab like [coinmarketcap.com/currencies/bitcoin/#ratings](https://coinmarketcap.com/currencies/bitcoin/#ratings).          ***NOTE:** Use this endpoint to request the latest FCAS score for specific cryptocurrencies. If you require FCAS for all supported cryptocurrencies use [/v1/partners/flipside-crypto/fcas/listings/latest](#operation/getV1PartnersFlipsidecryptoFcasListingsLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetV1PartnersFlipsidecryptoFcasQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1PartnersFlipsidecryptoFcasQuotesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;point_change_24h,percent_change_24h&#x60; to include all auxiliary fields. | [default to point_change_24h,percent_change_24h]

### Return type

[**FcasQuoteLatestResponseModel**](FCAS Quote Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ToolsPriceconversion**
> ToolsPriceConversionResponseModel GetV1ToolsPriceconversion(ctx, amount, optional)
Price Conversion v1 (deprecated)

Convert an amount of one cryptocurrency or fiat currency into one or more different currencies utilizing the latest market rate for each currency. You may optionally pass a historical timestamp as `time` to convert values based on historical rates (as your API plan supports).        **Technical Notes** - Latest market rate conversions are accurate to 1 minute of specificity. Historical conversions are accurate to 1 minute of specificity outside of non-USD fiat conversions which have 5 minute specificity.  - You may reference a current list of all supported cryptocurrencies via the <a href=\"/api/v1/#section/Standards-and-Conventions\" target=\"_blank\">cryptocurrency/map</a> endpoint. This endpoint also returns the supported date ranges for historical conversions via the `first_historical_data` and `last_historical_data` properties.    - Conversions are supported in 93 different fiat currencies and 4 precious metals <a href=\"/api/v1/#section/Standards-and-Conventions\" target=\"_blank\">as outlined here</a>. Historical fiat conversions are supported as far back as 2013-04-28. - A `last_updated` timestamp is included for both your source currency and each conversion currency. This is the timestamp of the closest market rate record referenced for each currency during the conversion.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic (Latest market price conversions) - Hobbyist (Latest market price conversions + 1 month historical) - Startup (Latest market price conversions + 1 month historical) - Standard (Latest market price conversions + 3 months historical) - Professional (Latest market price conversions + 12 months historical) - Enterprise (Latest market price conversions + up to 6 years historical)  **Cache / Update frequency:** Every 60 seconds for the lastest cryptocurrency and fiat currency rates.     **Plan credit use:** 1 call credit per call and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our cryptocurrency conversion page at [coinmarketcap.com/converter/](https://coinmarketcap.com/converter/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float32**| An amount of currency to convert. Example: 10.43 | 
 **optional** | ***DeprecatedApiGetV1ToolsPriceconversionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetV1ToolsPriceconversionOpts struct

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

