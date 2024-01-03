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
		InsertQ+CharacterClassT+CharacterClassesP+"values ($1, $2) RETURNING id", cc.Name, cc.Icon,
	).Scan(&cc.ID)
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
		UpdateQ+CharacterClassT+"SET name = $1, icon = $2  WHERE id = $3", cc.Name, cc.Icon,
		cc.ID,
	)

	return err
}

// Delete ...
func (r *CharacterClassRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+CharacterClassT+"WHERE id = $1", id)
	return err
}
