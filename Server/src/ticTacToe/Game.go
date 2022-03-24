package ticTacToe

import "fmt"

type GameBoard [3][3]PlayerTile

type Game struct {
	Board             GameBoard
	CurrentPlayerTurn PlayerTile
	TurnNumber        uint8
	History           GameHistory
	Winner            PlayerTile
	Ended             bool
}

func NewGame() Game {
	return Game{
		CurrentPlayerTurn: X,
	}
}

func CalculateWinner(board GameBoard) PlayerTile {
	for i := 0; i < 3; i++ {
		if board[i][0]&board[i][1]&board[i][2] != None {
			return board[0][0]
		}
		if board[0][i]&board[1][i]&board[2][i] != None {
			return board[0][0]
		}
	}
	if board[0][0]&board[1][1]&board[2][2] != None {
		return board[0][0]
	}
	if board[2][0]&board[1][1]&board[0][2] != None {
		return board[0][0]
	}
	return None
}

func (game *Game) TakeTurn(tile PlayerTile, x int8, y int8) GameError {
	if game.Winner != None {
		return NewValidationGameErr("the game is already over")
	}
	if game.TurnNumber == 8 {
		return NewValidationGameErr("the game is already over")
	}
	if int(y) >= len(game.Board) {
		return NewValidationGameErr("y value too high")
	}
	if int(y) < 0 {
		return NewValidationGameErr("y value must be greater than 0")
	}
	if int(x) < 0 {
		return NewValidationGameErr("x value must be greater than 0")
	}
	if int(x) >= len(game.Board[0]) {
		return NewValidationGameErr("x value too high")
	}
	if game.Board[y][x] != None {
		return NewValidationGameErr("this space already has a tile on it")
	}
	if game.CurrentPlayerTurn != tile {
		return NewValidationGameErr(fmt.Sprintf("it is currently %v's turn", PlayerTile(game.CurrentPlayerTurn)))
	}
	game.Board[y][x] = tile
	game.History.History[game.TurnNumber] = GameHistoryItem{
		Turn:     game.TurnNumber,
		Tile:     tile,
		Position: [2]uint8{uint8(x), uint8(y)},
	}
	game.TurnNumber++
	if game.CurrentPlayerTurn == X {
		game.CurrentPlayerTurn = O
	} else {
		game.CurrentPlayerTurn = X
	}
	game.Winner = CalculateWinner(game.Board)
	game.Ended = game.Winner != None || game.TurnNumber == 8

	return nil
}

type GameError interface {
	ValidationError() bool
	Error() string
}

type GameErr struct {
	IsValidationError bool
	ErrorMessage      string
}

func (error GameErr) ValidationError() bool {
	return error.IsValidationError
}

func (error GameErr) Error() string {
	return error.ErrorMessage
}

func NewInternalGameErr(errString string) *GameErr {
	return &GameErr{
		ErrorMessage:      errString,
		IsValidationError: false,
	}
}
func NewValidationGameErr(errString string) *GameErr {
	return &GameErr{
		ErrorMessage:      errString,
		IsValidationError: true,
	}
}
