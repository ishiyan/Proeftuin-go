@echo off
cd bz2csv
go build bz2csv.go
cd ..
cd gz2csv
go build gz2csv.go
cd ..
cd xz2csv
go build xz2csv.go
cd ..
cd csv2bz
go build csv2bz.go
cd ..
cd csv2gz
go build csv2gz.go
cd ..
cd csv2xz
go build csv2xz.go
cd ..
