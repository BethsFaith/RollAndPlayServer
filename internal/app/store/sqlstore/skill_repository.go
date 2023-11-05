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

// CreateSkill ...
func (r *SkillRepository) CreateSkill(s *model.Skill) error {
	return r.store.Create(
		insertQ+skillsT+skillsP+"values ($1, $2, $3) RETURNING id", s.Name, s.Icon, s.CategoryId,
	).Scan(&s.ID)
}

// CreateCategory ...
func (r *SkillRepository) CreateCategory(sc *model.SkillCategory) error {
	return r.store.Create(
		insertQ+skillCategoryT+skillCategoriesP+"values ($1, $2) RETURNING id",
		sc.Name, sc.Icon,
	).Scan(&sc.ID)
}

// FindSkill ...
func (r *SkillRepository) FindSkill(id int) (*model.Skill, error) {
	s := &model.Skill{}

	if err := r.store.SelectRow(
		selectQ+skillsT+"WHERE id = $1", id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.Icon,
		&s.CategoryId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

// FindCategory ...
func (r *SkillRepository) FindCategory(id int) (*model.SkillCategory, error) {
	sc := &model.SkillCategory{}

	if err := r.store.SelectRow(
		selectQ+skillCategoryT+"WHERE id = $1", id,
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
