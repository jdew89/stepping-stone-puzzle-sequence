package main

import "fmt"

type BoardState struct {
	Board   Board
	Y       int
	X       int
	NextNum int
}

type Board [5][6]int

func main() {
	board := Board{}

	board[1][2] = 1
	board[3][4] = 1

	next_num := 2
	boardStates := make([]BoardState, 0)

	//board[2][3] = 2
	//fmt.Println(SumNeighbors(&board, 2, 2))
	//next_num = 3

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

	//PlayBoardState(&boardStates[1], &boardStates)

	fmt.Println("====")
	fmt.Println("Total Board States:", len(boardStates))
	fmt.Println("Highest Num:", highest_num)
	fmt.Println("Highest Board State")
	PrintBoard(highestBoard)
	//PrintBoard(&board)
	//PrintBoard(&boardStates[2].Board)
}

//returns last placed num
func PlayBoardState(boardState *BoardState, boardStateList *[]BoardState) (int, *Board) {
	board := boardState.Board
	//fmt.Println("X:", boardState.X, ", Y:", boardState.Y, "next_num:", boardState.NextNum)

	next_num := boardState.NextNum

	placed_tile := false

	//fmt.Println("STARTING STATE\nNext Num:", next_num, "(", boardState.Y, ",", boardState.X, ")")
	//PrintBoard(&board)

	for next_num < 17 {
		for y := boardState.Y; y < len(board) && !placed_tile; y++ {
			for x := boardState.X; x < len(board[y]) && !placed_tile; x++ {
				if board[y][x] == 0 {
					//skip tiles without 0
					if SumNeighbors(&board, y, x) == next_num {

						tempState := BoardState{board, y, x + 1, next_num}
						*boardStateList = append(*boardStateList, tempState)

						board[y][x] = next_num
						//fmt.Println("placed!", next_num)
						next_num += 1
						placed_tile = true
					}
				}
			}
			boardState.X = 0
		}

		if !placed_tile {
			break
		}
		placed_tile = false
		//reset the x and y for state or it keep searching from the wrong stop
		boardState.Y = 0
		//fmt.Println(next_num)
	}
	//fmt.Println("END\nNext Num:", next_num)
	//PrintBoard(&board)
	return next_num - 1, &board

}

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

func PrintBoard(board *Board) {
	for y := 0; y < len(board); y++ {
		fmt.Println(board[y])
	}
}
