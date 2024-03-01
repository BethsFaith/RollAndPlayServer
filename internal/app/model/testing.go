package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestCharacteristic(t *testing.T) *Characteristic {
	t.Helper()
	return &Characteristic{
		Name:   "skill",
		Icon:   "path/icon",
		UserId: 1,
	}
}

func TestSkill(t *testing.T) *Skill {
	t.Helper()
	return &Skill{
		Name:       "skill",
		Icon:       "path/icon",
		CategoryId: -1,
		UserId:     1,
	}
}

func TestSkillCategory(t *testing.T) *SkillCategory {
	t.Helper()
	return &SkillCategory{
		Name:   "skill",
		Icon:   "path/icon",
		UserId: 1,
	}
}

func TestRace(t *testing.T) *Race {
	t.Helper()
	return &Race{
		Name:   "race",
		Model:  "path/model",
		UserId: 1,
	}
}

func TestAction(t *testing.T) *Action {
	t.Helper()
	return &Action{
		Name:    "race",
		Icon:    "path/icon",
		SkillId: 0,
		Points:  0,
		UserId:  1,
	}
}

func TestCharacterClass(t *testing.T) *CharacterClass {
	t.Helper()
	return &CharacterClass{
		Name:   "race",
		Icon:   "path/icon",
		UserId: 1,
	}
}

func TestRaceBonus(t *testing.T) *RaceBonus {
	t.Helper()
	return &RaceBonus{
		RaceId:  1,
		SkillId: 1,
		Bonus:   1,
	}
}

func TestCharacterClassBonus(t *testing.T) *CharacterClassBonus {
	t.Helper()
	return &CharacterClassBonus{
		ClassId: 1,
		SkillId: 1,
		Bonus:   1,
	}
}

func TestSystem(t *testing.T) *System {
	t.Helper()
	return &System{
		Name: "system",
		Icon: "path/icon",
	}
}
