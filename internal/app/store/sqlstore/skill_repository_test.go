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

	defer teardown("skills")

	s := sqlstore.New(db)
	category := model.TestSkillCategory(t)
	skill := model.TestSkill(t)
	skill.CategoryId = category.ID

	assert.NoError(t, s.Skill().CreateCategory(category))
	assert.NoError(t, s.Skill().CreateSkill(skill))
	assert.NotNil(t, skill)
}

func TestSkillRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("skills")

	s := sqlstore.New(db)

	id := 1
	_, err := s.Skill().FindSkill(id)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	category := model.TestSkillCategory(t)
	_ = s.Skill().CreateCategory(category)

	skill := model.TestSkill(t)
	skill.CategoryId = category.ID
	_ = s.Skill().CreateSkill(skill)

	skill, err = s.Skill().FindSkill(skill.ID)
	assert.NoError(t, err)
	assert.NotNil(t, skill)
}
