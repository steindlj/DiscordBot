package game

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Player1    *discordgo.User
	Player2    *discordgo.User
	CurrPlayer *discordgo.User
	Grid       [6][7]int
	RoundCount int
)

// Places a chip in the specified column in the lowest possible row
// and changes the value of the grid based on the current player.
func PlaceChip(col int) (r, c int) {
	RoundCount++
	var val int
	if strings.EqualFold(CurrPlayer.ID, Player1.ID) {
		val = 1
	} else {
		val = 2
	}
	for i := 5; i >= 0; i-- {
		if Grid[i][col] == 0 {
			Grid[i][col] = val
			return i, col
		}
	}
	return
}

// Alternately changes the current player.
func SetNextPlayer() {
	if strings.EqualFold(CurrPlayer.ID, Player1.ID) {
		CurrPlayer = Player2
	} else {
		CurrPlayer = Player1
	}
}

// Checks all possibilities for 4 adjacent chips.
func CheckWin() bool {
	return checkRows() || checkCols() || checkDiagonalsLeft() || checkDiagonalsRight()
}

// Checks all rows for 4 adjacent chips.
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

// Checks all columns for 4 adjacent chips.
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

// Checks all possible diagonals starting in the top left corner for 4 adjacent chips.
func checkDiagonalsLeft() bool {
	for i := 0; i <= 3; i++ {
		if i == 0 {
			if fromTopLeft(i, 0) {
				return true
			}
		} else if i == 3 {
			if fromTopLeft(0, i) {
				return true
			}
		} else {
			if fromTopLeft(i, 0) || fromTopLeft(0, i) {
				return true
			}
		}
	}
	return false
}

// Checks all possible diagonals starting in the top right corner for 4 adjacent chips.
func checkDiagonalsRight() bool {
	for i := 0; i < 6; i++ {
		if i < 3 {
			if fromTopRight(i, 6) {
				return true
			}
		} else {
			if fromTopRight(0, i) {
				return true
			}
		}
	}
	return false
}

// Checks for 4 adjacent chips in a diagonal starting at a point in the top left corner.
func fromTopLeft(i, j int) bool {
	for i+3 <= 5 && j+3 <= 6 {
		if Grid[i][j] == Grid[i+1][j+1] && Grid[i+1][j+1] == Grid[i+2][j+2] && Grid[i+2][j+2] == Grid[i+3][j+3] && Grid[i+3][j+3] != 0 {
			return true
		}
		i++
		j++
	}
	return false
}

// Checks for 4 adjacent chips in a diagonal starting at a point in the top right corner.
func fromTopRight(i, j int) bool {
	for i+3 <= 5 && j-3 >= 0 {
		if Grid[i][j] == Grid[i+1][j-1] && Grid[i+1][j-1] == Grid[i+2][j-2] && Grid[i+2][j-2] == Grid[i+3][j-3] && Grid[i+3][j-3] != 0 {
			return true
		}
		i++
		j--
	}
	return false
}
