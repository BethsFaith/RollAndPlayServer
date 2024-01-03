package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)

	assert.NoError(t, s.System().Create(system))
	assert.NotNil(t, system)
}

func TestSystemRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT)

	s := sqlstore.New(db)

	_, err := s.System().Find(1)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	system := model.TestSystem(t)
	_ = s.System().Create(system)

	system, err = s.System().Find(system.ID)
	assert.NoError(t, err)
	assert.NotNil(t, system)
}

func TestSystemRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)

	assert.NoError(t, s.System().Update(system))

	_ = s.System().Create(system)
	system.Name = "UpdatedName"
	system.Icon = "UpdatedIcon"

	assert.NoError(t, s.System().Update(system))

	updatedSystem, err := s.System().Find(system.ID)

	assert.NoError(t, err)
	assert.Equal(t, system, updatedSystem)
}

func TestSystemRepository_AddRace(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemRacesT, sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	race := model.TestRace(t)

	races, err := s.System().AddRace(system.ID, race.ID)
	assert.Error(t, err)
	assert.Nil(t, races)

	_ = s.Race().Create(race)
	_ = s.System().Create(system)

	races, err = s.System().AddRace(system.ID, race.ID)

	assert.NoError(t, err)
	assert.NotNil(t, system, races)
}

func TestSystemRepository_AddCharacterClass(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemClassesT, sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	class := model.TestCharacterClass(t)

	classes, err := s.System().AddCharacterClass(system.ID, class.ID)
	assert.Error(t, err)
	assert.Nil(t, classes)

	_ = s.CharacterClass().Create(class)
	_ = s.System().Create(system)

	classes, err = s.System().AddCharacterClass(system.ID, class.ID)

	assert.NoError(t, err)
	assert.NotNil(t, system, classes)
}

func TestSystemRepository_AddSkillCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemSkillsT, sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	sc := model.TestSkillCategory(t)

	categories, err := s.System().AddSkillCategory(system.ID, sc.ID)
	assert.Error(t, err)
	assert.Nil(t, categories)

	_ = s.Skill().CreateCategory(sc)
	_ = s.System().Create(system)

	categories, err = s.System().AddSkillCategory(system.ID, sc.ID)

	assert.NoError(t, err)
	assert.NotNil(t, system, categories)
}

func TestSystemRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)

	assert.NoError(t, s.System().Delete(system.ID))

	_ = s.System().Create(system)
	system, _ = s.System().Find(system.ID)
	assert.NotNil(t, system)

	assert.NoError(t, s.System().Delete(system.ID))
	deletedSystem, err := s.System().Find(system.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedSystem)
}

func TestSystemRepository_DeleteRace(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT, sqlstore.SystemRacesT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	race := model.TestRace(t)

	assert.NoError(t, s.System().Delete(system.ID))

	_ = s.System().Create(system)
	_ = s.Race().Create(race)
	_, _ = s.System().AddRace(system.ID, race.ID)

	assert.NoError(t, s.System().DeleteRace(system.ID, race.ID))
	races, err := s.System().GetRaces(system.ID)

	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
	assert.Equal(t, len(races), 0)
}

func TestSystemRepository_DeleteCharacterClass(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT, sqlstore.SystemClassesT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	class := model.TestCharacterClass(t)

	assert.NoError(t, s.System().Delete(system.ID))

	_ = s.System().Create(system)
	_ = s.CharacterClass().Create(class)
	_, _ = s.System().AddRace(system.ID, class.ID)

	assert.NoError(t, s.System().DeleteCharacterClass(system.ID, class.ID))
	classes, err := s.System().GetCharacterClasses(system.ID)

	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
	assert.Equal(t, len(classes), 0)
}

func TestSystemRepository_DeleteSkillCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SystemsT, sqlstore.SystemSkillsT)

	s := sqlstore.New(db)
	system := model.TestSystem(t)
	skillCategory := model.TestSkillCategory(t)

	assert.NoError(t, s.System().Delete(system.ID))

	_ = s.System().Create(system)
	_ = s.Skill().CreateCategory(skillCategory)
	_, _ = s.System().AddRace(system.ID, skillCategory.ID)

	assert.NoError(t, s.System().DeleteSkillCategory(system.ID, skillCategory.ID))
	skills, err := s.System().GetSkillCategories(system.ID)

	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
	assert.Equal(t, len(skills), 0)
}
