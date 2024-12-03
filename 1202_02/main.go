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

	var answer1 int64
	var answer2 int64
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
			answer1++
			answer2++
		} else {
			reduced := reducedElementsList(input)
			for _, v := range reduced {
				if checkSafe(v) {
					answer2++
					break
				}
			}
		}

	}
	fmt.Printf("Answer1: %d, Answer2; %d, diff: %d\n", answer1, answer2, answer1-answer2)
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

func reducedElementsList(arr []int64) [][]int64 {
	var result [][]int64
	for i := 0; i < len(arr); i++ {
		tmp := make([]int64, len(arr)-1)
		for j, k := 0, 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			tmp[k] = arr[j]
			k++
		}
		result = append(result, tmp)
	}
	return result
}
