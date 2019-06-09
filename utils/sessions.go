package utils

import (
	"github.com/lexsalg/goweb/models"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"
)

const (
	cookieName    = "go_session"
	cookieExpires = 24 * 2 * time.Hour
)

var Sessions = struct {
	m map[string]*models.User
	sync.RWMutex
}{m: make(map[string]*models.User)}

func GetUser(r *http.Request) *models.User {
	Sessions.Lock()
	defer Sessions.Unlock()

	uid := getValueCookie(r)
	if user, ok := Sessions.m[uid]; ok {
		return user
	}
	return &models.User{}
}

func SetSession(user *models.User, w http.ResponseWriter) {
	Sessions.Lock()
	defer Sessions.Unlock()

	uid, _ := uuid.NewV4()
	Sessions.m[uid.String()] = user

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   uid.String(),
		Path:    "/",
		//Expires: time.Now().Add(cookieExpires),
	}
	http.SetCookie(w, cookie)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	Sessions.Lock()
	defer Sessions.Unlock()
	delete(Sessions.m, getValueCookie(r))

	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getValueCookie(r *http.Request) string {
	if c, err := r.Cookie(cookieName); err == nil {
		return c.Value
	}
	return ""
}

func IsAuth(r *http.Request) bool {
	return getValueCookie(r) != ""
}
