package model_test

import (
	"RnpServer/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystem_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.System
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.System {
				s := model.TestSystem(t)
				return s
			},
			isValid: true,
		},
		{
			name: "no valid name",
			s: func() *model.System {
				s := model.TestSystem(t)
				s.Name = "d"
				return s
			},
			isValid: false,
		},
		{
			name: "empty icon",
			s: func() *model.System {
				s := model.TestSystem(t)
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

func TestSystemComponent_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.SystemComponent
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.SystemComponent {
				s := &model.SystemComponent{
					SystemId:    1,
					ComponentId: 1,
				}
				return s
			},
			isValid: true,
		},
		{
			name: "no valid system id",
			s: func() *model.SystemComponent {
				s := &model.SystemComponent{
					SystemId:    -1,
					ComponentId: 1,
				}
				return s
			},
			isValid: false,
		},
		{
			name: "no valid component id",
			s: func() *model.SystemComponent {
				s := &model.SystemComponent{
					SystemId:    1,
					ComponentId: -1,
				}
				return s
			},
			isValid: false,
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
