package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func split_sides(line string) (left int, right int) {
	splitFn := func(c rune) bool { return c == ' ' }
	parts := strings.FieldsFunc(line, splitFn)

	left, _ = strconv.Atoi(parts[0])
	right, _ = strconv.Atoi(parts[1])

	return left, right
}

func calculate_distance(left []int, right []int) (distance int) {
	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		distance += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	return distance

}

func calculate_similarity(left []int, right []int) (similarity int) {
	count := make(map[int]int)

	for _, value := range right {
		count[value] += 1
	}

	for _, value := range left {
		similarity += value * count[value]
	}

	return similarity
}

func split_lists(file io.Reader) (left []int, right []int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l, r := split_sides(scanner.Text())
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func main() {
	file, _ := os.Open("input")
	left, right := split_lists(file)
	distance := calculate_distance(left, right)

	fmt.Printf("Distance is %d\n", distance)

	similarity := calculate_similarity(left, right)

	fmt.Printf("Similarity is %d\n", similarity)
}
