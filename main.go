package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
	chars := map[string]int{}
	// Open file
	file, err := os.Open("indonesian.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		for _, eachchar := range scanner.Text() {
			chars[string(eachchar)]++
		}
	}

	file.Close()
	// Loop each line
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
	fmt.Println()
	fmt.Println("Top 5 most frequent characters\n ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: %d times mentioned or %.2f %%\n", p[i].Key, p[i].Value, float64(p[i].Value)/float64(totalChars)*100)
	}
	fmt.Println()

	chars = map[string]int{}
	// Open file
	file, err = os.Open("javanese.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		for _, eachchar := range scanner.Text() {
			if string(eachchar) != " " {
				chars[strings.ToLower(string(eachchar))]++
			}

		}
	}

	file.Close()
	// Loop each line
	p = make(PairList, len(chars))

	i = 0
	for k, v := range chars {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))
	fmt.Println("List of chars mentioned in Javanese Words\n ")
	totalChars = 0
	for _, v := range p {
		fmt.Printf("%s: %d\n", v.Key, v.Value)
		totalChars += v.Value
	}
	fmt.Println()
	fmt.Println("Top 5 most frequent characters\n ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: %d times mentioned or %.2f %%\n", p[i].Key, p[i].Value, float64(p[i].Value)/float64(totalChars)*100)
	}
}
