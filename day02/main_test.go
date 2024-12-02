package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	example_data := strings.NewReader(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)
	split := split_lines(example_data)
	expected := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	if !reflect.DeepEqual(split, expected) {
		t.Errorf("Expected %#v but got %#v", expected, split)
	}

}

type TestOrdering struct {
	Slice          []int
	OrderingType   OrderValidation
	RemovableValue []int
}

func TestOrderingFunc(t *testing.T) {
	tests := []TestOrdering{
		{[]int{1, 2, 3, 4, 5}, Increasing, nil},
		{[]int{5, 4, 3, 2, 1}, Decreasing, nil},
		{[]int{1, 4, 2, 3, 5}, Increasing, []int{4}},
		{[]int{1, 4, 2, 6, 5}, Unordered, nil},
	}

	for _, v := range tests {
		result := validate_order(v.Slice, true)
		if result.Ordering != v.OrderingType {
			t.Errorf("Expected %v but got %v for slice %#v", v.OrderingType, result.Ordering, v.Slice)
		}
		if !reflect.DeepEqual(result.RemovableValue, v.RemovableValue) {
			t.Errorf("Expected removable value %v but got %v for slice %#v", v.RemovableValue, result.RemovableValue, v.Slice)
		}
	}

}

type TestAdjacency struct {
	Slice          []int
	Adjacency      bool
	RemovableValue []int
}

func TestAdjacencyFunc(t *testing.T) {
	tests := []TestAdjacency{
		{[]int{1, 2, 3, 4, 5}, true, nil},
		{[]int{1, 3, 6, 7, 9}, true, nil},
		{[]int{1, 2, 7, 8, 9}, false, nil},
		{[]int{1, 2, 2, 3, 4}, true, []int{2}},
	}

	for _, v := range tests {
		result := validate_adjacency_levels(v.Slice, true)
		if result.Adjacency != v.Adjacency {
			t.Errorf("Expected %v but got %v for slice %#v", v.Adjacency, result.Adjacency, v.Slice)
		}
		if !reflect.DeepEqual(result.RemovableValue, v.RemovableValue) {
			t.Errorf("Expected removable value %v but got %v for slice %#v", v.RemovableValue, result.RemovableValue, v.Slice)
		}
	}
}

func TestSafeReports(t *testing.T) {
	lines := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	expected := 2

	safe_count := count_safe_reports(lines, false)

	if safe_count != expected {
		t.Errorf("Expected %v, got %v", expected, safe_count)
	}

}

func TestToleratedReports(t *testing.T) {
	lines := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	expected := 4

	safe_count := count_safe_reports(lines, true)

	if safe_count != expected {
		t.Errorf("Expected %v, got %v", expected, safe_count)
	}

}

// func TestPart1(t *testing.T) {
// 	example_left := []int{3, 4, 2, 1, 3, 3}
// 	example_right := []int{4, 3, 5, 3, 9, 3}
// 	expected_answer := 11
//
// 	answer := calculate_distance(example_left, example_right)
// 	if expected_answer != answer {
// 		t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
// 	}
// }
//
// func TestPart2(t *testing.T) {
// 	example_left := []int{3, 4, 2, 1, 3, 3}
// 	example_right := []int{4, 3, 5, 3, 9, 3}
// 	expected_answer := 31
//
// 	answer := calculate_similarity(example_left, example_right)
// 	if expected_answer != answer {
// 		t.Errorf("Expected '%d' but got '%d'", expected_answer, answer)
// 	}
// }
