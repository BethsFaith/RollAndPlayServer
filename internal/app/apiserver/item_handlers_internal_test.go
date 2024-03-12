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

func TestServer_handleItemCreate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		item         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			item: map[string]interface{}{
				"name": "Name",
				"icon": "Model.path",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no valid",
			item: map[string]interface{}{
				"name": "",
				"icon": "icon.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no valid",
			item: map[string]interface{}{
				"icon": "icon.path",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			item: map[string]interface{}{
				"name": "Name",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.item)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPost, "/private/items", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleItemGet(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	rec := httptest.NewRecorder()

	b := &bytes.Buffer{}

	req, _ := http.NewRequest(http.MethodGet, "/items", b)

	TestSetCookie(req, u, sc)

	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_handleItemUpdate(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()

	u := model.TestUser(t)
	_ = store.User().Create(u)

	Item := model.TestItem(t)
	assert.NoError(t, store.Item().Create(Item))

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		Item         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			Item: map[string]interface{}{
				"id":   Item.ID,
				"name": "NewName",
				"icon": "Newicon.path",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			Item: map[string]interface{}{
				"id":   Item.ID,
				"name": "NewName",
				"icon": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			Item: map[string]interface{}{
				"id":   Item.ID,
				"name": "",
				"icon": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid",
			Item: map[string]interface{}{
				"id":   Item.ID,
				"name": "11",
				"icon": "",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			Item: map[string]interface{}{
				"id":   1000,
				"name": "11",
				"icon": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.Item)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodPut, "/private/items", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleItemDelete(t *testing.T) {
	logger := log.TestLogger()

	store := teststore.New()
	u := model.TestUser(t)
	_ = store.User().Create(u)

	Item := model.TestItem(t)
	_ = store.Item().Create(Item)

	cookieStore, sc := TestCookie()
	s := newServer(store, cookieStore, logger)

	TestAuthUser(s, u)

	testCases := []struct {
		name         string
		Item         map[string]interface{}
		expectedCode int
	}{
		{
			name: "valid",
			Item: map[string]interface{}{
				"id": Item.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "no valid",
			Item: map[string]interface{}{
				"id": -1,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.Item)
			assert.NoError(t, err)

			req, _ := http.NewRequest(http.MethodDelete, "/private/items", b)

			TestSetCookie(req, u, sc)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
