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

	assert.NoError(t, s.Skill().Create(sk))
	assert.NotNil(t, sk)

	sk.CategoryId = 10
	assert.EqualError(t, s.Skill().Create(sk), store.ErrorNotExistRef.Error())
}

func TestSkillRepository_CreateCategory(t *testing.T) {
	s := teststore.New()
	c := model.TestSkillCategory(t)

	assert.NoError(t, s.Skill().CreateCategory(c))
	assert.NotNil(t, c)
}

func TestSkillRepository_Get(t *testing.T) {
	s := teststore.New()
	skill := model.TestSkill(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	skill.UserId = u.ID

	assert.NoError(t, s.Skill().Create(skill))
	assert.NotNil(t, skill)

	skill.Name = "test2"
	assert.NoError(t, s.Skill().Create(skill))

	skills, err := s.Skill().Get()
	assert.NoError(t, err)
	assert.NotNil(t, skills)

	assert.Equal(t, 2, len(skills))
}

func TestSkillRepository_GetByCategory(t *testing.T) {
	s := teststore.New()
	skill := model.TestSkill(t)
	c := model.TestSkillCategory(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	skill.UserId = u.ID
	c.UserId = u.ID

	assert.NoError(t, s.Skill().CreateCategory(c))
	assert.NotNil(t, c)
	skill.CategoryId = c.ID

	assert.NoError(t, s.Skill().Create(skill))
	assert.NotNil(t, skill)

	skill2 := model.TestSkill(t)
	skill2.Name = "test2"
	skill2.CategoryId = 0
	assert.NoError(t, s.Skill().Create(skill2))

	skills, err := s.Skill().GetByCategory(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, skills)

	assert.Equal(t, 1, len(skills))
}

func TestSkillRepository_GetCategories(t *testing.T) {
	s := teststore.New()
	c := model.TestSkillCategory(t)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	c.UserId = u.ID

	assert.NoError(t, s.Skill().CreateCategory(c))
	assert.NotNil(t, c)

	c.Name = "test2"
	assert.NoError(t, s.Skill().CreateCategory(c))

	categories, err := s.Skill().GetCategories()
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	assert.Equal(t, 2, len(categories))
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

func TestSkillRepository_Update(t *testing.T) {
	s := teststore.New()

	skill := model.TestSkill(t)
	_ = s.Skill().Create(skill)

	category := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(category)

	icon := "NewIcon"
	name := "NewName"
	skill.Icon = icon
	skill.Name = name
	skill.ID = category.ID

	err := s.Skill().Update(skill)
	assert.NoError(t, err)

	var updatedSkill *model.Skill
	updatedSkill, err = s.Skill().Find(skill.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedSkill)
	assert.Equal(t, updatedSkill.Icon, icon)
	assert.Equal(t, updatedSkill.Name, name)
	assert.Equal(t, updatedSkill.ID, category.ID)
}

func TestSkillRepository_UpdateCategory(t *testing.T) {
	s := teststore.New()

	c := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(c)

	icon := "NewIcon"
	name := "NewName"
	c.Icon = icon
	c.Name = name

	err := s.Skill().UpdateCategory(c)
	assert.NoError(t, err)

	var updatedCategory *model.SkillCategory
	updatedCategory, err = s.Skill().FindCategory(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedCategory)
	assert.Equal(t, updatedCategory.Icon, icon)
	assert.Equal(t, updatedCategory.Name, name)
}

func TestSkillRepository_Delete(t *testing.T) {
	s := teststore.New()

	skill := model.TestSkill(t)
	_ = s.Skill().Create(skill)

	err := s.Skill().Delete(skill.ID)
	assert.NoError(t, err)

	_, err = s.Skill().Find(skill.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}

func TestSkillRepository_DeleteCategory(t *testing.T) {
	s := teststore.New()

	category := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(category)

	err := s.Skill().DeleteCategory(category.ID)
	assert.NoError(t, err)

	_, err = s.Skill().FindCategory(category.ID)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())
}
