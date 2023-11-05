package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type SkillRepository struct {
	store      *Store
	skills     map[int]*model.Skill
	categories map[int]*model.SkillCategory
}

func (r *SkillRepository) Create(s *model.Skill) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, ok := r.categories[s.CategoryId]
	if !ok {
		return store.ErrorNotExistRef
	}

	s.ID = len(r.skills) + 1
	r.skills[s.ID] = s

	return nil
}

func (r *SkillRepository) CreateCategory(c *model.SkillCategory) error {
	if err := c.Validate(); err != nil {
		return err
	}

	c.ID = len(r.categories) + 1
	r.categories[c.ID] = c

	return nil
}

func (r *SkillRepository) Find(id int) (*model.Skill, error) {
	s, ok := r.skills[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}
	return s, nil
}

func (r *SkillRepository) FindCategory(id int) (*model.SkillCategory, error) {
	c, ok := r.categories[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}
	return c, nil
}
