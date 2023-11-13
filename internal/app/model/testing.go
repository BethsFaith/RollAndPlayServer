package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestSkill(t *testing.T) *Skill {
	t.Helper()
	return &Skill{
		Name:       "skill",
		Icon:       "path/icon",
		CategoryId: 0,
	}
}

func TestSkillCategory(t *testing.T) *SkillCategory {
	t.Helper()
	return &SkillCategory{
		Name: "skill",
		Icon: "path/icon",
	}
}

func TestRace(t *testing.T) *Race {
	t.Helper()
	return &Race{
		Name:  "race",
		Model: "path/model",
	}
}

func TestAction(t *testing.T) *Action {
	t.Helper()
	return &Action{
		Name:    "race",
		Icon:    "path/icon",
		SkillId: 0,
		Points:  0,
	}
}
