
package session

import (
    "net/http"
    "time"
)

func delCookie(w http.ResponseWriter, r *http.Request, name string) {
    cookie, err := r.Cookie(name)
    if err != nil || cookie.Value == "" {
	return
    } else {
	cookie := http.Cookie{Name: name, Path: "/", HttpOnly: true, Expires: time.Now(), MaxAge: -1}
	http.SetCookie(w, &cookie)
    }
}
