package model

import validation "github.com/go-ozzo/ozzo-validation"

// System ...
type System struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Validate ...
func (sc *System) Validate() error {
	return validation.ValidateStruct(
		sc,
		validation.Field(&sc.Name, validation.Required, validation.Length(6, 255)),
		validation.Field(&sc.Icon, validation.Length(1, 1024)),
	)
}
