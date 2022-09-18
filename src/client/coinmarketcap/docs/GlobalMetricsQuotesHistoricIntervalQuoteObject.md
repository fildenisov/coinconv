# GlobalMetricsQuotesHistoricIntervalQuoteObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Timestamp (ISO 8601) of when this historical quote was recorded. | [default to null]
**SearchInterval** | **string** | The interval timestamp for the search period that this historical quote was located against. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**BtcDominance** | **float32** | Percent of BTC market dominance by marketcap at this interval. | [default to null]
**ActiveCryptocurrencies** | **float32** | Number of active cryptocurrencies tracked by CoinMarketCap at the given point in time. This includes all cryptocurrencies with a &#x60;listing_status&#x60; of \&quot;active\&quot; or \&quot;untracked\&quot; as returned from our /cryptocurrency/map call. *Note: This field is only available after 2019-05-10 and will return &#x60;null&#x60; prior to that time.* | [default to null]
**ActiveExchanges** | **float32** | Number of active exchanges tracked by CoinMarketCap at the given point in time. This includes all exchanges with a &#x60;listing_status&#x60; of \&quot;active\&quot; or \&quot;untracked\&quot; as returned by our /exchange/map call. *Note: This field is only available after 2019-06-18 and will return &#x60;null&#x60; prior to that time.* | [default to null]
**ActiveMarketPairs** | **float32** | Number of active market pairs tracked by CoinMarketCap across all exchanges at the given point in time. *Note: This field is only available after 2019-05-10 and will return &#x60;null&#x60; prior to that time.* | [default to null]
**Quote** | [***GlobalMetricsQuotesHistoricQuoteCurrencyMap**](Global Metrics Quotes Historic - Quote currency map.md) | An object containing market data for this interval by currency option. The default currency mapped is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


