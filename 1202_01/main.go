package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var answer int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		if arr == nil {
			log.Fatal("invalid line: ", line)
		}

		input := make([]int64, len(arr))
		for i, v := range arr {
			input[i], err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal("invalid number: ", v)
			}
		}

		if checkSafe(input) {
			answer++
		}

	}
	fmt.Printf("Answer: %d\n", answer)
}

func checkSafe(arr []int64) bool {
	var before *int64
	var order *bool
	for _, v := range arr {
		if before != nil {
			if *before == v {
				return false
			}

			diff := *before - v
			if diff < 0 {
				diff *= -1
			}
			if diff > 3 {
				return false
			}
		}

		if order != nil {
			if *order && *before > v {
				return false
			}
			if !*order && *before < v {
				return false
			}
		}

		if order == nil && before != nil {
			if *before < v {
				order = new(bool)
				*order = true
			} else {
				order = new(bool)
				*order = false
			}
		}

		before = &v
	}

	return true
}
