
package signin

import (
    "database/sql"
    "log"
    "regexp"
    "github.com/Get-High-Team/config"
)

func checkErr(err error) {
    if err != nil {
	log.Fatal(err)
    }
}
func countRow(rows *sql.Rows) (ans int) {
    for rows.Next() {
	err := rows.Scan(&ans)
	if err != nil {
	    log.Fatal(err)
	}
    }
    return ans
}

func validate(user string, pwd string) bool {
    if !validateUsername(user) || !validatePassword(pwd) {
	return false
    }

    db, err := config.InitDB()
    if err != nil {
	log.Fatal("Cannot connect to Database", err)
    }
    defer db.Close()

    stmt, err := db.Prepare(`SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2;`)
    checkErr(err)

    rows, err := stmt.Query(user, pwd);
    checkErr(err)

    if (countRow(rows) == 0) {
	return false
    }

    return true
}

func validateUsername(user string) bool {
    if m, _ := regexp.MatchString("^[a-zA-Z0-9._]+[@][a-z0-9-]+[.][a-z0-9.]+$", user); !m {
	return false
    }

    return true
}

func validatePassword(pwd string) bool {
    if len(pwd) < 6 || len(pwd) > 30 {
	return false
    }
    return true
}
