# \ExchangeApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ExchangeInfo**](ExchangeApi.md#GetV1ExchangeInfo) | **Get** /v1/exchange/info | Metadata
[**GetV1ExchangeListingsLatest**](ExchangeApi.md#GetV1ExchangeListingsLatest) | **Get** /v1/exchange/listings/latest | Listings Latest
[**GetV1ExchangeMap**](ExchangeApi.md#GetV1ExchangeMap) | **Get** /v1/exchange/map | CoinMarketCap ID Map
[**GetV1ExchangeMarketpairsLatest**](ExchangeApi.md#GetV1ExchangeMarketpairsLatest) | **Get** /v1/exchange/market-pairs/latest | Market Pairs Latest
[**GetV1ExchangeQuotesHistorical**](ExchangeApi.md#GetV1ExchangeQuotesHistorical) | **Get** /v1/exchange/quotes/historical | Quotes Historical
[**GetV1ExchangeQuotesLatest**](ExchangeApi.md#GetV1ExchangeQuotesLatest) | **Get** /v1/exchange/quotes/latest | Quotes Latest


# **GetV1ExchangeInfo**
> ExchangesInfoResponseModel GetV1ExchangeInfo(ctx, optional)
Metadata

Returns all static metadata for one or more exchanges. This information includes details like launch date, logo, official website URL, social links, and market fee documentation URL.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Static data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 call credit per 100 exchanges returned (rounded up).   **CMC equivalent pages:** Exchange detail page metadata like [coinmarketcap.com/exchanges/binance/](https://coinmarketcap.com/exchanges/binance/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated CoinMarketCap cryptocurrency exchange ids. Example: \&quot;1,2\&quot; | 
 **slug** | **optional.String**| Alternatively, one or more comma-separated exchange names in URL friendly shorthand \&quot;slug\&quot; format (all lowercase, spaces replaced with hyphens). Example: \&quot;binance,gdax\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; is required. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;urls,logo,description,date_launched,notice,status&#x60; to include all auxiliary fields. | [default to urls,logo,description,date_launched,notice]

### Return type

[**ExchangesInfoResponseModel**](Exchanges Info - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ExchangeListingsLatest**
> ExchangeListingsLatestResponseModel GetV1ExchangeListingsLatest(ctx, optional)
Listings Latest

Returns a paginated list of all cryptocurrency exchanges including the latest aggregate market data for each exchange. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 1 minute.   **Plan credit use:** 1 call credit per 100 exchanges returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our latest exchange listing and ranking pages like [coinmarketcap.com/rankings/exchanges/](https://coinmarketcap.com/rankings/exchanges/).          ***NOTE:** Use this endpoint if you need a sorted and paginated list of exchanges. If you want to query for market data on a few specific exchanges use /v1/exchange/quotes/latest which is optimized for that purpose. The response data between these endpoints is otherwise the same.*  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeListingsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeListingsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **sort** | **optional.String**| What field to sort the list of exchanges by. | [default to volume_24h]
 **sortDir** | **optional.String**| The direction in which to order exchanges against the specified sort. | 
 **marketType** | **optional.String**| The type of exchange markets to include in rankings. This field is deprecated. Please use \&quot;all\&quot; for accurate sorting. | [default to all]
 **category** | **optional.String**| The category for this exchange. | [default to all]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,traffic_score,rank,exchange_score,effective_liquidity_24h,date_launched,fiats&#x60; to include all auxiliary fields. | [default to num_market_pairs,traffic_score,rank,exchange_score,effective_liquidity_24h]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**ExchangeListingsLatestResponseModel**](Exchange Listings Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ExchangeMap**
> ExchangeMapResponseModel GetV1ExchangeMap(ctx, optional)
CoinMarketCap ID Map

Returns a paginated list of all active cryptocurrency exchanges by CoinMarketCap ID. We recommend using this convenience endpoint to lookup and utilize our unique exchange `id` across all endpoints as typical exchange identifiers may change over time. As a convenience you may pass a comma-separated list of exchanges by `slug` to filter this list to only those you require or the `aux` parameter to slim down the payload.  By default this endpoint returns exchanges that have at least 1 actively tracked market. You may receive a map of all inactive cryptocurrencies by passing `listing_status=inactive`. You may also receive a map of registered exchanges that are listed but do not yet meet methodology requirements to have tracked markets available via `listing_status=untracked`. Please review **(3) Listing Tiers** in our <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology documentation</a> for additional details on listing states.   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Mapping data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 call credit per call.   **CMC equivalent pages:** No equivalent, this data is only available via API.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeMapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listingStatus** | **optional.String**| Only active exchanges are returned by default. Pass &#x60;inactive&#x60; to get a list of exchanges that are no longer active. Pass &#x60;untracked&#x60; to get a list of exchanges that are registered but do not currently meet methodology requirements to have active markets tracked. You may pass one or more comma-separated values. | [default to active]
 **slug** | **optional.String**| Optionally pass a comma-separated list of exchange slugs (lowercase URL friendly shorthand name with spaces replaced with dashes) to return CoinMarketCap IDs for. If this option is passed, other options will be ignored. | 
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | 
 **sort** | **optional.String**| What field to sort the list of exchanges by. | [default to id]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;first_historical_data,last_historical_data,is_active,status&#x60; to include all auxiliary fields. | [default to first_historical_data,last_historical_data,is_active]
 **cryptoId** | **optional.String**| Optionally include one fiat or cryptocurrency IDs to filter market pairs by. For example &#x60;?crypto_id&#x3D;1&#x60; would only return exchanges that have BTC. | 

### Return type

[**ExchangeMapResponseModel**](Exchange Map - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ExchangeMarketpairsLatest**
> ExchangeMarketPairsLatestResponseModel GetV1ExchangeMarketpairsLatest(ctx, optional)
Market Pairs Latest

Returns all active market pairs that CoinMarketCap tracks for a given exchange. The latest price and volume information is returned for each market. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.'    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 market pairs returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our exchange level active markets pages like [coinmarketcap.com/exchanges/binance/](https://coinmarketcap.com/exchanges/binance/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeMarketpairsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeMarketpairsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| A CoinMarketCap exchange ID. Example: \&quot;1\&quot; | 
 **slug** | **optional.String**| Alternatively pass an exchange \&quot;slug\&quot; (URL friendly all lowercase shorthand version of name with spaces replaced with hyphens). Example: \&quot;binance\&quot;. One \&quot;id\&quot; *or* \&quot;slug\&quot; is required. | 
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,category,fee_type,market_url,currency_name,currency_slug,price_quote,effective_liquidity,market_score,market_reputation&#x60; to include all auxiliary fields. | [default to num_market_pairs,category,fee_type]
 **matchedId** | **optional.String**| Optionally include one or more comma-delimited fiat or cryptocurrency IDs to filter market pairs by. For example &#x60;?matched_id&#x3D;2781&#x60; would only return BTC markets that matched: \&quot;BTC/USD\&quot; or \&quot;USD/BTC\&quot; for the requested exchange. This parameter cannot be used when &#x60;matched_symbol&#x60; is used. | 
 **matchedSymbol** | **optional.String**| Optionally include one or more comma-delimited fiat or cryptocurrency symbols to filter market pairs by. For example &#x60;?matched_symbol&#x3D;USD&#x60; would only return BTC markets that matched: \&quot;BTC/USD\&quot; or \&quot;USD/BTC\&quot; for the requested exchange. This parameter cannot be used when &#x60;matched_id&#x60; is used. | 
 **category** | **optional.String**| The category of trading this market falls under. Spot markets are the most common but options include derivatives and OTC. | [default to all]
 **feeType** | **optional.String**| The fee type the exchange enforces for this market. | [default to all]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**ExchangeMarketPairsLatestResponseModel**](Exchange Market Pairs Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ExchangeQuotesHistorical**
> ExchangeHistoricalQuotesResponseModel GetV1ExchangeQuotesHistorical(ctx, optional)
Quotes Historical

Returns an interval of historic quotes for any exchange based on time and interval parameters.  **Technical Notes** - A historic quote for every \"interval\" period between your \"time_start\" and \"time_end\" will be returned.   - If a \"time_start\" is not supplied, the \"interval\" will be applied in reverse from \"time_end\".   - If \"time_end\" is not supplied, it defaults to the current time.   - At each \"interval\" period, the historic quote that is closest in time to the requested time will be returned.   - If no historic quotes are available in a given \"interval\" period up until the next interval period, it will be skipped.  - This endpoint supports requesting multiple exchanges in the same call. Please note the API response will be wrapped in an additional object in this case.     **Interval Options**   There are 2 types of time interval formats that may be used for \"interval\".    The first are calendar year and time constants in UTC time:   **\"hourly\"** - Get the first quote available at the beginning of each calendar hour.   **\"daily\"** - Get the first quote available at the beginning of each calendar day.   **\"weekly\"** - Get the first quote available at the beginning of each calendar week.   **\"monthly\"** - Get the first quote available at the beginning of each calendar month.   **\"yearly\"** - Get the first quote available at the beginning of each calendar year.    The second are relative time intervals.   **\"m\"**: Get the first quote available every \"m\" minutes (60 second intervals). Supported minutes are: \"5m\", \"10m\", \"15m\", \"30m\", \"45m\".   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Get the first quote available every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard (3 months)   - Professional (Up to 12 months)   - Enterprise (Up to 6 years)  **Note:** You may use the /exchange/map endpoint to receive a list of earliest historical dates that may be fetched for each exchange as  `first_historical_data`. This timestamp will either be the date CoinMarketCap first started tracking the exchange or 2018-04-26T00:45:00.000Z, the earliest date this type of historical data is available for.    **Cache / Update frequency:** Every 5 minutes.   **Plan credit use:** 1 call credit per 100 historical data points returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** No equivalent, this data is only available via API outside of our volume sparkline charts in [coinmarketcap.com/rankings/exchanges/](https://coinmarketcap.com/rankings/exchanges/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeQuotesHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeQuotesHistoricalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated exchange CoinMarketCap ids. Example: \&quot;24,270\&quot; | 
 **slug** | **optional.String**| Alternatively, one or more comma-separated exchange names in URL friendly shorthand \&quot;slug\&quot; format (all lowercase, spaces replaced with hyphens). Example: \&quot;binance,kraken\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; is required. | 
 **timeStart** | **optional.String**| Timestamp (Unix or ISO 8601) to start returning quotes for. Optional, if not passed, we&#39;ll return quotes calculated in reverse from \&quot;time_end\&quot;. | 
 **timeEnd** | **optional.String**| Timestamp (Unix or ISO 8601) to stop returning quotes for (inclusive). Optional, if not passed, we&#39;ll default to the current time. If no \&quot;time_start\&quot; is passed, we return quotes in reverse order starting from this time. | 
 **count** | **optional.Float32**| The number of interval periods to return results for. Optional, required if both \&quot;time_start\&quot; and \&quot;time_end\&quot; aren&#39;t supplied. The default is 10 items. The current query limit is 10000. | [default to 10]
 **interval** | **optional.String**| Interval of time to return data points for. See details in endpoint description. | [default to 5m]
 **convert** | **optional.String**| By default market quotes are returned in USD. Optionally calculate market quotes in up to 3 other fiat currencies or cryptocurrencies. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**ExchangeHistoricalQuotesResponseModel**](Exchange Historical Quotes - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1ExchangeQuotesLatest**
> ExchangeQuotesLatestResponseModel GetV1ExchangeQuotesLatest(ctx, optional)
Quotes Latest

Returns the latest aggregate market data for 1 or more exchanges. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.  **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - ~~Basic~~ - ~~Hobbyist~~ - ~~Startup~~ - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 exchanges returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Latest market data summary for specific exchanges like [coinmarketcap.com/rankings/exchanges/](https://coinmarketcap.com/rankings/exchanges/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetV1ExchangeQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetV1ExchangeQuotesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated CoinMarketCap exchange IDs. Example: \&quot;1,2\&quot; | 
 **slug** | **optional.String**| Alternatively, pass a comma-separated list of exchange \&quot;slugs\&quot; (URL friendly all lowercase shorthand version of name with spaces replaced with hyphens). Example: \&quot;binance,gdax\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; is required. | 
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,traffic_score,rank,exchange_score,liquidity_score,effective_liquidity_24h&#x60; to include all auxiliary fields. | [default to num_market_pairs,traffic_score,rank,exchange_score,liquidity_score,effective_liquidity_24h]

### Return type

[**ExchangeQuotesLatestResponseModel**](Exchange Quotes Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

