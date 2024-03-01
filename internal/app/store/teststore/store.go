package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type Store struct {
	userRepository                *UserRepository
	characteristicRepository      *CharacteristicRepository
	skillRepository               *SkillRepository
	raceRepository                *RaceRepository
	actionRepository              *ActionRepository
	classRepository               *CharacterClassRepository
	raceBonusRepository           *RaceBonusRepository
	characterClassBonusRepository *CharacterClassBonusRepository
	systemRepository              *SystemRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Characteristic() store.CharacteristicRepository {
	if s.characteristicRepository != nil {
		return s.characteristicRepository
	}

	s.characteristicRepository = &CharacteristicRepository{
		store:           s,
		characteristics: make(map[int]*model.Characteristic),
	}

	return s.characteristicRepository
}

func (s *Store) Skill() store.SkillRepository {
	if s.skillRepository != nil {
		return s.skillRepository
	}

	s.skillRepository = &SkillRepository{
		store:      s,
		skills:     make(map[int]*model.Skill),
		categories: make(map[int]*model.SkillCategory),
	}

	return s.skillRepository
}

func (s *Store) Race() store.RaceRepository {
	if s.raceRepository != nil {
		return s.raceRepository
	}

	s.raceRepository = &RaceRepository{
		store: s,
		races: make(map[int]*model.Race),
	}

	return s.raceRepository
}

func (s *Store) Action() store.ActionRepository {
	if s.actionRepository != nil {
		return s.actionRepository
	}

	s.actionRepository = &ActionRepository{
		store:   s,
		actions: make(map[int]*model.Action),
	}

	return s.actionRepository
}

func (s *Store) CharacterClass() store.CharacterClassRepository {
	if s.classRepository != nil {
		return s.classRepository
	}

	s.classRepository = &CharacterClassRepository{
		store:   s,
		classes: make(map[int]*model.CharacterClass),
	}

	return s.classRepository
}

func (s *Store) RaceBonus() store.RaceBonusRepository {
	if s.raceBonusRepository != nil {
		return s.raceBonusRepository
	}

	s.raceBonusRepository = &RaceBonusRepository{
		store:   s,
		bonuses: make(map[RaceBonusKey]*model.RaceBonus),
	}

	return s.raceBonusRepository
}

func (s *Store) CharacterClassBonus() store.CharacterClassBonusRepository {
	if s.characterClassBonusRepository != nil {
		return s.characterClassBonusRepository
	}

	s.characterClassBonusRepository = &CharacterClassBonusRepository{
		store:   s,
		bonuses: make(map[CharacterClassBonusKey]*model.CharacterClassBonus),
	}

	return s.characterClassBonusRepository
}

func (s *Store) System() store.SystemRepository {
	if s.systemRepository != nil {
		return s.systemRepository
	}

	s.systemRepository = &SystemRepository{
		store:      s,
		systems:    make(map[int]*model.System),
		races:      make(map[int][]*model.SystemComponent),
		classes:    make(map[int][]*model.SystemComponent),
		categories: make(map[int][]*model.SystemComponent),
	}

	return s.systemRepository
}
