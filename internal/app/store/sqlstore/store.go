package sqlstore

import (
	"RnpServer/internal/app/store"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db                            *sql.DB
	userRepository                *UserRepository
	skillRepository               *SkillRepository
	raceRepository                *RaceRepository
	actionRepository              *ActionRepository
	characterClassRepository      *CharacterClassRepository
	raceBonusRepository           *RaceBonusRepository
	characterClassBonusRepository *CharacterClassBonusRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters...)
}

func (s *Store) CreateRetId(queryStr string, parameters ...any) *sql.Row {
	return s.db.QueryRow(queryStr, parameters...)
}

func (s *Store) SelectRow(queryStr string, parameters ...any) *sql.Row {
	return s.db.QueryRow(queryStr, parameters...)
}

func (s *Store) SelectRows(queryStr string, parameters ...any) (*sql.Rows, error) {
	rows, err := s.db.Query(queryStr, parameters...)
	if err != nil {
		return rows, err
	}

	return rows, err
}

func (s *Store) Update(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters...)
}

func (s *Store) Delete(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters...)
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Skill() store.SkillRepository {
	if s.skillRepository != nil {
		return s.skillRepository
	}

	s.skillRepository = &SkillRepository{
		store: s,
	}

	return s.skillRepository
}

func (s *Store) Race() store.RaceRepository {
	if s.raceRepository != nil {
		return s.raceRepository
	}

	s.raceRepository = &RaceRepository{
		store: s,
	}

	return s.raceRepository
}

func (s *Store) Action() store.ActionRepository {
	if s.actionRepository != nil {
		return s.actionRepository
	}

	s.actionRepository = &ActionRepository{
		store: s,
	}

	return s.actionRepository
}

func (s *Store) CharacterClass() store.CharacterClassRepository {
	if s.characterClassRepository != nil {
		return s.characterClassRepository
	}

	s.characterClassRepository = &CharacterClassRepository{
		store: s,
	}

	return s.characterClassRepository
}

func (s *Store) RaceBonus() store.RaceBonusRepository {
	if s.raceBonusRepository != nil {
		return s.raceBonusRepository
	}

	s.raceBonusRepository = &RaceBonusRepository{
		store: s,
	}

	return s.raceBonusRepository
}

func (s *Store) CharacterClassBonus() store.CharacterClassBonusRepository {
	if s.characterClassBonusRepository != nil {
		return s.characterClassBonusRepository
	}

	s.characterClassBonusRepository = &CharacterClassBonusRepository{
		store: s,
	}

	return s.characterClassBonusRepository
}
