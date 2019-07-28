package APIs

import (
	"encoding/json"
	"github.com/Get-High-Team/config"
	sqlFunc "github.com/Get-High-Team/func/sql"
	"log"
	"net/http"
)

type RespondData struct {
	Data int `json:"Data"`
}

func getNumberOfPosts(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT COUNT(*) FROM hostel; `)
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)

	var dat RespondData
	dat.Data = sqlFunc.CountRow(rows)

	res, err := json.Marshal(dat)
	checkErr(err)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(res))
}
