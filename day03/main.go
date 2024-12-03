package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func split_lines(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

type CalculationTypes int8

const (
	Mul CalculationTypes = 0
)

type Calculation struct {
	CalculationType CalculationTypes
	Left            int
	Right           int
}

func find_calculations(line string) (calculations []Calculation) {
	multiplication_regex, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	multiplication_matches := multiplication_regex.FindAllStringSubmatch(line, -1)
	for _, v := range multiplication_matches {
		left, _ := strconv.Atoi(v[1])
		right, _ := strconv.Atoi(v[2])
		calculations = append(calculations, Calculation{Mul, left, right})
	}
	return calculations
}

func perform_calculations(calculations []Calculation) (total int) {
	for _, v := range calculations {
		switch v.CalculationType {
		case Mul:
			total += v.Left * v.Right
		}
	}
	return total
}

func get_total(lines []string) (total int) {
	for _, line := range lines {
		line_calcs := find_calculations(line)
		total += perform_calculations(line_calcs)
	}
	return total
}

func main() {
	file, _ := os.Open("input")
	lines := split_lines(file)
	mul_total := get_total(lines)
	fmt.Printf("Total is %d\n", mul_total)
	//
	// similarity := calculate_similarity(left, right)
	//
	// fmt.Printf("Similarity is %d\n", similarity)
}
