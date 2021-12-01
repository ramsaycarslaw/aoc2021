package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func threeSum(lines []int64) int {
	var f1, f2, f3 int64 = 0, 0, 0

	var prev int64 = 0
	inc := 0

	for i := 0; i < len(lines)-2; i++ {
		f1 = lines[i]
		f2 = lines[i+1]
		f3 = lines[i+2]

		var curr int64 = f1 + f2 + f3

		if i == 0 {
			prev = curr
			continue
		}

		if curr > prev {
			inc++
		}

		prev = curr
	}

	return inc
}

func day1() {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var lines []int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, x)
	}

	var prev int64 = 0
	inc := 0

	for i, v := range lines {
		if i == 0 {
			prev = v
			continue
		}

		if v > prev {
			inc++
		}

		prev = v
	}

	fmt.Println(inc)
	fmt.Println(threeSum(lines))
}
