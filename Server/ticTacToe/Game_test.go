package ticTacToe

import (
	"testing"
)

func TestTakeTurnThrowsValidationExceptionOnGameEnd(t *testing.T) {
	game := NewGame()

	ShouldNotErrorWhenTakeTurn(&game, t, X, 0, 0)
	ShouldNotErrorWhenTakeTurn(&game, t, O, 1, 0)
	ShouldNotErrorWhenTakeTurn(&game, t, X, 2, 0)
	ShouldNotErrorWhenTakeTurn(&game, t, O, 2, 1)
	ShouldNotErrorWhenTakeTurn(&game, t, X, 1, 2)
	ShouldNotErrorWhenTakeTurn(&game, t, O, 0, 2)
	ShouldNotErrorWhenTakeTurn(&game, t, X, 1, 1)

	err := game.TakeTurn(O, 0, 1)
	if err != nil {
		t.Fatalf("8th call should have succeeded but didn't, %s", err.Error())
	}

	if !game.Ended {
		t.Fatalf("game should have ended but didn't")
	}

	err = game.TakeTurn(O, 2, 2)
	if err == nil {
		t.Fatalf("9th call should have thrown error but didn't")
	}
}

func TestTakeTurnPositionsOutOfRangeThrowsErr(t *testing.T) {
	game := NewGame()

	ShouldErrorWhenTakeTurn(game, t, X, 4, 0)
	ShouldErrorWhenTakeTurn(game, t, X, -1, 0)
	ShouldErrorWhenTakeTurn(game, t, X, 0, -1)
	ShouldErrorWhenTakeTurn(game, t, X, 0, 4)
}

func TestTakeTurnOutOfPlaceThrowsErr(t *testing.T) {
	game := NewGame()

	ShouldErrorWhenTakeTurn(game, t, O, 1, 0)

	// test that X also fails when it's O's turn
	ShouldNotErrorWhenTakeTurn(&game, t, X, 0, 0)
	ShouldErrorWhenTakeTurn(game, t, X, 2, 1)
}

func TestPlaceTwoTokensInSameSpot(t *testing.T) {
	game := NewGame()
	game.TakeTurn(X, 0, 0)

	err := game.TakeTurn(O, 0, 0)
	if err == nil {
		t.Fatalf("should not have errored but did")
	}
}

func TestTakeTurnUpdatesGameState(t *testing.T) {
	game := NewGame()
	game.TakeTurn(X, 0, 0)
	if game.Board[0][0] != X {
		t.Fatalf("game board not updated correctly")
	}
}

func TestGameEndsWithWinner(t *testing.T) {
	game := NewGame()
	ShouldNotErrorWhenTakeTurn(&game, t, X, 0, 0)
	ShouldNotErrorWhenTakeTurn(&game, t, O, 0, 1)
	ShouldNotErrorWhenTakeTurn(&game, t, X, 1, 0)
	ShouldNotErrorWhenTakeTurn(&game, t, O, 0, 2)
	ShouldNotErrorWhenTakeTurn(&game, t, X, 2, 0)
	// x x x
	// 0 n n
	// 0 n n

	if game.Winner != X {
		t.Fatalf("game winner should have been X but wasn't")
	}
	if !game.Ended {
		t.Fatalf("game should have ended but didnt")
	}

	ShouldErrorWhenTakeTurn(game, t, O, 2, 2)
}

func ShouldErrorWhenTakeTurn(game Game, t *testing.T, tile PlayerTile, x int8, y int8) {
	err := game.TakeTurn(tile, x, y)

	if err == nil {
		t.Fatalf("should have not succeed with (%d, %d) coords, but did", x, y)
	}
}

func ShouldNotErrorWhenTakeTurn(game *Game, t *testing.T, tile PlayerTile, x int8, y int8) {
	err := game.TakeTurn(tile, x, y)

	if err != nil {
		t.Fatalf("should have succeeded with (%d, %d) coords, but didnt, %s", x, y, err.Error())
	}
}
