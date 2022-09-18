/*
 * CoinMarketCap Cryptocurrency API Documentation
 *
 * # Introduction The CoinMarketCap API is a suite of high-performance RESTful JSON endpoints that are specifically designed to meet the mission-critical demands of application developers, data scientists, and enterprise business platforms.  This API reference includes all technical documentation developers need to integrate third-party applications and platforms. Additional answers to common questions can be found in the [CoinMarketCap API FAQ](https://coinmarketcap.com/api/faq).  # Quick Start Guide  For developers eager to hit the ground running with the CoinMarketCap API here are a few quick steps to make your first call with the API.  1. **Sign up for a free Developer Portal account.** You can sign up at [pro.coinmarketcap.com](https://pro.coinmarketcap.com) - This is our live production environment with the latest market data. Select the free `Basic` plan if it meets your needs or upgrade to a paid tier. 2. **Copy your API Key.** Once you sign up you'll land on your Developer Portal account dashboard. Copy your API from the `API Key` box in the top left panel. 3. **Make a test call using your key.** You may use the code examples provided below to make a test call with your programming language of choice. This example [fetches all active cryptocurrencies by market cap and return market values in USD](https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?start=1&limit=5000&convert=USD).     *Be sure to replace the API Key in sample code with your own and use API domain `pro-api.coinmarketcap.com` or use the test API Key `b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c` for `sandbox-api.coinmarketcap.com` testing with our sandbox environment. Please note that our sandbox api has mock data and should not be used in your application.* 4. **Postman Collection** To help with development, we provide a fully featured postman collection that you can import and use immediately! [Read more here](https://coinmarketcap.com/alexandria/article/register-for-coinmarketcap-api). 5. **Implement your application.** Now that you've confirmed your API Key is working, get familiar with the API by reading the rest of this API Reference and commence building your application!  ***Note:** Making HTTP requests on the client side with Javascript is currently prohibited through CORS configuration. This is to protect your API Key which should not be visible to users of your application so your API Key is not stolen. Secure your API Key by routing calls through your own backend service.*  <details open=\"open\"> <summary>**View Quick Start Code Examples**</summary> <details open=\"open\"> <summary>cURL command line</summary> ```bash  curl -H \"X-CMC_PRO_API_KEY: b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c\" -H \"Accept: application/json\" -d \"start=1&limit=5000&convert=USD\" -G https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest  ```     </details><details> <summary>Node.js</summary> ```javascript   /_* Example in Node.js *_/ const axios = require('axios');  let response = null; new Promise(async (resolve, reject) => {   try {     response = await axios.get('https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest', {       headers: {         'X-CMC_PRO_API_KEY': 'b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c',       },     });   } catch(ex) {     response = null;     // error     console.log(ex);     reject(ex);   }   if (response) {     // success     const json = response.data;     console.log(json);     resolve(json);   } });  ```   </details><details> <summary>Python</summary> ```python     #This example uses Python 2.7 and the python-request library.  from requests import Request, Session from requests.exceptions import ConnectionError, Timeout, TooManyRedirects import json  url = 'https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest' parameters = {   'start':'1',   'limit':'5000',   'convert':'USD' } headers = {   'Accepts': 'application/json',   'X-CMC_PRO_API_KEY': 'b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c', }  session = Session() session.headers.update(headers)  try:   response = session.get(url, params=parameters)   data = json.loads(response.text)   print(data) except (ConnectionError, Timeout, TooManyRedirects) as e:   print(e)    ``` </details><details> <summary>PHP</summary> ```php  /_**  * Requires curl enabled in php.ini  **_/  <?php $url = 'https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest'; $parameters = [   'start' => '1',   'limit' => '5000',   'convert' => 'USD' ];  $headers = [   'Accepts: application/json',   'X-CMC_PRO_API_KEY: b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c' ]; $qs = http_build_query($parameters); // query string encode the parameters $request = \"{$url}?{$qs}\"; // create the request URL   $curl = curl_init(); // Get cURL resource // Set cURL options curl_setopt_array($curl, array(   CURLOPT_URL => $request,            // set the request URL   CURLOPT_HTTPHEADER => $headers,     // set the headers    CURLOPT_RETURNTRANSFER => 1         // ask for raw response instead of bool ));  $response = curl_exec($curl); // Send the request, save the response print_r(json_decode($response)); // print json decoded response curl_close($curl); // Close request ?>  ``` </details><details> <summary>Java</summary> ```java  /_**   * This example uses the Apache HTTPComponents library.   *_/  import org.apache.http.HttpEntity; import org.apache.http.HttpHeaders; import org.apache.http.NameValuePair; import org.apache.http.client.methods.CloseableHttpResponse; import org.apache.http.client.methods.HttpGet; import org.apache.http.client.utils.URIBuilder; import org.apache.http.impl.client.CloseableHttpClient; import org.apache.http.impl.client.HttpClients; import org.apache.http.message.BasicNameValuePair; import org.apache.http.util.EntityUtils;  import java.io.IOException; import java.net.URISyntaxException; import java.util.ArrayList; import java.util.List;  public class JavaExample {    private static String apiKey = \"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c\";    public static void main(String[] args) {     String uri = \"https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest\";     List<NameValuePair> paratmers = new ArrayList<NameValuePair>();     paratmers.add(new BasicNameValuePair(\"start\",\"1\"));     paratmers.add(new BasicNameValuePair(\"limit\",\"5000\"));     paratmers.add(new BasicNameValuePair(\"convert\",\"USD\"));      try {       String result = makeAPICall(uri, paratmers);       System.out.println(result);     } catch (IOException e) {       System.out.println(\"Error: cannont access content - \" + e.toString());     } catch (URISyntaxException e) {       System.out.println(\"Error: Invalid URL \" + e.toString());     }   }    public static String makeAPICall(String uri, List<NameValuePair> parameters)       throws URISyntaxException, IOException {     String response_content = \"\";      URIBuilder query = new URIBuilder(uri);     query.addParameters(parameters);      CloseableHttpClient client = HttpClients.createDefault();     HttpGet request = new HttpGet(query.build());      request.setHeader(HttpHeaders.ACCEPT, \"application/json\");     request.addHeader(\"X-CMC_PRO_API_KEY\", apiKey);      CloseableHttpResponse response = client.execute(request);      try {       System.out.println(response.getStatusLine());       HttpEntity entity = response.getEntity();       response_content = EntityUtils.toString(entity);       EntityUtils.consume(entity);     } finally {       response.close();     }      return response_content;   }  }  ``` </details><details> <summary>C#</summary> ```csharp  using System; using System.Net; using System.Web;  class CSharpExample {   private static string API_KEY = \"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c\";    public static void Main(string[] args)   {     try     {     Console.WriteLine(makeAPICall());     }     catch (WebException e)     {     Console.WriteLine(e.Message);     }   }    static string makeAPICall()   {     var URL = new UriBuilder(\"https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest\");      var queryString = HttpUtility.ParseQueryString(string.Empty);     queryString[\"start\"] = \"1\";     queryString[\"limit\"] = \"5000\";     queryString[\"convert\"] = \"USD\";      URL.Query = queryString.ToString();      var client = new WebClient();     client.Headers.Add(\"X-CMC_PRO_API_KEY\", API_KEY);     client.Headers.Add(\"Accepts\", \"application/json\");     return client.DownloadString(URL.ToString());    }  }      ``` </details><details> <summary>Go</summary> ```go  package main  import (   \"fmt\"   \"io/ioutil\"   \"log\"   \"net/http\"   \"net/url\"   \"os\" )  func main() {   client := &http.Client{}   req, err := http.NewRequest(\"GET\",\"https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest\", nil)   if err != nil {     log.Print(err)     os.Exit(1)   }    q := url.Values{}   q.Add(\"start\", \"1\")   q.Add(\"limit\", \"5000\")   q.Add(\"convert\", \"USD\")    req.Header.Set(\"Accepts\", \"application/json\")   req.Header.Add(\"X-CMC_PRO_API_KEY\", \"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c\")   req.URL.RawQuery = q.Encode()     resp, err := client.Do(req);   if err != nil {     fmt.Println(\"Error sending request to server\")     os.Exit(1)   }   fmt.Println(resp.Status);   respBody, _ := ioutil.ReadAll(resp.Body)   fmt.Println(string(respBody));  }  ``` </details></details>  # Authentication  ### Acquiring an API Key All HTTP requests made against the CoinMarketCap API must be validated with an API Key. If you don't have an API Key yet visit the [API Developer Portal](https://coinmarketcap.com/api/) to register for one.  ### Using Your API Key You may use any *server side* programming language that can make HTTP requests to target the CoinMarketCap API. All requests should target domain `https://pro-api.coinmarketcap.com`.  You can supply your API Key in REST API calls in one of two ways:  1. Preferred method: Via a custom header named `X-CMC_PRO_API_KEY` 2. Convenience method: Via a query string parameter named `CMC_PRO_API_KEY`  ***Security Warning:** It's important to secure your API Key against public access. The custom header option is strongly recommended over the querystring option for passing your API Key in a production environment.*  ### API Key Usage Credits  Most API plans include a daily and monthly limit or \"hard cap\" to the number of data calls that can be made. This usage is tracked as API \"call credits\" which are incremented 1:1 against successful (HTTP Status 200) data calls made with your key with these exceptions: - Account management endpoints, usage stats endpoints, and error responses are not included in this limit.   - **Paginated endpoints:** List-based endpoints track an additional call credit for every 100 data points returned (rounded up) beyond our 100 data point defaults. Our lightweight `/map` endpoints are not included in this limit and always count as 1 credit. See individual endpoint documentation for more details.    - **Bundled API calls:** Many endpoints support [resource and currency conversion bundling](#section/Standards-and-Conventions). Bundled resources are also tracked as 1 call credit for every 100 resources returned (rounded up). Optional currency conversion bundling using the `convert` parameter also increment an additional API call credit for every conversion requested beyond the first.  You can log in to the [Developer Portal](https://coinmarketcap.com/api/) to view live stats on your API Key usage and limits including the number of credits used for each call. You can also find call credit usage in the JSON response for each API call. See the [`status` object](#section/Standards-and-Conventions) for details. You may also use the [/key/info](#operation/getV1KeyInfo) endpoint to quickly review your usage and when daily/monthly credits reset directly from the API.    ***Note:** \"day\" and \"month\" credit usage periods are defined relative to your API subscription. For example, if your monthly subscription started on the 5th at 5:30am, this billing anchor is also when your monthly credits refresh each month. The free Basic tier resets each day at UTC midnight and each calendar month at UTC midnight.*     # Endpoint Overview  ### The CoinMarketCap API is divided into 8 top-level categories Endpoint Category  | Description -------------------|--------------- [/cryptocurrency/_*](#tag/cryptocurrency) | Endpoints that return data around cryptocurrencies such as ordered cryptocurrency lists or price and volume data. [/exchange/_*](#tag/exchange) | Endpoints that return data around cryptocurrency exchanges such as ordered exchange lists and market pair data. [/global-metrics/_*](#tag/global-metrics) | Endpoints that return aggregate market data such as global market cap and BTC dominance. [/tools/_*](#tag/tools)  | Useful utilities such as cryptocurrency and fiat price conversions. [/blockchain/_*](#tag/blockchain)  | Endpoints that return block explorer related data for blockchains. [/fiat/_*](#tag/fiat) | Endpoints that return data around fiats currencies including mapping to CMC IDs. [/partners/_*](#tag/partners)  | Endpoints for convenient access to 3rd party crypto data. [/key/_*](#tag/key)  | API key administration endpoints to review and manage your usage. [/content/_*](#tag/content)  | Endpoints that return cryptocurrency-related news, headlines, articles, posts, and comments.  ### Endpoint paths follow a pattern matching the type of data provided  Endpoint Path  | Endpoint Type | Description ----------------------|-------------|--------- *_/latest | Latest Market Data | Latest market ticker quotes and averages for cryptocurrencies and exchanges. *_/historical | Historical Market Data | Intervals of historic market data like OHLCV data or data for use in charting libraries. *_/info | Metadata | Cryptocurrency and exchange metadata like block explorer URLs and logos. *_/map | ID Maps | Utility endpoints to get a map of resources to CoinMarketCap IDs.  ###  Cryptocurrency and exchange endpoints provide 2 different ways of accessing data depending on purpose  - **Listing endpoints:** Flexible paginated `*_/listings/_*` endpoints allow you to sort and filter lists of data like cryptocurrencies by market cap or exchanges by volume. - **Item endpoints:** Convenient ID-based resource endpoints like `*_/quotes/_*` and `*_/market-pairs/_*` allow you to bundle several IDs; for example, this allows you to get latest market quotes for a specific set of cryptocurrencies in one call.  # Standards and Conventions  Each HTTP request must contain the header `Accept: application/json`. You should also send an `Accept-Encoding: deflate, gzip` header to receive data fast and efficiently.  ### Endpoint Response Payload Format All endpoints return data in JSON format with the results of your query under `data` if the call is successful.  A `Status` object is always included for both successful calls and failures when possible. The `Status` object always includes the current time on the server when the call was executed as `timestamp`, the number of API call credits this call utilized as `credit_count`, and the number of milliseconds it took to process the request as `elapsed`. Any details about errors encountered can be found under the `error_code` and `error_message`. See [Errors and Rate Limits](#section/Errors-and-Rate-Limits) for details on errors.  ``` {   \"data\" : {     ...   },   \"status\": {     \"timestamp\": \"2018-06-06T07:52:27.273Z\",     \"error_code\": 400,     \"error_message\": \"Invalid value for \\\"id\\\"\",     \"elapsed\": 0,     \"credit_count\": 0   } } ```  ### Cryptocurrency, Exchange, and Fiat currency identifiers - Cryptocurrencies may be identified in endpoints using either the cryptocurrency's unique CoinMarketCap ID as `id` (eg. `id=1` for Bitcoin) or the cryptocurrency's symbol (eg. `symbol=BTC` for Bitcoin). For a current list of supported cryptocurrencies use our [`/cryptocurrency/map` call](#operation/getV1CryptocurrencyMap). - Exchanges may be identified in endpoints using either the exchange's unique CoinMarketCap ID as `id` (eg. `id=270` for Binance) or the exchange's web slug (eg. `slug=binance` for Binance). For a current list of supported exchanges use our [`/exchange/map` call](#operation/getV1ExchangeMap). - All fiat currency options use the standard [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) currency code (eg. `USD` for the US Dollar). For a current list of supported fiat currencies use our [`/fiat/map`](#operation/getV1FiatMap) endpoint. Unless otherwise stated, endpoints with fiat currency options like our `convert` parameter support these 93 major currency codes:  Currency | Currency Code | CoinMarketCap ID ---------|---------------|------------- United States Dollar ($) | USD | 2781 Albanian Lek (L) | ALL | 3526 Algerian Dinar (د.ج) | DZD | 3537 Argentine Peso ($) | ARS | 2821 Armenian Dram (֏) | AMD | 3527 Australian Dollar ($) | AUD | 2782 Azerbaijani Manat (₼) | AZN | 3528 Bahraini Dinar (.د.ب) | BHD | 3531 Bangladeshi Taka (৳) | BDT | 3530 Belarusian Ruble (Br) | BYN | 3533 Bermudan Dollar ($) | BMD | 3532 Bolivian Boliviano (Bs.) | BOB | 2832 Bosnia-Herzegovina Convertible Mark (KM) | BAM | 3529 Brazilian Real (R$) | BRL | 2783 Bulgarian Lev (лв) | BGN | 2814 Cambodian Riel (៛) | KHR | 3549 Canadian Dollar ($) | CAD | 2784 Chilean Peso ($) | CLP | 2786 Chinese Yuan (¥) | CNY | 2787 Colombian Peso ($) | COP | 2820 Costa Rican Colón (₡) | CRC | 3534 Croatian Kuna (kn) | HRK | 2815 Cuban Peso ($) | CUP | 3535 Czech Koruna (Kč) | CZK | 2788 Danish Krone (kr) | DKK | 2789 Dominican Peso ($) | DOP | 3536 Egyptian Pound (£) | EGP | 3538 Euro (€) | EUR | 2790 Georgian Lari (₾) | GEL | 3539 Ghanaian Cedi (₵) | GHS | 3540 Guatemalan Quetzal (Q) | GTQ | 3541 Honduran Lempira (L) | HNL | 3542 Hong Kong Dollar ($) | HKD | 2792 Hungarian Forint (Ft) | HUF | 2793 Icelandic Króna (kr) | ISK | 2818 Indian Rupee (₹) | INR | 2796 Indonesian Rupiah (Rp) | IDR | 2794 Iranian Rial (﷼) | IRR | 3544 Iraqi Dinar (ع.د) | IQD | 3543 Israeli New Shekel (₪) | ILS | 2795 Jamaican Dollar ($) | JMD | 3545 Japanese Yen (¥) | JPY | 2797 Jordanian Dinar (د.ا) | JOD | 3546 Kazakhstani Tenge (₸) | KZT | 3551 Kenyan Shilling (Sh) | KES | 3547 Kuwaiti Dinar (د.ك) | KWD | 3550 Kyrgystani Som (с) | KGS | 3548 Lebanese Pound (ل.ل) | LBP | 3552 Macedonian Denar (ден) | MKD | 3556 Malaysian Ringgit (RM) | MYR | 2800 Mauritian Rupee (₨) | MUR | 2816 Mexican Peso ($) | MXN | 2799 Moldovan Leu (L) | MDL | 3555 Mongolian Tugrik (₮) | MNT | 3558 Moroccan Dirham (د.م.) | MAD | 3554 Myanma Kyat (Ks) | MMK | 3557 Namibian Dollar ($) | NAD | 3559 Nepalese Rupee (₨) | NPR | 3561 New Taiwan Dollar (NT$) | TWD | 2811 New Zealand Dollar ($) | NZD | 2802 Nicaraguan Córdoba (C$) | NIO | 3560 Nigerian Naira (₦) | NGN | 2819 Norwegian Krone (kr) | NOK | 2801 Omani Rial (ر.ع.) | OMR | 3562 Pakistani Rupee (₨) | PKR | 2804 Panamanian Balboa (B/.) | PAB | 3563 Peruvian Sol (S/.) | PEN | 2822 Philippine Peso (₱) | PHP | 2803 Polish Złoty (zł) | PLN | 2805 Pound Sterling (£) | GBP | 2791 Qatari Rial (ر.ق) | QAR | 3564 Romanian Leu (lei) | RON | 2817 Russian Ruble (₽) | RUB | 2806 Saudi Riyal (ر.س) | SAR | 3566 Serbian Dinar (дин.) | RSD | 3565 Singapore Dollar (S$) | SGD | 2808 South African Rand (R) | ZAR | 2812 South Korean Won (₩) | KRW | 2798 South Sudanese Pound (£) | SSP | 3567 Sovereign Bolivar (Bs.) | VES | 3573 Sri Lankan Rupee (Rs) | LKR | 3553 Swedish Krona (  kr) | SEK | 2807 Swiss Franc (Fr) | CHF | 2785 Thai Baht (฿) | THB | 2809 Trinidad and Tobago Dollar ($) | TTD | 3569 Tunisian Dinar (د.ت) | TND | 3568 Turkish Lira (₺) | TRY | 2810 Ugandan Shilling (Sh) | UGX | 3570 Ukrainian Hryvnia (₴) | UAH | 2824 United Arab Emirates Dirham (د.إ) | AED | 2813 Uruguayan Peso ($) | UYU | 3571 Uzbekistan Som (so'm) | UZS | 3572 Vietnamese Dong (₫) | VND | 2823    Along with these four precious metals:    Precious Metal | Currency Code | CoinMarketCap ID ---------|---------------|------------- Gold Troy Ounce | XAU | 3575 Silver Troy Ounce | XAG | 3574 Platinum Ounce | XPT | 3577 Palladium Ounce | XPD | 3576   ***Warning:** **Using CoinMarketCap IDs is always recommended as not all cryptocurrency symbols are unique. They can also change with a cryptocurrency rebrand.** If a symbol is used the API will always default to the cryptocurrency with the highest market cap if there are multiple matches. Our `convert` parameter also defaults to fiat if a cryptocurrency symbol also matches a supported fiat currency. You may use the convenient `/map` endpoints to quickly find the corresponding CoinMarketCap ID for a cryptocurrency or exchange.*  ### Bundling API Calls - Many endpoints support ID and crypto/fiat currency conversion bundling. This means you can pass multiple comma-separated values to an endpoint to query or convert several items at once. Check the `id`, `symbol`, `slug`, and `convert` query parameter descriptions in the endpoint documentation to see if this is supported for an endpoint. - Endpoints that support bundling return data as an object map instead of an array. Each key-value pair will use the identifier you passed in as the key.  For example, if you passed `symbol=BTC,ETH` to `/v1/cryptocurrency/quotes/latest` you would receive:  ``` \"data\" : {     \"BTC\" : {       ...     },     \"ETH\" : {       ...     } } ```  Or if you passed `id=1,1027` you would receive:  ``` \"data\" : {     \"1\" : {       ...     },     \"1027\" : {       ...     } } ```  Price conversions that are returned inside endpoint responses behave in the same fashion. These are enclosed in a `quote` object.  ### Date and Time Formats - All endpoints that require date/time parameters allow timestamps to be passed in either [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format (eg. `2018-06-06T01:46:40Z`) or in Unix time (eg. `1528249600`). Timestamps that are passed in ISO 8601 format support basic and extended notations; if a timezone is not included, UTC will be the default. - All timestamps returned in JSON payloads are returned in UTC time using human-readable ISO 8601 format which follows this pattern: `yyyy-mm-ddThh:mm:ss.mmmZ`. The final `.mmm` designates milliseconds. Per the ISO 8601 spec the final `Z` is a constant that represents UTC time. - Data is collected, recorded, and reported in UTC time unless otherwise specified.  ### Versioning The CoinMarketCap API is versioned to guarantee new features and updates are non-breaking. The latest version of this API is `/v1/`.  # Errors and Rate Limits  ### API Request Throttling Use of the CoinMarketCap API is subject to API call rate limiting or \"request throttling\". This is the number of HTTP calls that can be made simultaneously or within the same minute with your API Key before receiving an HTTP 429 \"Too Many Requests\" throttling error. This limit scales with the <a rel=\"noopener noreferrer\" href=\"https://pro.coinmarketcap.com/features\" target=\"_blank\">usage tier</a> and resets every 60 seconds. Please review our <a rel=\"noopener noreferrer\" href=\"#section/Best-Practices\" target=\"_blank\">Best Practices</a> for implementation strategies that work well with rate limiting.  ### HTTP Status Codes The API uses standard HTTP status codes to indicate the success or failure of an API call.  - `400 (Bad Request)` The server could not process the request, likely due to an invalid argument. - `401 (Unauthorized)` Your request lacks valid authentication credentials, likely an issue with your API Key. - `402 (Payment Required)` Your API request was rejected due to it being a paid subscription plan with an overdue balance. Pay the balance in the [Developer Portal billing tab](https://pro.coinmarketcap.com/account/plan) and it will be enabled. - `403 (Forbidden)` Your request was rejected due to a permission issue, likely a restriction on the API Key's associated service plan. Here is a [convenient map](https://coinmarketcap.com/api/features) of service plans to endpoints. - `429 (Too Many Requests)` The API Key's rate limit was exceeded; consider slowing down your API Request frequency if this is an HTTP request throttling error. Consider upgrading your service plan if you have reached your monthly API call credit limit for the day/month. - `500 (Internal Server Error)` An unexpected server issue was encountered.  ### Error Response Codes A `Status` object is always included in the JSON response payload for both successful calls and failures when possible. During error scenarios you may reference the `error_code` and `error_message` properties of the Status object. One of the API error codes below will be returned if applicable otherwise the HTTP status code for the general error type is returned.     HTTP Status | Error Code | Error Message ------------|----------------|------------- 401 | 1001 [API_KEY_INVALID] | This API Key is invalid. 401 | 1002 [API_KEY_MISSING] | API key missing. 402 | 1003 [API_KEY_PLAN_REQUIRES_PAYEMENT] | Your API Key must be activated. Please go to pro.coinmarketcap.com/account/plan. 402 | 1004 [API_KEY_PLAN_PAYMENT_EXPIRED] | Your API Key's subscription plan has expired. 403 | 1005 [API_KEY_REQUIRED] | An API Key is required for this call. 403 | 1006 [API_KEY_PLAN_NOT_AUTHORIZED] | Your API Key subscription plan doesn't support this endpoint. 403 | 1007 [API_KEY_DISABLED] | This API Key has been disabled. Please contact support. 429 | 1008 [API_KEY_PLAN_MINUTE_RATE_LIMIT_REACHED] | You've exceeded your API Key's HTTP request rate limit. Rate limits reset every minute. 429 | 1009 [API_KEY_PLAN_DAILY_RATE_LIMIT_REACHED] | You've exceeded your API Key's daily rate limit. 429 | 1010 [API_KEY_PLAN_MONTHLY_RATE_LIMIT_REACHED] | You've exceeded your API Key's monthly rate limit. 429 | 1011 [IP_RATE_LIMIT_REACHED] | You've hit an IP rate limit.    # Best Practices  This section contains a few recommendations on how to efficiently utilize the CoinMarketCap API for your enterprise application, particularly if you already have a large base of users for your application.  ### Use CoinMarketCap ID Instead of Cryptocurrency Symbol  Utilizing common cryptocurrency symbols to reference cryptocurrencies on the API is easy and convenient but brittle. You should know that many cryptocurrencies have the same symbol, for example, there are currently three cryptocurrencies that commonly refer to themselves by the symbol HOT. Cryptocurrency symbols also often change with cryptocurrency rebrands. When fetching cryptocurrency by a symbol that matches several active cryptocurrencies we return the one with the highest market cap at the time of the query. To ensure you always target the cryptocurrency you expect, use our permanent CoinMarketCap IDs. These IDs are used reliably by numerous mission critical platforms and *never change*.      We make fetching a map of all active cryptocurrencies' CoinMarketCap IDs very easy. Just call our [`/cryptocurrency/map`](#operation/getV1CryptocurrencyMap) endpoint to receive a list of all active currencies mapped to the unique `id` property. This map also includes other typical identifiying properties like `name`, `symbol` and platform `token_address` that can be cross referenced. In cryptocurrency calls you would then send, for example `id=1027`, instead of `symbol=ETH`. **It's strongly recommended that any production code utilize these IDs for cryptocurrencies, exchanges, and markets to future-proof your code.**   ### Use the Right Endpoints for the Job  You may have noticed that `/cryptocurrency/listings/latest` and `/cryptocurrency/quotes/latest` return the same crypto data but in different formats. This is because the former is for requesting paginated and ordered lists of *all* cryptocurrencies while the latter is for selectively requesting only the specific cryptocurrencies you require. Many endpoints follow this pattern, allow the design of these endpoints to work for you!  ### Implement a Caching Strategy If Needed  There are standard legal data safeguards built into the <a rel=\"noopener noreferrer\" href=\"https://pro.coinmarketcap.com/user-agreement-commercial\" target=\"_blank\">Commercial User Terms</a> that application developers should keep in mind. These Terms help prevent unauthorized scraping and redistributing of CMC data but are intentionally worded to allow legitimate local caching of market data to support the operation of your application. If your application has a significant user base and you are concerned with staying within the call credit and API throttling limits of your subscription plan consider implementing a data caching strategy.  For example instead of making a `/cryptocurrency/quotes/latest` call every time one of your application's users needs to fetch market rates for specific cryptocurrencies, you could pre-fetch and cache the latest market data for every cryptocurrency in your application's local database every 60 seconds. This would only require 1 API call, `/cryptocurrency/listings/latest?limit=5000`, every 60 seconds. Then, anytime one of your application's users need to load a custom list of cryptocurrencies you could simply pull this latest market data from your local cache without the overhead of additional calls. This kind of optimization is practical for customers with large, demanding user bases.  ###  Code Defensively to Ensure a Robust REST API Integration  Whenever implementing any high availability REST API service for mission critical operations it's recommended to <a rel=\"noopener noreferrer\" href=\"https://en.wikipedia.org/wiki/Defensive_programming\" target=\"_blank\">code defensively</a>. Since the API is versioned, any breaking request or response format change would only be introduced through new versions of each endpoint, *however existing endpoints may still introduce new convenience properties over time*.   We suggest these best practices:  - You should parse the API response JSON as JSON and not through a regular expression or other means to avoid brittle parsing logic. - Your parsing code should explicitly parse only the response properties you require to guarantee new fields that may be returned in the future are ignored.  - You should add robust field validation to your response parsing logic. You can wrap complex field parsing, like dates, in try/catch statements to minimize the impact of unexpected parsing issues (like the unlikely return of a null value).  - Implement a \"Retry with exponential backoff\" coding pattern for your REST API call logic. This means if your HTTP request happens to get rate limited (HTTP 429) or encounters an unexpected server-side condition (HTTP 5xx) your code would automatically recover and try again using an intelligent recovery scheme. You may use one of the many libraries available; for example, <a rel=\"noopener noreferrer\" target=\"_blank\" href=\"https://github.com/tim-kos/node-retry\">this one</a> for Node or <a rel=\"noopener noreferrer\" target=\"_blank\" href=\"https://github.com/litl/backoff\">this one</a> for Python.  ### Reach Out and Upgrade Your Plan  If you're uncertain how to best implement the CoinMarketCap API in your application or your needs outgrow our current self-serve subscription tiers you can reach out to api@coinmarketcap.com. We'll review your needs and budget and may be able to tailor a custom enterprise plan that is right for you.    # Version History      The CoinMarketCap API utilizes <a rel=\"noopener noreferrer\" href=\"https://semver.org/\" target=\"_blank\">Semantic Versioning</a> in the format `major.minor.patch`. The current `major` version is incorporated into the API request path as `/v1/`. Non-breaking `minor` and `patch` updates to the API are released regularly. These may include new endpoints, data points, and API plan features which are always introduced in a non-breaking manner. *This means you can expect new properties to become available in our existing /v1/ endpoints however any breaking change will be introduced under a new major version of the API with legacy versions supported indefinitely unless otherwise stated*.      You can [subscribe to our API Newsletter](/#newsletter-signup) to get monthly email updates on CoinMarketCap API enhancements.  ### v1.0.0 on September 19, 2022  - [/v1/content/posts/top](#operation/getV1ContentPostsTop) now available to get cryptocurrency-related top posts. - [/v1/content/posts/latest](#operation/getV1ContentPostsLatest) now available to get cryptocurrency-related latest posts. - [/v1/content/posts/comments](#operation/getV1ContentPostsComments) now available to get comments of the post.  ### v1.0.0 on Augest 18, 2022  - [/v1/content/latest](#operation/getV1ContentLatest) now available to get news/headlines and Alexandria articles.  ### v1.0.0 on Augest 4, 2022  - [/v1/tools/postman](#operation/getV1ToolsPostman) now API postman collection is available.  ### v2.0.4 on October 11, 2021  - [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now includes `volume_change_24h`. - [/v2/cryptocurrency/quotes/latest](#operation/getV2CryptocurrencyQuotesLatest) now includes `volume_change_24h`.  ### v2.0.3 on October 6, 2021  - [/v1/cryptocurrency/trending/latest](#operation/getV1CryptocurrencyTrendingLatest) now supports `time_period` as an optional parameter.  ### v2.0.2 on September 13, 2021  - [/exchange/map](#operation/getV1ExchangeMap) now available to Free tier users. - [/exchange/info](#operation/getV1ExchangeInfo) now available to Free tier users.  ### v2.0.1 on September 8, 2021  - [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) now includes `volume_24h`, `depth_negative_two`, `depth_positive_two` and `volume_percentage`. - [/exchange/listings/latest](#operation/getV1ExchangeListingsLatest) now includes `open_interest`. ### v2.0.0 on August 17, 2021  - By popular request we have added a number of new useful endpoints ! - [/v1/cryptocurrency/categories](#operation/getV1CryptocurrencyCategories) can be used to access a list of categories and their associated coins. You can also filter the list of categories by one or more cryptocurrencies. - [/v1/cryptocurrency/category](#operation/getV1CryptocurrencyCategory) can be used to load only a single category of coins, listing the coins within that category. - [/v1/cryptocurrency/airdrops](#operation/getV1CryptocurrencyAirdrops) can be used to access a list of CoinMarketCap’s free airdrops. This defaults to a status of `ONGOING` but can be filtered to `UPCOMING` or `ENDED`. You can also query for a list of airdrops by cryptocurrency. - [/v1/cryptocurrency/airdrop](#operation/getV1CryptocurrencyAirdrop) can be used to load a single airdrop and its associated cryptocurrency. - [/v1/cryptocurrency/trending/latest](#operation/getV1CryptocurrencyTrendingLatest) can be used to load the most searched for cryptocurrencies within a period of time. This defaults to a `time_period` of the previous `24h`, but can be changed to `30d`, or `7d` for a larger window of time. - [/v1/cryptocurrency/trending/most-visited](#operation/getV1CryptocurrencyTrendingMostvisited) can be used to load the most visited cryptocurrencies within a period of time. This defaults to a `time_period` of the previous `24h`, but can be changed to `30d`, or `7d` for a larger window of time. - [/v1/cryptocurrency/trending/gainers-losers](#operation/getV1CryptocurrencyTrendingGainerslosers) can be used to load the biggest gainers & losers within a period of time. This defaults to a `time_period` of the previous `24h`, but can be changed to `30d`, or `7d` for a larger window of time.  ### v1.28.0 on August 9, 2021    - [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now includes `market_cap_dominance` and `fully_diluted_market_cap`. - [/v1/cryptocurrency/quotes/latest](#operation/getV2CryptocurrencyQuotesLatest) now includes `market_cap_dominance` and `fully_diluted_market_cap`.  ### v1.27.0 on January 27, 2021    - [/v2/cryptocurrency/info](#operation/getV2CryptocurrencyInfo) response format changed to allow for multiple coins per symbol. - [/v2/cryptocurrency/market-pairs/latest](#operation/getV2CryptocurrencyMarketpairsLatest) response format changed to allow for multiple coins per symbol. - [/v2/cryptocurrency/quotes/historical](#operation/getV2CryptocurrencyQuotesHistorical) response format changed to allow for multiple coins per symbol. - [/v2/cryptocurrency/ohlcv/historical](#operation/getV2CryptocurrencyOhlcvHistorical) response format changed to allow for multiple coins per symbol. - [/v2/tools/price-conversion](#operation/getV2ToolsPriceconversion) response format changed to allow for multiple coins per symbol. - [/v2/cryptocurrency/ohlcv/latest](#operation/getV2CryptocurrencyOhlcvLatest) response format changed to allow for multiple coins per symbol. - [/v2/cryptocurrency/price-performance-stats/latest](#operation/getV2CryptocurrencyPriceperformancestatsLatest) response format changed to allow for multiple coins per symbol.  ### v1.26.0 on January 21, 2021    - [/v2/cryptocurrency/quotes/latest](#operation/getV2CryptocurrencyQuotesLatest) response format changed to allow for multiple coins per symbol.  ### v1.25.0 on April 17, 2020    - [/v1.1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now includes a more robust `tags` response with slug, name, and category. - [/cryptocurrency/quotes/historical](#operation/getV1CryptocurrencyQuotesHistorical) and [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) now include `is_active` and `is_fiat` in the response.  ### v1.24.0 on Feb 24, 2020    - [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) has been modified to include the high and low timestamps. - [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) now includes `category` and `fee_type` market pair filtering options. - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now includes `category` and `fee_type` market pair filtering options.  ### v1.23.0 on Feb 3, 2020    - [/fiat/map](#operation/getV1FiatMap) is now available to fetch the latest mapping of supported fiat currencies to CMC IDs. - [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) now includes `matched_id` and `matched_symbol` market pair filtering options. - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now provides filter parameters `price_min`, `price_max`, `market_cap_min`, `market_cap_max`, `percent_change_24h_min`, `percent_change_24h_max`, `volume_24h_max`, `circulating_supply_min` and `circulating_supply_max` in addition to the existing `volume_24h_min` filter.    ### v1.22.0 on Oct 16, 2019    - [/global-metrics/quotes/latest](#operation/getV1GlobalmetricsQuotesLatest) now additionally returns `total_cryptocurrencies` and `total_exchanges` counts which include inactive projects who's data is still available via API.      ### v1.21.0 on Oct 1, 2019    - [/exchange/map](#operation/getV1ExchangeMap) now includes `sort` options including `volume_24h`. - [/cryptocurrency/map](#operation/getV1CryptocurrencyMap) fix for a scenario where `first_historical_data` and `last_historical_data` may not be populated. - Additional improvements to alphanumeric sorts.    ### v1.20.0 on Sep 25, 2019   - By popular request you may now configure API plan usage notifications and email alerts in the [Developer Portal](https://pro.coinmarketcap.com/account/notifications). - [/cryptocurrency/map](#operation/getV1CryptocurrencyMap) now includes `sort` options including `cmc_rank`.      ### v1.19.0 on Sep 19, 2019    - A new `/blockchain/` category of endpoints is now available with the introduction of our new [/v1/blockchain/statistics/latest](#operation/getV1BlockchainStatisticsLatest) endpoint. This endpoint can be used to poll blockchain statistics data as seen in our <a target=\"_blank\" href=\"https://blockchain.coinmarketcap.com/chain/bitcoin\">Blockchain Explorer</a>. - Additional platform error codes are now surfaced during HTTP Status Code 401, 402, 403, and 429 scenarios as documented in [Errors and Rate Limits](#section/Errors-and-Rate-Limits). - OHLCV endpoints using the `convert` option now match historical UTC open period exchange rates with greater accuracy. - [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) and [/exchange/info](#operation/getV1ExchangeInfo) now include the optional `aux` parameter where listing `status` can be requested in the list of supplemental properties. - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) and [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest): The accuracy of `percent_change_` conversions was improved when passing non-USD fiat `convert` options. - [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) and [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) now support relaxed request validation rules via the `skip_invalid` request parameter. - We also now return a helpful `notice` warning when API key usage is above 95% of daily and monthly API credit usage limits.    ### v1.18.0 on Aug 28, 2019    - [/key/info](#operation/getV1KeyInfo) has been added as a new endpoint. It may be used programmatically monitor your key usage compared to the rate limit and daily/monthly credit limits available to your API plan as an alternative to using the [Developer Portal Dashboard](https://pro.coinmarketcap.com/account).   - [/cryptocurrency/quotes/historical](#operation/getV1CryptocurrencyQuotesHistorical) and [/v1/global-metrics/quotes/historical](#operation/getV1GlobalmetricsQuotesHistorical) have new options to make charting tasks easier and more efficient. Use the new `aux` parameter to cut out response properties you don't need and include the new `search_interval` timestamp to normalize disparate historical records against the same `interval` time periods. - A 4 hour interval option `4h` was added to all historical time series data endpoints.      ### v1.17.0 on Aug 22, 2019    - [/cryptocurrency/price-performance-stats/latest](#operation/getV1CryptocurrencyPriceperformancestatsLatest) has been added as our 21st endpoint! It returns launch price ROI, all-time high / all-time low, and other price stats over several supported time periods. - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) now has the ability to filter all active markets for a cryptocurrency to specific base/quote pairs. Want to return only `BTC/USD` and `BTC/USDT` markets? Just pass `?symbol=BTC&matched_symbol=USD,USDT` or `?id=1&matched_id=2781,825`.   - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) now features `sort` options including `cmc_rank` to reproduce the <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology</a> based sort on pages like <a rel=\"noopener noreferrer\" href=\"https://coinmarketcap.com/currencies/bitcoin/#markets\" target=\"_blank\">Bitcoin Markets</a>.  - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) can now return any exchange level CMC notices affecting a market via the new `notice` `aux` parameter.   - [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) will now continue to return the last updated price data for cryptocurrency that have transitioned to an `inactive` state instead of returning an HTTP 400 error. These active coins that have gone inactive can easily be identified as having a `num_market_pairs` of `0` and a stale `last_updated` date. - [/exchange/info](#operation/getV1ExchangeInfo) now includes a brief text summary for most exchanges as `description`.    ### v1.16.0 on Aug 9, 2019  - We've introduced a new [partners](#tag/partners) category of endpoints for convenient access to 3rd party crypto data. <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/\" target=\"_blank\">FlipSide Crypto</a>'s <a rel=\"noopener noreferrer\" href=\"https://www.flipsidecrypto.com/fcas-explained\" target=\"_blank\">Fundamental Crypto Asset Score</a> (FCAS) is now available as the first partner integration.   - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now provides a `volume_24h_min` filter parameter. It can be used when a threshold of volume is required like in our <a target=\"_blank\" href=\"https://coinmarketcap.com/gainers-losers/\">Biggest Gainers and Losers</a> lists. - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) and [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) can now return rolling `volume_7d` and `volume_30d` via the supplemental `aux` parameter and sort options by these fields.  - `volume_24h_reported`, `volume_7d_reported`, `volume_30d_reported`, and `market_cap_by_total_supply` are also now available through the `aux` parameter with an additional sort option for the latter. - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) can now provide market price relative to the quote currency. Just pass `price_quote` to the supplemental `aux` parameter. This can be used to display consistent price data for a cryptocurrency across several markets no matter if it is the base or quote in each pair as seen in our <a target=\"_blank\" href=\"https://coinmarketcap.com/currencies/bitcoin/#markets\">Bitcoin markets</a> price column. - When requesting a custom `sort` on our list based endpoints, numeric fields like `percent_change_7d` now conveniently return non-applicable `null` values last regardless of sort order.   ### v1.15.0 on Jul 10, 2019  - [/cryptocurrency/map](#operation/getV1CryptocurrencyMap) and [/v1/exchange/map](#operation/getV1ExchangeMap) now expose a 3rd listing state of `untracked` between `active` and `inactive` as outlined in our <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology</a>. See endpoint documentation for additional details.     - [/cryptocurrency/quotes/historical](#operation/getV1CryptocurrencyQuotesHistorical), [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical), and [/exchange/quotes/latest](#operation/getV1ExchangeQuotesLatest) now support fetching multiple cryptocurrencies and exchanges in the same call. - [/global-metrics/quotes/latest](#operation/getV1GlobalmetricsQuotesLatest) now updates more frequently, every minute. It aslo now includes `total_volume_24h_reported`, `altcoin_volume_24h`, `altcoin_volume_24h_reported`, and `altcoin_market_cap`.  - [/global-metrics/quotes/historical](#operation/getV1GlobalmetricsQuotesHistorical) also includes these new dimensions along with historical `active_cryptocurrencies`, `active_exchanges`, and `active_market_pairs` counts.  - We've also added a new `aux` auxiliary parameter to many endpoints which can be used to customize your request. You may request new supplemental data properties that are not returned by default or slim down your response payload by excluding default `aux` fields you don't need in endpoints like [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest). [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) and [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) can now supply `market_url`, `currency_name`, and `currency_slug` for each market using this new parameter. [/exchange/listings/latest](#operation/getV1ExchangeListingsLatest) can now include the exchange `date_launched`.          ### v1.14.1 on Jun 14, 2019 - DATA: Phase 1 methodology updates    Per our <a target=\"_blank\" href=\"https://blog.coinmarketcap.com/2019/05/01/happy-6th-birthday-data-alliance-block-explorers-and-more/\">May 1 announcement</a> of the Data Accountability & Transparency Alliance (<a target=\"_blank\" href=\"https://coinmarketcap.com/data-transparency-alliance/\">DATA</a>), a platform <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology</a> update was published. No API changes are required but users should take note:  - Exchanges that are not compliant with mandatory transparency requirements (Ability to surface live trade and order book data) will be excluded from VWAP price and volume calculations returned from our `/cryptocurrency/` and `/global-metrics/` endpoints going forward. - These exchanges will also return a `volume_24h_adjusted` value of 0 from our `/exchange/` endpoints like the exclusions based on market category and fee type. Stale markets (24h or older) will also be excluded. All exchanges will continue to return `exchange_reported` values as reported.   - We welcome you to <a target=\"_blank\" href=\"https://coinmarketcap.com/data-transparency-alliance/\">learn more about the DATA alliance and become a partner</a>.     ### v1.14.0 on Jun 3, 2019  - [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) now include up to 5 block explorer URLs for each cryptocurrency including our brand new <a target=\"_blank\" href=\"https://blockchain.coinmarketcap.com\">Bitcoin and Ethereum Explorers</a>. - [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) now provides links to most cryptocurrency white papers and technical documentation! Just reference the `technical_doc` array. - [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) now returns a `notice` property that may highlight a significant event or condition that is impacting the cryptocurrency or how it is displayed. See the endpoint property description for more details.  - [/exchange/info](#operation/getV1ExchangeInfo) also includes a `notice` property. This one may highlight a condition that is impacting the availability of an exchange's market data or the use of the exchange. See the endpoint property description for more details.  - [/exchange/info](#operation/getV1ExchangeInfo) now includes the official launch date for each exchange as `date_launched`. - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) and [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) now include market `category` (Spot, Derivatives, or OTC) and `fee_type` (Percentage, No Fees, Transactional Mining, or Unknown) for every market returned. - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) now supports querying by cryptocurrency `slug`. - [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) now includes a `market_cap_strict` sort option to apply a strict numeric sort on this field.    ### v1.13.0 on May 17, 2019  - You may now leverage CoinMarketCap IDs for currency `quote` conversions across all endpoints! Just utilize the new `convert_id` parameter instead of the `convert` parameter. Learn more about creating robust integrations with CMC IDs in our [Best Practices](#section/Best-Practices). - We've updated requesting cryptocurrencies by `slug` to support legacy names from past cryptocurrency rebrands. For example, a request to `/cryptocurrency/quotes/latest?slug=antshares` successfully returns the cryptocurrency by current slug `neo`. - We've extended the brief text summary included as `description` in [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) to now cover all cryptocurrencies! - We've added the fetch-by-slug option to [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical). - Premium subscription users: On your next billing period we'll conveniently switch to displaying monthly/daily credit usage relative to your monthly billing period instead of calendar month and UTC midnight. Click the `?` on our updated <a rel=\"noopener noreferrer\" target=\"_blank\" href=\"https://pro.coinmarketcap.com/account\">API Key Usage</a> panel for more details.    ### v1.12.1 on May 1, 2019    - To celebrate CoinMarketCap's 6th anniversary we've upgraded the crypto API to make more of our data available at each tier!  - Our free Basic tier may now access live price conversions via [/tools/price-conversion](#operation/getV1ToolsPriceconversion). - Our Hobbyist tier now supports a month of historical price conversions with [/tools/price-conversion](#operation/getV1ToolsPriceconversion) using the `time` parameter. We've also made this plan 12% cheaper at $29/mo with a yearly subscription or $35/mo month-to-month. - Our Startup tier can now access a month of cryptocurrency OHLCV data via [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) along with [/tools/price-conversion](#operation/getV1ToolsPriceconversion). - Our Standard tier has been upgraded from 1 month to now 3 months of historical market data access across all historical endpoints.  - Our Enterprise, Professional, and Standard tiers now get access to a new #18th endpoint [/cryptocurrency/listings/historical](#operation/getV1CryptocurrencyListingsHistorical)! Utilize this endpoint to fetch daily historical crypto rankings from the past. We've made historical ranking snapshots available all the way back to 2013!   - All existing accounts and subscribers may take advantage of these updates. If you haven't signed up yet you can check out our updated plans on our <a rel=\"noopener noreferrer\" href=\"https://coinmarketcap.com/api/features/\" target=\"_blank\">feature comparison page</a>.    ### v1.12.0 on Apr 28, 2019    - Our API docs now supply API request examples in 7 languages for every endpoint: cURL, Node.js, Python, PHP, Java, C#, and Go. - Many customer sites format cryptocurrency data page URLs by SEO friendly names like we do here: <a rel=\"noopener noreferrer\" href=\"https://coinmarketcap.com/currencies/binance-coin/\" target=\"_blank\">coinmarketcap.com/currencies/binance-coin</a>. We've made it much easier for these kinds of pages to dynamically reference data from our API. You may now request cryptocurrencies from our [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo) and [/cryptocurrency/quotes/latest](#operation/getV1CryptocurrencyQuotesLatest) endpoints by `slug` as an alternative to `symbol` or `id`. As always, you can retrieve a quick list of every cryptocurrency we support and it's `id`, `symbol`, and `slug` via our [/cryptocurrency/map](#operation/getV1CryptocurrencyMap) endpoint.   - We've increased `convert` limits on historical endpoints once more. You can now request historical market data in up to 3 conversion options at a time like we do internally to display line charts <a rel=\"noopener noreferrer\" href=\"https://coinmarketcap.com/currencies/0x/#charts\" target=\"_blank\">like this</a>. You can now fetch market data converted into your primary cryptocurrency, fiat currency, and a parent platform cryptocurrency (Ethereum in this case) all in one call!     ### v1.11.0 on Mar 25, 2019  - We now supply a brief text summary for each cryptocurrency in the `description` field of [/cryptocurrency/info](#operation/getV1CryptocurrencyInfo). The majority of top cryptocurrencies include this field with more coming in the future.   - We've made `convert` limits on some endpoints and plans more flexible. Historical endpoints are now allowed 2 price conversion options instead of 1. Professional plan convert limit has doubled from 40 to 80. Enterprise has tripled from 40 to 120.   - CoinMarketCap Market ID: We now return `market_id` in /market-pairs/latest endpoints. Like our cryptocurrency and exchange IDs, this ID can reliably be used to uniquely identify each market *permanently* as this ID never changes.   - Market symbol overrides: We now supply an `exchange_symbol` in addition to `currency_symbol` for each market pair returned in our /market-pairs/latest endpoints. This allows you to reference the currency symbol provided by the exchange in case it differs from the CoinMarketCap identified symbol that the majority of markets use.      ### v1.10.1 on Jan 30, 2019  - Our API health status dashboard is now public at http://status.coinmarketcap.com.   - We now conveniently return `market_cap` in our [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) endpoint so you don't have to make a separately query when fetching historic OHLCV data.  - We've improved the accuracy of percent_change_1h / 24h / 7d calculations when using the `convert` option with our latest cryptocurrency endpoints.  - [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) now updates more frequently, every 1 minute.   - Contract Address and parent platform metadata changes are reflected on the API much more quickly.      ### v1.9.0 on Jan 8, 2019    - Did you know there are currently 684 active USD market pairs tracked by CoinMarketCap? You can now pass any [fiat CoinMarketCap ID](#section/Standards-and-Conventions) to the [/cryptocurrency/market-pairs/latest](#operation/getV1CryptocurrencyMarketpairsLatest) `id` parameter to list all active markets across all exchanges for a given fiat currency.    - We've added a new dedicated migration FAQ page for users migrating from our old Public API to the new API [here](https://pro.coinmarketcap.com/migrate). It includes a helpful tutorial link for Excel and Google Sheets users who need help migrating. - Cryptocurrency and exchange symbol and name rebrands are now reflected in the API much more quickly.    ### v1.8.0 on Dec 27, 2018   - We now supply the contract address for all cryptocurrencies on token platforms like Ethereum! Look for  `token_address` in the `platform` property of our cryptocurrency endpoints like [/cryptocurrency/map](#operation/getV1CryptocurrencyMap) and [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest). - All 96 non-USD fiat conversion rates now update every 1 minute like our USD rates! This includes using the `convert` option for all /latest market data endpoints as well as our [/tools/price-conversion](#operation/getV1ToolsPriceconversion) endpoint.       ### v1.7.0 on Dec 18, 2018  - We've upgraded our fiat (government) currency conversion support from our original 32 to now cover 93 fiat currencies!   - We've also introduced currency conversions for four precious metals: Gold, Silver, Platinum, and Palladium!  - You may pass all 97 fiat currency options to our [/tools/price-conversion](#operation/getV1ToolsPriceconversion) endpoint using either the `symbol` or `id` parameter. Using CMC `id` is always the most robust option. CMC IDs are now included in the full list of fiat options located [here](#section/Standards-and-Conventions).  - All historical endpoints including our price conversion endpoint with \"time\" parameter now support historical fiat conversions back to 2013!     ### v1.6.0 on Dec 4, 2018  - We've rolled out another top requested feature, giving you access to platform metadata for cryptocurrencies that are tokens built on other cryptocurrencies like Ethereum. Look for the new `platform` property on our cryptocurrency endpoints like [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) and [/cryptocurrency/map](#operation/getV1CryptocurrencyMap).  - We've also added a **CMC equivalent pages** section to our endpoint docs so you can easily determine which endpoints to use to reproduce functionality on the main coinmarketcap.com website.   - Welcome Public API users! With the migration of our legacy Public API into the Professional API we now have 1 unified API at CMC. This API is now known as the CoinMarketCap API and can always be accessed at [coinmarketcap.com/api](https://coinmarketcap.com/api).       ### v1.5.0 on Nov 28, 2018  - [/cryptocurrency/ohlcv/historical](#operation/getV1CryptocurrencyOhlcvHistorical) now supports hourly OHLCV! Use time_period=\"hourly\" and don't forget to set the \"interval\" parameter to \"hourly\" or one of the new hourly interval options. - [/tools/price-conversion](#operation/getV1ToolsPriceconversion) now supports historical USD conversions.   - We've increased the minute based rate limits for several plans. Standard plan has been upgraded from 30 to 60 calls per minute. Professional from 60 to 90. Enterprise from 90 to 120.   - We now include some customer and data partner logos and testimonials on the CoinMarketCap API site. Visit pro.coinmarketcap.com to check out what our enterprise customers are saying and contact us at api@coinmarketcap.com if you'd like to get added to the list!       ### v1.4.0 on Nov 20, 2018  - [/tools/price-conversion](#operation/getV1ToolsPriceconversion) can now provide the latest crypto-to-crypto conversions at 1 minute accuracy with extended decimal precision upwards of 8 decimal places. - [/tools/price-conversion](#operation/getV1ToolsPriceconversion) now supports historical crypto-to-crypto conversions leveraging our closest averages to the specified \"time\" parameter.  - All of our historical data endpoints now support historical cryptocurrency conversions using the \"convert\" parameter. The closest reference price for each \"convert\" option against each historical datapoint is used for each conversion. - [/global-metrics/quotes/historical](#operation/getV1GlobalmetricsQuotesHistorical) now supports the \"convert\" parameter.       ### v1.3.0 on Nov 9, 2018  - The latest UTC day's OHLCV record is now available sooner. 5-10 minutes after each UTC midnight. - We're now returning a new `vol_24h_adjusted` property on  [/exchange/quotes/latest](#operation/getV1ExchangeQuotesLatest) and [/exchange/listings/latest](#operation/getV1ExchangeListingsLatest) and a sort option for the latter so you may now list exchange rankings by CMC adjusted volume as well as exchange reported.  - We are now returning a `tags` property with [/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) with our first tag `mineable` so you know which currencies are mineable. Additional tags will be introduced in the future. - We've increased the \"convert\" parameter limit from 32 to 40 for plans that support max conversion limits.     ### v1.2.0 on Oct 30, 2018  - Our exchange [listing](#operation/getV1ExchangeListingsLatest) and [quotes](#operation/getV1ExchangeQuotesLatest) endpoints now update much more frequently! Every 1 minute instead of every 5 minutes. - These latest exchange data endpoints also now return `volume_7d / 30d` and `percent_change_volume_24h / 7d / 30d` along with existing data. - We've updated our documentation for [/exchange/market-pairs/latest](#operation/getV1ExchangeMarketpairsLatest) to reflect that it receives updates every 1 minute, not 5, since June.    ### v1.1.4 on Oct 19, 2018  - We've improved our tiered support inboxes by plan type to answer support requests even faster. - You may now opt-in to our API mailing list on signup. If you haven't signed up you can [here](/#newsletter-signup).    ### v1.1.3 on Oct 12, 2018  - We've increased the rate limit of our free Basic plan from 10 calls a minute to 30. - We've increased the rate limit of our Hobbyist plan from 15 to 30.  ### v1.1.2 on Oct 5, 2018    - We've updated our most popular /cryptocurrency/listings/latest endpoint to cost 1 credit per 200 data points instead of 100 to give customers more flexibility.  - By popular request we've introduced a new $33 personal use Hobbyist tier with access to our currency conversion calculator endpoint.  - Our existing commercial use Hobbyist tier has been renamed to Startup. Our free Starter tier has been renamed to Basic.    ### v1.1.1 on Sept 28, 2018    - We've increased our monthly credit limits for our smaller plans! Existing customers plans have also been updated.  - Our free Starter plan has been upgraded from 6 to 10k monthly credits (66% increase). - Our Hobbyist plan has been upgraded from 60k to 120k monthly credits (100% increase).  - Our Standard plan has been upgraded from 300 to 500k monthly credits (66% increase).   ### v1.1.0 on Sept 14, 2018    - We've introduced our first new endpoint since rollout, active day OHLCV for Standard plan and above with [/v1/cryptocurrency/ohlcv/latest](#operation/getV1CryptocurrencyOhlcvLatest)    ### v1.0.4 on Sept 7, 2018  - Subscription customers with billing renewal issues now receive an alert from our API during usage and an unpublished grace period before access is restricted. - API Documentation has been improved including an outline of credit usage cost outlined on each endpoint documentation page.    ### v1.0.3 on Aug 24, 2018 - /v1/tools/price-conversion floating point conversion accuracy was improved. - Added ability to query for non-alphanumeric crypto symbols like $PAC - Customers may now update their billing card on file with an active Stripe subscription at pro.coinmarketcap.com/account/plan 
 *
 * API version: 2.0.4
 * Contact: api@coinmarketcap.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

// APIClient manages communication with the CoinMarketCap Cryptocurrency API Documentation API v2.0.4
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	BlockchainApi *BlockchainApiService

	CryptocurrencyApi *CryptocurrencyApiService

	DefaultApi *DefaultApiService

	DeprecatedApi *DeprecatedApiService

	ExchangeApi *ExchangeApiService

	FiatApi *FiatApiService

	GlobalMetricsApi *GlobalMetricsApiService

	KeyApi *KeyApiService

	ToolsApi *ToolsApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.BlockchainApi = (*BlockchainApiService)(&c.common)
	c.CryptocurrencyApi = (*CryptocurrencyApiService)(&c.common)
	c.DefaultApi = (*DefaultApiService)(&c.common)
	c.DeprecatedApi = (*DeprecatedApiService)(&c.common)
	c.ExchangeApi = (*ExchangeApiService)(&c.common)
	c.FiatApi = (*FiatApiService)(&c.common)
	c.GlobalMetricsApi = (*GlobalMetricsApiService)(&c.common)
	c.KeyApi = (*KeyApiService)(&c.common)
	c.ToolsApi = (*ToolsApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
