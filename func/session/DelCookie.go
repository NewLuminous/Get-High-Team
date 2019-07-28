package session

import (
	"net/http"
	"time"
)

func DelCookie(w http.ResponseWriter, r *http.Request, name string) {
	_, err := r.Cookie(name)
	if err != nil {
		return
	} else {
		cookie := http.Cookie{Name: name, Path: "/", HttpOnly: true, Expires: time.Now(), MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}
