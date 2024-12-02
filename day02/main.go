package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type OrderValidation int8

const (
	Increasing OrderValidation = 0
	Decreasing OrderValidation = 1
	Unordered  OrderValidation = 2
)

func split_lines(file io.Reader) (lines [][]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		var split_line []int
		for _, v := range split {
			converted, _ := strconv.Atoi(v)
			split_line = append(split_line, converted)
		}
		lines = append(lines, split_line)
	}

	return lines
}

type OrderingInformation struct {
	Ordering       OrderValidation
	RemovableValue []int
}

func validate_order(line []int, tolerance bool) OrderingInformation {
	if slices.IsSorted(line) {
		return OrderingInformation{Increasing, nil}
	}
	reversed := slices.Clone(line)
	slices.Reverse(reversed)
	if slices.IsSorted(reversed) {
		return OrderingInformation{Decreasing, nil}
	}
	if tolerance {
		for i := range line {
			tmp := slices.Clone(line)
			reduced_ordering := validate_order(slices.Delete(tmp, i, i+1), false)
			if reduced_ordering.Ordering != Unordered {
				return OrderingInformation{reduced_ordering.Ordering, []int{i}}
			}
		}
	}
	return OrderingInformation{Unordered, nil}
}

type AdjacencyInformation struct {
	Adjacency      bool
	RemovableValue []int
}

func validate_adjacency_levels(line []int, tolerance bool) AdjacencyInformation {
	for i := range line {
		if i == len(line)-1 {
			return AdjacencyInformation{true, nil}
		}
		difference := int(math.Abs(float64(line[i]) - float64(line[i+1])))
		if (1 > difference) || (difference > 3) {
			if tolerance {
				for i := range line {
					tmp := slices.Clone(line)
					reduced_adjacency := validate_adjacency_levels(slices.Delete(tmp, i, i+1), false)
					if reduced_adjacency.Adjacency {
						return AdjacencyInformation{true, []int{i}}
					}
				}
			}

			return AdjacencyInformation{false, nil}
		}
	}

	return AdjacencyInformation{true, nil}
}

func count_safe_reports(lines [][]int, tolerance bool) (count int) {
	for _, line := range lines {
		order := validate_order(line, tolerance)
		adjacent := validate_adjacency_levels(line, tolerance)

		if (order.Ordering == Unordered || !adjacent.Adjacency) || (!reflect.DeepEqual(order.RemovableValue, adjacent.RemovableValue) && (order.RemovableValue != nil && adjacent.RemovableValue != nil)) {
			continue
		}

		count++
	}

	return count
}

func main() {
	file, _ := os.Open("input")
	lines := split_lines(file)
	safe_reports := count_safe_reports(lines, false)

	fmt.Printf("There are %d safe reports\n", safe_reports)

	file, _ = os.Open("input")
	lines = split_lines(file)
	tolerated_reports := count_safe_reports(lines, true)

	fmt.Printf("There are %d tolerated reports\n", tolerated_reports)
}
