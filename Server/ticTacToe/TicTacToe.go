package ticTacToe

type PlayerTile byte

const (
	None PlayerTile = 0
	X    PlayerTile = 1
	O    PlayerTile = 2
)

func (e PlayerTile) String() string {
	switch e {
	case None:
		return ""
	case X:
		return "X"
	case O:
		return "O"
	}
	return "err"
}
