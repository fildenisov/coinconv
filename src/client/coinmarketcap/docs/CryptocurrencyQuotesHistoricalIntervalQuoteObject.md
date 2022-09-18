# CryptocurrencyQuotesHistoricalIntervalQuoteObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Timestamp of when this historical quote was recorded. | [default to null]
**SearchInterval** | **string** | The interval timestamp for the search period that this historical quote was located against. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**Quote** | [***CryptocurrencyQuotesHistoricalQuoteCurrencyMap**](Cryptocurrency Quotes Historical - Quote currency map.md) | A map of market details for this quote in different currency conversions. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


