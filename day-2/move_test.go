package main

import (
	"reflect"
	"testing"
)

func TestReadMoves(t *testing.T) {
	var tables = []struct {
		path     string
		expected []move
	}{
		{"course.txt", []move{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}},
		{"empty.txt", []move{}},
	}
	for _, table := range tables {
		moves, err := readMoves(table.path)
		if err != nil {
			t.Errorf("Unexpected error, got %s", err.Error())
		}
		if !reflect.DeepEqual(table.expected, moves) {
			t.Errorf("Unexpected results, got %v, want %v", moves, table.expected)
		}
	}
}

func TestMovesProduct(t *testing.T) {
	var tables = []struct {
		moves    []move
		expected int
	}{
		{[]move{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}, 150},
		{[]move{}, 0},
		{[]move{{"forward", 5}}, 0},
		{[]move{{"down", 5}}, 0},
		{[]move{{"down", 5}, {"forward", 5}}, 25},
		{[]move{{"up", 5}, {"up", 5}}, 0},
	}
	for _, table := range tables {
		product, err := movesProduct(table.moves)
		if err != nil {
			t.Errorf("Unexpected error, got %s", err.Error())
		}
		if product != table.expected {
			t.Errorf("Unexpected results, got %v, want %v", product, table.expected)
		}
	}
}
