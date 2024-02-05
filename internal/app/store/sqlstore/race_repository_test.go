package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaceRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RacesT, sqlstore.UsersT)

	s := sqlstore.New(db)
	race := model.TestRace(t)

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	race.UserId = u.ID
	assert.NoError(t, s.Race().Create(race))
	assert.NotNil(t, race)
}

func TestRaceRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RacesT, sqlstore.UsersT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Race().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	race := model.TestRace(t)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	race.UserId = u.ID
	_ = s.Race().Create(race)

	race, err = s.Race().Find(race.ID)
	assert.NoError(t, err)
	assert.NotNil(t, race)
}

func TestRaceRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RacesT, sqlstore.UsersT)

	s := sqlstore.New(db)
	race := model.TestRace(t)

	assert.NoError(t, s.Race().Update(race))

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	race.UserId = u.ID
	_ = s.Race().Create(race)
	race.Name = "UpdatedName"

	assert.NoError(t, s.Race().Update(race))

	updatedRace, err := s.Race().Find(race.ID)

	assert.NoError(t, err)
	assert.Equal(t, race, updatedRace)
}

func TestRaceRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.RacesT, sqlstore.UsersT)

	s := sqlstore.New(db)
	race := model.TestRace(t)

	assert.NoError(t, s.Race().Delete(race.ID))

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	race.UserId = u.ID
	_ = s.Race().Create(race)
	race, _ = s.Race().Find(race.ID)
	assert.NotNil(t, race)

	assert.NoError(t, s.Race().Delete(race.ID))
	deletedRace, err := s.Race().Find(race.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedRace)
}
