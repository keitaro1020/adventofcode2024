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
	var cnt int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		ret, err := line.determineTrueLine()
		if err != nil {
			log.Fatal(err)
		}
		if ret {
			answer += line.TestValue
			fmt.Println(answer)
			cnt++
		}
	}

	fmt.Printf("Answer: %d (%d)\n", answer, cnt)
}

type Line struct {
	TestValue int64
	Numbers   []int64
}

func parseLine(line string) (*Line, error) {
	var l Line

	for _, v := range strings.Split(line, " ") {
		if strings.HasSuffix(v, ":") {
			i, err := strconv.ParseInt(strings.TrimRight(v, ":"), 10, 64)
			if err != nil {
				return nil, err
			}
			l.TestValue = i
		} else {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			l.Numbers = append(l.Numbers, i)
		}
	}

	return &l, nil
}

func (l *Line) determineTrueLine() (bool, error) {
	operatorTable, err := l.getOperatorCombination()
	if err != nil {
		return false, err
	}
	for _, operators := range operatorTable {
		result, err := calc(l.Numbers, operators)
		if err != nil {
			return false, err
		}
		if result == l.TestValue {
			return true, nil
		}
	}
	return false, nil
}

func calc(numbers []int64, operators []string) (int64, error) {
	result := numbers[0]
	for i, o := range operators {
		switch o {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		case "||":
			//v, err := strconv.ParseInt(fmt.Sprintf("%d%d", numbers[i], numbers[i+1]), 10, 64)
			//if err != nil {
			//	return 0, err
			//}
			//if i > 0 {
			//	switch operators[i-1] {
			//	case "+":
			//		result -= numbers[i]
			//		result += v
			//	case "*":
			//		result /= numbers[i]
			//		result *= v
			//	case "||":
			//		v, err = strconv.ParseInt(fmt.Sprintf("%d%d", result, numbers[i+1]), 10, 64)
			//		if err != nil {
			//			return 0, err
			//		}
			//		result = v
			//	}
			//} else {
			//	result = v
			//}
			v, err := strconv.ParseInt(fmt.Sprintf("%d%d", result, numbers[i+1]), 10, 64)
			if err != nil {
				return 0, err
			}
			result = v
		}
	}
	return result, nil
}

func (l *Line) getOperatorCombination() ([][]string, error) {
	max, err := strconv.ParseInt(strings.Repeat("2", len(l.Numbers)-1), 3, 64)
	if err != nil {
		return nil, err
	}

	operators := make([][]string, max+1)
	for i := int64(0); i <= max; i++ {
		operators[i] = make([]string, len(l.Numbers)-1)
		for j, c := range fmt.Sprintf("%0"+strconv.FormatInt(int64(len(l.Numbers))-1, 10)+"s", strconv.FormatInt(i, 3)) {
			switch c {
			case '0':
				operators[i][j] = "+"
			case '1':
				operators[i][j] = "*"
			case '2':
				operators[i][j] = "||"
			}
		}
	}

	return operators, nil
}
