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

func TestServer_handleSkillCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

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
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"name":        "Name",
				"icon":        "Icon",
				"category_id": -2,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"name":        "Name",
				"icon":        "Icon",
				"category_id": 2,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
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

			req, _ := http.NewRequest(http.MethodPost, "/private/skills", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSkillCategoryCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		category     map[string]string
		expectedCode int
	}{
		{
			name: "valid",
			category: map[string]string{
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			category: map[string]string{
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.category)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/skill-categories", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSkillsGet(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	rec := httptest.NewRecorder()

	b := &bytes.Buffer{}

	req, _ := http.NewRequest(http.MethodGet, "/skills", b)

	TestSetCookie(req, u, sc)

	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_handleSkillCategoryGet(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	rec := httptest.NewRecorder()

	b := &bytes.Buffer{}

	req, _ := http.NewRequest(http.MethodGet, "/skill-categories", b)

	TestSetCookie(req, u, sc)

	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_handleSkillUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	skill := model.TestSkill(t)
	_ = store.Skill().Create(skill)

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
				"id":          skill.ID,
				"name":        "Name",
				"icon":        "Icon",
				"category_id": skill.CategoryId,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			skill: map[string]interface{}{
				"id":          skill.ID,
				"category_id": category.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":          skill.ID,
				"name":        "Name",
				"icon":        "Icon",
				"category_id": -2,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":          skill.ID,
				"name":        "Name",
				"icon":        "Icon",
				"category_id": 202,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			skill: map[string]interface{}{
				"id":   skill.ID,
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

			req, _ := http.NewRequest(http.MethodPut, "/private/skills", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSkillCategoryUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	category := model.TestSkillCategory(t)
	_ = store.Skill().CreateCategory(category)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		category     map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			category: map[string]interface{}{
				"id":   category.ID,
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			category: map[string]interface{}{
				"id": category.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			category: map[string]interface{}{
				"id":   1000,
				"name": category.Name,
				"icon": "Icon",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.category)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/skill-categories", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSkillDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	skill := model.TestSkill(t)
	_ = store.Skill().Create(skill)

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
				"id": skill.ID,
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

			req, _ := http.NewRequest(http.MethodDelete, "/private/skills", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleSkillCategoriesDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	category := model.TestSkillCategory(t)
	_ = store.Skill().CreateCategory(category)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		category     map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			category: map[string]interface{}{
				"id":   category.ID,
				"name": "Name",
				"icon": "Icon",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			category: map[string]interface{}{
				"id": 1000,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.category)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/skill-categories", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
