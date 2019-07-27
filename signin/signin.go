
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
    w.Header().Set("content-type", "application/json")
    w.Write(res)

    if (!vld.Match) {
	fmt.Println("false")
	return
    }

    //Generate SSID and SetCookie
    newSSID := session.GenerateSSID()
    var cookie1, cookie2 http.Cookie
    if data.Rem {
	cookie1 = http.Cookie{Name: "SSID", Value: newSSID}
	cookie2 = http.Cookie{Name: "usr", Value: usr}
    } else {
	cookie1= http.Cookie{Name: "SSID", Value: newSSID, Expires: time.Now().Add(365*24*time.Hour)}
	cookie2 = http.Cookie{Name: "usr", Value: usr}
    }
    http.SetCookie(w, &cookie1)
    http.SetCookie(w, &cookie2)

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