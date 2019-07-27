
package session

import (
    "log"
    "github.com/Get-High-Team/config"
)

func checkErr(err error) {
    if err != nil {
	log.Fatal(err)
    }
}

func delSession(ssid string) {
    db, err := config.InitDB()
    if err != nil {
	log.Fatal("Cannot connect to Database", err)
    }
    defer db.Close()

    stmt, err := db.Prepare(`
	UPDATE users
	SET ssid = ''
	WHERE ssid = $1;
    `);
    checkErr(err)

    _, err = stmt.Query(ssid);

    checkErr(err)
}
