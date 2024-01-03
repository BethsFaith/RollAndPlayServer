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

func TestServer_handleActionCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		action       map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			action: map[string]interface{}{
				"name":     "Name",
				"icon":     "icon.path",
				"skill_id": 0,
				"points":   10,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"name":     "",
				"icon":     "icon.path",
				"skill_id": 0,
				"points":   10,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"name":   "fddkdk",
				"icon":   "icon.path",
				"points": -10,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"name":     "fddkdk",
				"icon":     "icon.path",
				"skill_id": -10,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"name": "fddkdk",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.action)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/actions", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleActionUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	a := model.TestAction(t)
	_ = store.Action().Create(a)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		action       map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			action: map[string]interface{}{
				"id":   a.ID,
				"name": "NewName",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			action: map[string]interface{}{
				"id":   a.ID,
				"name": "NewName",
				"icon": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"id":   a.ID,
				"name": "",
				"icon": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			action: map[string]interface{}{
				"id":   a.ID,
				"name": "11",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"id":   a.ID + 1,
				"name": "11",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			action: map[string]interface{}{
				"id":       a.ID,
				"name":     "11",
				"icon":     "",
				"skill_id": -10,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.action)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/actions", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleActionDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	a := model.TestAction(t)
	_ = store.Action().Create(a)

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
				"id": a.ID,
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

			req, _ := http.NewRequest(http.MethodDelete, "/private/actions", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
