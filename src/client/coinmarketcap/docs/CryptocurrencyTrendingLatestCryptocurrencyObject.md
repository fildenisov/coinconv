# CryptocurrencyTrendingLatestCryptocurrencyObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this cryptocurrency. | [default to null]
**Name** | **string** | The name of this cryptocurrency. | [default to null]
**Symbol** | **string** | The ticker symbol for this cryptocurrency. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this cryptocurrency name. | [default to null]
**CmcRank** | **int32** | The cryptocurrency&#39;s CoinMarketCap rank by market cap. | [optional] [default to null]
**IsFiat** | **int32** | 1 if this is a fiat | [optional] [default to null]
**SelfReportedCirculatingSupply** | **float32** | The self reported number of coins circulating for this cryptocurrency. | [optional] [default to null]
**SelfReportedMarketCap** | **float32** | The self reported market cap for this cryptocurrency. | [optional] [default to null]
**IsActive** | **int32** | 1 if this cryptocurrency has at least 1 active market currently being tracked by the platform, otherwise 0. A value of 1 is analogous with &#x60;listing_status&#x3D;active&#x60;. | [optional] [default to null]
**NumMarketPairs** | **int32** | The number of active trading pairs available for this cryptocurrency across supported exchanges. | [optional] [default to null]
**CirculatingSupply** | **float32** | The approximate number of coins circulating for this cryptocurrency. | [optional] [default to null]
**TotalSupply** | **float32** | The approximate total amount of coins in existence right now (minus any coins that have been verifiably burned). | [optional] [default to null]
**MarketCapByTotalSupply** | **float32** | The market cap by total supply. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**MaxSupply** | **float32** | The expected maximum limit of coins ever to be available for this cryptocurrency. | [optional] [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of the last time this cryptocurrency&#39;s market data was updated. | [default to null]
**DateAdded** | **string** | Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap. | [optional] [default to null]
**Tags** | [***Tags**](tags.md) | Array of tags associated with this cryptocurrency. Currently only a mineable tag will be returned if the cryptocurrency is mineable. Additional tags will be returned in the future. | [optional] [default to null]
**Platform** | [***Platform**](platform.md) | Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token, otherwise null. | [optional] [default to null]
**Quote** | [***CryptocurrencyQuoteMap**](Cryptocurrency - Quote map.md) | A map of market quotes in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


