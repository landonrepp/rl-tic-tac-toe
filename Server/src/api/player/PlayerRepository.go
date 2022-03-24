package repository

type repository interface {
	GetPlayerById(id string)
	SavePlayer(id string)
}
