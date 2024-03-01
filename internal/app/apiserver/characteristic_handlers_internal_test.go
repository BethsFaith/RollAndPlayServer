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

func TestServer_handleCharacteristicCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name           string
		characteristic map[string]interface{}
		expectedCode   int
	}{
		{
			name: "valid",
			characteristic: map[string]interface{}{
				"name": "Name",
				"icon": "icon.path",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"name": "",
				"icon": "icon.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"name": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.characteristic)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/characteristics", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleCharacteristicGet(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	rec := httptest.NewRecorder()

	b := &bytes.Buffer{}

	req, _ := http.NewRequest(http.MethodGet, "/characteristics", b)

	TestSetCookie(req, u, sc)

	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_handleCharacteristicUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	c := model.TestCharacteristic(t)
	_ = store.Characteristic().Create(c)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name           string
		characteristic map[string]interface{}
		expectedCode   int
	}{
		{
			name: "valid",
			characteristic: map[string]interface{}{
				"id":   c.ID,
				"name": "NewName",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			characteristic: map[string]interface{}{
				"id":   c.ID,
				"name": "NewName",
				"icon": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"id":   c.ID,
				"name": "",
				"icon": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			characteristic: map[string]interface{}{
				"id":   c.ID,
				"name": "11",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"id":   c.ID + 1,
				"name": "11",
				"icon": "NewIcon.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"id":   c.ID,
				"name": "",
				"icon": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.characteristic)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/characteristics", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleCharacteristicDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	c := model.TestCharacteristic(t)
	_ = store.Characteristic().Create(c)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name           string
		characteristic map[string]interface{}
		expectedCode   int
	}{
		{
			name: "valid",
			characteristic: map[string]interface{}{
				"id": c.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			characteristic: map[string]interface{}{
				"id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.characteristic)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/characteristics", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
