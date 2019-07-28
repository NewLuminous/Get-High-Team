package session

import (
	"github.com/Get-High-Team/config"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func DelSession(ssid string) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
	UPDATE users
	SET ssid = ''
	WHERE ssid = $1;
    `)
	checkErr(err)

	_, err = stmt.Query(ssid)

	checkErr(err)
}
