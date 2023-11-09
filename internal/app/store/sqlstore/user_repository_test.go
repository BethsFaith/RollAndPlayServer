package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.UsersT)

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.UsersT)

	s := sqlstore.New(db)

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	_ = s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.UsersT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	u := model.TestUser(t)
	_ = s.User().Create(u)

	u, err = s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
