package model

import validation "github.com/go-ozzo/ozzo-validation"

type CharacterClass struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Validate ...
func (c *CharacterClass) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 100)),
	)
}
