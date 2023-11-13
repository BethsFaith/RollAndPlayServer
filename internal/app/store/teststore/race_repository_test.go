package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaceRepository_Create(t *testing.T) {
	s := teststore.New()
	r := model.TestRace(t)

	assert.NoError(t, s.Race().Create(r))
	assert.NotNil(t, r)
}

func TestRaceRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.Race().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	r := model.TestRace(t)
	_ = s.Race().Create(r)

	r, err = s.Race().Find(r.ID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestRaceRepository_Update(t *testing.T) {
	s := teststore.New()

	race := model.TestRace(t)
	_ = s.Race().Create(race)

	name := "NewName"
	modelPath := "NewModel"
	race.Name = name
	race.Model = modelPath

	err := s.Race().Update(race)
	assert.NoError(t, err)

	var updatedRace *model.Race
	updatedRace, err = s.Race().Find(race.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedRace)
	assert.Equal(t, updatedRace.Model, modelPath)
	assert.Equal(t, updatedRace.Name, name)
}

func TestRaceRepository_Delete(t *testing.T) {
	s := teststore.New()

	race := model.TestRace(t)
	_ = s.Race().Create(race)

	err := s.Race().Delete(race.ID)
	assert.NoError(t, err)

	_, err = s.Race().Find(race.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
