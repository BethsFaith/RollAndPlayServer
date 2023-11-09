package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkillCategory_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.SkillCategory
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.SkillCategory {
				return model.TestSkillCategory(t)
			},
			isValid: true,
		},
		{
			name: "empty name and icon",
			s: func() *model.SkillCategory {
				s := model.TestSkillCategory(t)
				s.Name = ""
				s.Icon = ""
				return s
			},
			isValid: false,
		},
		{
			name: "empty icon",
			s: func() *model.SkillCategory {
				s := model.TestSkillCategory(t)
				s.Name = "SkillCategory"
				s.Icon = ""
				return s
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}

func TestSkill_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.Skill
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Skill {
				s := model.TestSkill(t)
				return s
			},
			isValid: true,
		},
		{
			name: "no valid category",
			s: func() *model.Skill {
				s := model.TestSkill(t)
				s.CategoryId = -2
				return s
			},
			isValid: false,
		},
		{
			name: "no valid name",
			s: func() *model.Skill {
				s := model.TestSkill(t)
				s.Name = ""
				return s
			},
			isValid: false,
		},
		{
			name: "empty icon",
			s: func() *model.Skill {
				s := model.TestSkill(t)
				s.Icon = ""
				return s
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
