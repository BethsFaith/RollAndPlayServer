package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Item struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Icon        string        `json:"icon"`
	TypeId      int           `json:"type_id"`
	RefTypeId   sql.NullInt64 `json:"-"`
	UserId      int           `json:"user_id"`
}

// Validate ...
func (i *Item) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.TypeId, validation.Min(0)),
		validation.Field(&i.UserId, validation.Required, validation.Min(1)),
	)
}

// BeforeInsertOrUpdate ...
func (i *Item) BeforeInsertOrUpdate() error {
	var err error
	if i.TypeId > 0 {
		err = i.RefTypeId.Scan(i.TypeId)
	} else {
		err = i.RefTypeId.Scan(nil)
	}
	return err
}

// AfterScan ...
func (i *Item) AfterScan() {
	i.TypeId = getDefaultOrValue(0, i.RefTypeId)
}

type ItemType struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserId int    `json:"user_id"`
}

// Validate ...
func (i *ItemType) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.UserId, validation.Required, validation.Min(1)),
	)
}

type ItemDescriptor struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TypeId int    `json:"type_id"`
}

// Validate ...
func (i *ItemDescriptor) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.TypeId, validation.Required, validation.Min(1)),
	)
}

type ItemDescriptorLine struct {
	ID           int `json:"id"`
	DescriptorId int `json:"descriptor_id"`
	ItemId       int `json:"item_id"`
}

// Validate ...
func (i *ItemDescriptorLine) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.DescriptorId, validation.Required, validation.Min(1)),
		validation.Field(&i.ItemId, validation.Required, validation.Min(1)),
	)
}
