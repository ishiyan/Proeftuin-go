package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"nq/nasdaq"
	"nq/nyse"
)

type symbols struct {
	Updated string                `json:"updated"`
	Symbols []nasdaq.NasdaqSymbol `json:"symbols"`
}

type IsinSedolSymbol struct {
	Mnemonic string `json:"mnemonic"`
	ISIN     string `json:"ISIN"`
	SEDOL    string `json:"SEDOL"`
}

func main() {
	t := time.Now().Format("2006-01-02_15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	categoryPtr := flag.String("category", "stock", "category: [stock, stock-nasdaq, stock-nyse, stock-amex, etf]")
	flag.Parse()

	symbolsFileName := "nasdaq-stock.json"
	isinSedolFileName := "nasdaq-stock.isin-sedol.json"
	switch *categoryPtr {
	case "stock-nasdaq":
		symbolsFileName = "nasdaq-stock-nasdaq.json"
		isinSedolFileName = "nasdaq-stock-nasdaq.isin-sedol.json"
	case "stock-nyse":
		symbolsFileName = "nasdaq-stock-nyse.json"
		isinSedolFileName = "nasdaq-stock-nyse.isin-sedol.json"
	case "stock-amex":
		symbolsFileName = "nasdaq-stock-amex.json"
		isinSedolFileName = "nasdaq-stock-amex.isin-sedol.json"
	case "etf":
		symbolsFileName = "nasdaq-etf.json"
		isinSedolFileName = "nasdaq-etf.isin-sedol.json"
	}

	fmt.Println("reading " + symbolsFileName)

	syms, err := readSymbols(symbolsFileName)
	if err != nil {
		panic(fmt.Sprintf("cannot read '%s' symbols: %s", symbolsFileName, err))
	}
	fmt.Printf("%d symbols read\n", len(syms.Symbols))

	nyseSyms := make([]IsinSedolSymbol, 0)
	for _, s := range syms.Symbols {
		v, err := nyse.GetSymbol(s.Mnemonic)
		if err != nil {
			fmt.Println(err)
			nyseSyms = append(nyseSyms, IsinSedolSymbol{
				Mnemonic: s.Mnemonic,
				ISIN:     "",
				SEDOL:    "",
			})

			continue
		}

		nyseSyms = append(nyseSyms, IsinSedolSymbol{
			Mnemonic: s.Mnemonic,
			ISIN:     v.ISIN,
			SEDOL:    v.SEDOL,
		})
	}

	suffix := "." + t
	err = os.Rename(isinSedolFileName, isinSedolFileName+suffix)
	if err != nil {
		panic(fmt.Sprintf("Cannot rename original JSON file: %s", err))
	}

	f, err := os.OpenFile(isinSedolFileName, os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("cannot create new JSON file: %s", err))
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	err = enc.Encode(nyseSyms)
	if err != nil {
		panic(fmt.Sprintf("cannot encode enriched symbols: %s", err))
	}

	fmt.Println("finished: " + time.Now().Format("2006-01-02_15-04-05"))
}

func readSymbols(fileName string) (*symbols, error) {
	var s symbols

	f, err := os.Open(fileName)
	if err != nil {
		return &s, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&s)
	if err != nil {
		return &s, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	return &s, nil
}
