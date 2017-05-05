package systems

import (
// "log"
// "math/rand"
// "engo.io/ecs"
// "engo.io/engo"
// "engo.io/engo/common"
)

// game logic
// =====================================
type player int
type gameState int

const (
	boardSize int = 3

	noone   player = 0
	player1 player = 1
	player2 player = 2

	inProgress gameState = 1
	finished   gameState = 2
)

type Game struct {
	Board         [boardSize][boardSize]player
	CurrentPlayer player
	State         gameState
	Winner        player
	currentTurn   int
}

func NewGame() Game {
	game := Game{Winner: noone, State: inProgress, currentTurn: 0}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			game.Board[i][j] = noone
		}
	}
	// firstTurn := rand.Intn(2) + 1
	// game.CurrentPlayer = player(firstTurn)
	game.CurrentPlayer = player1
	return game
}

func (game *Game) MakeMove(x, y int) (err bool) {
	// check for invalid params
	if x > boardSize-1 ||
		y > boardSize-1 ||
		game.Board[x][y] != noone ||
		game.State != inProgress {
		return true
	}
	game.Board[x][y] = game.CurrentPlayer

	// player won
	if isPlayerWon(game, x, y) {
		game.State = finished
		game.Winner = game.CurrentPlayer
	} else if isLastTurn(game) {
		// last turn, game draw
		game.State = finished
	} else {
		// next turn
		game.currentTurn++
		game.CurrentPlayer = nextPlayer(game.CurrentPlayer)
	}
	return false
}

func isPlayerWon(game *Game, x, y int) bool {
	var col, row, diag, rdiag int

	for i := 0; i < boardSize; i++ {
		if game.Board[x][i] == game.CurrentPlayer {
			col++
		}
		if game.Board[i][y] == game.CurrentPlayer {
			row++
		}
		if game.Board[i][i] == game.CurrentPlayer {
			diag++
		}
		if game.Board[i][boardSize-i-1] == game.CurrentPlayer {
			rdiag++
		}
	}
	return col == boardSize || row == boardSize || diag == boardSize || rdiag == boardSize
}

func nextPlayer(current player) player {
	if current == player1 {
		return player2
	} else {
		return player1
	}
}

func isLastTurn(game *Game) bool {
	return game.currentTurn == boardSize*boardSize-1
}
