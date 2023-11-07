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

func TestSkillRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillsT)

	s := sqlstore.New(db)
	skill := model.TestSkill(t)

	assert.NoError(t, s.Skill().Update(skill))

	_ = s.Skill().Create(skill)
	skill.Name = "UpdatedName"

	assert.NoError(t, s.Skill().Update(skill))

	updatedSkill, err := s.Skill().Find(skill.ID)

	assert.NoError(t, err)
	assert.Equal(t, skill, updatedSkill)
}

func TestSkillRepository_UpdateCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillCategoriesT)

	s := sqlstore.New(db)
	category := model.TestSkillCategory(t)

	assert.NoError(t, s.Skill().UpdateCategory(category))

	_ = s.Skill().CreateCategory(category)
	category.Name = "UpdatedName"

	assert.NoError(t, s.Skill().UpdateCategory(category))

	updatedCategory, err := s.Skill().FindCategory(category.ID)

	assert.NoError(t, err)
	assert.Equal(t, category, updatedCategory)
}

func TestSkillRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillsT)

	s := sqlstore.New(db)
	skill := model.TestSkill(t)

	assert.NoError(t, s.Skill().Delete(skill.ID))

	_ = s.Skill().Create(skill)
	skill, _ = s.Skill().Find(skill.ID)
	assert.NotNil(t, skill)

	assert.NoError(t, s.Skill().Delete(skill.ID))
	deletedSkill, err := s.Skill().Find(skill.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedSkill)
}

func TestSkillRepository_DeleteCategory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown(sqlstore.SkillsT)

	s := sqlstore.New(db)
	category := model.TestSkillCategory(t)

	assert.NoError(t, s.Skill().DeleteCategory(category.ID))

	_ = s.Skill().CreateCategory(category)
	category, _ = s.Skill().FindCategory(category.ID)
	assert.NotNil(t, category)

	assert.NoError(t, s.Skill().DeleteCategory(category.ID))
	deletedSkill, err := s.Skill().FindCategory(category.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedSkill)
}
