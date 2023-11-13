package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type RaceRepository struct {
	store *Store
}

// Create ...
func (r *RaceRepository) Create(s *model.Race) error {
	if err := s.Validate(); err != nil {
		return err
	}

	return r.store.Create(
		InsertQ+RacesT+RacesP+"values ($1, $2) RETURNING id", s.Name, s.Model,
	).Scan(&s.ID)
}

// Find ...
func (r *RaceRepository) Find(id int) (*model.Race, error) {
	race := &model.Race{}

	if err := r.store.SelectRow(
		SelectQ+RacesT+"WHERE id = $1", id,
	).Scan(
		&race.ID,
		&race.Name,
		&race.Model,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return race, nil
}

// Update ...
func (r *RaceRepository) Update(s *model.Race) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+RacesT+"SET name = $1, model = $2 WHERE id = $3", s.Name, s.Model, s.ID,
	)

	return err
}

// Delete ...
func (r *RaceRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+RacesT+"WHERE id = $1", id)
	return err
}
