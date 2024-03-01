package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacteristicRepository_Create(t *testing.T) {
	s := teststore.New()
	c := model.TestCharacteristic(t)

	assert.NoError(t, s.Characteristic().Create(c))
	assert.NotNil(t, c)
}

func TestCharacteristicRepository_Get(t *testing.T) {
	s := teststore.New()

	c := model.TestCharacteristic(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	c.UserId = u.ID

	assert.NoError(t, s.Characteristic().Create(c))
	assert.NotNil(t, c)

	a2 := *c
	a2.Name = "test2"
	assert.NoError(t, s.Characteristic().Create(&a2))

	characteristics, err := s.Characteristic().Get()
	assert.NoError(t, err)
	assert.NotNil(t, characteristics)
	assert.Equal(t, *c, *characteristics[0])
	assert.Equal(t, a2, *characteristics[1])
}

func TestCharacteristicRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.Characteristic().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestCharacteristic(t)
	_ = s.Characteristic().Create(c)

	c, err = s.Characteristic().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCharacteristicRepository_Update(t *testing.T) {
	s := teststore.New()

	c := model.TestCharacteristic(t)
	_ = s.Characteristic().Create(c)

	name := "NewName"
	icon := "NewIconPath"

	c.Name = name
	c.Icon = icon

	err := s.Characteristic().Update(c)
	assert.NoError(t, err)

	var updatedC *model.Characteristic
	updatedC, err = s.Characteristic().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedC)
	assert.Equal(t, updatedC.Icon, icon)
	assert.Equal(t, updatedC.Name, name)
}

func TestCharacteristicRepository_Delete(t *testing.T) {
	s := teststore.New()

	a := model.TestCharacteristic(t)
	_ = s.Characteristic().Create(a)

	err := s.Characteristic().Delete(a.ID)
	assert.NoError(t, err)

	_, err = s.Characteristic().Find(a.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
