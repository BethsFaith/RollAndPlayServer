package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestCharacterClassBonusRepository_FindByClassId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.CharacterClassBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	class := model.TestCharacterClass(t)
	class2 := model.TestCharacterClass(t)
	class2.Name = strconv.Itoa(2)
	_ = s.CharacterClass().Create(class)
	_ = s.CharacterClass().Create(class2)

	var bonuses1 []*model.CharacterClassBonus
	var bonuses2 []*model.CharacterClassBonus

	number := 10
	for i := 0; i < number; i++ {
		skill := model.TestSkill(t)
		skill.Name = strconv.Itoa(i)
		s.Skill().Create(skill)

		bonus := model.TestCharacterClassBonus(t)
		bonus.ClassId = class.ID
		bonus.SkillId = skill.ID
		bonuses1 = append(bonuses1, bonus)
		_ = s.CharacterClassBonus().Create(bonus)

		skill2 := model.TestSkill(t)
		skill2.Name = strconv.Itoa(i + 10)
		s.Skill().Create(skill2)

		bonus2 := model.TestCharacterClassBonus(t)
		bonus2.ClassId = class2.ID
		bonus2.SkillId = skill2.ID
		bonuses2 = append(bonuses2, bonus2)
		_ = s.CharacterClassBonus().Create(bonus2)
	}

	bonuses, err := s.CharacterClassBonus().FindByClassId(class.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses1))

	if len(bonuses) == len(bonuses1) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses1[i])
		}
	}

	bonuses, err = s.CharacterClassBonus().FindByClassId(class2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses2))

	if len(bonuses) == len(bonuses2) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses2[i])
		}
	}
}

//
//func TestCharacterClassBonusRepository_FindBySkillId(t *testing.T) {
//	db, teardown := sqlstore.TestDB(t, databaseURL)
//
//	defer teardown(sqlstore.CharacterClassBonusesT)
//
//	s := sqlstore.New(db)
//
//	id := 1
//	_, err := s.CharacterClassBonus().Find(id, id)
//	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
//
//	skill := model.TestSkill(t)
//	skill2 := model.TestSkill(t)
//	skill2.Name = strconv.Itoa(1)
//	s.Skill().Create(skill)
//	s.Skill().Create(skill2)
//
//	var bonuses1 []*model.CharacterClassBonus
//	var bonuses2 []*model.CharacterClassBonus
//
//	number := 10
//	for i := 0; i < number; i++ {
//		class := model.TestCharacterClass(t)
//		class.Name = strconv.Itoa(i)
//		_ = s.CharacterClass().Create(class)
//
//		bonus := model.TestCharacterClassBonus(t)
//		bonus.ClassId = class.ID
//		bonus.CategoryId = skill.ID
//		bonuses1 = append(bonuses1, bonus)
//		_ = s.CharacterClassBonus().Create(bonus)
//
//		class2 := model.TestCharacterClass(t)
//		class2.Name = strconv.Itoa(i + 10)
//		_ = s.CharacterClass().Create(class2)
//
//		bonus2 := model.TestCharacterClassBonus(t)
//		bonus2.ClassId = class2.ID
//		bonus2.CategoryId = skill2.ID
//		bonuses2 = append(bonuses2, bonus2)
//		_ = s.CharacterClassBonus().Create(bonus2)
//	}
//
//	bonuses, err := s.CharacterClassBonus().FindBySkillId(skill.ID)
//	assert.NoError(t, err)
//	assert.NotNil(t, bonuses)
//	assert.Equal(t, len(bonuses), len(bonuses1))
//
//	if len(bonuses) == len(bonuses1) {
//		for i := 0; i < len(bonuses); i++ {
//			assert.Equal(t, bonuses[i], bonuses1[i])
//		}
//	}
//
//	bonuses, err = s.CharacterClassBonus().FindBySkillId(skill2.ID)
//	assert.NoError(t, err)
//	assert.NotNil(t, bonuses)
//	assert.Equal(t, len(bonuses), len(bonuses2))
//
//	if len(bonuses) == len(bonuses2) {
//		for i := 0; i < len(bonuses); i++ {
//			assert.Equal(t, bonuses[i], bonuses2[i])
//		}
//	}
//}

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
