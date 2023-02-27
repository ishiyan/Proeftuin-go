https://api.nasdaq.com/api/news/topic/articlebysymbol?q=adbe|stocks&offset=0&limit=7&fallback=true
https://api.nasdaq.com/api/company/ADBE/company-profile
https://api.nasdaq.com/api/quote/ADBE/summary?assetclass=stocks

https://www.nasdaq.com/market-activity/stocks/symbol-change-history
https://www.nasdaq.com/market-activity/quotes/dividend-history
https://www.nasdaq.com/market-activity/stocks/screener
https://www.nasdaq.com/market-activity/quotes/nasdaq-ndx-index

https://api.nasdaq.com/api/quote/TSLA/realtime-trades?&limit=99999999&fromTime=09:30

09:30
10:00
10:30
11:00
11:30
12:00
12:30
13:00
13:30
14:00
14:30
15:00
15:30

https://api.nasdaq.com/api/quote/TSLA/extended-trading?markettype=post&assetclass=stocks&limit=99999999&time=1

4:00 - 4:29 = 1
4:30 - 4:59 = 2
5:00 - 5:29 = 3
5:30 - 5:59 = 4
6:00 - 6:29 = 5
6:30 - 6:59 = 6
7:00 - 7:29 = 7
7:30 - 7:59 = 8

https://api.nasdaq.com/api/quote/TSLA/extended-trading?markettype=pre&assetclass=stocks&limit=99999999&time=1


4:00 - 4:29 = 1
4:30 - 4:59 = 2
5:00 - 5:29 = 3
5:30 - 5:59 = 4
6:00 - 6:29 = 5
6:30 - 6:59 = 6
7:00 - 7:29 = 7
7:30 - 7:59 = 8
8:00 - 8:29 = 9
8:30 - 8:59 = 10
9:00 - 9:29 = 11




04:00 EST = 10.00 CET
09:30 EST = 15:30 CET
15:30 EST = 21:30 CET
16:00 EST = 22:00 CET
20:00 EST = 02:00 CET

https://api.nasdaq.com/api/market-info

https://api.nasdaq.com/api/quote/TSLA/extended-trading?markettype=pre&assetclass=stocks&time=


Nasdaq-GS is XNGS -- segment of XNAS
https://api.nasdaq.com/api/screener/stocks?tableonly=true&limit=999999&offset=0&download=false
https://api.nasdaq.com/api/quote/AAPL/chart?assetclass=stocks

https://api.nasdaq.com/api/quote/ADBE/info?assetclass=stocks


https://api.nasdaq.com/api/screener/stocks?tableonly=true&limit=999999&offset=0&download=false
{
    "data":{
        "filters":null
        "table":{
            "headers":{
                "symbol":"Symbol",
                "name":"Name",
                "lastsale":"Last Sale",
                "netchange":"Net Change",
                "pctchange":"% Change",
                "marketCap":"Market Cap"},
            "rows":[
                {"symbol":"AAPL",
                "name":"Apple Inc. Common Stock",
                "lastsale":"$148.965",
                "netchange":"0.055",
                "pctchange":"0.037%",
                "marketCap":"2,582,656,853,100",
                "url":"/market-activity/stocks/aapl"},

                {"symbol":"ZTAQW",
                "name":"Zimmer Energy Transition Acquisition Corp. Warrants",
                "lastsale":"$0.2916",
                "netchange":"0.1027",
                "pctchange":"54.367%",
                "marketCap":"NA",
                "url":"/market-activity/stocks/ztaqw"}
            ]
        },
        "totalrecords":7802,
        "asof":"Last price as of Feb 23, 2023"
    },
    "message":null,
    "status":{
        "rCode":200,
        "bCodeMessage":null,
        "developerMessage":null}
}

https://api.nasdaq.com/api/quote/ADBE/info?assetclass=stocks
{
    "data":{
        "symbol":"ADBE",
        "companyName":"Adobe Inc. Common Stock",
        "stockType":"Common Stock",
        "exchange":"NASDAQ-GS",
        "isNasdaqListed":true,
        "isNasdaq100":true,
        "isHeld":false,
        "primaryData":{
            "lastSalePrice":"$336.06",
            "netChange":"-10.96",
            "percentageChange":"-3.16%",
            "deltaIndicator":"down",
            "lastTradeTimestamp":"Feb 24, 2023 6:09 AM ET - PRE-MARKET",
            "isRealTime":true,
            "bidPrice":"$335.80",
            "askPrice":"$337.00",
            "bidSize":"1",
            "askSize":"2",
            "volume":"2,082"},
        "secondaryData":{
            "lastSalePrice":"$347.02",
            "netChange":"-1.70",
            "percentageChange":"-0.49%",
            "deltaIndicator":"down",
            "lastTradeTimestamp":"CLOSED AT 4:00 PM ET ON Feb 23, 2023",
            "isRealTime":false,
            "bidPrice":"",
            "askPrice":"",
            "bidSize":"",
            "askSize":"",
            "volume":""},
        "marketStatus":"Pre-Market",
        "assetClass":"STOCKS",
        "keyStats":null,
        "notifications":[]
    },
    "message":null,
    "status":{
        "rCode":200,
        "bCodeMessage":null,
        "developerMessage":null}
}

