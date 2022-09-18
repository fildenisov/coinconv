# GlobalMetricsQuotesLatestResultsObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtcDominance** | **float32** | Bitcoin&#39;s market dominance percentage by market cap. | [optional] [default to null]
**EthDominance** | **float32** | Ethereum&#39;s market dominance percentage by market cap. | [optional] [default to null]
**ActiveCryptocurrencies** | **float32** | Count of active cryptocurrencies tracked by CoinMarketCap. This includes all cryptocurrencies with a &#x60;listing_status&#x60; of \&quot;active\&quot; or \&quot;listed\&quot; as returned from our /cryptocurrency/map call. | [default to null]
**TotalCryptocurrencies** | **float32** | Count of all cryptocurrencies tracked by CoinMarketCap. This includes \&quot;inactive\&quot; &#x60;listing_status&#x60; cryptocurrencies. | [default to null]
**ActiveMarketPairs** | **float32** | Count of active market pairs tracked by CoinMarketCap across all exchanges. | [default to null]
**ActiveExchanges** | **float32** | Count of active exchanges tracked by CoinMarketCap. This includes all exchanges with a &#x60;listing_status&#x60; of \&quot;active\&quot; or \&quot;listed\&quot; as returned by our /exchange/map call. | [default to null]
**TotalExchanges** | **float32** | Count of all exchanges tracked by CoinMarketCap. This includes \&quot;inactive\&quot; &#x60;listing_status&#x60; exchanges. | [default to null]
**LastUpdated** | **string** | Timestamp of when this record was last updated. | [default to null]
**Quote** | [***GlobalMetricsQuotesLatestQuoteMap**](Global Metrics Quotes Latest - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


