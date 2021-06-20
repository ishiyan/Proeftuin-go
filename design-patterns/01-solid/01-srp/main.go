package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type journal struct {
	entries []string
}

func (j *journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *journal) RemoveEntry(index int) {
	// ...
}

// breaks srp

func (j *journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *journal) Load(filename string) {

}

func (j *journal) LoadFromWeb(url *url.URL) {

}

var lineSeparator = "\n"

func saveToFile(j *journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

type persistence struct {
	lineSeparator string
}

func (p *persistence) saveToFile(j *journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))

	// separate function
	saveToFile(&j, "journal.txt")

	//
	p := persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
