package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type SystemRepository struct {
	store      *Store
	systems    map[int]*model.System
	races      map[int][]*model.SystemComponent
	classes    map[int][]*model.SystemComponent
	categories map[int][]*model.SystemComponent
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

func (r *SystemRepository) GetRaces(id int) ([]*model.SystemComponent, error) {
	return r.races[id], nil
}

func (r *SystemRepository) GetSkillCategories(id int) ([]*model.SystemComponent, error) {
	return r.categories[id], nil
}

func (r *SystemRepository) GetCharacterClasses(id int) ([]*model.SystemComponent, error) {
	return r.classes[id], nil
}

func (r *SystemRepository) AddRace(id int, raceId int) ([]*model.SystemComponent, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: raceId,
	}

	r.races[id] = append(r.races[id], s)

	return r.races[id], nil
}

func (r *SystemRepository) AddSkillCategory(id int, categoryId int) ([]*model.SystemComponent, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: categoryId,
	}

	r.categories[id] = append(r.categories[id], s)

	return r.categories[id], nil
}

func (r *SystemRepository) AddCharacterClass(id int, classId int) ([]*model.SystemComponent, error) {
	s := &model.SystemComponent{
		SystemId:    id,
		ComponentId: classId,
	}

	r.classes[id] = append(r.classes[id], s)

	return r.classes[id], nil
}

func (r *SystemRepository) Update(system *model.System) error {
	s, ok := r.systems[system.ID]
	if !ok {
		return store.ErrorRecordNotFound
	}

	err := system.Validate()
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

func (r *SystemRepository) DeleteRace(id int, raceId int) error {
	races, ok := r.races[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	for i := range races {
		if races[i].ComponentId == raceId {
			var t = races[i]
			races[i] = races[len(races)-1]
			races[len(races)-1] = t

			break
		}
	}
	races = races[:len(races)-1]

	return nil
}

func (r *SystemRepository) DeleteCharacterClass(id int, classId int) error {
	classes, ok := r.classes[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	for i := range classes {
		if classes[i].ComponentId == classId {
			var t = classes[i]
			classes[i] = classes[len(classes)-1]
			classes[len(classes)-1] = t

			break
		}
	}
	classes = classes[:len(classes)-1]

	return nil
}

func (r *SystemRepository) DeleteSkillCategory(id int, categoryId int) error {
	categories, ok := r.categories[id]
	if !ok {
		return store.ErrorRecordNotFound
	}

	for i := range categories {
		if categories[i].ComponentId == categoryId {
			var t = categories[i]
			categories[i] = categories[len(categories)-1]
			categories[len(categories)-1] = t

			break
		}
	}
	categories = categories[:len(categories)-1]

	return nil
}
