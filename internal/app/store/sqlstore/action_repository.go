package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type ActionRepository struct {
	store *Store
}

// Create ...
func (r *ActionRepository) Create(a *model.Action) error {
	if err := a.Validate(); err != nil {
		return err
	}
	if err := a.BeforeInsertOrUpdate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+ActionsT+ActionsP+"values ($1, $2, $3, $4, $5) RETURNING id",
		a.Name, a.Icon, a.RefSkillId, a.Points, a.UserId,
	).Scan(&a.ID)
}

// Get ...
func (r *ActionRepository) Get() ([]*model.Action, error) {
	var actions []*model.Action

	bRows, err := r.store.SelectRows(
		SelectQ + ActionsT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		a := &model.Action{}

		err := bRows.Scan(&a.ID, &a.Name, &a.Icon, &a.RefSkillId, &a.Points, &a.UserId)
		if err != nil {
			return nil, err
		}

		a.AfterScan()

		actions = append(actions, a)
	}

	return actions, nil
}

// Find ...
func (r *ActionRepository) Find(id int) (*model.Action, error) {
	a := &model.Action{}

	if err := r.store.SelectRow(
		SelectQ+ActionsT+"WHERE id = $1", id,
	).Scan(
		&a.ID,
		&a.Name,
		&a.Icon,
		&a.RefSkillId,
		&a.Points,
		&a.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	a.AfterScan()

	return a, nil
}

// Update ...
func (r *ActionRepository) Update(a *model.Action) error {
	if err := a.Validate(); err != nil {
		return err
	}
	if err := a.BeforeInsertOrUpdate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+ActionsT+"SET name = $1, icon = $2, skill_id = $3, points = $4, user_id = $5 WHERE id = $6",
		a.Name, a.Icon, a.RefSkillId, a.Points, a.UserId, a.ID,
	)

	return err
}

// Delete ...
func (r *ActionRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+ActionsT+"WHERE id = $1", id)
	return err
}
