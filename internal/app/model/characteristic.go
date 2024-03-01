package model

import validation "github.com/go-ozzo/ozzo-validation"

// Characteristic ...
type Characteristic struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserId int    `json:"user_id"`
}

// Validate ...
func (c *Characteristic) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.UserId, validation.Required, validation.Min(0)),
	)
}
