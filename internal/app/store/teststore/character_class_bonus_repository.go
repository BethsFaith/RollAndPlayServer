package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type CharacterClassBonusKey struct {
	classId int
	skillId int
}

type CharacterClassBonusRepository struct {
	store   *Store
	bonuses map[CharacterClassBonusKey]*model.CharacterClassBonus
}

// Create ...
func (r *CharacterClassBonusRepository) Create(rb *model.CharacterClassBonus) error {
	if err := rb.Validate(); err != nil {
		return err
	}

	key := CharacterClassBonusKey{
		classId: rb.ClassId,
		skillId: rb.SkillId,
	}

	r.bonuses[key] = rb

	return nil
}

// Find ...
func (r *CharacterClassBonusRepository) Find(classId int, skillId int) (*model.CharacterClassBonus, error) {
	cb, ok := r.bonuses[CharacterClassBonusKey{classId: classId, skillId: skillId}]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return cb, nil
}

func (r *CharacterClassBonusRepository) FindByClassId(classId int) ([]*model.CharacterClassBonus, error) {
	var bonuses []*model.CharacterClassBonus

	for key, value := range r.bonuses {
		if key.classId == classId {
			bonuses = append(bonuses, value)
		}
	}

	if len(bonuses) == 0 {
		return nil, store.ErrorRecordNotFound
	}

	return bonuses, nil
}

// Update ...
func (r *CharacterClassBonusRepository) Update(cb *model.CharacterClassBonus) error {
	if err := cb.Validate(); err != nil {
		return err
	}

	source, ok := r.bonuses[CharacterClassBonusKey{classId: cb.ClassId, skillId: cb.SkillId}]
	if !ok {
		return store.ErrorRecordNotFound
	}

	source.Bonus = cb.Bonus

	return nil
}

// Delete ...
func (r *CharacterClassBonusRepository) Delete(classId int, skillId int) error {
	key := CharacterClassBonusKey{classId: classId, skillId: skillId}
	_, ok := r.bonuses[key]
	if !ok {
		return store.ErrorRecordNotFound
	}

	delete(r.bonuses, key)

	return nil
}
