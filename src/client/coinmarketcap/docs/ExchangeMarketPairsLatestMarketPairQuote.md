# ExchangeMarketPairsLatestMarketPairQuote

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** | The last reported exchange price for this market pair converted into the requested convert currency. | [default to null]
**PriceQuote** | **float32** | The latest exchange reported price in base units converted into the requested convert currency. *This field is only returned if requested through the &#x60;aux&#x60; request parameter.* | [optional] [default to null]
**Volume24h** | **float32** | The last reported exchange volume for this market pair converted into the requested convert currency. | [default to null]
**DepthNegativeTwo** | **float32** | -2% Depth in the specified currency. | [optional] [default to null]
**DepthPositiveTwo** | **float32** | +2% Depth in the specified currency. | [optional] [default to null]
**LastUpdated** | **string** | Timestamp (ISO 8601) of when the conversion currency&#39;s current value was referenced for this conversion. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


