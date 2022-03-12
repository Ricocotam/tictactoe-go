package main

import (
	"fmt"
	"tictactoe/tictactoe"
)

func main() {
	fmt.Println("Hello")

	p1 := tictactoe.Player("Toto")
	p2 := tictactoe.Player("Tata")
	game := tictactoe.NewGame(p1, p2)
	var x, y int
	var move *tictactoe.Pos

	for !game.IsFinished() {
		fmt.Println("Next move for Player", game.Players[game.CurrentPlayer])
		fmt.Scanf("%d %d", &x, &y)
		move = tictactoe.NewPos(x, y)
		game.Play(move)
		if game.Err != nil {
			fmt.Println(game.Err)
		}
		fmt.Println(game)
	}

}
