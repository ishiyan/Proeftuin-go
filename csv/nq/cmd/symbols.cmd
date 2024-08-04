@echo off

nqsym.exe -category=stock-nasdaq >>nasdaq-stock-nasdaq.symbols.log
nqsym.exe -category=stock-amex >>nasdaq-stock-amex.symbols.log
nqsym.exe -category=stock-nyse >>nasdaq-stock-nyse.symbols.log
nqsym.exe -category=etf >>nasdaq-etf.symbols.log
