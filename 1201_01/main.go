package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var leftArray []int64
	var rightArray []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "   ")
		if arr == nil || len(arr) != 2 {
			log.Fatal("invalid line: ", line)
		}

		left, err := strconv.ParseInt(arr[0], 10, 64)
		if err != nil {
			log.Fatal("invalid left number: ", arr[0])
		}
		right, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			log.Fatal("invalid right number: ", arr[1])
		}
		leftArray = append(leftArray, left)
		rightArray = append(rightArray, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(leftArray) != len(rightArray) {
		log.Fatal("invalid list size")
	}

	slices.Sort(leftArray)
	slices.Sort(rightArray)

	var answer int64
	for i := 0; i < len(leftArray); i++ {
		tmp := leftArray[i] - rightArray[i]
		if tmp < 0 {
			tmp *= -1
		}
		answer += tmp
	}

	fmt.Printf("Answer: %d\n", answer)
}
