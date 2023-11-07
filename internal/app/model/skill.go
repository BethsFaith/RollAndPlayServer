package model

import (
	"database/sql"
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
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Icon          string        `json:"icon"`
	CategoryId    int           `json:"category_id"`
	RefCategoryId sql.NullInt64 `json:"-"`
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
		validation.Field(&s.CategoryId, validation.Min(0)),
	)
}

func (s *Skill) BeforeCreate() error {
	if s.CategoryId > 0 {
		err := s.RefCategoryId.Scan(s.CategoryId)
		return err
	} else {
		err := s.RefCategoryId.Scan(nil)
		return err
	}
}

func (s *Skill) AfterScan() {
	getDefaultOrValue(0, s.RefCategoryId)
}
