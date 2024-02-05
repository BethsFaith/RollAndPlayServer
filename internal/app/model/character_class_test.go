package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClass_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		c       func() *model.CharacterClass
		isValid bool
	}{
		{
			name: "valid",
			c: func() *model.CharacterClass {
				c := model.TestCharacterClass(t)
				return c
			},
			isValid: true,
		},
		{
			name: "no valid name",
			c: func() *model.CharacterClass {
				c := model.TestCharacterClass(t)
				c.Name = ""
				return c
			},
			isValid: false,
		},
		{
			name: "empty icon",
			c: func() *model.CharacterClass {
				c := model.TestCharacterClass(t)
				c.Icon = ""
				return c
			},
			isValid: true,
		},
		{
			name: "not valid user id",
			c: func() *model.CharacterClass {
				c := model.TestCharacterClass(t)
				c.UserId = 0
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.c().Validate())
			} else {
				assert.Error(t, tc.c().Validate())
			}
		})
	}
}
