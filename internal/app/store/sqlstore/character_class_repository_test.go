package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClassRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassT, sqlstore.UsersT)

	s := sqlstore.New(db)
	characterClass := model.TestCharacterClass(t)

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	characterClass.UserId = u.ID

	assert.NoError(t, s.CharacterClass().Create(characterClass))
	assert.NotNil(t, characterClass)
}

func TestCharacterClassRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassT, sqlstore.UsersT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.CharacterClass().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	a := model.TestCharacterClass(t)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	a.UserId = u.ID
	_ = s.CharacterClass().Create(a)

	a, err = s.CharacterClass().Find(a.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestCharacterClassRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassT, sqlstore.UsersT)

	s := sqlstore.New(db)
	characterClass := model.TestCharacterClass(t)

	assert.NoError(t, s.CharacterClass().Update(characterClass))

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	characterClass.UserId = u.ID
	_ = s.CharacterClass().Create(characterClass)
	characterClass.Name = "UpdatedName"

	assert.NoError(t, s.CharacterClass().Update(characterClass))

	updatedCharacterClass, err := s.CharacterClass().Find(characterClass.ID)

	assert.NoError(t, err)
	assert.Equal(t, characterClass, updatedCharacterClass)
}

func TestCharacterClassRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.CharacterClassT, sqlstore.UsersT)

	s := sqlstore.New(db)
	characterClass := model.TestCharacterClass(t)

	assert.NoError(t, s.CharacterClass().Delete(characterClass.ID))

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	characterClass.UserId = u.ID
	_ = s.CharacterClass().Create(characterClass)
	characterClass, _ = s.CharacterClass().Find(characterClass.ID)
	assert.NotNil(t, characterClass)

	assert.NoError(t, s.CharacterClass().Delete(characterClass.ID))
	deletedClass, err := s.CharacterClass().Find(characterClass.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedClass)
}
