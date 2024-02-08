package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type ActionRepository struct {
	store   *Store
	actions map[int]*model.Action
}

// Create ...
func (r *ActionRepository) Create(a *model.Action) error {
	if err := a.Validate(); err != nil {
		return err
	}

	a.ID = len(r.actions) + 1

	r.actions[a.ID] = a

	return nil
}

// Get ...
func (r *ActionRepository) Get() ([]*model.Action, error) {
	var actions []*model.Action

	for i := range r.actions {
		actions = append(actions, r.actions[i])
	}

	return actions, nil
}

// Find ...
func (r *ActionRepository) Find(id int) (*model.Action, error) {
	a, ok := r.actions[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return a, nil
}

// Update ...
func (r *ActionRepository) Update(a *model.Action) error {
	if err := a.Validate(); err != nil {
		return err
	}

	source, ok := r.actions[a.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Name = a.Name
	source.Icon = a.Icon
	source.Points = a.Points
	source.SkillId = a.SkillId
	source.UserId = a.UserId

	return nil
}

// Delete ...
func (r *ActionRepository) Delete(id int) error {
	_, ok := r.actions[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.actions, id)

	return nil
}
