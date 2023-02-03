package zigzagconversion

import (
	"fmt"
)

func Convert(s string, numRows int) string {

	grid := make([][]string, numRows)
	appendToAllRows(grid)

	row := 0
	col := 0
	diagonal := false

	if numRows == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		strChar := string([]byte{s[i]})

		if row == 0 {
			diagonal = false
		} else if row == numRows-1 {
			diagonal = true
		}

		grid[row][col] = strChar
		if diagonal {
			row = row - 1
			appendToAllRows(grid)
			col++
		} else {
			row++
		}
	}

	return prettyPrint(grid)
}

func prettyPrint(grid [][]string) string {
	prettyStr := ""
	for i, row := range grid {
		fmt.Println(row)
		for j := range row {
			if grid[i][j] == " " {
				continue
			}
			prettyStr = fmt.Sprintf("%s%s", prettyStr, grid[i][j])
		}
	}
	return prettyStr
}

func appendToAllRows(grid [][]string) [][]string {
	for i, row := range grid {
		row = append(row, " ")
		grid[i] = row
	}
	return grid
}
