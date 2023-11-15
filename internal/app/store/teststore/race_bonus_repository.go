package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type RaceBonusKey struct {
	raceId  int
	skillId int
}

type RaceBonusRepository struct {
	store   *Store
	bonuses map[RaceBonusKey]*model.RaceBonus
}

// Create ...
func (r *RaceBonusRepository) Create(rb *model.RaceBonus) error {
	if err := rb.Validate(); err != nil {
		return err
	}

	key := RaceBonusKey{
		raceId:  rb.RaceId,
		skillId: rb.SkillId,
	}

	r.bonuses[key] = rb

	return nil
}

// Find ...
func (r *RaceBonusRepository) Find(raceId int, skillId int) (*model.RaceBonus, error) {
	rb, ok := r.bonuses[RaceBonusKey{raceId: raceId, skillId: skillId}]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return rb, nil
}

// Update ...
func (r *RaceBonusRepository) Update(rb *model.RaceBonus) error {
	if err := rb.Validate(); err != nil {
		return err
	}

	source, ok := r.bonuses[RaceBonusKey{raceId: rb.RaceId, skillId: rb.SkillId}]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Bonus = rb.Bonus

	return nil
}

// Delete ...
func (r *RaceBonusRepository) Delete(raceId int, skillId int) error {
	key := RaceBonusKey{raceId: raceId, skillId: skillId}
	_, ok := r.bonuses[key]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.bonuses, key)

	return nil
}
