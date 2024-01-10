//Commandline Tic-Tac-Toe game
//Author: Austin Mais
//Date: 1/9/2024
//Version: 1.0
//It is a two player game where each player takes turns placing their
//respective symbol on the board. The first player to get three in a row wins.
//The game will end in a draw if all spaces are filled and no player has three
//in a row.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	name string
	xOrO int
}

type movePlacement struct {
	row    int
	column int
}

func printBoard(board [][]int) {
	for _, h := range board {
		for _, i := range h {
			if i == -1 {
				fmt.Printf("|___")
			} else if i == 1 {
				fmt.Printf("|_X_")
			} else {
				fmt.Printf("|_O_")
			}
		}
		fmt.Println("|")
	}
}

func move(p player, m movePlacement, b *[][]int) error {
    //Check for valid move
    if m.row > 2 || m.row < 0 || m.column > 2 || m.column < 0 {
        return fmt.Errorf("Invalid move. Please try again.")
    }
    //Check for space already taken
	if (*b)[m.row][m.column] != -1 {
		return fmt.Errorf("Space already taken. Please try again.")
	}
    //Place move
	(*b)[m.row][m.column] = p.xOrO
	return nil
}

func checkWin(b [][]int) bool {
	//Check for horizontal win
	for _, h := range b {
		if h[0] == h[1] && h[1] == h[2] && h[0] != -1 {
			return true
		}
	}
	//Check for vertical win
	for i := 0; i < 3; i++ {
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] && b[0][i] != -1 {
			return true
		}
	}
	//Check for diagonal win
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[0][0] != -1 {
		return true
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[0][2] != -1 {
		return true
	}
	return false
}

func main() {
    //Variable declarations
	reader := bufio.NewReader(os.Stdin)
	board := [][]int{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	p1 := player{"", 1}
	p2 := player{"", 0}
	moveCount := 1
	input := ""

    //Game setup
	printBoard(board)
	fmt.Println("Welcome to Tic-Tac-Toe")
	fmt.Println("Player 1, enter your name:")
	p1.name, _ = reader.ReadString('\n')
	fmt.Println("Player 2, enter your name:")
	p2.name, _ = reader.ReadString('\n')
	fmt.Println("Player X is " + p1.name + " and Player O is " + p2.name)

    //Game loop
	for !checkWin(board) {
		printBoard(board)
		fmt.Println("Select row: ")
		input, _ = reader.ReadString('\n')
		inputRow, _ := strconv.Atoi(strings.TrimSpace(input))
		fmt.Println("Select column:")
		input, _ = reader.ReadString('\n')
		inputCol, _ := strconv.Atoi(strings.TrimSpace(input))
		currentMove := movePlacement{inputRow - 1, inputCol - 1}
		if moveCount%2 == 0 {
			err := move(p2, currentMove, &board)
			if err != nil {
                fmt.Println()
				fmt.Println(err)
                fmt.Println()
				continue
			}
		} else {
			err := move(p1, currentMove, &board)
			if err != nil {
                fmt.Println()
				fmt.Println(err)
                fmt.Println()
				continue
			}
		}
		if moveCount == 9 {
			break
		}
		moveCount++
	}

    //Game Completion
    fmt.Println()
    fmt.Println("***GAME OVER***")
    fmt.Println()
	if moveCount == 9 && !checkWin(board) {
		fmt.Println("Draw!")
	} else if moveCount%2 == 0 {
		fmt.Println(strings.TrimSpace(p1.name) + " WINS!")
	} else {
        fmt.Println(strings.TrimSpace(p2.name) + " WINS!")
	}
    fmt.Println()
    fmt.Println("Final Board:")
	printBoard(board)
}
