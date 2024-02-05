package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type CharacterClassRepository struct {
	store   *Store
	classes map[int]*model.CharacterClass
}

// Create ...
func (r *CharacterClassRepository) Create(cc *model.CharacterClass) error {
	if err := cc.Validate(); err != nil {
		return err
	}

	cc.ID = len(r.classes) + 1

	r.classes[cc.ID] = cc

	return nil
}

// Find ...
func (r *CharacterClassRepository) Find(id int) (*model.CharacterClass, error) {
	class, ok := r.classes[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return class, nil
}

// Update ...
func (r *CharacterClassRepository) Update(cc *model.CharacterClass) error {
	if err := cc.Validate(); err != nil {
		return err
	}

	source, ok := r.classes[cc.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Name = cc.Name
	source.Icon = cc.Icon
	source.UserId = cc.UserId

	return nil
}

// Delete ...
func (r *CharacterClassRepository) Delete(id int) error {
	_, ok := r.classes[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.classes, id)

	return nil
}
