package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	fmt.Println("Program to count statistics of Characters in Indonesian Words")
	// Initialize map
	//make map of each alphabet character
	chars := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	// Open file
	file, err := os.Open("dictionary.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	// Loop each line
	for _, eachline := range txtlines {
		for _, eachchar := range eachline {
			chars[string(eachchar)]++
		}
	}
	p := make(PairList, len(chars))

	i := 0
	for k, v := range chars {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))
	fmt.Println("List of chars mentioned in Indonesian Words\n ")
	totalChars := 0
	for _, v := range p {
		fmt.Printf("%s: %d\n", v.Key, v.Value)
		totalChars += v.Value
	}
	fmt.Println("Top 5 most frequent characters\n ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: %d times mentioned or %.2f %%\n", p[i].Key, p[i].Value, float64(p[i].Value)/float64(totalChars)*100)
	}

}
