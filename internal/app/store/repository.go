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
}
