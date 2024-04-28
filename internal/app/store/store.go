package store

// Store ...
type Store interface {
	User() UserRepository
	Characteristic() CharacteristicRepository
	Skill() SkillRepository
	Race() RaceRepository
	Action() ActionRepository
	Item() ItemRepository
	CharacterClass() CharacterClassRepository
	RaceBonus() RaceBonusRepository
	CharacterClassBonus() CharacterClassBonusRepository
	System() SystemRepository
}
