package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type CharacteristicRepository struct {
	store           *Store
	characteristics map[int]*model.Characteristic
}

// Create ...
func (r *CharacteristicRepository) Create(a *model.Characteristic) error {
	if err := a.Validate(); err != nil {
		return err
	}

	a.ID = len(r.characteristics) + 1

	r.characteristics[a.ID] = a

	return nil
}

// Get ...
func (r *CharacteristicRepository) Get() ([]*model.Characteristic, error) {
	var characteristics []*model.Characteristic

	for i := range r.characteristics {
		characteristics = append(characteristics, r.characteristics[i])
	}

	return characteristics, nil
}

// Find ...
func (r *CharacteristicRepository) Find(id int) (*model.Characteristic, error) {
	c, ok := r.characteristics[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return c, nil
}

// Update ...
func (r *CharacteristicRepository) Update(c *model.Characteristic) error {
	if err := c.Validate(); err != nil {
		return err
	}

	source, ok := r.characteristics[c.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Name = c.Name
	source.Icon = c.Icon
	source.UserId = c.UserId

	return nil
}

// Delete ...
func (r *CharacteristicRepository) Delete(id int) error {
	_, ok := r.characteristics[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.characteristics, id)

	return nil
}
