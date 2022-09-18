# GlobalMetricsQuotesHistoricCurrencyQuoteObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMarketCap** | **float32** | The sum of all individual cryptocurrency market capitalizations at the given point in time, historically converted into units of the requested currency. | [default to null]
**TotalVolume24h** | **float32** | The sum of rolling 24 hour adjusted volume (as outlined in our methodology) for all cryptocurrencies at the given point in time, historically converted into units of the requested currency. | [default to null]
**TotalVolume24hReported** | **float32** | The sum of rolling 24 hour reported volume for all cryptocurrencies at the given point in time, historically converted into units of the requested currency. *Note: This field is only available after 2019-05-10 and will return &#x60;null&#x60; prior to that time.* | [default to null]
**AltcoinMarketCap** | **float32** | The sum of rolling 24 hour adjusted volume (as outlined in our methodology) for all cryptocurrencies excluding Bitcoin at the given point in time, historically converted into units of the requested currency. | [default to null]
**AltcoinVolume24h** | **float32** | The sum of all individual cryptocurrency market capitalizations excluding Bitcoin at the given point in time, historically converted into units of the requested currency. | [default to null]
**AltcoinVolume24hReported** | **float32** | The sum of rolling 24 hour reported volume for all cryptocurrencies excluding Bitcoin at the given point in time, historically converted into units of the requested currency. *Note: This field is only available after 2019-05-10 and will return &#x60;null&#x60; prior to that time.* | [default to null]
**Timestamp** | **string** | Timestamp (ISO 8601) of when the conversion currency&#39;s current value was referenced for this conversion. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


