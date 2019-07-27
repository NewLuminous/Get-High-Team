package main

import (
	"fmt"
	"github.com/Get-High-Team/APIs"
	"github.com/Get-High-Team/signin"
	"github.com/Get-High-Team/signup"
	"log"
	"net/http"
	"time"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signup" {
		signup.Handler(w, r)
		return
	}
	if r.URL.Path == "/signin" {
		signin.Handler(w, r)
		return
	}

	if r.URL.Path == "/APIs" {
		APIs.Handler(w, r)
		return
	}

	http.ServeFile(w, r, "/home/hieutm211/workspace/go/src/github.com/Get-High-Team/" + r.URL.Path)
}

func main() {
	server := &http.Server{
		Handler:      http.HandlerFunc(customHandler),
		Addr:         ":1234",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Server is listening on port 1234")
	log.Fatal(server.ListenAndServe())
}
