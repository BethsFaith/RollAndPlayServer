package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
)

// SkillCategory ...
type SkillCategory struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserId int    `json:"user_id"`
}

// Skill ...
type Skill struct {
	ID                  int           `json:"id"`
	Name                string        `json:"name"`
	Icon                string        `json:"icon"`
	CategoryId          int           `json:"category_id"`
	RefCategoryId       sql.NullInt64 `json:"-"`
	CharacteristicId    int           `json:"characteristic_id"`
	RefCharacteristicId sql.NullInt64 `json:"-"`
	UserId              int           `json:"user_id"`
}

// Validate ...
func (sc *SkillCategory) Validate() error {
	return validation.ValidateStruct(
		sc,
		validation.Field(&sc.Name, validation.Required),
		validation.Field(&sc.UserId, validation.Required, validation.Min(1)),
	)
}

// Validate ...
func (s *Skill) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.CategoryId, validation.Min(0)),
		validation.Field(&s.CharacteristicId, validation.Min(0)),
		validation.Field(&s.UserId, validation.Required, validation.Min(1)),
	)
}

func (s *Skill) BeforeInsertOrUpdate() error {
	if s.CategoryId > 0 {
		err := s.RefCategoryId.Scan(s.CategoryId)
		if err != nil {
			return err
		}
	} else {
		err := s.RefCategoryId.Scan(nil)
		if err != nil {
			return err
		}
	}
	if s.CharacteristicId > 0 {
		err := s.RefCharacteristicId.Scan(s.CharacteristicId)
		if err != nil {
			return err
		}
	} else {
		err := s.RefCharacteristicId.Scan(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Skill) AfterScan() {
	s.CategoryId = getDefaultOrValue(0, s.RefCategoryId)
	s.CharacteristicId = getDefaultOrValue(0, s.RefCharacteristicId)
}
