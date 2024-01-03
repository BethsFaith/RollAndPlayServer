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

func TestServer_handleSystemCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		system       map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			system: map[string]interface{}{
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "valid",
			system: map[string]interface{}{
				"name": "Name",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			system: map[string]interface{}{
				"name": "N",
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			system: map[string]interface{}{
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.system)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/systems", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemGetRaces(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	race := model.TestRace(t)
	_ = store.Race().Create(race)

	_, _ = store.System().AddRace(sys.ID, race.ID)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id": sys.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodGet, "/private/systems/races", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemGetClasses(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	class := model.TestCharacterClass(t)
	_ = store.CharacterClass().Create(class)

	_, _ = store.System().AddCharacterClass(sys.ID, class.ID)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id": sys.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodGet, "/private/systems/classes", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemGetCategories(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	category := model.TestSkillCategory(t)
	_ = store.Skill().CreateCategory(category)

	_, _ = store.System().AddSkillCategory(sys.ID, category.ID)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id": sys.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodGet, "/private/systems/classes", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":   sys.ID,
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":   sys.ID,
				"name": "Nameee",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":   sys.ID,
				"name": "",
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":   sys.ID,
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":   1000,
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/systems", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemBindRace(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	race := model.TestRace(t)
	_ = store.Race().Create(race)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":      sys.ID,
				"race_id": race.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":      1000,
				"race_id": race.ID,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":      sys.ID,
				"race_id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/systems/races", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemBindClass(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	class := model.TestCharacterClass(t)
	_ = store.CharacterClass().Create(class)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":       sys.ID,
				"class_id": class.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":       1000,
				"class_id": class.ID,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":       sys.ID,
				"class_id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/systems/classes", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemBindCategory(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	category := model.TestSkillCategory(t)
	_ = store.Skill().CreateCategory(category)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":                sys.ID,
				"skill_category_id": category.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":                1000,
				"skill_category_id": category.ID,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":                sys.ID,
				"skill_category_id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/systems/skills", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSystemDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	sys := model.TestSystem(t)
	_ = store.System().Create(sys)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		skill        map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			skill: map[string]interface{}{
				"id": sys.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.skill)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/systems", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
