# ExchangeListingsLatestExchangeObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this exchange. | [default to null]
**Name** | **string** | The name of this exchange. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this exchange name. | [default to null]
**NumMarketPairs** | **string** | The number of trading pairs actively tracked on this exchange. | [optional] [default to null]
**DateLaunched** | **string** | Timestamp (ISO 8601) of the date this exchange launched. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**ExchangeScore** | **float32** | The exchange score. | [optional] [default to null]
**LiquidityScore** | **float32** | The liquidity score. | [optional] [default to null]
**Rank** | **int32** | The exchange rank. | [optional] [default to null]
**TrafficScore** | **float32** | The traffic score. | [optional] [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of the last time this record was upated. | [default to null]
**Quote** | [***ExchangeListingsLatestQuoteMap**](Exchange Listings Latest - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


