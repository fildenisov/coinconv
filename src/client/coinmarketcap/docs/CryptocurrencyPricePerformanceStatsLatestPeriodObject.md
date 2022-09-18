# CryptocurrencyPricePerformanceStatsLatestPeriodObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenTimestamp** | **string** | Timestamp (ISO 8601) of the start of this time period. Please note that this is a rolling period back from current time for time periods outside of &#x60;yesterday&#x60;. | [default to null]
**HighTimestamp** | **string** | Timestamp (ISO 8601) of when this cryptocurrency achieved it&#39;s highest USD price during the requested time period. *Note: The &#x60;yesterday&#x60; period currently doesn&#39;t support this field and will return &#x60;null&#x60;.* | [default to null]
**LowTimestamp** | **string** | Timestamp (ISO 8601) of when this cryptocurrency achieved it&#39;s lowest USD price during the requested time period. *Note: The &#x60;yesterday&#x60; period currently doesn&#39;t support this field and will return &#x60;null&#x60;.* | [default to null]
**CloseTimestamp** | **string** | Timestamp (ISO 8601) of the end of this time period. Please note that this is a rolling period back from current time for time periods outside of &#x60;yesterday&#x60;. | [default to null]
**Quote** | [***CryptocurrencyPricePerformanceStatsLatestQuoteMap**](Cryptocurrency Price Performance Stats Latest - Quote map.md) | An object map of time period quotes for each convert option requested. The default map included is USD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


