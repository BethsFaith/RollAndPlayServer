package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ItemsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	i := model.TestItem(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	i.UserId = u.ID

	assert.NoError(t, s.Item().Create(i))
	assert.NotNil(t, i)

	i.TypeId = -1
	assert.Error(t, s.Item().Create(i))

	i.TypeId = 10
	assert.Error(t, s.Item().Create(i))
}

func TestItemRepository_Get(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ItemsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	a := model.TestItem(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	a.UserId = u.ID

	assert.NoError(t, s.Item().Create(a))
	assert.NotNil(t, a)

	a.Name = "test2"
	assert.NoError(t, s.Item().Create(a))

	items, err := s.Item().Get()
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, len(items), 2)
}

func TestItemRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ItemsT, sqlstore.UsersT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Item().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	a := model.TestItem(t)
	u := model.TestUser(t)

	_ = s.User().Create(u)
	a.UserId = u.ID
	_ = s.Item().Create(a)

	a, err = s.Item().Find(a.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestItemRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ItemsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	a := model.TestItem(t)

	assert.NoError(t, s.Item().Update(a))

	u := model.TestUser(t)

	_ = s.User().Create(u)
	a.UserId = u.ID
	_ = s.Item().Create(a)

	a.Name = "UpdatedName"

	assert.NoError(t, s.Item().Update(a))

	updatedA, err := s.Item().Find(a.ID)

	assert.NoError(t, err)
	assert.Equal(t, a, updatedA)
}

func TestItemRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ItemsT, sqlstore.UsersT)

	s := sqlstore.New(db)
	a := model.TestItem(t)
	u := model.TestUser(t)

	assert.NoError(t, s.Item().Delete(a.ID))

	_ = s.User().Create(u)
	a.UserId = u.ID
	_ = s.Item().Create(a)
	a, _ = s.Item().Find(a.ID)
	assert.NotNil(t, a)

	assert.NoError(t, s.Item().Delete(a.ID))
	deletedI, err := s.Item().Find(a.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedI)
}
