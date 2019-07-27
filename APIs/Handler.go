package APIs

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, ok := r.URL.Query()["id"]
		if !ok {
			log.Println("URL Param 'id' is missing")
		}

		fmt.Println(id)

		switch id[0] {
		case "getUser":
			getUser(w, r)
		}
	}
}
