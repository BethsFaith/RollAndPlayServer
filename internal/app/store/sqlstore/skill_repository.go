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
		_, err := r.Find(s.CategoryId)
		if err != nil {
			return store.ErrorNotExistRef
		}
	}

	return r.store.Create(
		InsertQ+SkillsT+SkillsP+"values ($1, $2, $3) RETURNING id", s.Name, s.Icon, s.RefCategoryId,
	).Scan(&s.ID)
}

// CreateCategory ...
func (r *SkillRepository) CreateCategory(sc *model.SkillCategory) error {
	return r.store.Create(
		InsertQ+SkillCategoriesT+SkillCategoriesP+"values ($1, $2) RETURNING id",
		sc.Name, sc.Icon,
	).Scan(&sc.ID)
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
		UpdateQ+SkillsT+"SET name = $1, icon = $2, category_id = $3 WHERE id = $4", s.Name, s.Icon,
		s.RefCategoryId, s.ID,
	)

	return err
}

// UpdateCategory ...
func (r *SkillRepository) UpdateCategory(sc *model.SkillCategory) error {
	_, err := r.store.Update(
		UpdateQ+SkillCategoriesT+"SET name = $1, icon = $2 WHERE id = $3", sc.Name, sc.Icon, sc.ID,
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
