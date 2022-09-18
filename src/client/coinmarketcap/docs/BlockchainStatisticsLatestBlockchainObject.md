# BlockchainStatisticsLatestBlockchainObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this blockchain&#39;s cryptocurrency. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of the cryptocurrency&#39;s name. | [default to null]
**Symbol** | **string** | The ticker symbol for the cryptocurrency. | [default to null]
**BlockRewardStatic** | **float32** | The reward assigned to the miner of a block excluding fees. | [default to null]
**ConsensusMechanism** | **string** | The consensus mechanism used by the blockchain, for example, \&quot;proof-of-work\&quot; or \&quot;proof-of-stake\&quot;. | [default to null]
**Difficulty** | **string** | The global block difficulty determining how hard to find a hash on this blockchain. *Note: This integer is returned as a string to use with BigInt libraries as it may exceed the max safe integer size for many programming languages.* | [default to null]
**Hashrate24h** | **string** | The average hashrate over the past 24 hours. *Note: This integer is returned as a string to use with BigInt libraries as it may exceed the max safe integer size for many programming languages.* | [default to null]
**PendingTransactions** | **int32** | The number of pending transactions. | [default to null]
**ReductionRate** | **string** | The rate the block reward is adjusted at a specified interval. | [default to null]
**TotalBlocks** | **int32** | The total number of blocks. | [default to null]
**TotalTransactions** | **string** | The total number of transactions. *Note: This integer is returned as a string to use with BigInt libraries as it may exceed the max safe integer size for many programming languages.* | [default to null]
**Tps24h** | **float32** | The average transactions per second over the past 24 hours. | [default to null]
**FirstBlockTimestamp** | **string** | Timestamp (ISO 8601) of the time the first block was mined on this chain. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


