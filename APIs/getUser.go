package APIs

import (
	"github.com/Get-High-Team/config"
	"log"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	ssid := GetSSID(r)

	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT name FROM users WHERE ssid = $1; `)
	checkErr(err)

	rows, err := stmt.Query(ssid)
	checkErr(err)

	username := ""
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Println(err)
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(username))
}
