# ExchangeMarketPairsLatestMarketPairInfoObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketId** | **int32** | The CoinMarketCap ID for this market pair. This ID can reliably be used to identify this unique market as the ID never changes. | [default to null]
**MarketPair** | **string** | The name of this market pair. Example: \&quot;BTC/USD\&quot; | [default to null]
**Category** | **string** | The category of trading this market falls under. Spot markets are the most common but options include derivatives and OTC. | [default to null]
**FeeType** | **string** | The fee type the exchange enforces for this market. | [optional] [default to null]
**MarketUrl** | **string** | The URL to this market&#39;s trading page on the exchange if available. If not available the exchange&#39;s homepage URL is returned. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**MarkPairBase** | [***ExchangeMarketPairsLatestPairBaseCurrencyInfoObject**](Exchange Market Pairs Latest - Pair Base Currency Info object.md) | Base currency details object for this market pair. | [default to null]
**MarkPairQuote** | [***ExchangeMarketPairsLatestPairBaseCurrencyInfoObject1**](Exchange Market Pairs Latest - Pair Base Currency Info object 1.md) | Quote (secondary) currency details object for this market pair | [default to null]
**Quote** | [***ExchangeMarketPairsLatestMarketPairQuoteObject**](Exchange Market Pairs Latest - Market Pair Quote object.md) | Market Pair quotes object containing key-&gt;quote objects for each convert option requested. USD and \&quot;exchange_reported\&quot; are defaults. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


