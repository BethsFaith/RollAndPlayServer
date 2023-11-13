package store

// Store ...
type Store interface {
	User() UserRepository
	Skill() SkillRepository
	Race() RaceRepository
	Action() ActionRepository
}
