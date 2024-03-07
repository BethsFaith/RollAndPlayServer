package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type ItemRepository struct {
	store *Store
}

// Create ...
func (r *ItemRepository) Create(i *model.Item) error {
	if err := i.Validate(); err != nil {
		return err
	}

	err := i.BeforeInsertOrUpdate()
	if err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+ItemsT+ItemsP+"values ($1, $2, $3, $4, $5, $6) RETURNING id",
		i.Name, i.Description, i.Icon, i.Count, i.RefTypeId, i.UserId,
	).Scan(&i.ID)
}

// Get ...
func (r *ItemRepository) Get() ([]*model.Item, error) {
	var items []*model.Item

	bRows, err := r.store.SelectRows(
		SelectQ + ItemsT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		i := &model.Item{}

		err := bRows.Scan(&i.ID, &i.Name, &i.Description, &i.Icon, &i.Count,
			&i.RefTypeId, &i.UserId)
		if err != nil {
			return nil, err
		}

		i.AfterScan()

		items = append(items, i)
	}

	return items, nil
}

// Find ...
func (r *ItemRepository) Find(id int) (*model.Item, error) {
	i := &model.Item{}

	if err := r.store.SelectRow(
		SelectQ+ItemsT+"WHERE id = $1", id,
	).Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Icon,
		&i.Count,
		&i.RefTypeId,
		&i.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	i.AfterScan()

	return i, nil
}

// Update ...
func (r *ItemRepository) Update(i *model.Item) error {
	if err := i.Validate(); err != nil {
		return err
	}

	err := i.BeforeInsertOrUpdate()
	if err != nil {
		return err
	}

	_, err = r.store.Update(
		UpdateQ+ItemsT+"SET name = $1, description = $2, icon = $3, "+
			"count = $4, type_id = $5, user_id = $6 WHERE id = $7",
		i.Name, i.Description, i.Icon, i.Count, i.RefTypeId, i.UserId, i.ID,
	)

	return err
}

// Delete ...
func (r *ItemRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+ItemsT+"WHERE id = $1", id)
	return err
}

// CreateType ...
func (r *ItemRepository) CreateType(i *model.ItemType) error {
	if err := i.Validate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+ItemTypesT+ItemTypesP+"values ($1, $2, $3) RETURNING id",
		i.Name, i.Icon, i.UserId,
	).Scan(&i.ID)
}

// GetTypes ...
func (r *ItemRepository) GetTypes() ([]*model.ItemType, error) {
	var types []*model.ItemType

	bRows, err := r.store.SelectRows(
		SelectQ + ItemTypesT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		i := &model.ItemType{}

		err := bRows.Scan(&i.ID, &i.Name, &i.Icon, &i.UserId)
		if err != nil {
			return nil, err
		}

		types = append(types, i)
	}

	return types, nil
}

// FindType ...
func (r *ItemRepository) FindType(id int) (*model.ItemType, error) {
	i := &model.ItemType{}

	if err := r.store.SelectRow(
		SelectQ+ItemTypesT+"WHERE id = $1", id,
	).Scan(
		&i.ID,
		&i.Name,
		&i.Icon,
		&i.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return i, nil
}

// UpdateType ...
func (r *ItemRepository) UpdateType(i *model.ItemType) error {
	if err := i.Validate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+ItemTypesT+"SET name = $1, icon = $2, user_id = $3 WHERE id = $4",
		i.Name, i.Icon, i.UserId, i.ID,
	)

	return err
}

// DeleteType ...
func (r *ItemRepository) DeleteType(id int) error {
	_, err := r.store.Delete(DeleteQ+ItemTypesT+"WHERE id = $1", id)
	return err
}

// CreateDescriptor ...
func (r *ItemRepository) CreateDescriptor(i *model.ItemDescriptor) error {
	if err := i.Validate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+ItemDescriptorsT+ItemDescriptorsP+"values ($1, $2) RETURNING id",
		i.Name, i.TypeId,
	).Scan(&i.ID)
}

// GetDescriptors ...
func (r *ItemRepository) GetDescriptors() ([]*model.ItemDescriptor, error) {
	var descriptors []*model.ItemDescriptor

	bRows, err := r.store.SelectRows(
		SelectQ + ItemDescriptorsT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		i := &model.ItemDescriptor{}

		err := bRows.Scan(&i.ID, &i.Name, &i.TypeId)
		if err != nil {
			return nil, err
		}

		descriptors = append(descriptors, i)
	}

	return descriptors, nil
}

// FindDescriptor ...
func (r *ItemRepository) FindDescriptor(id int) (*model.ItemDescriptor, error) {
	i := &model.ItemDescriptor{}

	if err := r.store.SelectRow(
		SelectQ+ItemDescriptorsT+"WHERE id = $1", id,
	).Scan(
		&i.ID,
		&i.Name,
		&i.TypeId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return i, nil
}

// UpdateDescriptor ...
func (r *ItemRepository) UpdateDescriptor(i *model.ItemDescriptor) error {
	if err := i.Validate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+ItemDescriptorsT+"SET name = $1, item_type_id = $2 WHERE id = $3",
		i.Name, i.TypeId, i.ID,
	)

	return err
}

// DeleteDescriptor ...
func (r *ItemRepository) DeleteDescriptor(id int) error {
	_, err := r.store.Delete(DeleteQ+ItemDescriptorsT+"WHERE id = $1", id)
	return err
}

//// CreateDescriptorLine ...
//func (r *ItemRepository) CreateDescriptorLine(i *model.ItemDescriptorLine) error {
//	if err := i.Validate(); err != nil {
//		return err
//	}
//
//	return r.store.CreateRetId(
//		InsertQ+ItemDescriptorLinesT+ItemDescriptorLinesP+"values ($1, $2) RETURNING id",
//		i.DescriptorId, i.ItemId,
//	).Scan(&i.ID)
//}

//// GetDescriptors ...
//func (r *ItemRepository) GetDescriptors() ([]*model.ItemDescriptor, error) {
//	var descriptors []*model.ItemDescriptor
//
//	bRows, err := r.store.SelectRows(
//		SelectQ + ItemDescriptorsT,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	for bRows.Next() {
//		i := &model.ItemDescriptor{}
//
//		err := bRows.Scan(&i.ID, &i.Name, &i.TypeId)
//		if err != nil {
//			return nil, err
//		}
//
//		descriptors = append(descriptors, i)
//	}
//
//	return descriptors, nil
//}
//
//// FindDescriptor ...
//func (r *ItemRepository) FindDescriptor(id int) (*model.ItemDescriptorLine, error) {
//	i := &model.ItemDescriptorLine{}
//
//	if err := r.store.SelectRow(
//		SelectQ+ItemDescriptorsT+"WHERE id = $1", id,
//	).Scan(
//		&i.ID,
//		&i.DescriptorId,
//		&i.TypeId,
//	); err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, store.ErrorRecordNotFound
//		}
//		return nil, err
//	}
//
//	return i, nil
//}
