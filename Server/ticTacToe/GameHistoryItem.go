package ticTacToe

type GameHistoryItem struct {
	Turn     uint8
	Tile     PlayerTile
	Position [2]uint8
}
