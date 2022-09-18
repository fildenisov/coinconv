# coinconv
Is a test task project.

## Task details
1. Create cli utility which converts one currency into another using [coinmarketcap API](https://coinmarketcap.com/api/v1/#section/Introduction).
    Example:
    ```
    $> ./coinconv 123.45 USD BTC
    654.05
    ```

2. Code should be uploaded to github.
3. Use the test API key `b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c`
4. Take care about clean architecture and SOLID principles. 

## Solution details
### /src/client/coinmarketcap 
Is an auto generated REST client. Downloaded from here -> [swagger](https://pro-api.coinmarketcap.com/swagger.json) 