https://api.nasdaq.com/api/screener/stocks?tableonly=false&limit=999999&offset=0&download=false
{
    "data":{
        "filters":{
            "region":[
                {"name":"Africa","value":"africa"},
                {"name":"Asia","value":"asia"},
                {"name":"Australia and South Pacific","value":"australia_and_south_pacific"},
                {"name":"Caribbean","value":"caribbean"},
                {"name":"Europe","value":"europe"},
                {"name":"Middle East","value":"middle_east"},
                {"name":"North America","value":"north_america"},
                {"name":"South America","value":"south_america"}],
            "country":[
                {"name":"Argentina","value":"argentina"},
                {"name":"Armenia","value":"armenia"},
                {"name":"Australia","value":"australia"},
                {"name":"Austria","value":"austria"},
                {"name":"Belgium","value":"belgium"},
                {"name":"Bermuda","value":"bermuda"},
                {"name":"Brazil","value":"brazil"},
                {"name":"Canada","value":"canada"},
                {"name":"Cayman Islands","value":"cayman_islands"},
                {"name":"Chile","value":"chile"},
                {"name":"Colombia","value":"colombia"},
                {"name":"Costa Rica","value":"costa_rica"},
                {"name":"Curacao","value":"curacao"},
                {"name":"Cyprus","value":"cyprus"},
                {"name":"Denmark","value":"denmark"},
                {"name":"Finland","value":"finland"},
                {"name":"France","value":"france"},
                {"name":"Germany","value":"germany"},
                {"name":"Greece","value":"greece"},
                {"name":"Guernsey","value":"guernsey"},
                {"name":"Hong Kong","value":"hong_kong"},
                {"name":"India","value":"india"},
                {"name":"Indonesia","value":"indonesia"},
                {"name":"Ireland","value":"ireland"},
                {"name":"Isle of Man","value":"isle_of_man"},
                {"name":"Israel","value":"israel"},
                {"name":"Italy","value":"italy"},
                {"name":"Japan","value":"japan"},
                {"name":"Jersey","value":"jersey"},
                {"name":"Luxembourg","value":"luxembourg"},
                {"name":"Macau","value":"macau"},
                {"name":"Mexico","value":"mexico"},
                {"name":"Monaco","value":"monaco"},
                {"name":"Netherlands","value":"netherlands"},
                {"name":"Norway","value":"norway"},
                {"name":"Panama","value":"panama"},
                {"name":"Peru","value":"peru"},
                {"name":"Philippines","value":"philippines"},
                {"name":"Puerto Rico","value":"puerto_rico"},
                {"name":"Russia","value":"russia"},
                {"name":"Singapore","value":"singapore"},
                {"name":"South Africa","value":"south_africa"},
                {"name":"South Korea","value":"south_korea"},
                {"name":"Spain","value":"spain"},
                {"name":"Sweden","value":"sweden"},
                {"name":"switzerland","value":"switzerland"},
                {"name":"Taiwan","value":"taiwan"},
                {"name":"Turkey","value":"turkey"},
                {"name":"United Kingdom","value":"united_kingdom"},
                {"name":"United States","value":"united_states"},
                {"name":"USA","value":"usa"}],
            "exchange":[
                {"name":"NASDAQ","value":"NASDAQ"},
                {"name":"NYSE","value":"NYSE"},
                {"name":"AMEX","value":"AMEX"}],
            "sector":[
                {"name":"Technology","value":"technology"},
                {"name":"Telecommunications","value":"telecommunications"},
                {"name":"Healthcare","value":"health_care"},
                {"name":"Financials","value":"finance"},
                {"name":"Real Estate","value":"real_estate"},
                {"name":"Consumer Discretionary","value":"consumer_discretionary"},
                {"name":"Consumer Staples","value":"consumer_staples"},
                {"name":"Industrials","value":"industrials"},
                {"name":"Basic Materials","value":"basic_materials"},
                {"name":"Energy","value":"energy"},
                {"name":"Utilities","value":"utilities"}],
            "recommendation":[
                {"name":"Strong Buy","value":"strong_buy"},
                {"name":"Hold","value":"hold"},
                {"name":"Buy","value":"buy"},
                {"name":"Sell","value":"sell"},
                {"name":"Strong Sell","value":"strong_sell"}],
            "marketcap":[
                {"name":"Mega (>$200B)","value":"mega"},
                {"name":"Large ($10B-$200B)","value":"large"},
                {"name":"Medium ($2B-$10B)","value":"mid"},
                {"name":"Small ($300M-$2B)","value":"small"},
                {"name":"Micro ($50M-$300M)","value":"micro"},
                {"name":"Nano (<$50M)","value":"nano"}],
            "exsubcategory":[
                {"name":"Global Select","value":"NGS"},
                {"name":"Global Market","value":"NGM"},
                {"name":"Capital Market","value":"NCM"},
                {"name":"ADR","value":"ADR"}]
        },
        "table":{
            "headers":{
                "symbol":"Symbol",
                "name":"Name",
                "lastsale":"Last Sale",
                "netchange":"Net Change",
                "pctchange":"% Change",
                "marketCap":"Market Cap"},
            "rows":[
                {"symbol":"AAPL",
                "name":"Apple Inc. Common Stock",
                "lastsale":"$149.40",
                "netchange":"0.49",
                "pctchange":"0.329%",
                "marketCap":"2,590,198,596,000",
                "url":"/market-activity/stocks/aapl"},

https://api.nasdaq.com/api/company/ADBE/company-profile
{
    "data":{
        "ModuleTitle":{"label":"Module Title","value":"Company Description"},
        "CompanyName":{"label":"Company Name","value":"Adobe Inc."},
        "Symbol":{"label":"Symbol","value":"ADBE"},
        "Address":{"label":"Address","value":"345 PARK AVENUE, SAN JOSE, California, 95110-2704, United States"},
        "Phone":{"label":"Phone","value":"+1 408 536-6000"},
        "Industry":{"label":"Industry","value":"Computer Software: Prepackaged Software"},
        "Sector":{"label":"Sector","value":"Technology"},
        "Region":{"label":"Region","value":"North America"},
        "CompanyDescription":{"label":"Company Description","value":"Adobe provides content creation, document management, and digital marketing and advertising software and services to creative professionals and marketers for creating, managing, delivering, measuring, optimizing and engaging with compelling content multiple operating systems, devices and media. The company operates with three segments: digital media content creation, digital experience for marketing solutions, and publishing for legacy products (less than 5% of revenue)."},
        "CompanyUrl":{"label":"Company Url","value":"https://www.adobe.com"},
        "KeyExecutives":{"label":"Key Executives","value":[{"name":"Cynthia A. Stoddard","title":"Chief Information Officer & Senior Vice President"},{"name":"Daniel J. Durn","title":"CFO, EVP-Finance, Technology Services & Operations"},{"name":"Shantanu Narayen","title":"Chairman, President & Chief Executive Officer"}]}},
    "message":null,
    "status":{"rCode":200,"bCodeMessage":null,"developerMessage":null}
}

https://api.nasdaq.com/api/quote/ADBE/summary?assetclass=stocks
{
    "data":{
        "symbol":"ADBE",
        "summaryData":{
            "Exchange":{"label":"Exchange","value":"NASDAQ-GS"},
            "Sector":{"label":"Sector","value":"Technology"},
            "Industry":{"label":"Industry","value":"Computer Software: Prepackaged Software"},
            "OneYrTarget":{"label":"1 Year Target","value":"$387.50"},
            "TodayHighLow":{"label":"Today's High/Low","value":"N/A"},
            "ShareVolume":{"label":"Share Volume","value":"2,918"},
            "AverageVolume":{"label":"Average Volume","value":"2,772,771"},
            "PreviousClose":{"label":"Previous Close","value":"$347.02"},
            "FiftTwoWeekHighLow":{"label":"52 Week High/Low","value":"$479.21/$274.73"},
            "MarketCap":{"label":"Market Cap","value":"153,848,268,000"},
            "PERatio":{"label":"P/E Ratio","value":34.36},
            "ForwardPE1Yr":{"label":"Forward P/E 1 Yr.","value":"28.56"},
            "EarningsPerShare":{"label":"Earnings Per Share(EPS)","value":"$10.10"},
            "AnnualizedDividend":{"label":"Annualized Dividend","value":"N/A"},
            "ExDividendDate":{"label":"Ex Dividend Date","value":"N/A"},
            "DividendPaymentDate":{"label":"Dividend Pay Date","value":"N/A"},
            "Yield":{"label":"Current Yield","value":"N/A"},
            "Beta":{"label":"Beta","value":1.0}},
            "assetClass":"STOCKS",
            "additionalData":null,
            "bidAsk":{
                "Bid * Size":{"label":"Bid * Size","value":"$335.80 * 1"},
                "Ask * Size":{"label":"Ask * Size","value":"$337.00 * 2"}}
        },
        "message":null,
        "status":{"rCode":200,"bCodeMessage":null,"developerMessage":null}
}

https://api.nasdaq.com/api/company/GNLX/company-profile
https://api.nasdaq.com/api/quote/GNLX/summary?assetclass=stocks
https://api.nasdaq.com/api/quote/GNLX/info?assetclass=stocks

American Depositary Share (ADS)
An American depositary share (ADS) is an equity share of a non-US company that is held by a US depositary bank and is available for purchase by US investors.
