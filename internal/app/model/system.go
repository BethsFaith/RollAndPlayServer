package model

import validation "github.com/go-ozzo/ozzo-validation"

// System ...
type System struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// SystemComponent ...
type SystemComponent struct {
	SystemId    int `json:"system_id"`
	ComponentId int `json:"component_id"`
}

// Validate ...
func (s *System) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Length(2, 255)),
		validation.Field(&s.Icon, validation.Length(1, 1024)),
	)
}

// Validate ...
func (sc *SystemComponent) Validate() error {
	return validation.ValidateStruct(
		sc,
		validation.Field(&sc.SystemId, validation.Required, validation.Min(1)),
		validation.Field(&sc.ComponentId, validation.Required, validation.Min(1)),
	)
}
