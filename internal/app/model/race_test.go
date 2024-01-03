package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRace_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		r       func() *model.Race
		isValid bool
	}{
		{
			name: "valid",
			r: func() *model.Race {
				r := model.TestRace(t)
				return r
			},
			isValid: true,
		},
		{
			name: "no valid name",
			r: func() *model.Race {
				r := model.TestRace(t)
				r.Name = ""
				return r
			},
			isValid: false,
		},
		{
			name: "empty model",
			r: func() *model.Race {
				r := model.TestRace(t)
				r.Model = ""
				return r
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.r().Validate())
			} else {
				assert.Error(t, tc.r().Validate())
			}
		})
	}
}
