package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer fin.Close()

	fout, err := os.Create("out/" + os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	w := bufio.NewWriter(fout)
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",\"|\",", "|", -1)
		line = strings.Replace(line, "\"", "", -1)
		_, _ = w.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
