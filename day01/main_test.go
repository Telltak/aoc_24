package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	example_data := strings.NewReader(
		`3   4
4   3
2   5
1   3
3   9
3   3`)
	left, right := split_lists(example_data)
	expected_left := []int{3, 4, 2, 1, 3, 3}
	expected_right := []int{4, 3, 5, 3, 9, 3}

	if !reflect.DeepEqual(left, expected_left) {
		t.Errorf("Expected %#v but got %#v", expected_left, left)
	}

	if !reflect.DeepEqual(right, expected_right) {
		t.Errorf("Expected %#v but got %#v", expected_right, right)
	}

}

func TestPart1(t *testing.T) {
	example_left := []int{3, 4, 2, 1, 3, 3}
	example_right := []int{4, 3, 5, 3, 9, 3}
	expected_answer := 11

	answer := calculate_distance(example_left, example_right)
	if expected_answer != answer {
		t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
	}
}

func TestPart2(t *testing.T) {
	example_left := []int{3, 4, 2, 1, 3, 3}
	example_right := []int{4, 3, 5, 3, 9, 3}
	expected_answer := 31

	answer := calculate_similarity(example_left, example_right)
	if expected_answer != answer {
		t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
	}
}
