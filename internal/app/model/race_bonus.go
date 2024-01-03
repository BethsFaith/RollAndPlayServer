package model

import validation "github.com/go-ozzo/ozzo-validation"

type RaceBonus struct {
	RaceId  int `json:"race_id"`
	SkillId int `json:"skill_id"`
	Bonus   int `json:"bonus"`
}

func (rb *RaceBonus) Validate() error {
	return validation.ValidateStruct(
		rb,
		validation.Field(&rb.SkillId, validation.Required, validation.Min(1)),
		validation.Field(&rb.RaceId, validation.Required, validation.Min(1)),
		validation.Field(&rb.Bonus, validation.Required, validation.Min(0)),
	)
}
