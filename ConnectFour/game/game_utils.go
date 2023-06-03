package game

import (
	"strings"
)

var (
	Player1Id string
	Player2Id string
	Grid [6][7]int
	CurrPlayer string
)

func CheckWin() bool {
	return checkRows() || checkCols() || checkDiagonalsLeft() || checkDiagonalsRight()
}

func checkRows() bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7-3; j++ {
			if Grid[i][j] == Grid[i][j+1] && Grid[i][j+1] == Grid[i][j+2] && Grid[i][j+2] == Grid[i][j+3] && Grid[i][j+3] != 0 {
				return true
			}
		}
	}
	return false
}

func checkCols() bool {
	for i := 0; i < 7; i++ {
		for j := 0; j < 6-3; j++ {
			if Grid[j][i] == Grid[j+1][i] && Grid[j+1][i] == Grid[j+2][i] && Grid[j+2][i] == Grid[j+3][i] && Grid[j+3][i] != 0 {
				return true
			}
		}
	}
	return false
}

func checkDiagonalsLeft() bool {
	for i := 0; i <= 3; i++ {
		if i == 0 {
			if fromUpperLeft(i, 0) {
				return true
			}
		} else if i == 3 {
			if fromUpperLeft(0, i) {
				return true
			}
		} else {
			if fromUpperLeft(i, 0) || fromUpperLeft(0, i) {
				return true
			}
		}
	}
	return false
}

func checkDiagonalsRight() bool {
	for i := 0; i < 6; i++ {
		if i < 3 {
			if fromUpperRight(i, 6) {
				return true
			}
		} else {
			if fromUpperRight(0, i) {
				return true
			}
		}
	}
	return false
}

func fromUpperLeft(i, j int) bool {
	for i+3 <= 5 && j+3 <= 6 {
		if Grid[i][j] == Grid[i+1][j+1] && Grid[i+1][j+1] == Grid[i+2][j+2] && Grid[i+2][j+2] == Grid[i+3][j+3] && Grid[i+3][j+3] != 0 {
			return true
		}
		i++
		j++
	}
	return false
}

func fromUpperRight(i, j int) bool {
	for i+3 <= 5 && j-3 >= 0 {
		if Grid[i][j] == Grid[i+1][j-1] && Grid[i+1][j-1] == Grid[i+2][j-2] && Grid[i+2][j-2] == Grid[i+3][j-3] && Grid[i+3][j-3] != 0 {
			return true
		}
		i++
		j--
	}
	return false
}

func SetChip(col int) (r, c int){
	var val int
	if strings.EqualFold(CurrPlayer, Player1Id) {
		val = 1;
	} else {
		val = 2;
	}
	for i := 5; i >= 0; i-- {
		if Grid[i][col] == 0 {
		 	Grid[i][col] = val
			 return i, col
		}
	}
	return 
}