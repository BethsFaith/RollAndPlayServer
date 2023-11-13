package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type RaceRepository struct {
	store *Store
	races map[int]*model.Race
}

// Create ...
func (r *RaceRepository) Create(race *model.Race) error {
	if err := race.Validate(); err != nil {
		return err
	}

	race.ID = len(r.races) + 1

	r.races[race.ID] = race

	return nil
}

// Find ...
func (r *RaceRepository) Find(id int) (*model.Race, error) {
	race, ok := r.races[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return race, nil
}

// Update ...
func (r *RaceRepository) Update(race *model.Race) error {
	if err := race.Validate(); err != nil {
		return err
	}

	source, ok := r.races[race.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Name = race.Name
	source.Model = race.Model

	return nil
}

// Delete ...
func (r *RaceRepository) Delete(id int) error {
	_, ok := r.races[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.races, id)

	return nil
}
