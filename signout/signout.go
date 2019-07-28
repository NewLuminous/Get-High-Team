package signout

import (
	"github.com/Get-High-Team/func/session"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("SSID")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Println(err)
		}
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	session.DelCookie(w, r, "SSID")
}
