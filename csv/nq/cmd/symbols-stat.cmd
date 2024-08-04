@echo off

nqsymstat.exe -symbols=nasdaq-stock-nasdaq.json >>nasdaq-stock-nasdaq.stat
nqsymstat.exe -symbols=nasdaq-stock-amex.json >>nasdaq-stock-amex.stat
nqsymstat.exe -symbols=nasdaq-stock-nyse.json >>nasdaq-stock-nyse.stat
nqsymstat.exe -symbols=nasdaq-etf.json >>nasdaq-etf.stat
