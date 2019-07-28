package sql

import (
	"database/sql"
	"log"
)

func CountRow(rows *sql.Rows) (ans int) {
	for rows.Next() {
		err := rows.Scan(&ans)
		if err != nil {
			log.Println(err)
		}
	}
	return ans
}
