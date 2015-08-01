package main

import "net/http"
import "fmt"
import "log"

type httpHandler struct {
}

func handle_request(ch chan *http.ResponseWriter) {
	w := <-ch
	fmt.Fprintf(*w, "Hello, welcome!!! to here")
}

func (v httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch := make(chan *http.ResponseWriter)
	go handle_request(ch)
	ch <- &w

}

func start_httpd() {
	http.Handle("/", httpHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
