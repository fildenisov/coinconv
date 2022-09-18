# ExchangesInfoExchangeInfoObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this exchange. | [default to null]
**Name** | **string** | The name of this exchange. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of the exchange name. | [default to null]
**Logo** | **string** | Link to a CoinMarketCap hosted logo png for this exchange. 64px is default size returned. Replace \&quot;64x64\&quot; in the image path with these alternative sizes: 16, 32, 64, 128, 200 | [default to null]
**Description** | **string** | A CoinMarketCap supplied brief description of this cryptocurrency exchange. This field will return null if a description is not available. | [default to null]
**DateLaunched** | **string** | Timestamp (ISO 8601) of the launch date for this exchange. | [default to null]
**Notice** | **string** | A [Markdown](https://commonmark.org/help/) formatted message outlining a condition that is impacting the availability of the exchange&#39;s market data or the secure use of the exchange, otherwise null. This may include a maintenance event on the exchange&#39;s end or CoinMarketCap&#39;s end, an alert about reported issues with withdrawls from this exchange, or another condition that may be impacting the exchange and it&#39;s markets. If present, this notice is also displayed in an alert banner at the top of the exchange&#39;s page on coinmarketcap.com. | [default to null]
**WeeklyVisits** | **float32** | The number of weekly visitors. | [optional] [default to null]
**SpotVolumeUsd** | **float32** | Reported all time spot volume in the specified currency. | [optional] [default to null]
**Urls** | [***ExchangesInfoUrLsObject**](Exchanges Info - URLs object.md) | An object containing various resource URLs for this exchange. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


