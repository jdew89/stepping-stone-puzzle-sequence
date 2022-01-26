package main

import "fmt"

// Stores states of the board for iteration.
type BoardState struct {
	Board   Board // Reference to the board
	NextY   int   // Next Y position needing cheked
	NextX   int   // Next X position needing cheked
	NextNum int   // Next number to place
}

type Board [5][6]int

func main() {
	board := Board{}

	//initialize 1's on the board
	board[1][2] = 1
	board[3][4] = 1

	next_num := 2
	boardStates := make([]BoardState, 0)

	tempState := BoardState{board, 0, 0, next_num}
	boardStates = append(boardStates, tempState)
	var highestBoard *Board

	highest_num := 0
	i := 0
	for {
		//check len every loop because it changes
		if i >= len(boardStates) {
			break
		}

		returned_num, finalBoard := PlayBoardState(&boardStates[i], &boardStates)

		if returned_num > highest_num {
			highest_num = returned_num
			highestBoard = finalBoard
		}

		//fmt.Println("i:", i, " - highest_num:", highest_num, " - len(boardStates):", len(boardStates))
		i++
	}

	fmt.Println("Total Board States:", len(boardStates))
	fmt.Println("Highest Num:", highest_num)
	fmt.Println("Highest Board State")
	PrintBoard(highestBoard)
}

// Returns the last placed number and board configuration where it stopped.
func PlayBoardState(boardState *BoardState, boardStateList *[]BoardState) (int, *Board) {
	board := boardState.Board
	//fmt.Println("X:", boardState.NextX, ", Y:", boardState.NextY, "next_num:", boardState.NextNum)

	next_num := boardState.NextNum

	placed_tile := false

	//fmt.Println("STARTING STATE\nNext Num:", next_num, "(", boardState.NextY, ",", boardState.NextX, ")")
	//PrintBoard(&board)

	startingX := boardState.NextX
	startingY := boardState.NextY

	for next_num < 17 {
		for y := startingY; y < len(board) && !placed_tile; y++ {
			for x := startingX; x < len(board[y]) && !placed_tile; x++ {
				if board[y][x] == 0 {
					//skip tiles without 0
					if SumNeighbors(&board, y, x) == next_num {
						*boardStateList = append(*boardStateList, BoardState{board, y, x + 1, next_num})

						board[y][x] = next_num
						next_num += 1
						placed_tile = true
					}
				}
			}
			//reset the x to 0 or it keeps searching from the wrong spot
			startingX = 0
		}

		if !placed_tile {
			break
		}
		placed_tile = false
		//reset the y to 0 or it keeps searching from the wrong spot
		startingY = 0
	}
	//fmt.Println("END\nNext Num:", next_num)
	//PrintBoard(&board)
	return next_num - 1, &board

}

// Calcs the sum of a position on the board.
// Looks at each of it's eight neighbors and sums them.
func SumNeighbors(board *Board, y int, x int) int {
	sum := 0

	if y < len(board)-1 && x > 0 {
		sum += board[y+1][x-1]
	}
	if y < len(board)-1 {
		sum += board[y+1][x]
	}
	if y < len(board)-1 && x < len(board[y])-1 {
		sum += board[y+1][x+1]
	}
	if x > 0 {
		sum += board[y][x-1]
	}
	if x < len(board[y])-1 {
		sum += board[y][x+1]
	}
	if y > 0 && x > 0 {
		sum += board[y-1][x-1]
	}
	if y > 0 {

		sum += board[y-1][x]
	}
	if y > 0 && x < len(board[y])-1 {
		sum += board[y-1][x+1]
	}
	//fmt.Println("x:", x, ", y:", y, " : sum:", sum)

	return sum
}

// Prints 1 row of the board on each line
func PrintBoard(board *Board) {
	for y := 0; y < len(board); y++ {
		fmt.Println(board[y])
	}
}
