package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type CharacteristicRepository struct {
	store *Store
}

// Create ...
func (r *CharacteristicRepository) Create(a *model.Characteristic) error {
	if err := a.Validate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+CharacteristicsT+CharacteristicsP+"values ($1, $2, $3) RETURNING id",
		a.Name, a.Icon, a.UserId,
	).Scan(&a.ID)
}

// Get ...
func (r *CharacteristicRepository) Get() ([]*model.Characteristic, error) {
	var characteristics []*model.Characteristic

	bRows, err := r.store.SelectRows(
		SelectQ + CharacteristicsT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		a := &model.Characteristic{}

		err := bRows.Scan(&a.ID, &a.Name, &a.Icon, &a.UserId)
		if err != nil {
			return nil, err
		}

		characteristics = append(characteristics, a)
	}

	return characteristics, nil
}

// Find ...
func (r *CharacteristicRepository) Find(id int) (*model.Characteristic, error) {
	c := &model.Characteristic{}

	if err := r.store.SelectRow(
		SelectQ+CharacteristicsT+"WHERE id = $1", id,
	).Scan(
		&c.ID,
		&c.Name,
		&c.Icon,
		&c.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return c, nil
}

// Update ...
func (r *CharacteristicRepository) Update(a *model.Characteristic) error {
	if err := a.Validate(); err != nil {
		return err
	}
	_, err := r.store.Update(
		UpdateQ+CharacteristicsT+"SET name = $1, icon = $2, user_id = $3 WHERE id = $4",
		a.Name, a.Icon, a.UserId, a.ID,
	)

	return err
}

// Delete ...
func (r *CharacteristicRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+CharacteristicsT+"WHERE id = $1", id)
	return err
}
