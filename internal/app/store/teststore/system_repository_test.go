package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemRepository_Create(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)

	assert.NoError(t, s.System().Create(sys))
	assert.NotNil(t, sys)
}

func TestSystemRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.System().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	system := model.TestSystem(t)
	_ = s.System().Create(system)

	system, err = s.System().Find(system.ID)
	assert.NoError(t, err)
	assert.NotNil(t, system)
}

func TestSystemRepository_Update(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)
	_ = s.System().Create(sys)

	sys.Icon = "NewIcon"
	sys.Name = "NewName"

	err := s.System().Update(sys)
	assert.NoError(t, err)

	updatedSystem, err := s.System().Find(sys.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedSystem)
	assert.Equal(t, sys, updatedSystem)
}

func TestSystemRepository_AddRace(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)
	r := model.TestRace(t)
	_ = s.System().Create(sys)

	races, err := s.System().AddRace(sys.ID, r.ID)
	assert.NoError(t, err)
	assert.NotNil(t, races)

	assert.Equal(t, races[0].ID, r.ID)
}

func TestSystemRepository_AddSkillCategory(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)
	category := model.TestSkillCategory(t)
	_ = s.System().Create(sys)

	categories, err := s.System().AddSkillCategory(sys.ID, category.ID)
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	assert.Equal(t, categories[0].ID, category.ID)
}

func TestSystemRepository_AddCharacterClass(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)
	class := model.TestCharacterClass(t)
	_ = s.System().Create(sys)

	classes, err := s.System().AddCharacterClass(sys.ID, class.ID)
	assert.NoError(t, err)
	assert.NotNil(t, classes)

	assert.Equal(t, classes[0].ID, class.ID)
}

func TestSystemRepository_Delete(t *testing.T) {
	s := teststore.New()

	sys := model.TestSystem(t)
	_ = s.System().Create(sys)

	err := s.System().Delete(sys.ID)
	assert.NoError(t, err)

	_, err = s.System().Find(sys.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
