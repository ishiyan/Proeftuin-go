package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		caps, err := readData("capitals.txt")
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	} //        ^^^^^^^^^^^^^^^^^^^^ breaks Dependency Inversion Principle
	return result
}

func main() {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	ok := tp == (17500000 + 17400000)
	fmt.Println(ok)
}
