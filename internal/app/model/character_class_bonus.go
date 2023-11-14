package model

import validation "github.com/go-ozzo/ozzo-validation"

type CharacterClassBonus struct {
	ClassId int `json:"class_id"`
	SkillId int `json:"skill_id"`
	Bonus   int `json:"bonus"`
}

func (cb *CharacterClassBonus) Validate() error {
	return validation.ValidateStruct(
		cb,
		validation.Field(&cb.SkillId, validation.Required, validation.Min(1)),
		validation.Field(&cb.ClassId, validation.Required, validation.Min(1)),
		validation.Field(&cb.Bonus, validation.Required, validation.Min(0)),
	)
}
