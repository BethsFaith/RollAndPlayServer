package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaceBonusRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)

	r := model.TestRace(t)
	sk := model.TestSkill(t)
	rb := model.TestRaceBonus(t)

	assert.NoError(t, s.Race().Create(r))
	assert.NoError(t, s.Skill().Create(sk))

	rb.RaceId = r.ID
	rb.SkillId = sk.ID

	assert.NoError(t, s.RaceBonus().Create(rb))
	assert.NotNil(t, rb)
}

func TestRaceBonusRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.RaceBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	r := model.TestRace(t)
	sk := model.TestSkill(t)
	rb := model.TestRaceBonus(t)

	assert.NoError(t, s.Race().Create(r))
	assert.NoError(t, s.Skill().Create(sk))

	rb.RaceId = r.ID
	rb.SkillId = sk.ID

	_ = s.RaceBonus().Create(rb)

	rb, err = s.RaceBonus().Find(rb.RaceId, rb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, rb)
}

func TestRaceBonusRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)
	r := model.TestRace(t)
	sk := model.TestSkill(t)
	rb := model.TestRaceBonus(t)

	assert.NoError(t, s.Race().Create(r))
	assert.NoError(t, s.Skill().Create(sk))

	rb.RaceId = r.ID
	rb.SkillId = sk.ID

	_ = s.RaceBonus().Create(rb)
	rb.Bonus = rb.Bonus + 10

	assert.NoError(t, s.RaceBonus().Update(rb))

	updatedRb, err := s.RaceBonus().Find(rb.RaceId, rb.SkillId)

	assert.NoError(t, err)
	assert.Equal(t, rb, updatedRb)
}

func TestRaceBonusRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)
	r := model.TestRace(t)
	sk := model.TestSkill(t)
	rb := model.TestRaceBonus(t)

	assert.NoError(t, s.RaceBonus().Delete(rb.RaceId, rb.SkillId))

	assert.NoError(t, s.Race().Create(r))
	assert.NoError(t, s.Skill().Create(sk))
	rb.RaceId = r.ID
	rb.SkillId = sk.ID
	_ = s.RaceBonus().Create(rb)

	rb, err := s.RaceBonus().Find(rb.RaceId, rb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, rb)

	assert.NoError(t, s.RaceBonus().Delete(rb.RaceId, rb.SkillId))
	deletedRace, err := s.RaceBonus().Find(rb.RaceId, rb.SkillId)

	assert.Error(t, err)
	assert.Nil(t, deletedRace)
}
