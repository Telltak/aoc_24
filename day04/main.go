package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func split_lines(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func create_grid(lines []string) (grid [][]string) {
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func reverse_string(s string) (reversed string) {
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return reversed
}

func search_horizontal(grid [][]string, search_string string) (count int) {
	reversed_string := reverse_string(search_string)

	for _, line := range grid {
		for i := 0; i < len(line)-3; i++ {
			candidate := strings.Join(line[i:i+4], "")
			if candidate == search_string || candidate == reversed_string {
				count++
			}
		}
	}
	return count
}

func search_vertical(grid [][]string, search_string string) (count int) {
	reversed_string := reverse_string(search_string)

	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid)-3; row++ {
			var candidate string
			for i := 0; i < 4; i++ {
				candidate += string(grid[row+i][col])
			}
			if candidate == search_string || candidate == reversed_string {
				count++
			}
		}
	}
	return count
}

func search_diagonal(grid [][]string, search_string string) (count int) {
	reversed_string := reverse_string(search_string)
	// down right
	for col := 0; col < len(grid[0])-3; col++ {
		for row := 0; row < len(grid)-3; row++ {
			var candidate string
			for i := 0; i < 4; i++ {
				candidate += string(grid[row+i][col+i])
			}
			if candidate == search_string || candidate == reversed_string {
				count++
			}
		}
	}

	// down left
	for col := 3; col < len(grid[0]); col++ {
		for row := 0; row < len(grid)-3; row++ {
			var candidate string
			for i := 0; i < 4; i++ {
				candidate += string(grid[row+i][col-i])
			}
			if candidate == search_string || candidate == reversed_string {
				count++
			}
		}
	}

	return count
}

func search_xmas(grid [][]string, search_string string) (count int) {
	count += search_horizontal(grid, search_string)
	count += search_vertical(grid, search_string)
	count += search_diagonal(grid, search_string)

	return count
}

func search_mas(grid [][]string) (count int) {
	for col := 0; col < len(grid[0])-2; col++ {
		for row := 0; row < len(grid)-2; row++ {
			var left string
			var right string
			for i := 0; i < 3; i++ {
				left += string(grid[row+i][col-i+2])
				right += string(grid[row+i][col+i])

				if (left == "MAS" || left == "SAM") && (right == "MAS" || right == "SAM") {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	file, _ := os.Open("input")
	lines := split_lines(file)
	grid := create_grid(lines)

	xmas_total := search_xmas(grid, "XMAS")

	fmt.Printf("Xmas total: %d\n", xmas_total)

	mas_total := search_mas(grid)

	fmt.Printf("mas total: %d\n", mas_total)

}
