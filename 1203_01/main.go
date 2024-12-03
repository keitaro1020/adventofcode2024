package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var regexpMul = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var answer int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		answer += calcLine(line)
	}
	fmt.Printf("Answer: %d\n", answer)
}

func calcLine(line string) int64 {
	muls := regexpMul.FindAllString(line, -1)
	println(len(muls))
	var result int64
	for _, v := range muls {
		m := regexpMul.FindStringSubmatch(v)
		if len(m) != 3 {
			log.Fatal("invalid value: ", v)
		}
		a, err := strconv.ParseInt(m[1], 10, 64)
		if err != nil {
			log.Fatal("invalid number: ", m[1])
		}
		b, err := strconv.ParseInt(m[2], 10, 64)
		if err != nil {
			log.Fatal("invalid number: ", m[2])
		}
		result += a * b
	}
	return result
}
