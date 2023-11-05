package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SkillCategory ...
type SkillCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Skill ...
type Skill struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	CategoryId int    `json:"category_id"`
}

// Validate ...
func (sc *SkillCategory) Validate() error {
	return validation.ValidateStruct(
		sc,
		validation.Field(&sc.Name, validation.Required),
	)
}

// Validate ...
func (s *Skill) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.CategoryId, validation.Required, validation.Min(0)),
	)
}
