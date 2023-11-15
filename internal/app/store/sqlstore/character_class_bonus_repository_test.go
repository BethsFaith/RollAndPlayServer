package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClassBonusRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassBonusesT)

	s := sqlstore.New(db)

	c := model.TestCharacterClass(t)
	sk := model.TestSkill(t)
	cb := model.TestCharacterClassBonus(t)

	assert.NoError(t, s.CharacterClass().Create(c))
	assert.NoError(t, s.Skill().Create(sk))

	cb.ClassId = c.ID
	cb.SkillId = sk.ID

	assert.NoError(t, s.CharacterClassBonus().Create(cb))
	assert.NotNil(t, cb)
}

func TestCharacterClassBonusRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.CharacterClassBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestCharacterClass(t)
	sk := model.TestSkill(t)
	cb := model.TestCharacterClassBonus(t)

	assert.NoError(t, s.CharacterClass().Create(c))
	assert.NoError(t, s.Skill().Create(sk))

	cb.ClassId = c.ID
	cb.SkillId = sk.ID

	_ = s.CharacterClassBonus().Create(cb)

	cb, err = s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, cb)
}

func TestCharacterClassBonusRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.CharacterClassBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestCharacterClass(t)
	sk := model.TestSkill(t)
	cb := model.TestCharacterClassBonus(t)

	assert.NoError(t, s.CharacterClass().Create(c))
	assert.NoError(t, s.Skill().Create(sk))

	cb.ClassId = c.ID
	cb.SkillId = sk.ID

	_ = s.CharacterClassBonus().Create(cb)

	cb.Bonus = cb.Bonus + 10

	assert.NoError(t, s.CharacterClassBonus().Update(cb))

	updatedCb, err := s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)

	assert.NoError(t, err)
	assert.Equal(t, cb, updatedCb)
}

func TestCharacterClassBonusRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassBonusesT)

	s := sqlstore.New(db)
	c := model.TestCharacterClass(t)
	sk := model.TestSkill(t)
	cb := model.TestCharacterClassBonus(t)

	assert.NoError(t, s.CharacterClassBonus().Delete(cb.ClassId, cb.SkillId))

	assert.NoError(t, s.CharacterClass().Create(c))
	assert.NoError(t, s.Skill().Create(sk))
	cb.ClassId = c.ID
	cb.SkillId = sk.ID
	_ = s.CharacterClassBonus().Create(cb)

	cb, err := s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)
	assert.NoError(t, err)
	assert.NotNil(t, cb)

	assert.NoError(t, s.CharacterClassBonus().Delete(cb.ClassId, cb.SkillId))
	deletedCb, err := s.CharacterClassBonus().Find(cb.ClassId, cb.SkillId)

	assert.Error(t, err)
	assert.Nil(t, deletedCb)
}
