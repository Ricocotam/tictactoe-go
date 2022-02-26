package tictactoe

import (
	"errors"
	"fmt"
)

type Grid [3][3]string

type Pos struct {
	x, y    int
	isValid bool
}

type Player string

type Game struct {
	board         Grid
	Players       [2]Player
	winner        int
	CurrentPlayer int
	signs         [2]string
	finished      bool
	err           error
}

func Play(move Pos, game *Game, player Player) {
	checkMove(&move)
	game.Play(&move)
}

func checkMove(move *Pos) {
	if move.x < 0 || move.y < 0 {
		move.isValid = false
	}
	if move.x > 3 || move.y > 3 {
		move.isValid = false
	}
}

func NewPos(x, y int) *Pos {
	return &Pos{x, y, true}
}

func NewGame(p1, p2 Player) *Game {
	var game Game
	game.Players = [2]Player{p1, p2}
	game.winner = -1
	game.signs = [2]string{"x", "o"}

	return &game
}

func (game *Game) IsFinished() bool {
	return game.finished
}

func (game *Game) Play(move *Pos) {
	if !move.isValid {
		game.err = errors.New("proposed move is invalid")
		return
	}

	game.board.Play(move, game.signs[game.CurrentPlayer])
	game.next_player()

	game.Finished()
}

func (game *Game) Finished() {
	won := game.board.Winner()

	if won {
		game.finished = true
		game.winner = game.CurrentPlayer
		return
	} else {
		game.finished = game.board.Finished()
	}
}

func (board *Grid) Play(move *Pos, sign string) {
	board[move.x][move.y] = sign
}

func (game *Game) next_player() {
	game.CurrentPlayer = (game.CurrentPlayer + 1) % 2
}

func (board *Grid) Finished() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func (board *Grid) Winner() (won bool) {
	won = board.checkRows()
	won = won || board.checkColumns()
	won = won || board.checkDiagonals()

	return won
}

func checkLines(plays []string) (won bool) {
	won = false
	for _, play := range plays {
		won = won || (len(play) == 3 && allEqual(play[:]))
		if won {
			return won
		}
	}
	return won
}

func allEqual(arr string) bool {
	if len(arr) <= 1 {
		return true
	}

	first := rune(arr[0])
	for _, e := range arr[1:] {
		if e != first {
			return false
		}
	}
	return true
}

func (board *Grid) checkRows() (won bool) {
	var rows [3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rows[i] += board[i][j]
		}
	}
	won = checkLines(rows[:])
	return won
}

func (board *Grid) checkColumns() (won bool) {
	var cols [3]string
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			cols[j] += board[i][j]
		}
	}
	won = checkLines(cols[:])
	return won
}

func (board *Grid) checkDiagonals() (won bool) {
	var diags [2]string
	for i := 0; i < 3; i++ {
		diags[0] += board[i][i]
	}
	for i := 2; i >= 0; i-- {
		diags[1] += board[2-i][i]
	}
	won = checkLines(diags[:])
	return won
}

func (game *Game) String() string {
	return fmt.Sprintf("%s", game.board)
}
