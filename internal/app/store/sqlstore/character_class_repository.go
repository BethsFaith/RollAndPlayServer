package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type CharacterClassRepository struct {
	store *Store
}

// Create ...
func (r *CharacterClassRepository) Create(cc *model.CharacterClass) error {
	if err := cc.Validate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+CharacterClassT+CharacterClassesP+"values ($1, $2, $3) RETURNING id",
		cc.Name, cc.Icon, cc.UserId,
	).Scan(&cc.ID)
}

// Get ...
func (r *CharacterClassRepository) Get() ([]*model.CharacterClass, error) {
	var classes []*model.CharacterClass

	bRows, err := r.store.SelectRows(
		SelectQ + CharacterClassT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		cc := &model.CharacterClass{}

		err := bRows.Scan(&cc.ID, &cc.Name, &cc.Icon, &cc.UserId)
		if err != nil {
			return nil, err
		}

		classes = append(classes, cc)
	}

	return classes, nil
}

// Find ...
func (r *CharacterClassRepository) Find(id int) (*model.CharacterClass, error) {
	cc := &model.CharacterClass{}

	if err := r.store.SelectRow(
		SelectQ+CharacterClassT+"WHERE id = $1", id,
	).Scan(
		&cc.ID,
		&cc.Name,
		&cc.Icon,
		&cc.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return cc, nil
}

// Update ...
func (r *CharacterClassRepository) Update(cc *model.CharacterClass) error {
	if err := cc.Validate(); err != nil {
		return err
	}
	_, err := r.store.Update(
		UpdateQ+CharacterClassT+"SET name = $1, icon = $2, user_id = $3 WHERE id = $4", cc.Name, cc.Icon,
		cc.UserId, cc.ID,
	)

	return err
}

// Delete ...
func (r *CharacterClassRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+CharacterClassT+"WHERE id = $1", id)
	return err
}
