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

	fieldMap := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fieldMap = append(fieldMap, strings.Split(line, ""))
	}
	workedFieldMap, err := work(fieldMap)
	if err != nil {
		log.Fatal(err)
	}
	printMap(workedFieldMap)
	fmt.Printf("Answer: %d\n", countVistCount(workedFieldMap))
}

type Direction int

const (
	InvalidDirection Direction = iota
	North
	East
	South
	West
)

type Position struct {
	X         int
	Y         int
	Direction Direction
}

func work(fieldMap [][]string) ([][]string, error) {
	guardPosition, err := searchGuardPosition(fieldMap)
	if err != nil {
		return nil, err
	}
	log.Printf("initial guard position: %v", *guardPosition)

	switchDirectionCount := 0
	for {
		cellType, x, y := getNextCellType(fieldMap, guardPosition)
		switch cellType {
		case Out:
			fieldMap[guardPosition.Y][guardPosition.X] = VisitedCell
			return fieldMap, nil
		case Obstructions:
			if switchDirectionCount == 4 {
				return nil, fmt.Errorf("invalid position: %v", *guardPosition)
			}
			nextDirection := switchDirection(guardPosition.Direction)
			if nextDirection == InvalidDirection {
				return nil, fmt.Errorf("invalid direction: %v", guardPosition.Direction)
			}
			guardPosition.Direction = nextDirection
			switch nextDirection {
			case North:
				fieldMap[guardPosition.Y][guardPosition.X] = GuardNorthCell
			case East:
				fieldMap[guardPosition.Y][guardPosition.X] = GuardEastCell
			case South:
				fieldMap[guardPosition.Y][guardPosition.X] = GuardSouthCell
			case West:
				fieldMap[guardPosition.Y][guardPosition.X] = GuardWestCell
			}
			switchDirectionCount++
			continue
		default:
			switchDirectionCount = 0
			fieldMap[y][x] = fieldMap[guardPosition.Y][guardPosition.X]
			fieldMap[guardPosition.Y][guardPosition.X] = VisitedCell
			guardPosition.X = x
			guardPosition.Y = y
			continue
		}
	}

	return nil, nil
}

type CellType int

const (
	InvalidCell CellType = iota
	Empty
	Obstructions
	Visited
	Guard
	Out
)
const (
	EmptyCell        = "."
	ObstructionsCell = "#"
	VisitedCell      = "X"
	GuardNorthCell   = "^"
	GuardEastCell    = ">"
	GuardSouthCell   = "v"
	GuardWestCell    = "<"
)

func cellCheck(cell string) (CellType, Direction) {
	switch cell {
	case EmptyCell:
		return Empty, InvalidDirection
	case ObstructionsCell:
		return Obstructions, InvalidDirection
	case VisitedCell:
		return Visited, InvalidDirection
	case GuardNorthCell:
		return Guard, North
	case GuardEastCell:
		return Guard, East
	case GuardSouthCell:
		return Guard, South
	case GuardWestCell:
		return Guard, West
	}
	return InvalidCell, InvalidDirection
}

func searchGuardPosition(fieldMap [][]string) (*Position, error) {
	position := &Position{}

	for y, line := range fieldMap {
		for x, cell := range line {
			cellType, direction := cellCheck(cell)
			if cellType == Guard {
				if direction == InvalidDirection {
					return nil, fmt.Errorf("invalid direction: %v", cell)
				}
				position.X = x
				position.Y = y
				position.Direction = direction
			}
		}
	}
	return position, nil
}

func getNextCellType(fieldMap [][]string, position *Position) (cellType CellType, x int, y int) {
	cellType, x, y = InvalidCell, position.X, position.Y
	switch position.Direction {
	case North:
		y--
	case East:
		x++
	case South:
		y++
	case West:
		x--
	}
	if x == -1 || x == len(fieldMap[0]) || y == -1 || y == len(fieldMap) {
		cellType = Out
		return
	}
	cellType, _ = cellCheck(fieldMap[y][x])
	return
}

func switchDirection(direction Direction) Direction {
	switch direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	return InvalidDirection
}

func countVistCount(fieldMap [][]string) int64 {
	var count int64
	for _, line := range fieldMap {
		for _, cell := range line {
			if cellType, _ := cellCheck(cell); cellType == Visited {
				count++
			}
		}
	}
	return count
}

func printMap(fieldMap [][]string) {
	for _, line := range fieldMap {
		for _, cell := range line {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}