package teststore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkillRepository_Create(t *testing.T) {
	s := teststore.New()

	sk := model.TestSkill(t)
	err := s.Skill().Create(sk)
	assert.EqualError(t, err, store.ErrorNotExistRef.Error())

	c := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(c)

	sk.CategoryId = c.ID

	assert.NoError(t, s.Skill().Create(sk))
	assert.NotNil(t, sk)
}

func TestSkillRepository_CreateCategory(t *testing.T) {
	s := teststore.New()
	c := model.TestSkillCategory(t)

	assert.NoError(t, s.Skill().CreateCategory(c))
	assert.NotNil(t, c)
}

func TestSkillRepository_Find(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.Skill().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(c)

	sk := model.TestSkill(t)
	sk.ID = id
	sk.CategoryId = c.ID
	_ = s.Skill().Create(sk)

	sk, err = s.Skill().Find(id)
	assert.NoError(t, err)
	assert.NotNil(t, sk)
}

func TestSkillRepository_FindCategory(t *testing.T) {
	s := teststore.New()

	id := 1
	_, err := s.Skill().FindCategory(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	c := model.TestSkillCategory(t)
	c.ID = id
	_ = s.Skill().CreateCategory(c)

	c, err = s.Skill().FindCategory(id)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
