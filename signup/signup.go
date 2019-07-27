package signup

import (
	//    "crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/Get-High-Team/config"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientData struct {
	Usr  string `json:"Username"`
	Pwd  string `json:"Password"`
	Name string `json:"Name"`
}

type Validation struct {
	Usr  string `json:Username`
	Pwd  string `json:Password`
	Name string `json:Name`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
	bodyReq, _ := ioutil.ReadAll(r.Body)

	var data ClientData
	err := json.Unmarshal(bodyReq, &data)
	if err != nil {
		log.Fatal(err)
		return
	}

	usr := data.Usr
	pwd := data.Pwd
	name := data.Name

	fmt.Println(string(bodyReq))
	fmt.Println(usr)
	fmt.Println(pwd)
	fmt.Println(name)

	var vld Validation
	vld.Usr = validateUsername(usr)
	vld.Pwd = validatePassword(pwd)
	vld.Name = validateName(name)

	res, err := json.Marshal(vld)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(res)

	if vld.Usr != "ok" || vld.Pwd != "ok" || vld.Name != "ok" {
		return
	}

	//encrypt password
	//encryptedPwd := sha256.Sum256([]byte(pwd))
	//pwd = string(encryptedPwd[:])

	//add user to database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}
	defer db.Close()

	sqlStmt := `
    	INSERT INTO users (username, password, name)
    	VALUES($1, $2, $3);`
	_, err = db.Exec(sqlStmt, usr, pwd, name)
	if err != nil {
		log.Fatal(err)
		return
	}
}
