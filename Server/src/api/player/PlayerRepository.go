package player

type playerRepository interface {
	GetPlayerById(id string)
	SavePlayer(Player string)
}
