package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type SystemRepository struct {
	store      *Store
	systems    map[int]*model.System
	components map[int]*model.SystemComponent
}

func (r *SystemRepository) Create(s *model.System) error {
	if err := s.Validate(); err != nil {
		return err
	}

	s.ID = len(r.systems) + 1
	r.systems[s.ID] = s

	return nil
}

func (r *SystemRepository) Find(id int) (*model.System, error) {
	s, ok := r.systems[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}
	return s, nil
}

func (r *SystemRepository) AddRace(id int, raceId int) ([]*model.Race, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: raceId,
	}

	r.components[len(r.components)+1] = s

	var races []*model.Race
	for _, value := range r.components {
		r := &model.Race{
			ID: value.ComponentId,
		}

		races = append(races, r)
	}

	if len(races) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return races, nil
}

func (r *SystemRepository) AddSkillCategory(id int, categoryId int) ([]*model.SkillCategory, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: categoryId,
	}

	r.components[len(r.components)+1] = s

	var categories []*model.SkillCategory
	for _, value := range r.components {
		r := &model.SkillCategory{
			ID: value.ComponentId,
		}

		categories = append(categories, r)
	}

	if len(categories) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return categories, nil
}

func (r *SystemRepository) AddCharacterClass(id int, classId int) ([]*model.CharacterClass, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: classId,
	}

	r.components[len(r.components)+1] = s

	var categories []*model.CharacterClass
	for _, value := range r.components {
		r := &model.CharacterClass{
			ID: value.ComponentId,
		}

		categories = append(categories, r)
	}

	if len(categories) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return categories, nil
}

func (r *SystemRepository) Update(system *model.System) error {
	s, ok := r.systems[system.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	err := s.Validate()
	if err != nil {
		return err
	}

	s.Name = system.Name
	s.Icon = system.Icon

	return nil
}

func (r *SystemRepository) Delete(id int) error {
	_, ok := r.systems[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.systems, id)

	return nil
}
