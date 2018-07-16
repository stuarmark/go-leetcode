package main

import (
	"fmt"
)

func main() {
	var ddbyte = [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(ddbyte))
}

func isValidSudoku(board [][]byte) bool {
	boardmap := make(map[string]string)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			boardvalue := string(board[i][j])
			if boardvalue != "." {
				currentblock := (i / 3 * 3) + (j / 3)
				_, okr := boardmap[fmt.Sprintf("r%d%s", i, boardvalue)]
				_, okc := boardmap[fmt.Sprintf("c%d%s", j, boardvalue)]
				_, okb := boardmap[fmt.Sprintf("b%d%s", currentblock, boardvalue)]
				if okr || okc || okb {
					return false
				}
				boardmap[fmt.Sprintf("r%d%s", i, boardvalue)] = fmt.Sprintf("r%d%s", i, boardvalue)
				boardmap[fmt.Sprintf("c%d%s", j, boardvalue)] = fmt.Sprintf("c%d%s", j, boardvalue)
				boardmap[fmt.Sprintf("b%d%s", currentblock, boardvalue)] = fmt.Sprintf("b%d%s", currentblock, boardvalue)
			}
		}
	}
	return true
}
