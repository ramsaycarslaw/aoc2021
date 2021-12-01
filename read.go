package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	// while there is strill input
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Return the contents of a file as a string
func AsString(path string) []string {
	return readFile(path)
}

// return the contents of a file as an array of int64
func AsInt64(path string) []int64 {
	lines := readFile(path)

	var ints []int64
	for _, v := range lines {
		x, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, x)
	}
	return ints
}

// return the contents of a file as a float
func AsFloat(path string) []float64 {
	lines := readFile(path)

	var ints []float64
	for _, v := range lines {
		x, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, x)
	}
	return ints
}
