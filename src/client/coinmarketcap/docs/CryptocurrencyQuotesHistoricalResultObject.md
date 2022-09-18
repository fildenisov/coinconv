# CryptocurrencyQuotesHistoricalResultObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The CoinMarketCap cryptocurrency ID. | [default to null]
**Name** | **string** | The cryptocurrency name. | [default to null]
**Symbol** | **string** | The cryptocurrency symbol. | [default to null]
**IsActive** | **int32** | 1 if this cryptocurrency has at least 1 active market currently being tracked by the platform, otherwise 0. A value of 1 is analogous with &#x60;listing_status&#x3D;active&#x60;. | [optional] [default to null]
**IsFiat** | **int32** | 1 if this is a fiat | [optional] [default to null]
**Quotes** | [***CryptocurrencyQuotesHistoricalIntervalQuotesArray**](Cryptocurrency Quotes Historical - Interval Quotes array.md) | An array of quotes for each interval for this cryptocurrency. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


