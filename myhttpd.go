package main

import "net/http"
import "fmt"
import "log"

type httpHandler struct {
}

func (v httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ur := UserRequest{}
    ur.init(&w, r)
	ur.handle_request()
}

func start_httpd() {
	http.Handle("/", httpHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test123() {
    fmt.Println("hi")
}