package main

import (
	"fmt"
)

type Board [3][3]int

func main() {
	board := Board{
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
	}

	renderBoard(board)

	board[0][0] = 1
	board[1][2] = -1
	renderBoard(board)
}

func renderBoard(b Board) {
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
		for j := 0; j < 3; j++ {
			switch b[i][j] {
			case 0:
				fmt.Print("■ ")
			case 1:
				fmt.Print("✗ ")
			case -1:
				fmt.Print("○ ")
			default:
				fmt.Print("? ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

}
