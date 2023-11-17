package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store/teststore"
	"RnpServer/internal/log"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleRaceCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		race         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			race: map[string]interface{}{
				"name":  "Name",
				"model": "Model.path",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			race: map[string]interface{}{
				"name":  "",
				"model": "Model.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			race: map[string]interface{}{
				"model": "Model.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			race: map[string]interface{}{
				"name": "Name",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.race)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/races", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	race := model.TestRace(t)
	_ = store.Race().Create(race)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		race         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			race: map[string]interface{}{
				"id":    race.ID,
				"name":  "NewName",
				"model": "NewModel.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			race: map[string]interface{}{
				"id":    race.ID,
				"name":  "NewName",
				"model": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			race: map[string]interface{}{
				"id":    race.ID,
				"name":  "",
				"model": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			race: map[string]interface{}{
				"id":    race.ID,
				"name":  "11",
				"model": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			race: map[string]interface{}{
				"id":    1000,
				"name":  "11",
				"model": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.race)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/races", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	race := model.TestRace(t)
	_ = store.Race().Create(race)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		race         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			race: map[string]interface{}{
				"id": race.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			race: map[string]interface{}{
				"id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.race)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/races", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceBonusCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	race := model.TestRace(t)
	skill := model.TestSkill(t)

	_ = store.Race().Create(race)
	_ = store.Skill().Create(skill)

	bonus := model.TestRaceBonus(t)
	bonus.RaceId = race.ID
	bonus.SkillId = skill.ID

	testCases := []struct {
		name         string
		bonus        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": bonus.SkillId,
				"bonus":    bonus.Bonus,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": bonus.SkillId,
				"bonus":    0,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": -1,
				"bonus":    bonus.Bonus,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id":  -1,
				"skill_id": bonus.SkillId,
				"bonus":    bonus.Bonus,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.bonus)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/races/bonuses", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceBonuses(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	race := model.TestRace(t)
	skill := model.TestSkill(t)
	skill2 := model.TestSkill(t)
	skill2.Name = "nAME2"

	_ = store.Race().Create(race)
	_ = store.Skill().Create(skill)
	_ = store.Skill().Create(skill2)

	bonus := model.TestRaceBonus(t)
	bonus2 := model.TestRaceBonus(t)
	bonus.RaceId = race.ID
	bonus.SkillId = skill.ID
	bonus2.RaceId = race.ID
	bonus2.SkillId = skill2.ID
	_ = store.RaceBonus().Create(bonus)
	_ = store.RaceBonus().Create(bonus2)

	testCases := []struct {
		name         string
		bonus        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			bonus: map[string]interface{}{
				"race_id": bonus.RaceId,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.bonus)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodGet, "/private/races/bonuses", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceBonusUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	race := model.TestRace(t)
	skill := model.TestSkill(t)

	_ = store.Race().Create(race)
	_ = store.Skill().Create(skill)

	bonus := model.TestRaceBonus(t)
	bonus.RaceId = race.ID
	bonus.SkillId = skill.ID

	_ = store.RaceBonus().Create(bonus)

	testCases := []struct {
		name         string
		bonus        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": bonus.SkillId,
				"bonus":    bonus.Bonus + 2,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": bonus.SkillId,
				"bonus":    0,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.bonus)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/races/bonuses", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleRaceBonusDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	race := model.TestRace(t)
	skill := model.TestSkill(t)

	_ = store.Race().Create(race)
	_ = store.Skill().Create(skill)

	bonus := model.TestRaceBonus(t)
	bonus.RaceId = race.ID
	bonus.SkillId = skill.ID

	_ = store.RaceBonus().Create(bonus)

	testCases := []struct {
		name         string
		bonus        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": bonus.SkillId,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			bonus: map[string]interface{}{
				"race_id":  bonus.RaceId,
				"skill_id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.bonus)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/races/bonuses", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
