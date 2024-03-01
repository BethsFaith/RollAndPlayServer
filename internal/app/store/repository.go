package store

import "RnpServer/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	Update(u *model.User) error
}

// SkillRepository ...
type SkillRepository interface {
	Create(*model.Skill) error
	CreateCategory(*model.SkillCategory) error
	Get() ([]*model.Skill, error)
	GetByCategory(id int) ([]*model.Skill, error)
	GetCategories() ([]*model.SkillCategory, error)
	Find(int) (*model.Skill, error)
	FindCategory(int) (*model.SkillCategory, error)
	Update(*model.Skill) error
	UpdateCategory(category *model.SkillCategory) error
	Delete(int) error
	DeleteCategory(int) error
}

type RaceRepository interface {
	Create(*model.Race) error
	Get() ([]*model.Race, error)
	Find(int) (*model.Race, error)
	Update(*model.Race) error
	Delete(int) error
}

type ActionRepository interface {
	Create(*model.Action) error
	Get() ([]*model.Action, error)
	Find(int) (*model.Action, error)
	Update(*model.Action) error
	Delete(int) error
}

type CharacterClassRepository interface {
	Create(*model.CharacterClass) error
	Get() ([]*model.CharacterClass, error)
	Find(int) (*model.CharacterClass, error)
	Update(*model.CharacterClass) error
	Delete(int) error
}

type RaceBonusRepository interface {
	Create(*model.RaceBonus) error
	Find(int, int) (*model.RaceBonus, error)
	FindBySkillId(int) ([]*model.RaceBonus, error)
	FindByRaceId(int) ([]*model.RaceBonus, error)
	Update(*model.RaceBonus) error
	Delete(int, int) error
}

type CharacterClassBonusRepository interface {
	Create(*model.CharacterClassBonus) error
	Find(int, int) (*model.CharacterClassBonus, error)
	FindByClassId(int) ([]*model.CharacterClassBonus, error)
	FindBySkillId(skillId int) ([]*model.CharacterClassBonus, error)
	Update(*model.CharacterClassBonus) error
	Delete(int, int) error
}

type SystemRepository interface {
	Create(*model.System) error
	Find(int) (*model.System, error)
	GetRaces(int) ([]*model.SystemComponent, error)
	GetSkillCategories(int) ([]*model.SystemComponent, error)
	GetCharacterClasses(int) ([]*model.SystemComponent, error)
	AddRace(int, int) ([]*model.SystemComponent, error)
	AddSkillCategory(int, int) ([]*model.SystemComponent, error)
	AddCharacterClass(int, int) ([]*model.SystemComponent, error)
	Update(*model.System) error
	Delete(int) error
	DeleteRace(int, int) error
	DeleteCharacterClass(int, int) error
	DeleteSkillCategory(int, int) error
}

type CharacteristicRepository interface {
	Create(characteristic *model.Characteristic) error
	Find(int) (*model.Characteristic, error)
	Get() ([]*model.Characteristic, error)
	Update(*model.Characteristic) error
	Delete(int) error
}
