
package signin

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "fmt"
    "net/http"
    "github.com/Get-High-Team/func/session"
    "github.com/Get-High-Team/config"
    "time"
)

type ClientData struct {
    Usr	    string  `json:"Username"`
    Pwd	    string  `json:"Password"`
    Rem	    bool    `json:"Remenber"`
}

type Validation struct {
    Match   bool    `json:"Match"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("\nsignin")
    bodyReq, _ := ioutil.ReadAll(r.Body)

    var data ClientData
    err := json.Unmarshal(bodyReq, &data)
    if err != nil {
	log.Fatal(err)
	return
    }

    usr := data.Usr
    pwd := data.Pwd

    var vld Validation
    vld.Match = validate(usr, pwd)

    res, err := json.Marshal(vld)
    if err != nil {
	log.Fatal(err)
	return
    }
    w.Header().Set("Access-Control-Allow-Origin", "*")

    if (!vld.Match) {
	fmt.Println("false")
	w.Write(res)
	return
    }

    //Generate SSID and SetCookie
    newSSID := session.GenerateSSID()

    cookie := http.Cookie{Name: "SSID", Value: newSSID, Expires: time.Now().Add(365*24*time.Hour)}
    http.SetCookie(w, &cookie)

    w.Write(res)

    //Save the SSID to DB
    db, err := config.InitDB()
    if err != nil {
	log.Fatal("Cannot connect to Database", err)
    }
    defer db.Close()

    stmt, err := db.Prepare(`
	UPDATE users
	SET ssid = $1,
	    lastLogin = $2
	WHERE username = $3
    `);
    checkErr(err)

    _, err = stmt.Query(newSSID, time.Now().String(), usr);
    checkErr(err)

    fmt.Println("ok")
}
