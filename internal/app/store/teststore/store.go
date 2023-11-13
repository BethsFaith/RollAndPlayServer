package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type Store struct {
	userRepository  *UserRepository
	skillRepository *SkillRepository
	raceRepository  *RaceRepository
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
