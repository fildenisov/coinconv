# CryptocurrencyMarketPairsLatestExchangeInfoObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the exchange this market pair is under. | [default to null]
**Slug** | **string** | The slug of the exchange this market pair is under. | [default to null]
**Name** | **string** | The name of the exchange this market pair is under. | [default to null]
**Notice** | **string** | A [Markdown](https://commonmark.org/help/) formatted message outlining a condition that is impacting the availability of this exchange&#39;s market data or the secure use of the exchange, otherwise null. This may include a maintenance event on the exchange&#39;s end or CoinMarketCap&#39;s end, an alert about reported issues with withdrawls from this exchange, or another condition that may be impacting this exchange and it&#39;s markets. If present, this notice is also displayed in an alert banner at the top of the exchange&#39;s page on coinmarketcap.com. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


