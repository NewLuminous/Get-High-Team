package APIs

import (
	"encoding/json"
	"github.com/Get-High-Team/config"
	"io/ioutil"
	"log"
	"net/http"
)

type PostDataClient struct {
	PostId   int `json:"PostId"`
}

type PostRespondData struct {
    Email string `json:"Email"`
    Date string `json:"Date"`
    City string `json:"City"`
    District string `json:"District"`
    Title string `json:"Title"`
    Phone string `json:"Phone"`
    Price int	`json:"Price"`
    Street string `json:"Street"`
    Area int `json:"Area"`
    MaxPeople string `json:"MaxPeople"`
    LiveWithOwner bool `json:"LiveWithOwner"`
    Description string `json:"Description"`
}

func getPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reqData PostDataClient
	err := json.Unmarshal(reqBody, &reqData)
	checkErr(err)

	log.Println(string(reqBody))

	db, err := config.InitDB()
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	defer db.Close()

	var resData PostRespondData

	//get hostel information
	stmt, err := db.Prepare(`
	    SELECT locationid, owner, area, price, description, phone, title, date, livewithowner 
	    FROM hostel 
	    WHERE id = $1;
	`)
	checkErr(err)

	rows, err := stmt.Query(reqData.PostId)
	checkErr(err)

	rows.Next();

	var (
	    locationID int
	    userID int
	)

	rows.Scan(&locationID, &userID, &resData.Area, &resData.Price, &resData.Description, &resData.Phone, &resData.Title, &resData.Date, &resData.LiveWithOwner)

	//get location information

	stmt, err = db.Prepare(`
	    SELECT city, district, street 
	    FROM hostellocation
	    WHERE id = $1;
	`)
	checkErr(err)

	rows, err = stmt.Query(locationID)
	checkErr(err)

	rows.Next();
	rows.Scan(&resData.City, &resData.District, &resData.Street)

	//get email information

	stmt, err = db.Prepare(`
	    SELECT username 
	    FROM users
	    WHERE id = $1;
	`)
	checkErr(err)

	rows, err = stmt.Query(userID)
	checkErr(err)

	rows.Next();
	rows.Scan(&resData.Email)

	res, err := json.Marshal(resData)
	checkErr(err)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}
