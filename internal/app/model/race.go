package model

import validation "github.com/go-ozzo/ozzo-validation"

// Race ...
type Race struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	UserId int    `json:"user_id"`
}

// Validate ...
func (r *Race) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&r.UserId, validation.Required, validation.Min(1)),
	)
}
