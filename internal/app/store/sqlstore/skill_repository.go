package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type SkillRepository struct {
	store *Store
}

// Create ...
func (r *SkillRepository) Create(s *model.Skill) error {
	if err := s.Validate(); err != nil {
		return err
	}
	if err := s.BeforeInsertOrUpdate(); err != nil {
		return err
	}

	if s.RefCategoryId.Valid {
		_, err := r.FindCategory(s.CategoryId)
		if err != nil {
			return store.ErrorNotExistRef
		}
	}

	return r.store.CreateRetId(
		InsertQ+SkillsT+SkillsP+"values ($1, $2, $3, $4, $5) RETURNING id",
		s.Name, s.Icon, s.RefCategoryId, s.RefCharacteristicId, s.UserId,
	).Scan(&s.ID)
}

// CreateCategory ...
func (r *SkillRepository) CreateCategory(sc *model.SkillCategory) error {
	return r.store.CreateRetId(
		InsertQ+SkillCategoriesT+SkillCategoriesP+"values ($1, $2, $3) RETURNING id",
		sc.Name, sc.Icon, sc.UserId,
	).Scan(&sc.ID)
}

// Get ...
func (r *SkillRepository) Get() ([]*model.Skill, error) {
	var skills []*model.Skill

	bRows, err := r.store.SelectRows(
		SelectQ + SkillsT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		s := &model.Skill{}

		err := bRows.Scan(&s.ID, &s.Name, &s.Icon, &s.RefCategoryId, &s.RefCharacteristicId, &s.UserId)
		if err != nil {
			return nil, err
		}

		s.AfterScan()

		skills = append(skills, s)
	}

	return skills, nil
}

// GetByCategory ...
func (r *SkillRepository) GetByCategory(id int) ([]*model.Skill, error) {
	var skills []*model.Skill

	bRows, err := r.store.SelectRows(
		SelectQ+SkillsT+"WHERE category_id = $1", id,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		s := &model.Skill{}

		err := bRows.Scan(&s.ID, &s.Name, &s.Icon, &s.RefCategoryId, &s.UserId)
		if err != nil {
			return nil, err
		}

		s.AfterScan()

		skills = append(skills, s)
	}

	return skills, nil
}

// GetCategories ...
func (r *SkillRepository) GetCategories() ([]*model.SkillCategory, error) {
	var categories []*model.SkillCategory

	bRows, err := r.store.SelectRows(
		SelectQ + SkillCategoriesT,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		sc := &model.SkillCategory{}

		err := bRows.Scan(&sc.ID, &sc.Name, &sc.Icon, &sc.UserId)
		if err != nil {
			return nil, err
		}

		categories = append(categories, sc)
	}

	return categories, nil
}

// Find ...
func (r *SkillRepository) Find(id int) (*model.Skill, error) {
	s := &model.Skill{}

	if err := r.store.SelectRow(
		SelectQ+SkillsT+"WHERE id = $1", id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.Icon,
		&s.RefCategoryId,
		&s.RefCharacteristicId,
		&s.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	s.AfterScan()

	return s, nil
}

// FindCategory ...
func (r *SkillRepository) FindCategory(id int) (*model.SkillCategory, error) {
	sc := &model.SkillCategory{}

	if err := r.store.SelectRow(
		SelectQ+SkillCategoriesT+"WHERE id = $1", id,
	).Scan(
		&sc.ID,
		&sc.Name,
		&sc.Icon,
		&sc.UserId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return sc, nil
}

// Update ...
func (r *SkillRepository) Update(s *model.Skill) error {
	if err := s.Validate(); err != nil {
		return err
	}
	if err := s.BeforeInsertOrUpdate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+SkillsT+"SET name = $1, icon = $2, category_id = $3, characteristic_id = $4, user_id = $5"+
			" WHERE id = $5",
		s.Name, s.Icon, s.RefCategoryId, s.UserId, s.ID,
	)

	return err
}

// UpdateCategory ...
func (r *SkillRepository) UpdateCategory(sc *model.SkillCategory) error {
	_, err := r.store.Update(
		UpdateQ+SkillCategoriesT+"SET name = $1, icon = $2, user_id = $3 WHERE id = $4",
		sc.Name, sc.Icon, sc.UserId, sc.ID,
	)

	return err
}

// Delete ...
func (r *SkillRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+SkillsT+"WHERE id = $1", id)
	return err
}

func (r *SkillRepository) DeleteCategory(id int) error {
	_, err := r.store.Delete(DeleteQ+SkillCategoriesT+"WHERE id = $1", id)
	return err
}
