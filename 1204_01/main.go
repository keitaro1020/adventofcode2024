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
			if cell == "X" {
				answer += int64(check(i, j, table))
			}
		}
	}

	fmt.Printf("Answer: %d\n", answer)
}

func check(i, j int, table [][]string) int {
	answer := 0
	// 上方向チェック
	if i >= 3 && table[i-1][j] == "M" && table[i-2][j] == "A" && table[i-3][j] == "S" {
		answer++
	}

	// 下方向チェック
	if len(table) >= i+4 && table[i+1][j] == "M" && table[i+2][j] == "A" && table[i+3][j] == "S" {
		answer++
	}

	// 左方向チェック
	if j >= 3 && table[i][j-1] == "M" && table[i][j-2] == "A" && table[i][j-3] == "S" {
		answer++
	}

	// 右方向チェック
	if len(table[i]) >= j+4 && table[i][j+1] == "M" && table[i][j+2] == "A" && table[i][j+3] == "S" {
		answer++
	}

	// 左斜め上方向チェック
	if i >= 3 && j >= 3 && table[i-1][j-1] == "M" && table[i-2][j-2] == "A" && table[i-3][j-3] == "S" {
		answer++
	}

	// 左斜め下方向チェック
	if len(table) >= i+4 && j >= 3 && table[i+1][j-1] == "M" && table[i+2][j-2] == "A" && table[i+3][j-3] == "S" {
		answer++
	}

	// 右斜め下方向チェック
	if len(table) >= i+4 && len(table[i]) >= j+4 && table[i+1][j+1] == "M" && table[i+2][j+2] == "A" && table[i+3][j+3] == "S" {
		answer++
	}

	// 右斜め上方向チェック
	if i >= 3 && len(table[i]) >= j+4 && table[i-1][j+1] == "M" && table[i-2][j+2] == "A" && table[i-3][j+3] == "S" {
		answer++
	}
	return answer
}
