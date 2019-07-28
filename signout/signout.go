package signout

import (
	"github.com/Get-High-Team/func/session"
	//"github.com/Get-High-Team/config"
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

	/*
		//Save the SSID to DB
		db, err := config.InitDB()
		if err != nil {
			log.Println("Cannot connect to Database", err)
		}
		defer db.Close()

		stmt, err := db.Prepare(`
		UPDATE users
		SET ssid = $1,
		    lastLogin = $2
		WHERE username = $3
	    `)
		checkErr(err)

		_, err = stmt.Query(newSSID, time.Now().String(), usr)
		checkErr(err)

		fmt.Println("ok")
	    */
	log.Println("signed out")
}
