# ExchangeMapExchangeObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this exchange. | [default to null]
**Name** | **string** | The name of this exchange. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this exchange name. | [default to null]
**IsActive** | **int32** | 1 if this exchange is still being actively tracked and updated, otherwise 0. | [optional] [default to null]
**Status** | **string** | The listing status of the exchange. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**FirstHistoricalData** | **string** | Timestamp (ISO 8601) of the earliest market data record available to query using our historical endpoints. &#x60;null&#x60; if there is no historical data currently available for this exchange. | [optional] [default to null]
**LastHistoricalData** | **string** | Timestamp (ISO 8601) of the latest market data record available to query using our historical endpoints. &#x60;null&#x60; if there is no historical data currently available for this exchange. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


