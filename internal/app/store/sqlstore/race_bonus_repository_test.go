package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestRaceBonusRepository_FindByRaceId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.CharacterClassBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	race := model.TestRace(t)
	race2 := model.TestRace(t)
	race2.Name = strconv.Itoa(2)
	s.Race().Create(race)
	s.Race().Create(race2)

	var bonuses1 []*model.RaceBonus
	var bonuses2 []*model.RaceBonus

	number := 10
	for i := 0; i < number; i++ {
		skill := model.TestSkill(t)
		skill.Name = strconv.Itoa(i)
		s.Skill().Create(skill)

		bonus := model.TestRaceBonus(t)
		bonus.RaceId = race.ID
		bonus.SkillId = skill.ID
		bonuses1 = append(bonuses1, bonus)
		s.RaceBonus().Create(bonus)

		skill2 := model.TestSkill(t)
		skill2.Name = strconv.Itoa(i + 10)
		s.Skill().Create(skill2)

		bonus2 := model.TestRaceBonus(t)
		bonus2.RaceId = race2.ID
		bonus2.SkillId = skill2.ID
		bonuses2 = append(bonuses2, bonus2)
		s.RaceBonus().Create(bonus2)
	}

	bonuses, err := s.RaceBonus().FindByRaceId(race.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses1))

	if len(bonuses) == len(bonuses1) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses1[i])
		}
	}

	bonuses, err = s.RaceBonus().FindByRaceId(race2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses2))

	if len(bonuses) == len(bonuses2) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses2[i])
		}
	}
}

func TestRaceBonusRepository_FindBySkillId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RaceBonusesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.RaceBonus().Find(id, id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	skill := model.TestSkill(t)
	skill2 := model.TestSkill(t)
	skill2.Name = strconv.Itoa(1)
	s.Skill().Create(skill)
	s.Skill().Create(skill2)

	var bonuses1 []*model.RaceBonus
	var bonuses2 []*model.RaceBonus

	number := 10
	for i := 0; i < number; i++ {
		race := model.TestRace(t)
		race.Name = strconv.Itoa(i)
		s.Race().Create(race)

		bonus := model.TestRaceBonus(t)
		bonus.RaceId = race.ID
		bonus.SkillId = skill.ID
		bonuses1 = append(bonuses1, bonus)
		s.RaceBonus().Create(bonus)

		race2 := model.TestRace(t)
		race2.Name = strconv.Itoa(i + 10)
		s.Race().Create(race2)

		bonus2 := model.TestRaceBonus(t)
		bonus2.RaceId = race2.ID
		bonus2.SkillId = skill2.ID
		bonuses2 = append(bonuses2, bonus2)
		s.RaceBonus().Create(bonus2)
	}

	bonuses, err := s.RaceBonus().FindBySkillId(skill.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses1))

	if len(bonuses) == len(bonuses1) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses1[i])
		}
	}

	bonuses, err = s.RaceBonus().FindBySkillId(skill2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bonuses)
	assert.Equal(t, len(bonuses), len(bonuses2))

	if len(bonuses) == len(bonuses2) {
		for i := 0; i < len(bonuses); i++ {
			assert.Equal(t, bonuses[i], bonuses2[i])
		}
	}
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
