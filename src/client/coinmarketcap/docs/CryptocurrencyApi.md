# \CryptocurrencyApi

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CryptocurrencyAirdrop**](CryptocurrencyApi.md#GetV1CryptocurrencyAirdrop) | **Get** /v1/cryptocurrency/airdrop | Airdrop
[**GetV1CryptocurrencyAirdrops**](CryptocurrencyApi.md#GetV1CryptocurrencyAirdrops) | **Get** /v1/cryptocurrency/airdrops | Airdrops
[**GetV1CryptocurrencyCategories**](CryptocurrencyApi.md#GetV1CryptocurrencyCategories) | **Get** /v1/cryptocurrency/categories | Categories
[**GetV1CryptocurrencyCategory**](CryptocurrencyApi.md#GetV1CryptocurrencyCategory) | **Get** /v1/cryptocurrency/category | Category
[**GetV1CryptocurrencyListingsHistorical**](CryptocurrencyApi.md#GetV1CryptocurrencyListingsHistorical) | **Get** /v1/cryptocurrency/listings/historical | Listings Historical
[**GetV1CryptocurrencyListingsLatest**](CryptocurrencyApi.md#GetV1CryptocurrencyListingsLatest) | **Get** /v1/cryptocurrency/listings/latest | Listings Latest
[**GetV1CryptocurrencyListingsNew**](CryptocurrencyApi.md#GetV1CryptocurrencyListingsNew) | **Get** /v1/cryptocurrency/listings/new | Listings New
[**GetV1CryptocurrencyMap**](CryptocurrencyApi.md#GetV1CryptocurrencyMap) | **Get** /v1/cryptocurrency/map | CoinMarketCap ID Map
[**GetV1CryptocurrencyTrendingGainerslosers**](CryptocurrencyApi.md#GetV1CryptocurrencyTrendingGainerslosers) | **Get** /v1/cryptocurrency/trending/gainers-losers | Trending Gainers &amp; Losers
[**GetV1CryptocurrencyTrendingLatest**](CryptocurrencyApi.md#GetV1CryptocurrencyTrendingLatest) | **Get** /v1/cryptocurrency/trending/latest | Trending Latest
[**GetV1CryptocurrencyTrendingMostvisited**](CryptocurrencyApi.md#GetV1CryptocurrencyTrendingMostvisited) | **Get** /v1/cryptocurrency/trending/most-visited | Trending Most Visited
[**GetV2CryptocurrencyInfo**](CryptocurrencyApi.md#GetV2CryptocurrencyInfo) | **Get** /v2/cryptocurrency/info | Metadata v2
[**GetV2CryptocurrencyMarketpairsLatest**](CryptocurrencyApi.md#GetV2CryptocurrencyMarketpairsLatest) | **Get** /v2/cryptocurrency/market-pairs/latest | Market Pairs Latest v2
[**GetV2CryptocurrencyOhlcvHistorical**](CryptocurrencyApi.md#GetV2CryptocurrencyOhlcvHistorical) | **Get** /v2/cryptocurrency/ohlcv/historical | OHLCV Historical v2
[**GetV2CryptocurrencyOhlcvLatest**](CryptocurrencyApi.md#GetV2CryptocurrencyOhlcvLatest) | **Get** /v2/cryptocurrency/ohlcv/latest | OHLCV Latest v2
[**GetV2CryptocurrencyPriceperformancestatsLatest**](CryptocurrencyApi.md#GetV2CryptocurrencyPriceperformancestatsLatest) | **Get** /v2/cryptocurrency/price-performance-stats/latest | Price Performance Stats v2
[**GetV2CryptocurrencyQuotesHistorical**](CryptocurrencyApi.md#GetV2CryptocurrencyQuotesHistorical) | **Get** /v2/cryptocurrency/quotes/historical | Quotes Historical v2
[**GetV2CryptocurrencyQuotesLatest**](CryptocurrencyApi.md#GetV2CryptocurrencyQuotesLatest) | **Get** /v2/cryptocurrency/quotes/latest | Quotes Latest v2


# **GetV1CryptocurrencyAirdrop**
> AirdropResponseModel GetV1CryptocurrencyAirdrop(ctx, id)
Airdrop

Returns information about a single airdrop available on CoinMarketCap. Includes the cryptocurrency metadata.     **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request no matter query size.   **CMC equivalent pages:** Our free airdrops page [coinmarketcap.com/airdrop/](https://coinmarketcap.com/airdrop/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Airdrop Unique ID. This can be found using the Airdrops API. | 

### Return type

[**AirdropResponseModel**](Airdrop - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyAirdrops**
> AirdropsResponseModel GetV1CryptocurrencyAirdrops(ctx, optional)
Airdrops

Returns a list of past, present, or future airdrops which have run on CoinMarketCap.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request no matter query size.   **CMC equivalent pages:** Our free airdrops page [coinmarketcap.com/airdrop/](https://coinmarketcap.com/airdrop/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyAirdropsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyAirdropsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **status** | **optional.String**| What status of airdrops. | [default to ONGOING]
 **id** | **optional.String**| Filtered airdrops by one cryptocurrency CoinMarketCap IDs. Example: 1 | 
 **slug** | **optional.String**| Alternatively filter airdrops by a cryptocurrency slug. Example: \&quot;bitcoin\&quot; | 
 **symbol** | **optional.String**| Alternatively filter airdrops one cryptocurrency symbol. Example: \&quot;BTC\&quot;. | 

### Return type

[**AirdropsResponseModel**](Airdrops - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyCategories**
> CategoriesResponseModel GetV1CryptocurrencyCategories(ctx, optional)
Categories

Returns information about all coin categories available on CoinMarketCap. Includes a paginated list of cryptocurrency quotes and metadata from each category.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Free   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request + 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our free airdrops page [coinmarketcap.com/cryptocurrency-category/](https://coinmarketcap.com/cryptocurrency-category/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | 
 **id** | **optional.String**| Filtered categories by one or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **slug** | **optional.String**| Alternatively filter categories by a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively filter categories one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. | 

### Return type

[**CategoriesResponseModel**](Categories - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyCategory**
> CategoryResponseModel GetV1CryptocurrencyCategory(ctx, id, optional)
Category

Returns information about a single coin category available on CoinMarketCap. Includes a paginated list of the cryptocurrency quotes and metadata for the category.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Free   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request + 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our Cryptocurrency Category page [coinmarketcap.com/cryptocurrency-category/](https://coinmarketcap.com/cryptocurrency-category/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The Category ID. This can be found using the Categories API. | 
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyCategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of coins to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of coins to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**CategoryResponseModel**](Category - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyListingsHistorical**
> CryptocurrencyListingsLatestResponseModel GetV1CryptocurrencyListingsHistorical(ctx, date, optional)
Listings Historical

Returns a ranked and sorted list of all cryptocurrencies for a historical UTC date.     **Technical Notes** - This endpoint is identical in format to our [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) endpoint but is used to retrieve historical daily ranking snapshots from the end of each UTC day.   - Daily snapshots reflect market data at the end of each UTC day and may be requested as far back as 2013-04-28 (as supported by your plan's historical limits).   - The required \"date\" parameter can be passed as a Unix timestamp or ISO 8601 date but only the date portion of the timestamp will be referenced. It is recommended to send an ISO date format like \"2019-10-10\" without time. - This endpoint is for retrieving paginated and sorted lists of all currencies. If you require historical market data on specific cryptocurrencies you should use [/cryptocurrency/quotes/historical](#operation/getV1CryptocurrencyQuotesHistorical).       Cryptocurrencies are listed by cmc_rank by default. You may optionally sort against any of the following:   **cmc_rank**: CoinMarketCap's market cap rank as outlined in <a href=\"https://coinmarketcap.com/methodology/\" target=\"_blank\">our methodology</a>.   **name**: The cryptocurrency name.   **symbol**: The cryptocurrency symbol.   **market_cap**: market cap (latest trade price x circulating supply).   **price**: latest average trade price across markets.   **circulating_supply**: approximate number of coins currently in circulation.   **total_supply**: approximate total amount of coins in existence right now (minus any coins that have been verifiably burned).   **max_supply**: our best approximation of the maximum amount of coins that will ever exist in the lifetime of the currency.   **num_market_pairs**: number of market pairs across all exchanges trading each currency.   **volume_24h**: 24 hour trading volume for each currency.   **percent_change_1h**: 1 hour trading price percentage change for each currency.   **percent_change_24h**: 24 hour trading price percentage change for each currency.   **percent_change_7d**: 7 day trading price percentage change for each currency.       **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard (3 months)   - Professional (12 months)   - Enterprise (Up to 6 years)  **Cache / Update frequency:** The last completed UTC day is available 30 minutes after midnight on the next UTC day.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our historical daily crypto ranking snapshot pages like this one on [February 02, 2014](https://coinmarketcap.com/historical/20140202/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**| date (Unix or ISO 8601) to reference day of snapshot. | 
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyListingsHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyListingsHistoricalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **sort** | **optional.String**| What field to sort the list of cryptocurrencies by. | [default to cmc_rank]
 **sortDir** | **optional.String**| The direction in which to order cryptocurrencies against the specified sort. | 
 **cryptocurrencyType** | **optional.String**| The type of cryptocurrency to include. | [default to all]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;platform,tags,date_added,circulating_supply,total_supply,max_supply,cmc_rank,num_market_pairs&#x60; to include all auxiliary fields. | [default to platform,tags,date_added,circulating_supply,total_supply,max_supply,cmc_rank,num_market_pairs]

### Return type

[**CryptocurrencyListingsLatestResponseModel**](Cryptocurrency Listings Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyListingsLatest**
> CryptocurrencyListingsLatestResponseModel1 GetV1CryptocurrencyListingsLatest(ctx, optional)
Listings Latest

Returns a paginated list of all active cryptocurrencies with latest market data. The default \"market_cap\" sort returns cryptocurrency in order of CoinMarketCap's market cap rank (as outlined in <a href=\"https://coinmarketcap.com/methodology/\" target=\"_blank\">our methodology</a>) but you may configure this call to order by another market ranking field. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.   You may sort against any of the following:   **market_cap**: CoinMarketCap's market cap rank as outlined in <a href=\"https://coinmarketcap.com/methodology/\" target=\"_blank\">our methodology</a>.   **market_cap_strict**: A strict market cap sort (latest trade price x circulating supply).   **name**: The cryptocurrency name.   **symbol**: The cryptocurrency symbol.   **date_added**: Date cryptocurrency was added to the system.   **price**: latest average trade price across markets.   **circulating_supply**: approximate number of coins currently in circulation.   **total_supply**: approximate total amount of coins in existence right now (minus any coins that have been verifiably burned).   **max_supply**: our best approximation of the maximum amount of coins that will ever exist in the lifetime of the currency.   **num_market_pairs**: number of market pairs across all exchanges trading each currency.   **market_cap_by_total_supply_strict**: market cap by total supply.   **volume_24h**: rolling 24 hour adjusted trading volume.   **volume_7d**: rolling 24 hour adjusted trading volume.   **volume_30d**: rolling 24 hour adjusted trading volume.   **percent_change_1h**: 1 hour trading price percentage change for each currency.   **percent_change_24h**: 24 hour trading price percentage change for each currency.   **percent_change_7d**: 7 day trading price percentage change for each currency.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first. **CMC equivalent pages:** Our latest cryptocurrency listing and ranking pages like [coinmarketcap.com/all/views/all/](https://coinmarketcap.com/all/views/all/), [coinmarketcap.com/tokens/](https://coinmarketcap.com/tokens/), [coinmarketcap.com/gainers-losers/](https://coinmarketcap.com/gainers-losers/), [coinmarketcap.com/new/](https://coinmarketcap.com/new/).         ***NOTE:** Use this endpoint if you need a sorted and paginated list of all cryptocurrencies. If you want to query for market data on a few specific cryptocurrencies use [/v1/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.* 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyListingsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyListingsLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **priceMin** | **optional.Float32**| Optionally specify a threshold of minimum USD price to filter results by. | 
 **priceMax** | **optional.Float32**| Optionally specify a threshold of maximum USD price to filter results by. | 
 **marketCapMin** | **optional.Float32**| Optionally specify a threshold of minimum market cap to filter results by. | 
 **marketCapMax** | **optional.Float32**| Optionally specify a threshold of maximum market cap to filter results by. | 
 **volume24hMin** | **optional.Float32**| Optionally specify a threshold of minimum 24 hour USD volume to filter results by. | 
 **volume24hMax** | **optional.Float32**| Optionally specify a threshold of maximum 24 hour USD volume to filter results by. | 
 **circulatingSupplyMin** | **optional.Float32**| Optionally specify a threshold of minimum circulating supply to filter results by. | 
 **circulatingSupplyMax** | **optional.Float32**| Optionally specify a threshold of maximum circulating supply to filter results by. | 
 **percentChange24hMin** | **optional.Float32**| Optionally specify a threshold of minimum 24 hour percent change to filter results by. | 
 **percentChange24hMax** | **optional.Float32**| Optionally specify a threshold of maximum 24 hour percent change to filter results by. | 
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **sort** | **optional.String**| What field to sort the list of cryptocurrencies by. | [default to market_cap]
 **sortDir** | **optional.String**| The direction in which to order cryptocurrencies against the specified sort. | 
 **cryptocurrencyType** | **optional.String**| The type of cryptocurrency to include. | [default to all]
 **tag** | **optional.String**| The tag of cryptocurrency to include. | [default to all]
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc&#x60; to include all auxiliary fields. | [default to num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply]

### Return type

[**CryptocurrencyListingsLatestResponseModel1**](Cryptocurrency Listings Latest - Response Model 1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyListingsNew**
> CryptocurrencyListingsNewResponseModel GetV1CryptocurrencyListingsNew(ctx, optional)
Listings New

Returns a paginated list of most recently added cryptocurrencies.     **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our \"new\" cryptocurrency page [coinmarketcap.com/new/](https://coinmarketcap.com/new)    ***NOTE:** Use this endpoint if you need a sorted and paginated list of all recently added cryptocurrencies.* 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyListingsNewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyListingsNewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **sortDir** | **optional.String**| The direction in which to order cryptocurrencies against the specified sort. | 

### Return type

[**CryptocurrencyListingsNewResponseModel**](Cryptocurrency Listings New - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyMap**
> CryptocurrencyMapResponseModel GetV1CryptocurrencyMap(ctx, optional)
CoinMarketCap ID Map

Returns a mapping of all cryptocurrencies to unique CoinMarketCap `id`s. Per our <a href=\"#section/Best-Practices\" target=\"_blank\">Best Practices</a> we recommend utilizing CMC ID instead of cryptocurrency symbols to securely identify cryptocurrencies with our other endpoints and in your own application logic.  Each cryptocurrency returned includes typical identifiers such as `name`, `symbol`, and `token_address` for flexible mapping to `id`.         By default this endpoint returns cryptocurrencies that have actively tracked markets on supported exchanges. You may receive a map of all inactive cryptocurrencies by passing `listing_status=inactive`. You may also receive a map of registered cryptocurrency projects that are listed but do not yet meet methodology requirements to have tracked markets via `listing_status=untracked`. Please review our <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology documentation</a> for additional details on listing states.         Cryptocurrencies returned include `first_historical_data` and `last_historical_data` timestamps to conveniently reference historical date ranges available to query with historical time-series data endpoints. You may also use the `aux` parameter to only include properties you require to slim down the payload if calling this endpoint frequently.     **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Mapping data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 API call credit per request no matter query size.   **CMC equivalent pages:** No equivalent, this data is only available via API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyMapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listingStatus** | **optional.String**| Only active cryptocurrencies are returned by default. Pass &#x60;inactive&#x60; to get a list of cryptocurrencies that are no longer active. Pass &#x60;untracked&#x60; to get a list of cryptocurrencies that are listed but do not yet meet methodology requirements to have tracked markets available. You may pass one or more comma-separated values. | [default to active]
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | 
 **sort** | **optional.String**| What field to sort the list of cryptocurrencies by. | [default to id]
 **symbol** | **optional.String**| Optionally pass a comma-separated list of cryptocurrency symbols to return CoinMarketCap IDs for. If this option is passed, other options will be ignored. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;platform,first_historical_data,last_historical_data,is_active,status&#x60; to include all auxiliary fields. | [default to platform,first_historical_data,last_historical_data,is_active]

### Return type

[**CryptocurrencyMapResponseModel**](Cryptocurrency Map - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyTrendingGainerslosers**
> CryptocurrencyTrendingGainersLosersResponseModel GetV1CryptocurrencyTrendingGainerslosers(ctx, optional)
Trending Gainers & Losers

Returns a paginated list of all trending cryptocurrencies, determined and sorted by the largest price gains or losses.   You may sort against any of the following:   **percent_change_24h**: 24 hour trading price percentage change for each currency.  **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 10 minutes.   **Plan credit use:** 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our cryptocurrency Gainers & Losers page [coinmarketcap.com/gainers-losers/](https://coinmarketcap.com/gainers-losers/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyTrendingGainerslosersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyTrendingGainerslosersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **timePeriod** | **optional.String**| Adjusts the overall window of time for the biggest gainers and losers. | [default to 24h]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **sort** | **optional.String**| What field to sort the list of cryptocurrencies by. | [default to percent_change_24h]
 **sortDir** | **optional.String**| The direction in which to order cryptocurrencies against the specified sort. | 

### Return type

[**CryptocurrencyTrendingGainersLosersResponseModel**](Cryptocurrency Trending Gainers &amp; Losers - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyTrendingLatest**
> CryptocurrencyTrendingLatestResponseModel GetV1CryptocurrencyTrendingLatest(ctx, optional)
Trending Latest

Returns a paginated list of all trending cryptocurrency market data, determined and sorted by CoinMarketCap search volume.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 10 minutes.   **Plan credit use:** 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our cryptocurrency Trending page [coinmarketcap.com/trending-cryptocurrencies/](https://coinmarketcap.com/trending-cryptocurrencies/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyTrendingLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyTrendingLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **timePeriod** | **optional.String**| Adjusts the overall window of time for the latest trending coins. | [default to 24h]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**CryptocurrencyTrendingLatestResponseModel**](Cryptocurrency Trending Latest - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyTrendingMostvisited**
> CryptocurrencyTrendingMostVisitedResponseModel GetV1CryptocurrencyTrendingMostvisited(ctx, optional)
Trending Most Visited

Returns a paginated list of all trending cryptocurrency market data, determined and sorted by traffic to coin detail pages.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 10 minutes.   **Plan credit use:** 1 call credit per 200 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** The CoinMarketCap “Most Visited” trending list. [coinmarketcap.com/most-viewed-pages/](https://coinmarketcap.com/most-viewed-pages/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyTrendingMostvisitedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyTrendingMostvisitedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | [default to 1]
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | [default to 100]
 **timePeriod** | **optional.String**| Adjusts the overall window of time for most visited currencies. | [default to 24h]
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 

### Return type

[**CryptocurrencyTrendingMostVisitedResponseModel**](Cryptocurrency Trending Most Visited - Response Model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV2CryptocurrencyInfo**
> CryptocurrenciesInfoResponseModel GetV2CryptocurrencyInfo(ctx, optional)
Metadata v2

Returns all static metadata available for one or more cryptocurrencies. This information includes details like logo, description, official website URL, social links, and links to a cryptocurrency's technical documentation.  **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br> **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Static data is updated only as needed, every 30 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up).   **CMC equivalent pages:** Cryptocurrency detail page metadata like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyInfoOpts struct

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

# **GetV2CryptocurrencyMarketpairsLatest**
> CryptocurrencyMarketPairsLatestResponseModel GetV2CryptocurrencyMarketpairsLatest(ctx, optional)
Market Pairs Latest v2

Lists all active market pairs that CoinMarketCap tracks for a given cryptocurrency or fiat currency. All markets with this currency as the pair base *or* pair quote will be returned. The latest price and volume information is returned for each market. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.   **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br>   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - ~~Startup~~   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 1 minute.   **Plan credit use:** 1 call credit per 100 market pairs returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our active cryptocurrency markets pages like [coinmarketcap.com/currencies/bitcoin/#markets](https://coinmarketcap.com/currencies/bitcoin/#markets).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyMarketpairsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyMarketpairsLatestOpts struct

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

# **GetV2CryptocurrencyOhlcvHistorical**
> CryptocurrencyOhlcvHistoricalResponseModel GetV2CryptocurrencyOhlcvHistorical(ctx, optional)
OHLCV Historical v2

Returns historical OHLCV (Open, High, Low, Close, Volume) data along with market cap for any cryptocurrency using time interval parameters. Currently daily and hourly OHLCV periods are supported. Volume is not currently supported for hourly OHLCV intervals before 2020-09-22.  **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br>   **Technical Notes** - Only the date portion of the timestamp is used for daily OHLCV so it's recommended to send an ISO date format like \"2018-09-19\" without time for this \"time_period\".  - One OHLCV quote will be returned for every \"time_period\" between your \"time_start\" (exclusive) and \"time_end\" (inclusive).   - If a \"time_start\" is not supplied, the \"time_period\" will be calculated in reverse from \"time_end\" using the \"count\" parameter which defaults to 10 results.   - If \"time_end\" is not supplied, it defaults to the current time.    - If you don't need every \"time_period\" between your dates you may adjust the frequency that \"time_period\" is sampled using the \"interval\" parameter. For example with \"time_period\" set to \"daily\" you may set \"interval\" to \"2d\" to get the daily OHLCV for every other day. You could set \"interval\" to \"monthly\" to get the first daily OHLCV for each month, or set it to \"yearly\" to get the daily OHLCV value against the same date every year.    **Implementation Tips** - If querying for a specific OHLCV date your \"time_start\" should specify a timestamp of 1 interval prior as \"time_start\" is an exclusive time parameter (as opposed to \"time_end\" which is inclusive to the search). This means that when you pass a \"time_start\" results will be returned for the *next* complete \"time_period\". For example, if you are querying for a daily OHLCV datapoint for 2018-11-30 your \"time_start\" should be \"2018-11-29\".    - If only specifying a \"count\" parameter to return latest OHLCV periods, your \"count\" should be 1 number higher than the number of results you expect to receive. \"Count\" defines the number of \"time_period\" intervals queried, *not* the number of results to return, and this includes the currently active time period which is incomplete when working backwards from current time. For example, if you want the last daily OHLCV value available simply pass \"count=2\" to skip the incomplete active time period. - This endpoint supports requesting multiple cryptocurrencies in the same call. Please note the API response will be wrapped in an additional object in this case.      **Interval Options**      There are 2 types of time interval formats that may be used for \"time_period\" and \"interval\" parameters. For \"time_period\" these return aggregate OHLCV data from the beginning to end of each interval period. Apply these time intervals to \"interval\" to adjust how frequently \"time_period\" is sampled.      The first are calendar year and time constants in UTC time:   **\"hourly\"** - Hour intervals in UTC.   **\"daily\"** - Calendar day intervals for each UTC day.   **\"weekly\"** - Calendar week intervals for each calendar week.   **\"monthly\"** - Calendar month intervals for each calendar month.     **\"yearly\"** - Calendar year intervals for each calendar year.      The second are relative time intervals.   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Time periods that repeat every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".      Please note that \"time_period\" currently supports the \"daily\" and \"hourly\" options. \"interval\" supports all interval options.      **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~ - ~~Hobbyist~~ - Startup (1 month) - Standard (3 months) - Professional (12 months) - Enterprise (Up to 6 years)  **Cache / Update frequency:** Latest Daily OHLCV record is available ~5 to ~10 minutes after each midnight UTC. The latest hourly OHLCV record is available 5 minutes after each UTC hour.   **Plan credit use:** 1 call credit per 100 OHLCV data points returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our historical cryptocurrency data pages like [coinmarketcap.com/currencies/bitcoin/historical-data/](https://coinmarketcap.com/currencies/bitcoin/historical-data/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyOhlcvHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyOhlcvHistoricalOpts struct

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

# **GetV2CryptocurrencyOhlcvLatest**
> CryptocurrencyOhlcvLatestResponseModel GetV2CryptocurrencyOhlcvLatest(ctx, optional)
OHLCV Latest v2

Returns the latest OHLCV (Open, High, Low, Close, Volume) market values for one or more cryptocurrencies for the current UTC day. Since the current UTC day is still active these values are updated frequently. You can find the final calculated OHLCV values for the last completed UTC day along with all historic days using /cryptocurrency/ohlcv/historical.   **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br>   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - Startup   - Standard   - Professional   - Enterprise    **Cache / Update frequency:** Every 10 minutes. Additional OHLCV intervals and 1 minute updates will be available in the future.     **Plan credit use:** 1 call credit per 100 OHLCV values returned (rounded up) and 1 call credit per `convert` option beyond the first.     **CMC equivalent pages:** No equivalent, this data is only available via API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyOhlcvLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyOhlcvLatestOpts struct

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

# **GetV2CryptocurrencyPriceperformancestatsLatest**
> CryptocurrencyPricePerformanceStatsLatestResponseModel GetV2CryptocurrencyPriceperformancestatsLatest(ctx, optional)
Price Performance Stats v2

Returns price performance statistics for one or more cryptocurrencies including launch price ROI and all-time high / all-time low. Stats are returned for an `all_time` period by default. UTC `yesterday` and a number of *rolling time periods* may be requested using the `time_period` parameter. Utilize the `convert` parameter to translate values into multiple fiats or cryptocurrencies using historical rates.   **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br>   **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - ~~Basic~~   - ~~Hobbyist~~   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** The statistics module displayed on cryptocurrency pages like [Bitcoin](https://coinmarketcap.com/currencies/bitcoin/).         ***NOTE:** You may also use [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) for traditional OHLCV data at historical daily and hourly intervals. You may also use [/v1/cryptocurrency/ohlcv/latest](#operation/getV1CryptocurrencyOhlcvLatest) for OHLCV data for the current UTC day.* 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyPriceperformancestatsLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyPriceperformancestatsLatestOpts struct

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

# **GetV2CryptocurrencyQuotesHistorical**
> CryptocurrencyQuotesHistoricalResponseModel GetV2CryptocurrencyQuotesHistorical(ctx, optional)
Quotes Historical v2

Returns an interval of historic market quotes for any cryptocurrency based on time and interval parameters.  **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br> **Technical Notes**   - A historic quote for every \"interval\" period between your \"time_start\" and \"time_end\" will be returned.   - If a \"time_start\" is not supplied, the \"interval\" will be applied in reverse from \"time_end\".   - If \"time_end\" is not supplied, it defaults to the current time.   - At each \"interval\" period, the historic quote that is closest in time to the requested time will be returned.   - If no historic quotes are available in a given \"interval\" period up until the next interval period, it will be skipped.    **Implementation Tips** - Want to get the last quote of each UTC day? Don't use \"interval=daily\" as that returns the first quote. Instead use \"interval=24h\" to repeat a specific timestamp search every 24 hours and pass ex. \"time_start=2019-01-04T23:59:00.000Z\" to query for the last record of each UTC day. - This endpoint supports requesting multiple cryptocurrencies in the same call. Please note the API response will be wrapped in an additional object in this case.      **Interval Options**   There are 2 types of time interval formats that may be used for \"interval\".  The first are calendar year and time constants in UTC time:   **\"hourly\"** - Get the first quote available at the beginning of each calendar hour.   **\"daily\"** - Get the first quote available at the beginning of each calendar day.   **\"weekly\"** - Get the first quote available at the beginning of each calendar week.   **\"monthly\"** - Get the first quote available at the beginning of each calendar month.   **\"yearly\"** - Get the first quote available at the beginning of each calendar year.    The second are relative time intervals.   **\"m\"**: Get the first quote available every \"m\" minutes (60 second intervals). Supported minutes are: \"5m\", \"10m\", \"15m\", \"30m\", \"45m\".   **\"h\"**: Get the first quote available every \"h\" hours (3600 second intervals). Supported hour intervals are: \"1h\", \"2h\", \"3h\", \"4h\", \"6h\", \"12h\".   **\"d\"**: Get the first quote available every \"d\" days (86400 second intervals). Supported day intervals are: \"1d\", \"2d\", \"3d\", \"7d\", \"14d\", \"15d\", \"30d\", \"60d\", \"90d\", \"365d\".    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - ~~Basic~~ - ~~Hobbyist~~ - ~~Startup~~ - Standard (3 month) - Professional (12 months) - Enterprise (Up to 6 years)  **Cache / Update frequency:** Every 5 minutes.     **Plan credit use:** 1 call credit per 100 historical data points returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Our historical cryptocurrency charts like [coinmarketcap.com/currencies/bitcoin/#charts](https://coinmarketcap.com/currencies/bitcoin/#charts).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyQuotesHistoricalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyQuotesHistoricalOpts struct

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

# **GetV2CryptocurrencyQuotesLatest**
> CryptocurrencyQuotesLatestResponseModel GetV2CryptocurrencyQuotesLatest(ctx, optional)
Quotes Latest v2

Returns the latest market quote for 1 or more cryptocurrencies. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.  **Please note**: This documentation relates to our updated V2 endpoint, which may be incompatible with our V1 versions. Documentation for deprecated endpoints can be found <a href=\"#tag/deprecated\">here</a>.<br><br> **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 60 seconds.   **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first.   **CMC equivalent pages:** Latest market data pages for specific cryptocurrencies like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/).       ***NOTE:** Use this endpoint to request the latest quote for specific cryptocurrencies. If you need to request all cryptocurrencies use [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV2CryptocurrencyQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV2CryptocurrencyQuotesLatestOpts struct

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

