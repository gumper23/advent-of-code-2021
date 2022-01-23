package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	units     int
}

func readMoves(path string) (moves []move, err error) {
	moves = []move{}
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			log.Fatalf("Invalid lines: %s\n", scanner.Text())
		}
		units, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Invalid units: %s\n", parts[1])
		}
		moves = append(moves, move{parts[0], units})
	}
	return moves, scanner.Err()
}

func movesProduct(moves []move) (product int, err error) {
	var depth, horizontal int
	for _, move := range moves {
		switch move.direction {
		case "up":
			depth -= move.units
			if depth < 0 {
				depth = 0
			}
		case "down":
			depth += move.units
			if depth < 0 {
				depth = 0
			}
		case "forward":
			horizontal += move.units
			if horizontal < 0 {
				horizontal = 0
			}
		default:
			return -1, fmt.Errorf("unexpected direction %s", move.direction)
		}
	}

	return depth * horizontal, nil
}

func main() {
	moves, err := readMoves("course.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	product, err := movesProduct(moves)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("Product of %v = %d\n", moves, product)
}
