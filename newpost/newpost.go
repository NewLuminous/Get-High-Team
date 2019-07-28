
package newpost

import (
    "io/ioutil"
    "net/http"
    "log"
    "encoding/json"
)

type DataClient struct {
    Firstname string `Firstname`
    LastName string `LastName`
    OwnerAddr string `OwnerAddr`
    City string `City`
    District string `District`
    Street string `Street`
    Area int `Area`
    MaxPeople string `MaxPeoPle`
    LiveWithOwner bool `LiveWithOwner`
    Description string `Description`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    log.Println(string(reqBody))

    var reqData DataClient
    err := json.Unmarshal(reqBody, &reqData)
    if err != nil {
	log.Println(err)
    }
}
