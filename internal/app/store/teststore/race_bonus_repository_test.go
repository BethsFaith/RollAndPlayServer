package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaceBonusRepository_Create(t *testing.T) {
	s := teststore.New()
	r := model.TestRaceBonus(t)

	assert.NoError(t, s.RaceBonus().Create(r))
	assert.NotNil(t, r)
}

func TestRaceBonusRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.RaceBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	r := model.TestRaceBonus(t)
	_ = s.RaceBonus().Create(r)

	r, err = s.RaceBonus().Find(r.RaceId, r.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestRaceBonusRepository_Update(t *testing.T) {
	s := teststore.New()

	rb := model.TestRaceBonus(t)
	_ = s.RaceBonus().Create(rb)

	rb.Bonus = rb.Bonus + 10

	err := s.RaceBonus().Update(rb)
	assert.NoError(t, err)

	var updatedRb *model.RaceBonus
	updatedRb, err = s.RaceBonus().Find(rb.RaceId, rb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedRb)
	assert.Equal(t, updatedRb.Bonus, rb.Bonus)
}

func TestRaceBonusRepository_Delete(t *testing.T) {
	s := teststore.New()

	rb := model.TestRaceBonus(t)
	_ = s.RaceBonus().Create(rb)

	err := s.RaceBonus().Delete(rb.RaceId, rb.SkillId)
	assert.NoError(t, err)

	_, err = s.RaceBonus().Find(rb.RaceId, rb.SkillId)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
