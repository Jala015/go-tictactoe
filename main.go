package main

import (
	"fmt"
)

type Board [3][3]int

func main() {
	currentBoard := Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	player := 1 //track current player

	currentBoard.render()
	for currentBoard.victory() == 0 {
		if player == 1 {
			newX, newY := receiveInput(currentBoard)
			currentBoard[newX][newY] = 1
			player = -1
		} else {
			//todo lógica da máquina
			player = 1
		}
		currentBoard.render()
		fmt.Printf("Score: %v\n", currentBoard.calculateScore())
	}

	fmt.Println(currentBoard.victory())
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

//draw a small tictactoe board on terminal
func (b *Board) render() {
	fmt.Println()
	for i := 0; i < 3; i++ { //iterate rows
		fmt.Print(" ")
		for j := 0; j < 3; j++ { // iterate columns
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

//give a score to the current board configuration. The bigger the score, higher chances to win.
func (b *Board) calculateScore() float32 {
	var score float32 = 0
	winCombinations := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, // upper row
		{{1, 0}, {1, 1}, {1, 2}}, // middle row
		{{2, 0}, {2, 1}, {2, 2}}, // lower row
		{{0, 0}, {1, 0}, {2, 0}}, // left column
		{{0, 1}, {1, 1}, {2, 1}}, // middle column
		{{0, 2}, {1, 2}, {2, 2}}, // right column
		{{0, 2}, {1, 1}, {2, 0}}, // ascending diagonal
		{{0, 0}, {1, 1}, {2, 2}}, // descending diagonal
	}
	for _, line := range winCombinations {
		sum := 0
		for _, coord := range line {
			sum += b[coord[0]][coord[1]]
		}
		switch sum {
		case 3:
			score = 100.0
			return score
		case -3:
			score = -100.0
			return score
		case 2:
			score += (100.0 - score) / 2
		case -2:
			score += (-100.0 - score) / 2
		}
	}
	return score
}

//find coordinates with 0 values
func (b *Board) availableMoves() [][2]int {
	var result [][2]int
	for i := 0; i < 3; i++ { //iterate rows
		for j := 0; j < 3; j++ { //iterate columns
			if b[i][j] == 0 {
				result = append(result, [2]int{i, j})
			}
		}
	}
	return result
}

// check victory status
// 0: game not ended
// 1: ✗ won
// 2: ○ won
// 3: Draw
func (b *Board) victory() int {
	currentScore := b.calculateScore()
	switch {
	case currentScore >= 100:
		return 1
	case currentScore <= -100:
		return 2
	default:
		if len(b.availableMoves()) == 0 {
			return 3
		} else {
			return 0
		}
	}
}
