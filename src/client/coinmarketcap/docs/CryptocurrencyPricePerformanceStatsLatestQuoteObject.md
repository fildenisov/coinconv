# CryptocurrencyPricePerformanceStatsLatestQuoteObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Open** | **float32** | Cryptocurrency price at the start of the requested time period historically converted into units of the convert currency. | [default to null]
**OpenTimestamp** | **string** | Timestamp (ISO 8601) of the closest convert currency reference price used during &#x60;open&#x60; price conversion. | [default to null]
**High** | **float32** | Highest USD price achieved within the requested time period historically converted into units of the convert currency. | [default to null]
**HighTimestamp** | **string** | Timestamp (ISO 8601) of the closest convert currency reference price used during &#x60;high&#x60; price conversion. *For &#x60;yesterday&#x60; UTC close will be used.* | [default to null]
**Low** | **float32** | Lowest USD price achieved within the requested time period historically converted into units of the convert currency. | [default to null]
**LowTimestamp** | **string** | Timestamp (ISO 8601) of the closest convert currency reference price used during &#x60;low&#x60; price conversion. *For &#x60;yesterday&#x60; UTC close will be used.* | [default to null]
**Close** | **float32** | Cryptocurrency price at the end of the requested time period historically converted into units of the convert currency. | [default to null]
**CloseTimestamp** | **string** | Timestamp (ISO 8601) of the closest convert currency reference price used during &#x60;close&#x60; price conversion. | [default to null]
**PercentChange** | **float32** | The approximate percentage change (ROI) if purchased at the start of the time period. This is the time of launch or earliest known price for the &#x60;all_time&#x60; period. This value includes historical change in market rate for the specified convert currency. | [default to null]
**PriceChange** | **float32** | The actual price change between the start of the time period and end. This is the time of launch or earliest known price for the &#x60;all_time&#x60; period. This value includes historical change in market rate for the specified convert currency. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


