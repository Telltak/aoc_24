package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	example_data := strings.NewReader(
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)
	lines := split_lines(example_data)
	expected := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %#v but got %#v", expected, lines)
	}

}

func TestGrid(t *testing.T) {
	example_data := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
	}

	expected := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
	}

	actual := create_grid(example_data)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v but got %#v", expected, actual)
	}
}

func TestHorizontal(t *testing.T) {
	example_data := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
	}
	expected := 2

	actual := search_horizontal(example_data, "XMAS")

	if actual != expected {
		t.Errorf("Expected %#v but got %#v", expected, actual)
	}
}

func TestVertical(t *testing.T) {
	example_data := [][]string{
		{"X", "M", "S"},
		{"M", "X", "A"},
		{"A", "A", "M"},
		{"S", "S", "X"},
	}
	expected := 2

	actual := search_vertical(example_data, "XMAS")

	if actual != expected {
		t.Errorf("Expected %#v but got %#v", expected, actual)
	}
}

func TestDiagonal(t *testing.T) {
	example_data := [][]string{
		{"X", "M", "S", "S"},
		{"M", "M", "A", "S"},
		{"A", "M", "A", "M"},
		{"X", "S", "X", "S"},
	}
	expected := 2

	actual := search_diagonal(example_data, "XMAS")

	if actual != expected {
		t.Errorf("Expected %#v but got %#v", expected, actual)
	}
}

func TestPart1(t *testing.T) {
	// example_left := []int{3, 4, 2, 1, 3, 3}
	// example_right := []int{4, 3, 5, 3, 9, 3}
	// expected_answer := 11
	//
	// answer := calculate_distance(example_left, example_right)
	// if expected_answer != answer {
	// 	t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
	// }
}

func TestPart2(t *testing.T) {
	// example_left := []int{3, 4, 2, 1, 3, 3}
	// example_right := []int{4, 3, 5, 3, 9, 3}
	// expected_answer := 31
	//
	// answer := calculate_similarity(example_left, example_right)
	// if expected_answer != answer {
	// 	t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
	// }
}
