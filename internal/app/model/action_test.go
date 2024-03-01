package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAction_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		a       func() *model.Action
		isValid bool
	}{
		{
			name: "valid",
			a: func() *model.Action {
				return model.TestAction(t)
			},
			isValid: true,
		},
		{
			name: "empty icon",
			a: func() *model.Action {
				a := model.TestAction(t)
				a.Icon = ""

				return a
			},
			isValid: true,
		},
		{
			name: "empty name and icon",
			a: func() *model.Action {
				a := model.TestAction(t)
				a.Name = ""
				a.Icon = ""

				return a
			},
			isValid: false,
		},
		{
			name: "not valid skill id",
			a: func() *model.Action {
				a := model.TestAction(t)
				a.SkillId = -1

				return a
			},
			isValid: false,
		},
		{
			name: "not valid points",
			a: func() *model.Action {
				a := model.TestAction(t)
				a.Points = -1

				return a
			},
			isValid: false,
		},
		{
			name: "not valid user id",
			a: func() *model.Action {
				a := model.TestAction(t)
				a.UserId = 0

				return a
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.a().Validate())
			} else {
				assert.Error(t, tc.a().Validate())
			}
		})
	}
}

func TestAction_BeforeInsertOrUpdate(t *testing.T) {
	a := model.TestAction(t)
	assert.NoError(t, a.BeforeInsertOrUpdate())
}
