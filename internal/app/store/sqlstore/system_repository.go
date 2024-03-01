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

func (r *SystemRepository) GetRaces(systemId int) ([]*model.SystemComponent, error) {
	return r.selectComponent(systemId, SystemRacesT)
}
func (r *SystemRepository) GetSkillCategories(systemId int) ([]*model.SystemComponent, error) {
	return r.selectComponent(systemId, SystemSkillsT)
}
func (r *SystemRepository) GetCharacterClasses(systemId int) ([]*model.SystemComponent, error) {
	return r.selectComponent(systemId, SystemClassesT)
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

func (r *SystemRepository) AddRace(id int, raceId int) ([]*model.SystemComponent, error) {
	_, err := r.store.Create(
		InsertQ+SystemRacesT+SystemRacesP+"values ($1, $2) ", id, raceId,
	)

	if err == nil {
		return nil, err
	}

	return r.selectComponent(id, SystemRacesT)
}

func (r *SystemRepository) AddSkillCategory(id int, categoryId int) ([]*model.SystemComponent, error) {
	_, err := r.store.Create(
		InsertQ+SystemSkillsT+SystemSkillsP+"values ($1, $2) ", id, categoryId,
	)

	if err == nil {
		return nil, err
	}

	return r.selectComponent(id, SystemSkillsT)
}

func (r *SystemRepository) AddCharacterClass(id int, categoryId int) ([]*model.SystemComponent, error) {
	_, err := r.store.Create(
		InsertQ+SystemClassesT+SystemClassesP+"values ($1, $2) ", id, categoryId,
	)

	if err == nil {
		return nil, err
	}

	return r.selectComponent(id, SystemClassesT)
}

func (r *SystemRepository) Delete(id int) error {
	_, err := r.store.Delete(DeleteQ+SystemsT+"WHERE id = $1", id)

	return err
}

func (r *SystemRepository) DeleteRace(id int, raceId int) error {
	_, err := r.store.Delete(DeleteQ+SystemRacesT+"WHERE system_id = $1 AND race_id = $2", id, raceId)

	return err
}

func (r *SystemRepository) DeleteSkillCategory(id int, categoryId int) error {
	_, err := r.store.Delete(DeleteQ+SystemSkillsT+"WHERE system_id = $1 AND skill_category_id = $2", id, categoryId)

	return err
}

func (r *SystemRepository) DeleteCharacterClass(id int, class_id int) error {
	_, err := r.store.Delete(DeleteQ+SystemClassesT+"WHERE system_id = $1 AND class_id = $2", id, class_id)

	return err
}
