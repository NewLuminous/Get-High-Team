package signin

import (
	//    "crypto/sha256"
	"database/sql"
	"github.com/Get-High-Team/config"
	"log"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
func countRow(rows *sql.Rows) (ans int) {
	for rows.Next() {
		err := rows.Scan(&ans)
		if err != nil {
			log.Println(err)
		}
	}
	return ans
}

func validate(user string, pwd string) bool {
	//check syntax
	if !validateUsername(user) || !validatePassword(pwd) {
		return false
	}

	//encrypte password
	//encryptedPwd := sha256.Sum256([]byte(pwd))
	//pwd = string(encryptedPwd[:])

	//compare to database
	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2;`)
	checkErr(err)

	rows, err := stmt.Query(user, pwd)
	checkErr(err)

	if countRow(rows) == 0 {
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
