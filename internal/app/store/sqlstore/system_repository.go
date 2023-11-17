package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type SystemRepository struct {
	store *Store
}

func (r *SystemRepository) Create(s *model.System) error {
	if err := s.Validate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+SystemsT+SystemP+"values ($1, $2) RETURNING id", s.Name, s.Icon,
	).Scan(&s.ID)
}

func (r *SystemRepository) Find(id int) (*model.System, error) {
	s := &model.System{}

	if err := r.store.SelectRow(
		SelectQ+SystemsT+"WHERE id = $1", id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.Icon,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

func (r *SystemRepository) selectComponent(id int, table string) ([]*model.SystemComponent, error) {
	var components []*model.SystemComponent

	rRows, err := r.store.SelectRows(
		SelectQ+table+"Where system_id = $1", id,
	)

	if err != nil {
		return nil, err
	}

	for rRows.Next() {
		c := &model.SystemComponent{}

		err := rRows.Scan(&c.SystemId, &c.ComponentId)
		if err != nil {
			return nil, err
		}

		components = append(components, c)
	}

	if len(components) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return components, nil
}

func (r *SystemRepository) Update(s *model.System) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+SystemsT+"SET name = $1, icon = $2 WHERE id = $3", s.Name, s.Icon, s.ID,
	)

	return err
}

func (r *SystemRepository) AddRace(id int, raceId int) ([]*model.Race, error) {
	_, err := r.store.Create(
		InsertQ+SystemRacesT+SystemRacesP+"values ($1, $2) ", id, raceId,
	)

	if err == nil {
		return nil, err
	}

	components, err := r.selectComponent(id, SystemRacesT)

	var races []*model.Race
	for _, value := range components {
		r := &model.Race{}
		r.ID = value.ComponentId

		races = append(races, r)
	}

	if len(races) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return races, nil
}

func (r *SystemRepository) AddSkillCategory(id int, categoryId int) ([]*model.SkillCategory, error) {
	_, err := r.store.Create(
		InsertQ+SystemSkillsT+SystemSkillsP+"values ($1, $2) ", id, categoryId,
	)

	if err == nil {
		return nil, err
	}

	components, err := r.selectComponent(id, SystemSkillsT)

	var categories []*model.SkillCategory
	for _, value := range components {
		sc := &model.SkillCategory{}
		sc.ID = value.ComponentId

		categories = append(categories, sc)
	}

	if len(categories) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return categories, nil
}

func (r *SystemRepository) AddCharacterClass(id int, categoryId int) ([]*model.CharacterClass, error) {
	_, err := r.store.Create(
		InsertQ+SystemClassesT+SystemClassesP+"values ($1, $2) ", id, categoryId,
	)

	if err == nil {
		return nil, err
	}

	components, err := r.selectComponent(id, SystemClassesT)

	var categories []*model.CharacterClass
	for _, value := range components {
		cc := &model.CharacterClass{}
		cc.ID = value.ComponentId

		categories = append(categories, cc)
	}

	if len(categories) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return categories, nil
}

func (r *SystemRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+SystemsT+"WHERE id = $1", id)

	return err
}
