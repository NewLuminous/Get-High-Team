
package main 

import (
    "fmt"
    "net/http"
    "time"
    "log"
    "github.com/Get-High-Team/signup"
    "github.com/Get-High-Team/signin"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
     if (r.URL.Path == "/signup") {
	signup.Handler(w, r)
	return
     }
     if (r.URL.Path == "/signin") {
	signin.Handler(w, r)
	return
     }
}

func main() {
    server := &http.Server {
	Handler: http.HandlerFunc(customHandler),
	Addr: ":1234",
	ReadTimeout: 15*time.Second,
	WriteTimeout: 15*time.Second,
    }

    fmt.Println("Server is listening on port 1234");
    log.Fatal(server.ListenAndServe())
}
