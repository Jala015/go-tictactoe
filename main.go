package main

import (
	"fmt"
)

type Board [3][3]int

func main() {
	currentBoard := Board{
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
	}

	renderBoard(currentBoard)
	newX, newY := receiveInput(currentBoard)
	currentBoard[newX][newY] = 1
	renderBoard(currentBoard)
}

// receives a coordinate to make a move and return the values only when they are valid
func receiveInput(b Board) (int, int) {
	var boardX, boardY int

	for { //loop until valid coordinates
		fmt.Print("Type coordinates separated by space: ")
		_, err := fmt.Scan(&boardX, &boardY)

		if err != nil { //input is not valid
			fmt.Println("Error reading coordinates:", err)
			continue
		}
		if (boardX >= 0 && boardX <= 2) && (boardY >= 0 && boardY <= 2) { // check if user input coordinates are on a valid range
			if b[boardX][boardY] == 0 { // check if user input coordinates are on an empy cell
				break
			} else {
				fmt.Println("Coordinate is not empty")
			}
		} else {
			fmt.Println("Coordinates are not on a valid range (0,1 or 2 for each axis)")
		}
	}
	return boardX, boardY
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
