package APIs

import (
	"encoding/json"
	"github.com/Get-High-Team/config"
	"io/ioutil"
	"log"
	"net/http"
)

type DataClient struct {
	From int `json:"From"`
	To   int `json:"To"`
}

type DataRespond struct {
	Id	int	`json:"Id"`
	Title   string `json:"Title`
	Date    string `json:"Date"`
	Price   int64  `json:"Price"`
	Area    int    `json:"Area"`
	Address string `json: "Address"`
	Image   string `json:"Image"`
}

type DataRespondSlice struct {
	Data []DataRespond `json:"Data"`
}

func getIndexPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reqData DataClient
	err := json.Unmarshal(reqBody, &reqData)
	checkErr(err)

	log.Println(string(reqBody))

	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
	    SELECT hostel.id, title, date, price, area, city, district, street, address 
	    FROM hostel 
	    INNER JOIN hostellocation 
	    ON (hostel.locationId = hostellocation.id)
	    ORDER BY hostel.id DESC
	    LIMIT $1
	    OFFSET $2;
	`)
	checkErr(err)

	rows, err := stmt.Query(reqData.To-reqData.From+1, reqData.From)
	checkErr(err)

	var dataSlice DataRespondSlice
	for rows.Next() {
		var (
			datagram DataRespond
			city     string
			district string
			street   string
			address  string
		)
		err = rows.Scan(&datagram.Id, &datagram.Title, &datagram.Date, &datagram.Price, &datagram.Area, &city, &district, &street, &address)
		datagram.Address = address + ", " + street + ", " + district + ", " + city
		datagram.Image = ""
		dataSlice.Data = append(dataSlice.Data, datagram)
	}

	res, err := json.Marshal(dataSlice)
	checkErr(err)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}
