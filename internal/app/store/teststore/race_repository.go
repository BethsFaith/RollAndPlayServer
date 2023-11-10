package teststore

import "RnpServer/internal/app/model"

type RaceRepository struct {
	store *Store
	races map[int]*model.Race
}
