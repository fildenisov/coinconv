# ExchangeQuotesLatestExchangeObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The CoinMarketCap exchange ID. | [default to null]
**Name** | **string** | The exchange name. | [default to null]
**Slug** | **string** | The exchange slug. | [default to null]
**NumMarketPairs** | **int32** | The number of active trading pairs available for this exchange. | [default to null]
**ExchangeScore** | **float32** | The exchange score. | [optional] [default to null]
**LiquidityScore** | **float32** | The liquidity score. | [optional] [default to null]
**Rank** | **int32** | The exchange rank. | [optional] [default to null]
**TrafficScore** | **float32** | The traffic score. | [optional] [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of the last time this exchange&#39;s market data was updated. | [default to null]
**Quote** | [***ExchangeQuotesLatestQuoteMap**](Exchange Quotes Latest - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


