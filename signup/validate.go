package signup

import (
	"github.com/Get-High-Team/config"
	sqlFunc "github.com/Get-High-Team/func/sql"
	"log"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func validateUsername(user string) string {
	if m, _ := regexp.MatchString("^[a-zA-Z0-9._]+[@][a-z0-9-]+[.][a-z0-9.]+$", user); !m {
		return "not valid"
	}

	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT COUNT(*) FROM users WHERE username = $1; `)
	checkErr(err)

	rows, err := stmt.Query(user)
	checkErr(err)

	if sqlFunc.CountRow(rows) != 0 {
		return "duplicated"
	}

	return "ok"
}

func validatePassword(pwd string) string {
	if len(pwd) < 6 || len(pwd) > 30 {
		return "not valid"
	}
	return "ok"
}

func validateName(name string) string {
	if m, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z ]+[a-zA-Z]$", name); !m {
		return "not valid"
	}
	return "ok"
}
