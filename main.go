package main

import (
	"fmt"
	"os"
)

type Board [3][3]int
type Coordinate [2]int

var maxRecursionLevel int = 3

func main() {
	currentBoard := Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	//track current player. Set to 1 for the computer to start playing
	player := -1

	currentBoard.render(Coordinate{5, 5})
	for currentBoard.victory() == 0 {
		var newX, newY int
		if player == -1 {
			//
			// 👨‍💻 user plays
			//

			fmt.Println("👨‍💻User turn (✗):")
			newX, newY = receiveInput(currentBoard)
			currentBoard[newX][newY] = -1
			player *= -1 //invert player
		} else {
			//
			// 🤖computer plays
			//

			fmt.Println("🤖 Computer turn (○):")
			// list possible moves by ranking
			count := 8
			moves := currentBoard.rankMoves(1, &count)
			//get best ranked coordinate
			var maxRank float32 = -500000
			var bestCoord Coordinate
			for coord, rank := range moves {
				if rank > maxRank {
					maxRank = rank
					bestCoord = coord
				}
			}
			newX, newY = bestCoord[0], bestCoord[1]
			currentBoard[newX][newY] = 1
			fmt.Printf("Ranks: %v\n", moves)
			fmt.Printf("Computer has simulated %v moves\n", count)
			player *= -1 //invert player
		}
		currentBoard.render(Coordinate{newX, newY})
	}

	switch currentBoard.victory() {
	case 1:
		fmt.Println("✗ wins!")
	case 2:
		fmt.Println("○ wins!")
	case 3:
		fmt.Println("Draw!")
	}
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
		if (boardX >= 1 && boardX <= 3) && (boardY >= 1 && boardY <= 3) { // check if user input coordinates are on a valid range
			if b[boardX-1][boardY-1] == 0 { // check if user input coordinates are on an empy cell
				break
			} else {
				fmt.Println("Coordinate is not empty")
			}
		} else {
			fmt.Println("Coordinates are not on a valid range (1,2 or 3 for each axis)")
		}
	}
	return boardX - 1, boardY - 1
}

//draw a small tictactoe board on terminal
func (b *Board) render(lastMove Coordinate) {
	const colorRed = "\033[0;31m"
	const colorBlack = "\033[0;30m"
	const colorNone = "\033[0m"

	fmt.Println()
	for i := 0; i < 3; i++ { //iterate rows
		fmt.Print(" ")
		for j := 0; j < 3; j++ { // iterate columns
			switch b[i][j] {
			case 0:
				fmt.Fprintf(os.Stdout, "%s%s%s", colorBlack, "■ ", colorNone)
			case -1:
				if lastMove == [2]int{i, j} {
					fmt.Fprintf(os.Stdout, "%s%s%s", colorRed, "✗ ", colorNone)
				} else {
					fmt.Print("✗ ")
				}
			case 1:
				if lastMove == [2]int{i, j} {
					fmt.Fprintf(os.Stdout, "%s%s%s", colorRed, "○ ", colorNone)
				} else {
					fmt.Print("○ ")
				}
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
	pos2 := 0
	neg2 := 0
	winCombinations := [8][3]Coordinate{
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
			score = 1000.0
			return score
		case -3:
			score = -1000.0
			return score
		case 2:
			pos2++
			score += 20
		case -2:
			neg2++
			score -= 20
		}
	}
	if pos2 >= 2 {
		return 300
	}
	if neg2 >= 2 {
		return -300
	}
	return score
}

//find coordinates with 0 values
func (b Board) availableMoves() []Coordinate {
	var result []Coordinate
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
func (b Board) victory() int {
	currentScore := b.calculateScore()
	switch {
	case currentScore >= 1000:
		return 2
	case currentScore <= -1000:
		return 1
	default:
		if len(b.availableMoves()) == 0 {
			return 3
		} else {
			return 0
		}
	}
}

//simulate next moves and create a ranking. r represents how many moves ahead to consider.
func (b Board) rankMoves(r int, counter *int) map[Coordinate]float32 {
	available := b.availableMoves()
	moveScores := make(map[Coordinate]float32)

	for _, move := range available {
		var b2 Board = b
		var score float32 = 0.0

		// simulate player or user moves
		if r%2 == 0 {
			b2[move[0]][move[1]] = -1
		} else {
			b2[move[0]][move[1]] = 1
		}
		score = b2.calculateScore()

		if r < maxRecursionLevel && b2.victory() == 0 { //iterate trough next moves
			*counter++ //count how many times the function executed
			moveScores2 := b2.rankMoves(r+1, counter)
			//sum scores to evaluate how balanced is the scenario
			var sum float32
			for _, s := range moveScores2 {
				sum += s
			}
			averageScores2 := sum
			score = (score + averageScores2*0.9) / 1.9
		}
		moveScores[move] = score
	}

	return moveScores
}
