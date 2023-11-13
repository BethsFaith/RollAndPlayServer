package store

import "RnpServer/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

// SkillRepository ...
type SkillRepository interface {
	Create(*model.Skill) error
	CreateCategory(*model.SkillCategory) error
	Find(int) (*model.Skill, error)
	FindCategory(int) (*model.SkillCategory, error)
	Update(*model.Skill) error
	UpdateCategory(category *model.SkillCategory) error
	Delete(id int) error
	DeleteCategory(id int) error
}

type RaceRepository interface {
	Create(*model.Race) error
	Find(int) (*model.Race, error)
	Update(*model.Race) error
	Delete(id int) error
}

type ActionRepository interface {
	Create(*model.Action) error
	Find(int) (*model.Action, error)
	Update(*model.Action) error
	Delete(id int) error
}
