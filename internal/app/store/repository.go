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

type CharacterClassRepository interface {
	Create(*model.CharacterClass) error
	Find(int) (*model.CharacterClass, error)
	Update(*model.CharacterClass) error
	Delete(id int) error
}

type RaceBonusRepository interface {
	Create(*model.RaceBonus) error
	Find(int, int) (*model.RaceBonus, error)
	FindByRaceId(int) ([]*model.RaceBonus, error)
	Update(*model.RaceBonus) error
	Delete(int, int) error
}

type CharacterClassBonusRepository interface {
	Create(*model.CharacterClassBonus) error
	Find(int, int) (*model.CharacterClassBonus, error)
	FindByClassId(int) ([]*model.CharacterClassBonus, error)
	Update(*model.CharacterClassBonus) error
	Delete(int, int) error
}

type SystemRepository interface {
	Create(*model.System) error
	Find(int) (*model.System, error)
	AddRace(int, int) ([]*model.Race, error)
	AddSkillCategory(int, int) ([]*model.SkillCategory, error)
	AddCharacterClass(int, int) ([]*model.CharacterClass, error)
	Update(*model.System) error
	Delete(int) error
}
