package APIs

import (
	"log"
	"net/http"
)

func GetSSID(r *http.Request) string {
	cookie, err := r.Cookie("SSID")

	if err == http.ErrNoCookie {
		return ""
	} else if err != nil {
		log.Println(err)
	}
	return cookie.Value
}
