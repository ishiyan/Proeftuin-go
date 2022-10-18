package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filePath string, numFields int) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.TrimLeadingSpace = true
	r.Comment = '#'
	r.FieldsPerRecord = numFields

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

type managementFee struct {
	accountNumber string
	year          string
	gros          string
	net           string
	btw           string
	used          bool
}

type promotion struct {
	accountNumber string
	gros          float64
	net           float64
	used          bool
}

type correction struct {
	accountNumber string
	gros          float64
	net           float64
	used          bool
}

type mapping struct {
	accountNumber string
	accountId     string
	used          bool
}

func readManagementFee() map[string]*managementFee {
	csv := readCsvFile("management-fee-2021.csv", 5)
	m := map[string]*managementFee{}

	for _, e := range csv {
		v := managementFee{
			accountNumber: e[0],
			year:          e[1],
			gros:          e[2],
			net:           e[3],
			btw:           e[4],
		}

		if _, ok := m[v.accountNumber]; ok {
			log.Fatal("Duplicate account number in management fees: " + v.accountNumber)
		}

		m[v.accountNumber] = &v
	}

	return m
}

func readPromotions() map[string]*promotion {
	csv := readCsvFile("promotions-2021.csv", 3)
	m := map[string]*promotion{}

	for _, e := range csv {
		gros, _ := strconv.ParseFloat(e[1], 64)
		net, _ := strconv.ParseFloat(e[2], 64)

		if w, ok := m[e[0]]; ok {
			w.gros += gros
			w.net += net
		} else {
			v := promotion{
				accountNumber: e[0],
				gros:          gros,
				net:           net,
			}

			m[v.accountNumber] = &v
		}
	}

	return m
}

func readCorrections() map[string]*correction {
	csv := readCsvFile("corrections-2021.csv", 3)
	m := map[string]*correction{}

	for _, e := range csv {
		gros, _ := strconv.ParseFloat(e[1], 64)
		net, _ := strconv.ParseFloat(e[2], 64)

		if w, ok := m[e[0]]; ok {
			w.gros += gros
			w.net += net
		} else {
			v := correction{
				accountNumber: e[0],
				gros:          gros,
				net:           net,
			}

			m[v.accountNumber] = &v
		}
	}

	return m
}

func readMapping() map[string]*mapping {
	csv := readCsvFile("mapping-2021.csv", 2)
	m := map[string]*mapping{}

	for _, e := range csv {
		if _, ok := m[e[0]]; !ok {
			v := mapping{
				accountNumber: e[0],
				accountId:     e[1],
			}

			m[v.accountNumber] = &v
		}
	}

	return m
}

func main() {
	mp := readMapping()
	cor := readCorrections()
	pro := readPromotions()
	mf := readManagementFee()

	// fmt.Println(mp)
	// fmt.Println(cor)
	// fmt.Println(pro)
	// fmt.Println(mf)

	fmt.Println("AccountNumber;AccountId;Year;BrutoFee;NetFee;VAT;Correction;Promotion")

	for _, emf := range mf {
		var id = "xxx"
		if e1, ok := mp[emf.accountNumber]; ok {
			id = e1.accountId
			e1.used = true
		} else {
			log.Println("No account id for account number " + emf.accountNumber)
		}

		var p float64
		var c float64

		if e2, ok := pro[emf.accountNumber]; ok {
			p = e2.net
			e2.used = true
		}

		if e3, ok := cor[emf.accountNumber]; ok {
			c = e3.net
			e3.used = true
		}

		fmt.Printf("%v;%v;%v;%v;%v;%v;%v;%v\n", emf.accountNumber, id, emf.year, emf.gros, emf.net, emf.btw, c, p)
	}

	for _, e := range pro {
		if !e.used {
			log.Println("Promotion not used: account number " + e.accountNumber)
		}
	}

	for _, e := range cor {
		if !e.used {
			log.Println("Correction not used: account number " + e.accountNumber)
		}
	}

	for _, e := range mp {
		if !e.used {
			log.Println("Mapping not used: account number " + e.accountNumber)
		}
	}
}
