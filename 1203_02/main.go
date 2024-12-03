package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var regexpParse = regexp.MustCompile(`(mul\(\d+,\d+\)|don't\(\)|do\(\))`)
var regexpMul = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	muls := parseMuls(lines)

	fmt.Printf("Answer: %d\n", calcMuls(muls))
}

func parseMuls(lines []string) []string {
	var muls []string
	do := true
	for _, line := range lines {
		attrs := regexpParse.FindAllString(line, -1)
		for _, attr := range attrs {
			switch attr {
			case "don't()":
				do = false
			case "do()":
				do = true
			default:
				if do {
					muls = append(muls, attr)
				}
			}
		}
	}
	return muls
}

func calcMuls(muls []string) int64 {
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
