package main

import (
	"fmt"
	"github.com/Get-High-Team/APIs"
	"github.com/Get-High-Team/config"
	"github.com/Get-High-Team/signin"
	"github.com/Get-High-Team/signout"
	"github.com/Get-High-Team/signup"
	"github.com/Get-High-Team/newpost"
	"log"
	"net/http"
	"strings"
	"time"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/signup":
		signup.Handler(w, r)
		return
	case "/signin":
		signin.Handler(w, r)
		return
	case "/signout":
		signout.Handler(w, r)
		return
	case "/newpost":
		newpost.Handler(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/APIs") {
		APIs.Handler(w, r)
		return
	}

	http.ServeFile(w, r, config.Path+r.URL.Path)
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
