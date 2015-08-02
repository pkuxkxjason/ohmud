package main

import "net/http"
import "fmt"
import "log"

type httpHandler struct {
}

func handle_request(ch chan *UserRequest) {
	ur := <-ch
    ur.handle_request()
}

func (v httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ur := UserRequest{}
    ur.init(&w, r)
	ch := make(chan *UserRequest)
	go handle_request(ch)
	ch <- &ur
}

func start_httpd() {
	http.Handle("/", httpHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test123() {
    fmt.Println("hi")
}