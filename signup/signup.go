
package signup

import (
    "fmt"
    "net/http"
)

type validation struct {
    usr	    string  `json:username`
    pwd	    string  `json:password`
    name    string  `json:name`
    check   string  `json:check`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    usr := r.FormValue("username");
    pwd := r.FormValue("password");
    name := r.FormValue("name");
    check := r.FormValue("agreeCheck");
    submit := r.FormValue("submit");

    var vld validation
    vld.usr = validateUsername(usr)
    vld.pwd = validatePassword(pwd)
    vld.name = validateName(name)
    vld.check = validateCheck(check)

    if submit == "submit" {
	fmt.Println("submitted")
    }
}

