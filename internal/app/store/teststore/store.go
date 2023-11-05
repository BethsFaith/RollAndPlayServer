package teststore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
)

type Store struct {
	userRepository  *UserRepository
	skillRepository *SkillRepository
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
