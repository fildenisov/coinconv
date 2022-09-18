# Status

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Current ISO 8601 timestamp on the server. | [default to null]
**ErrorCode** | **int32** | An internal error code for the current error. If a unique platform error code is not available the HTTP status code is returned. | [default to null]
**ErrorMessage** | **string** | An error message to go along with the error code. | [default to null]
**Elapsed** | **int32** | Number of milliseconds taken to generate this response | [default to null]
**CreditCount** | **int32** | Number of API call credits required for this call. Always 0 for errors. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


