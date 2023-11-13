package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActionRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ActionsT)

	s := sqlstore.New(db)
	a := model.TestAction(t)

	assert.NoError(t, s.Action().Create(a))
	assert.NotNil(t, a)

	a.SkillId = -1
	assert.Error(t, s.Action().Create(a))

	a.SkillId = 10
	assert.Error(t, s.Action().Create(a))
}

func TestActionRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ActionsT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Action().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	a := model.TestAction(t)
	_ = s.Action().Create(a)

	a, err = s.Action().Find(a.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestActionRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ActionsT)

	s := sqlstore.New(db)
	a := model.TestAction(t)

	assert.NoError(t, s.Action().Update(a))

	_ = s.Action().Create(a)
	a.Name = "UpdatedName"

	assert.NoError(t, s.Action().Update(a))

	updatedA, err := s.Action().Find(a.ID)

	assert.NoError(t, err)
	assert.Equal(t, a, updatedA)
}

func TestActionRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.ActionsT)

	s := sqlstore.New(db)
	a := model.TestAction(t)

	assert.NoError(t, s.Action().Delete(a.ID))

	_ = s.Action().Create(a)
	a, _ = s.Action().Find(a.ID)
	assert.NotNil(t, a)

	assert.NoError(t, s.Action().Delete(a.ID))
	deletedA, err := s.Action().Find(a.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedA)
}
