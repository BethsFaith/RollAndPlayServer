package apiserver

import (
	"RnpServer/internal/app/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
)

func TestCookie() (*sessions.CookieStore, *securecookie.SecureCookie) {
	secretKey := []byte("secret")
	serv := sessions.NewCookieStore(secretKey)
	sc := securecookie.New(secretKey, nil)

	return serv, sc
}

func TestAuthUser(s *server, u *model.User) {
	payload := map[string]string{
		"email":    u.Email,
		"password": u.Password,
	}

	rec := httptest.NewRecorder()
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(payload)
	req := httptest.NewRequest(http.MethodPost, "/sessions", b)

	s.ServeHTTP(rec, req)
}

func TestSetCookie(req *http.Request, u *model.User, sc *securecookie.SecureCookie) {
	cookieValue := map[interface{}]interface{}{
		"user_id": u.ID,
	}
	cookieStr, _ := sc.Encode(sessionName, cookieValue)

	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
}
