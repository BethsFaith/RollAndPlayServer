package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaceBonus_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		rb      func() *model.RaceBonus
		isValid bool
	}{
		{
			name: "valid",
			rb: func() *model.RaceBonus {
				return model.TestRaceBonus(t)
			},
			isValid: true,
		},
		{
			name: "no valid skill id",
			rb: func() *model.RaceBonus {
				c := model.TestRaceBonus(t)
				c.SkillId = 0
				return c
			},
			isValid: false,
		},
		{
			name: "no valid class id",
			rb: func() *model.RaceBonus {
				c := model.TestRaceBonus(t)
				c.RaceId = 0
				return c
			},
			isValid: false,
		},
		{
			name: "no valid bonus",
			rb: func() *model.RaceBonus {
				c := model.TestRaceBonus(t)
				c.Bonus = 0
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.rb().Validate())
			} else {
				assert.Error(t, tc.rb().Validate())
			}
		})
	}
}
