package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var answer int64
	table := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, strings.Split(line, ""))
	}

	for i, line := range table {
		for j, cell := range line {
			if cell == "A" {
				answer += int64(check(i, j, table))
			}
		}
	}

	fmt.Printf("Answer: %d\n", answer)
}

func check(i, j int, table [][]string) int {
	answer := 0

	if i < 1 || j < 1 || len(table) < i+2 || len(table[i]) < j+2 {
		return 0
	}

	// M_S
	// _A_
	// M_S
	if table[i-1][j-1] == "M" && table[i-1][j+1] == "S" && table[i+1][j-1] == "M" && table[i+1][j+1] == "S" {
		answer++
	}

	// M_M
	// _A_
	// S_S
	if table[i-1][j-1] == "M" && table[i-1][j+1] == "M" && table[i+1][j-1] == "S" && table[i+1][j+1] == "S" {
		answer++
	}

	// S_M
	// _A_
	// S_M
	if table[i-1][j-1] == "S" && table[i-1][j+1] == "M" && table[i+1][j-1] == "S" && table[i+1][j+1] == "M" {
		answer++
	}

	// S_S
	// _A_
	// M_M
	if table[i-1][j-1] == "S" && table[i-1][j+1] == "S" && table[i+1][j-1] == "M" && table[i+1][j+1] == "M" {
		answer++
	}

	return answer
}
