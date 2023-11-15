package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
)

type RaceBonusRepository struct {
	store *Store
}

func (r *RaceBonusRepository) Create(rb *model.RaceBonus) error {
	if err := rb.Validate(); err != nil {
		return err
	}

	_, err := r.store.Create(
		InsertQ+RaceBonusesT+RaceBonusesP+"values ($1, $2, $3)", rb.RaceId,
		rb.SkillId, rb.Bonus,
	)

	return err
}

func (r *RaceBonusRepository) Find(raceId int, skillId int) (*model.RaceBonus, error) {
	rb := &model.RaceBonus{}

	if err := r.store.SelectRow(
		SelectQ+RaceBonusesT+"WHERE race_id = $1 AND skill_id = $2", raceId, skillId,
	).Scan(
		&rb.RaceId,
		&rb.SkillId,
		&rb.Bonus,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return rb, nil
}

func (r *RaceBonusRepository) Update(rb *model.RaceBonus) error {
	if err := rb.Validate(); err != nil {
		return err
	}

	_, err := r.store.Update(
		UpdateQ+RaceBonusesT+"SET bonus = $1 WHERE race_id = $2 AND skill_id = $3", rb.Bonus, rb.RaceId, rb.SkillId,
	)

	return err
}

func (r *RaceBonusRepository) Delete(raceId int, skillId int) error {
	_, err := r.store.Delete(DeleteQ+RaceBonusesT+"WHERE race_id = $1 AND skill_id = $2", raceId, skillId)
	return err
}
