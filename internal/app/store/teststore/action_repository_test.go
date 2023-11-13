package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActionRepository_Create(t *testing.T) {
	s := teststore.New()
	a := model.TestAction(t)

	assert.NoError(t, s.Action().Create(a))
	assert.NotNil(t, a)
}

func TestActionRepository_Find(t *testing.T) {
	s := teststore.New()

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
	s := teststore.New()

	a := model.TestAction(t)
	_ = s.Action().Create(a)

	skill := model.TestSkill(t)
	_ = s.Skill().Create(skill)

	name := "NewName"
	icon := "NewIconPath"
	points := 0

	a.Name = name
	a.Icon = icon
	a.Points = points
	a.SkillId = skill.ID

	err := s.Action().Update(a)
	assert.NoError(t, err)

	var updatedA *model.Action
	updatedA, err = s.Action().Find(a.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedA)
	assert.Equal(t, updatedA.Icon, icon)
	assert.Equal(t, updatedA.Name, name)
	assert.Equal(t, updatedA.SkillId, skill.ID)
	assert.Equal(t, updatedA.Points, points)
}

func TestActionRepository_Delete(t *testing.T) {
	s := teststore.New()

	a := model.TestAction(t)
	_ = s.Action().Create(a)

	err := s.Action().Delete(a.ID)
	assert.NoError(t, err)

	_, err = s.Action().Find(a.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
