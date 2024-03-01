package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClassRepository_Create(t *testing.T) {
	s := teststore.New()
	c := model.TestCharacterClass(t)

	assert.NoError(t, s.CharacterClass().Create(c))
	assert.NotNil(t, c)
}

func TestCharacterClassRepository_Get(t *testing.T) {
	s := teststore.New()

	cc := model.TestCharacterClass(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	cc.UserId = u.ID

	assert.NoError(t, s.CharacterClass().Create(cc))
	assert.NotNil(t, cc)

	cc2 := *cc
	cc2.Name = "test2"
	assert.NoError(t, s.CharacterClass().Create(&cc2))

	classes, err := s.CharacterClass().Get()
	assert.NoError(t, err)
	assert.NotNil(t, classes)
	assert.Equal(t, *cc, *classes[0])
	assert.Equal(t, cc2, *classes[1])
}

func TestCharacterClassRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.CharacterClass().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestCharacterClass(t)
	_ = s.CharacterClass().Create(c)

	c, err = s.CharacterClass().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCharacterClassRepository_Update(t *testing.T) {
	s := teststore.New()

	class := model.TestCharacterClass(t)
	_ = s.CharacterClass().Create(class)

	name := "NewName"
	icon := "NewIcon"

	class.Name = name
	class.Icon = icon

	err := s.CharacterClass().Update(class)
	assert.NoError(t, err)

	var updatedClass *model.CharacterClass
	updatedClass, err = s.CharacterClass().Find(class.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedClass)
	assert.Equal(t, updatedClass.Icon, icon)
	assert.Equal(t, updatedClass.Name, name)
}

func TestCharacterClassRepository_Delete(t *testing.T) {
	s := teststore.New()

	class := model.TestCharacterClass(t)
	_ = s.CharacterClass().Create(class)

	err := s.CharacterClass().Delete(class.ID)
	assert.NoError(t, err)

	_, err = s.CharacterClass().Find(class.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
