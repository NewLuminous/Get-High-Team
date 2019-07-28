package newpost

import (
	"encoding/json"
	"fmt"
	"github.com/Get-High-Team/APIs"
	"github.com/Get-High-Team/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DataClient struct {
	Firstname     string `json:"Firstname"`
	Lastname      string `json:"Lastname"`
	OwnerAddr     string `json:"OwnerAddress"`
	City          string `json:"City"`
	District      string `json:"District"`
	Street        string `json:"Street"`
	Area          int    `json:"Area"`
	MaxPeople     int    `json:"MaxPeople"`
	LiveWithOwner bool   `json:"LiveWithOwner"`
	Price         int    `json:"Price"`
	Description   string `json:"Description"`
	Title         string `json:"Title"`
	Phone         string `json:"Phone"`
}

type RespondData struct {
	Firstname     bool `json:"Firstname"`
	Lastname      bool `json:"OwnerAddress"`
	OwnerAddr     bool `json:"City"`
	City          bool `json:"District"`
	District      bool `json:"District"`
	Street        bool `json:"Street"`
	Area          bool `json:"Area"`
	MaxPeople     bool `json:"MaxPeople"`
	LiveWithOwner bool `json:"LiveWithOwner"`
	Price         bool `json:"Price"`
	Description   bool `json:"Description"`
	Title         bool `json:"Title"`
	Phone         bool `json:"Phone"`
	PostID        int  `json:"PostID"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Println(string(reqBody))

	var reqData DataClient
	err := json.Unmarshal(reqBody, &reqData)
	if err != nil {
		log.Println(err)
	}

	var resData RespondData
	resData.Firstname = true
	resData.Lastname = true
	resData.OwnerAddr = true
	resData.City = true
	resData.District = true
	resData.Street = true
	resData.Area = true
	resData.MaxPeople = true
	resData.LiveWithOwner = true
	resData.Price = true
	resData.Description = true
	resData.Title = true
	resData.PostID = 0

	res, err := json.Marshal(resData)
	if err != nil {
		log.Println(err)
	}

	if !resData.Firstname || !resData.Lastname || !resData.OwnerAddr || !resData.City || !resData.District || !resData.Street || !resData.Area || !resData.MaxPeople || !resData.LiveWithOwner || !resData.Price || !resData.Description || !resData.Title {
		w.Write([]byte(res))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(res)

	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	//add location
	stmt, err := db.Prepare(`
    	    INSERT INTO hostellocation (city, district, street)
    	    VALUES($1, $2, $3) RETURNING id; `)

	if err != nil {
		log.Println(err)
	}

	rows, err := stmt.Query(reqData.City, reqData.District, reqData.Street)
	if err != nil {
		log.Println(err)
		return
	}

	var locationID int
	rows.Next()
	rows.Scan(&locationID)

	log.Println(locationID)
	//find user id
	stmt, err = db.Prepare(`
	    SELECT id FROM users WHERE ssid = $1;
	`)
	if err != nil {
		log.Println(err)
		return
	}

	rows, err = stmt.Query(APIs.GetSSID(r))
	if err != nil {
		log.Println(err)
		return
	}

	rows.Next()

	var usrID int
	rows.Scan(&usrID)

	log.Println(usrID)

	//add hostel
	stmt, err = db.Prepare(`
    	    INSERT INTO hostel (locationid, owner, title, area, price, description, phone, date)
    	    VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`)
	if err != nil {
		log.Println(err)
	}

	rows, err = stmt.Query(locationID, usrID, reqData.Title, reqData.Area, reqData.Price, reqData.Description, reqData.Phone, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("huhuhuhuhu **** ", err)
		return
	}

	rows.Next()
	rows.Scan(&resData.PostID)

	res, err = json.Marshal(resData)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(res))

	fmt.Println("ok")
}
