# CryptocurrencyListingsLatestCryptocurrencyObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this cryptocurrency. | [default to null]
**Name** | **string** | The name of this cryptocurrency. | [default to null]
**Symbol** | **string** | The ticker symbol for this cryptocurrency. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this cryptocurrency name. | [default to null]
**CmcRank** | **int32** | The cryptocurrency&#39;s historic CoinMarketCap rank at the end of the requested UTC day. | [default to null]
**NumMarketPairs** | **int32** | The number of active trading pairs available for this cryptocurrency across supported exchanges. | [optional] [default to null]
**CirculatingSupply** | **float32** | The approximate number of coins circulating for this cryptocurrency at the end of the requested UTC day. | [default to null]
**TotalSupply** | **float32** | The approximate total amount of coins in existence right now (minus any coins that have been verifiably burned) at the end of the requested UTC day. | [default to null]
**MaxSupply** | **float32** | The expected maximum limit of coins ever to be available for this cryptocurrency. | [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of when this cryptocurrency&#39;s market data was referenced for this UTC date snapshot. This is always the last update available during the UTC date requested. | [default to null]
**DateAdded** | **string** | Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap. | [default to null]
**Tags** | [***Tags**](tags.md) | Array of tags associated with this cryptocurrency. Currently only a mineable tag will be returned if the cryptocurrency is mineable. Additional tags will be returned in the future. | [default to null]
**Platform** | [***Platform**](platform.md) | Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token, otherwise null. | [default to null]
**Quote** | [***CryptocurrencyListingsLatestQuoteMap**](Cryptocurrency Listings Latest - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


