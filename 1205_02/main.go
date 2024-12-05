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

	scanner := bufio.NewScanner(file)

	readRule := true
	pageRule := make(map[int64][]int64)
	pages := make([][]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readRule = false
			continue
		}
		if readRule {
			pageRule, err = parsePageRule(line, pageRule)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			parsedPages, err := parsePages(line)
			if err != nil {
				log.Fatal(err)
			}
			pages = append(pages, parsedPages)
		}
	}
	log.Printf("pageRule size: %d, pages size: %d\n", len(pageRule), len(pages))

	filteredPages := filterPages(pageRule, pages)
	log.Printf("filteredPages size: %d\n", len(filteredPages))

	sortedPages := sortPages(pageRule, filteredPages)

	answer, err := sumCenterPageNumber(sortedPages)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Answer: %d\n", answer)
}

func parsePageRule(line string, pageRule map[int64][]int64) (map[int64][]int64, error) {
	arr := strings.Split(line, "|")
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid line: %v", line)
	}

	left, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid left number: %v", arr[0])
	}
	right, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid right number: %v", arr[1])
	}

	pageRule[left] = append(pageRule[left], right)

	return pageRule, nil
}

func parsePages(line string) ([]int64, error) {
	arr := strings.Split(line, ",")

	var pages []int64
	for _, page := range arr {
		pageNum, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid page number: %v", page)
		}
		pages = append(pages, pageNum)
	}

	return pages, nil
}

func filterPages(pageRule map[int64][]int64, pages [][]int64) [][]int64 {
	var filteredPages [][]int64
	for _, pageLine := range pages {
		for i, page := range pageLine {
			if i == 0 {
				continue
			}
			if !checkPageRule(pageRule[page], pageLine[:i]) {
				filteredPages = append(filteredPages, pageLine)
				break
			}
		}
	}
	return filteredPages
}

func checkPageRule(pageRule []int64, inputedPages []int64) bool {
	for _, rule := range pageRule {
		if slices.Contains(inputedPages, rule) {
			return false
		}
	}
	return true
}

func sortPages(pageRule map[int64][]int64, pages [][]int64) [][]int64 {
	var sortedPages [][]int64
	for _, pageLine := range pages {
		sortedPages = append(sortedPages, sortPageLine(pageRule, pageLine))
	}

	return sortedPages
}

func sortPageLine(pageRule map[int64][]int64, pageLine []int64) []int64 {
	sortedPageLine := make([]int64, 0)

	for _, page := range pageLine {
		rule := pageRule[page]
		idx := 0
		for i := len(sortedPageLine); i >= 0; i-- {
			allOK := true
			for _, r := range rule {
				if slices.Contains(sortedPageLine[0:i], r) {
					allOK = false
					break
				}
			}
			if allOK {
				idx = i
				break
			}
		}
		sortedPageLine = slices.Insert(sortedPageLine, idx, page)
	}

	return sortedPageLine
}

func sumCenterPageNumber(pages [][]int64) (int64, error) {
	var sum int64
	for _, pageLine := range pages {
		if len(pageLine)%2 == 0 {
			return 0, fmt.Errorf("invalid page line: %v", pageLine)
		}

		centerIndex := len(pageLine) / 2
		sum += pageLine[centerIndex]
	}
	return sum, nil
}
