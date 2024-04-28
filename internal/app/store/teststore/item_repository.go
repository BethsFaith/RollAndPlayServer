package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type ItemRepository struct {
	store       *Store
	items       map[int]*model.Item
	types       map[int]*model.ItemType
	descriptors map[int]*model.ItemDescriptor
}

// Create ...
func (r *ItemRepository) Create(i *model.Item) error {
	if err := i.Validate(); err != nil {
		return err
	}

	i.ID = len(r.items) + 1

	r.items[i.ID] = i

	return nil
}

// Get ...
func (r *ItemRepository) Get() ([]*model.Item, error) {
	var items []*model.Item

	for i := range r.items {
		items = append(items, r.items[i])
	}

	return items, nil
}

// Find ...
func (r *ItemRepository) Find(id int) (*model.Item, error) {
	c, ok := r.items[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return c, nil
}

// Update ...
func (r *ItemRepository) Update(c *model.Item) error {
	if err := c.Validate(); err != nil {
		return err
	}

	source, ok := r.items[c.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Name = c.Name
	source.Icon = c.Icon
	source.TypeId = c.TypeId
	source.UserId = c.UserId

	return nil
}

// Delete ...
func (r *ItemRepository) Delete(id int) error {
	_, ok := r.items[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.items, id)

	return nil
}

func (r *ItemRepository) CreateType(itemType *model.ItemType) error {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) GetTypes() ([]*model.ItemType, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) FindType(i int) (*model.ItemType, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) UpdateType(itemType *model.ItemType) error {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) DeleteType(i int) error {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) CreateDescriptor(descriptor *model.ItemDescriptor) error {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) GetDescriptors() ([]*model.ItemDescriptor, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) FindDescriptor(i int) (*model.ItemDescriptor, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) UpdateDescriptor(descriptor *model.ItemDescriptor) error {
	//TODO implement me
	panic("implement me")
}

func (r *ItemRepository) DeleteDescriptor(i int) error {
	//TODO implement me
	panic("implement me")
}
