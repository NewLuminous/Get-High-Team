package APIs

import (
	"github.com/Get-High-Team/config"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	username := ""

	for _, cookie := range r.Cookies() {

		if cookie.Name == "SSID" {
			ssid := cookie.Value

			db, err := config.InitDB()
			if err != nil {
				log.Fatal("Cannot connect to Database", err)
			}
			defer db.Close()

			stmt, err := db.Prepare(`SELECT name FROM users WHERE ssid = $1; `)
			checkErr(err)

			rows, err := stmt.Query(ssid)
			checkErr(err)

			for rows.Next() {
				err := rows.Scan(&username)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(username))
}
