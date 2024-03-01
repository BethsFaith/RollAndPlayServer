package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Action ...
type Action struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Icon       string        `json:"icon"`
	Points     int           `json:"points"`
	SkillId    int           `json:"skill_id"`
	UserId     int           `json:"user_id"`
	RefSkillId sql.NullInt64 `json:"-"`
}

// Validate ...
func (a *Action) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&a.Points, validation.Min(0)),
		validation.Field(&a.SkillId, validation.Min(0)),
		validation.Field(&a.UserId, validation.Required, validation.Min(0)),
	)
}

func (a *Action) BeforeInsertOrUpdate() error {
	if a.SkillId >= 0 {
		err := a.RefSkillId.Scan(a.SkillId)
		return err
	} else {
		err := a.RefSkillId.Scan(nil)
		return err
	}
}

func (a *Action) AfterScan() {
	a.SkillId = getDefaultOrValue(-1, a.RefSkillId)
}
