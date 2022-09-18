# CryptocurrencyOhlcvLatestCryptocurrencyObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this cryptocurrency. | [default to null]
**Name** | **string** | The name of this cryptocurrency. | [default to null]
**Symbol** | **string** | The ticker symbol for this cryptocurrency. | [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of the lastest market value record included to generate the latest active day OHLCV values. | [default to null]
**TimeOpen** | **string** | Timestamp (ISO 8601) of the start of this OHLCV period. | [default to null]
**TimeHigh** | **string** | Timestamp (ISO 8601) of the high of this OHLCV period. | [default to null]
**TimeLow** | **string** | Timestamp (ISO 8601) of the low of this OHLCV period. | [default to null]
**TimeClose** | **string** | Timestamp (ISO 8601) of the end of this OHLCV period. Always &#x60;null&#x60; as the current day is incomplete. See &#x60;last_updated&#x60; for the last UTC time included in the current OHLCV calculation. | [default to null]
**Quote** | [***CryptocurrencyOhlcvLatestQuoteMap**](Cryptocurrency OHLCV Latest - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


