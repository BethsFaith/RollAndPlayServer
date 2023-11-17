package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type CharacterClassBonusRepository struct {
	store *Store
}

// Create ...
func (r *CharacterClassBonusRepository) Create(cb *model.CharacterClassBonus) error {
	if err := cb.Validate(); err != nil {
		return err
	}

	_, err := r.store.Create(
		InsertQ+CharacterClassBonusesT+CharacterClassBonusesP+
			"values ($1, $2, $3) RETURNING (class_id, skill_id) ",
		&cb.ClassId, &cb.SkillId, cb.Bonus,
	)

	return err
}

// Find ...
func (r *CharacterClassBonusRepository) Find(classId int, skillId int) (*model.CharacterClassBonus, error) {
	cc := &model.CharacterClassBonus{}

	if err := r.store.SelectRow(
		SelectQ+CharacterClassBonusesT+"WHERE class_id = $1 AND skill_id = $2", classId, skillId,
	).Scan(
		&cc.ClassId,
		&cc.SkillId,
		&cc.Bonus,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return cc, nil
}

// FindByClassId ...
func (r *CharacterClassBonusRepository) FindByClassId(classId int) ([]*model.CharacterClassBonus, error) {
	var bonuses []*model.CharacterClassBonus

	bRows, err := r.store.SelectRows(
		SelectQ+CharacterClassBonusesT+"WHERE class_id = $1", classId,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		cc := &model.CharacterClassBonus{}

		err := bRows.Scan(&cc.ClassId, &cc.SkillId, &cc.Bonus)
		if err != nil {
			return nil, err
		}

		bonuses = append(bonuses, cc)
	}

	return bonuses, nil
}

// FindBySkillId ...
func (r *CharacterClassBonusRepository) FindBySkillId(skillId int) ([]*model.CharacterClassBonus, error) {
	var bonuses []*model.CharacterClassBonus

	bRows, err := r.store.SelectRows(
		SelectQ+CharacterClassBonusesT+"WHERE skill_id = $1", skillId,
	)
	if err != nil {
		return nil, err
	}

	for bRows.Next() {
		cc := &model.CharacterClassBonus{}

		err := bRows.Scan(&cc.ClassId, &cc.SkillId, &cc.Bonus)
		if err != nil {
			return nil, err
		}

		bonuses = append(bonuses, cc)
	}

	return bonuses, nil
}

// Update ...
func (r *CharacterClassBonusRepository) Update(cb *model.CharacterClassBonus) error {
	if err := cb.Validate(); err != nil {
		return err
	}
	_, err := r.store.Update(
		UpdateQ+CharacterClassBonusesT+"SET bonus = $1 WHERE class_id = $2 AND skill_id = $3",
		cb.Bonus, cb.ClassId, cb.SkillId,
	)

	return err
}

// Delete ...
func (r *CharacterClassBonusRepository) Delete(classId int, skillId int) error {
	_, err := r.store.Delete(DeleteQ+CharacterClassBonusesT+"WHERE class_id = $1 AND skill_id = $2",
		classId, skillId)
	return err
}
