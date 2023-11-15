package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClassBonusRepository_Create(t *testing.T) {
	s := teststore.New()
	cb := model.TestCharacterClassBonus(t)

	assert.NoError(t, s.CharacterClassBonus().Create(cb))
	assert.NotNil(t, cb)
}

func TestCharacterClassBonusRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.CharacterClassBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	cb := model.TestCharacterClassBonus(t)
	_ = s.CharacterClassBonus().Create(cb)

	cb, err = s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, cb)
}

func TestCharacterClassBonusRepository_Update(t *testing.T) {
	s := teststore.New()

	cb := model.TestCharacterClassBonus(t)
	_ = s.CharacterClassBonus().Create(cb)

	cb.Bonus = cb.Bonus + 10

	err := s.CharacterClassBonus().Update(cb)
	assert.NoError(t, err)

	var updatedCb *model.CharacterClassBonus
	updatedCb, err = s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedCb)
	assert.Equal(t, updatedCb.Bonus, cb.Bonus)
}

func TestCharacterClassBonusRepository_Delete(t *testing.T) {
	s := teststore.New()

	cb := model.TestCharacterClassBonus(t)
	_ = s.CharacterClassBonus().Create(cb)

	err := s.CharacterClassBonus().Delete(cb.ClassId, cb.SkillId)
	assert.NoError(t, err)

	_, err = s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
