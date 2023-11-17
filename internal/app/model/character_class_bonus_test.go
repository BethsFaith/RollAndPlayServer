package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterClassBonus_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		cb      func() *model.CharacterClassBonus
		isValid bool
	}{
		{
			name: "valid",
			cb: func() *model.CharacterClassBonus {
				return model.TestCharacterClassBonus(t)
			},
			isValid: true,
		},
		{
			name: "no valid skill id",
			cb: func() *model.CharacterClassBonus {
				c := model.TestCharacterClassBonus(t)
				c.SkillId = 0
				return c
			},
			isValid: false,
		},
		{
			name: "no valid class id",
			cb: func() *model.CharacterClassBonus {
				c := model.TestCharacterClassBonus(t)
				c.ClassId = 0
				return c
			},
			isValid: false,
		},
		{
			name: "no valid bonus",
			cb: func() *model.CharacterClassBonus {
				c := model.TestCharacterClassBonus(t)
				c.Bonus = 0
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.cb().Validate())
			} else {
				assert.Error(t, tc.cb().Validate())
			}
		})
	}
}
