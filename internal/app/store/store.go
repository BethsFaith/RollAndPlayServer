package store

// Store ...
type Store interface {
	User() UserRepository
	Characteristic() CharacteristicRepository
	Skill() SkillRepository
	Race() RaceRepository
	Action() ActionRepository
	CharacterClass() CharacterClassRepository
	RaceBonus() RaceBonusRepository
	CharacterClassBonus() CharacterClassBonusRepository
	System() SystemRepository
}
