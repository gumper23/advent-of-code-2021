package main

import (
	"reflect"
	"testing"
)

func TestIncreasingSweeps(t *testing.T) {
	var tables = []struct {
		sweeps   []int
		expected int
	}{
		{[]int{2, 7, 11, 15}, 3},
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{[]int{269, 263, 260, 240, 210, 208, 207, 200, 199}, 0},
		{[]int{1, 1, 2, 3, 5, 8, 13}, 5},
		{[]int{}, 0},
		{[]int{1}, 0},
	}
	for _, table := range tables {
		increases := increasingSweeps(table.sweeps)
		if table.expected != increases {
			t.Errorf("Unexpected results, got %v, want %v", increases, table.expected)
		}
	}
}

func TestReadSweeps(t *testing.T) {
	var tables = []struct {
		path     string
		expected []int
	}{
		{"sweeps-1.txt", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}},
		{"sweeps-2.txt", []int{}},
	}
	for _, table := range tables {
		sweeps, err := readSweeps(table.path)
		if err != nil {
			t.Errorf("Unexpected error for %s, got %s", table.path, err.Error())
		}
		if !reflect.DeepEqual(sweeps, table.expected) {
			t.Errorf("Unexpected results, got [%v], want [%v]", sweeps, table.expected)
		}
	}
}
