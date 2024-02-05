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

	if err := s.BeforeInsertOrUpdate(); err != nil {
		return err
	}
	if s.RefCategoryId.Valid {
		_, ok := r.categories[s.CategoryId]
		if !ok {
			return store.ErrorNotExistRef
		}
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

func (r *SkillRepository) Update(skill *model.Skill) error {
	s, ok := r.skills[skill.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	if err := s.Validate(); err != nil {
		return err
	}
	if err := s.BeforeInsertOrUpdate(); err != nil {
		return err
	}
	if skill.RefCategoryId.Valid {
		_, ok = r.categories[skill.CategoryId]
		if !ok {
			return store.ErrorNotExistRef
		}
	}

	s.Name = skill.Name
	s.Icon = skill.Icon
	s.CategoryId = skill.CategoryId

	return nil
}

func (r *SkillRepository) UpdateCategory(category *model.SkillCategory) error {
	s, ok := r.categories[category.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	s.Name = category.Name
	s.Icon = category.Icon

	return nil
}

func (r *SkillRepository) Delete(id int) error {
	_, ok := r.skills[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.skills, id)

	return nil
}

func (r *SkillRepository) DeleteCategory(id int) error {
	_, ok := r.categories[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.categories, id)

	return nil
}
