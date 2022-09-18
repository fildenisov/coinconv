# Plan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditLimitDaily** | **float32** | The number of API credits that can be used each daily period before receiving a HTTP 429 rate limit error. This limit is based on the API plan tier. | [default to null]
**CreditLimitDailyReset** | **string** | A human readable countdown of when the API key daily credit limit will reset back to 0. | [default to null]
**CreditLimitDailyResetTimestamp** | **string** | Timestamp (ISO 8601) of when the daily credit limit will reset. This is based on your billing plan activation date for premium subscription based keys or UTC midnight for free Basic plan keys. | [default to null]
**CreditLimitMonthly** | **float32** | The number of API credits that can be used each monthly period before receiving a HTTP 429 rate limit error. This limit is based on the API plan tier. | [default to null]
**CreditLimitMonthlyReset** | **string** | A human readable countdown of when the API key monthly credit limit will reset back to 0. | [default to null]
**CreditLimitMonthlyResetTimestamp** | **string** | Timestamp (ISO 8601) of when the monthly credit limit will reset. This is based on your billing plan activation date for premium subscription based keys or calendar month UTC midnight for free Basic plan keys. | [default to null]
**RateLimitMinute** | **float32** | The number of API calls that can be made within the same UTC minute before receiving a HTTP 429 rate limit error. This limit is based on the API plan tier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


