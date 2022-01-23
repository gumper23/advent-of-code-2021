package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readSweeps(path string) (sweeps []int, err error) {
	sweeps = []int{}

	file, err := os.Open(path)
	if err != nil {
		return sweeps, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sweep, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return sweeps, err
		}
		sweeps = append(sweeps, sweep)
	}
	return sweeps, scanner.Err()
}

func increasingSweeps(sweeps []int) (increases int) {
	for i := 0; i < len(sweeps); i++ {
		if i == 0 {
			continue
		} else if sweeps[i] > sweeps[i-1] {
			increases++
		}
	}
	return increases
}

func main() {
	sweeps, err := readSweeps("sweeps-1.txt")
	if err != nil {
		log.Fatalf("Error reading sweeps-1.txt: %s", err.Error())
	}
	fmt.Printf("%v\n", sweeps)
	fmt.Printf("Number of increasing sweeps: [%d]\n", increasingSweeps(sweeps))

	sweeps, err = readSweeps("sweeps-2.txt")
	if err != nil {
		log.Fatalf("Error reading sweeps-1.txt: %s", err.Error())
	}
	fmt.Printf("%v\n", sweeps)
	fmt.Printf("Number of increasing sweeps: [%d]\n", increasingSweeps(sweeps))
}
