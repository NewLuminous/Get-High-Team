package APIs

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var apiName string

	if r.Method == "GET" {
		id, ok := r.URL.Query()["id"]
		if !ok {
			log.Println("URL Param 'id' is missing")
			return
		}

		apiName = id[0]

		switch apiName {
		case "getUser":
			getUser(w, r)
		case "getNumberOfPosts":
			getNumberOfPosts(w, r)
		}
	}
	if r.Method == "POST" {
		apiName = strings.TrimPrefix(r.URL.Path, "/APIs/")

		switch apiName {
		case "getIndexPost":
			getIndexPost(w, r)
		}
	}

	fmt.Println("\nAPIs:", apiName)
}
