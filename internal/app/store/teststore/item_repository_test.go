package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemRepository_Create(t *testing.T) {
	s := teststore.New()
	r := model.TestItem(t)

	assert.NoError(t, s.Item().Create(r))
	assert.NotNil(t, r)
}

func TestItemRepository_Get(t *testing.T) {
	s := teststore.New()

	r := model.TestItem(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	r.UserId = u.ID

	assert.NoError(t, s.Item().Create(r))
	assert.NotNil(t, r)

	r2 := model.TestItem(t)
	r2.Name = "test2"
	assert.NoError(t, s.Item().Create(r2))

	Items, err := s.Item().Get()
	assert.NoError(t, err)
	assert.NotNil(t, Items)
	assert.Equal(t, r, Items[0])
	assert.Equal(t, r2, Items[1])
}

func TestItemRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.Item().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	r := model.TestItem(t)
	_ = s.Item().Create(r)

	r, err = s.Item().Find(r.ID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestItemRepository_Update(t *testing.T) {
	s := teststore.New()

	Item := model.TestItem(t)
	_ = s.Item().Create(Item)

	name := "NewName"
	Item.Name = name

	err := s.Item().Update(Item)
	assert.NoError(t, err)

	var updatedItem *model.Item
	updatedItem, err = s.Item().Find(Item.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedItem)
	assert.Equal(t, updatedItem.Name, name)
}

func TestItemRepository_Delete(t *testing.T) {
	s := teststore.New()

	Item := model.TestItem(t)
	_ = s.Item().Create(Item)

	err := s.Item().Delete(Item.ID)
	assert.NoError(t, err)

	_, err = s.Item().Find(Item.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
