package sqlstore_test

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"RnpServer/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkillRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillsT)

	s := sqlstore.New(db)
	skill := model.TestSkill(t)

	assert.NoError(t, s.Skill().Create(skill))
	assert.NotNil(t, skill)

	skill.CategoryId = -1
	assert.Error(t, s.Skill().Create(skill))

	skill.CategoryId = 10
	assert.EqualError(t, s.Skill().Create(skill), store.ErrorNotExistRef.Error())
}

func TestSkillRepository_CreateCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillCategoriesT)

	s := sqlstore.New(db)
	cat := model.TestSkillCategory(t)

	assert.NoError(t, s.Skill().CreateCategory(cat))
	assert.NotNil(t, cat)
}

func TestSkillRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillsT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Skill().Find(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	skill := model.TestSkill(t)
	_ = s.Skill().Create(skill)

	skill, err = s.Skill().Find(skill.ID)
	assert.NoError(t, err)
	assert.NotNil(t, skill)
}

func TestSkillRepository_FindCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillCategoriesT)

	s := sqlstore.New(db)

	id := 1
	_, err := s.Skill().FindCategory(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	cat := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(cat)
	id = cat.ID

	cat, err = s.Skill().FindCategory(id)
	assert.NoError(t, err)
	assert.NotNil(t, cat)
}
