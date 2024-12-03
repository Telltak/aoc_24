package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	example_data := strings.NewReader(
		`where())wh
,who()>}wh
mul(888,16`)
	lines := split_lines(example_data)
	expected := []string{
		"where())wh",
		",who()>}wh",
		"mul(888,16",
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %#v but got %#v", expected, lines)
	}

}

func TestRegexr(t *testing.T) {
	example_data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	calcs := find_calculations(example_data)
	expected := []Calculation{
		{Mul, 2, 4},
		{Mul, 5, 5},
		{Mul, 11, 8},
		{Mul, 8, 5},
	}

	if !reflect.DeepEqual(expected, calcs) {
		t.Errorf("Expected %#v but got %#v", expected, calcs)
	}
}

func TestCalculations(t *testing.T) {
	example_calculations := []Calculation{
		{Mul, 2, 4},
		{Mul, 5, 5},
		{Mul, 11, 8},
		{Mul, 8, 5},
	}

	total := perform_calculations(example_calculations)
	expected := 161

	if total != expected {
		t.Errorf("Expected %d but got %d", expected, total)
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
