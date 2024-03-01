package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacteristicRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacteristicsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	c := model.TestCharacteristic(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	c.UserId = u.ID

	assert.NoError(t, s.Characteristic().Create(c))
	assert.NotNil(t, c)
}

func TestCharacteristicRepository_Get(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacteristicsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	c := model.TestCharacteristic(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	c.UserId = u.ID

	assert.NoError(t, s.Characteristic().Create(c))
	assert.NotNil(t, c)

	c.Name = "test2"
	assert.NoError(t, s.Characteristic().Create(c))

	characteristics, err := s.Characteristic().Get()
	assert.NoError(t, err)
	assert.NotNil(t, characteristics)
	assert.Equal(t, len(characteristics), 2)
}

func TestCharacteristicRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacteristicsT, sqlstore.UsersT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Characteristic().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestCharacteristic(t)
	u := model.TestUser(t)

	_ = s.User().Create(u)
	c.UserId = u.ID
	_ = s.Characteristic().Create(c)

	c, err = s.Characteristic().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCharacteristicRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacteristicsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	c := model.TestCharacteristic(t)

	assert.NoError(t, s.Characteristic().Update(c))

	u := model.TestUser(t)

	_ = s.User().Create(u)
	c.UserId = u.ID
	_ = s.Characteristic().Create(c)

	c.Name = "UpdatedName"

	assert.NoError(t, s.Characteristic().Update(c))

	updatedC, err := s.Characteristic().Find(c.ID)

	assert.NoError(t, err)
	assert.Equal(t, c, updatedC)
}

func TestCharacteristicRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacteristicsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	c := model.TestCharacteristic(t)
	u := model.TestUser(t)

	assert.NoError(t, s.Characteristic().Delete(c.ID))

	_ = s.User().Create(u)
	c.UserId = u.ID
	_ = s.Characteristic().Create(c)
	c, _ = s.Characteristic().Find(c.ID)
	assert.NotNil(t, c)

	assert.NoError(t, s.Characteristic().Delete(c.ID))
	deletedC, err := s.Characteristic().Find(c.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedC)
}
