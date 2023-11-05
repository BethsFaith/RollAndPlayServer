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
		Icon:       "10101001",
		CategoryId: 1,
	}
}

func TestSkillCategory(t *testing.T) *SkillCategory {
	t.Helper()
	return &SkillCategory{
		Name: "skill",
		Icon: "10101001",
	}
}
