package main
import "net/http"
import "fmt"


type UserRequest struct {
    w *http.ResponseWriter
    r *http.Request
}

func (s *UserRequest) init(w *http.ResponseWriter, r *http.Request) {
    s.w = w
    s.r = r
}

func (s *UserRequest) handle_request() {
    var response string

    s.r.ParseForm()
    userid := s.r.Form.Get("userid")
    passwd := s.r.Form.Get("passwd")
    command := s.r.Form.Get("command")

    if userobj := check_user_pass(userid, passwd); userobj != nil {
        response = userobj.process_command(command)
    } else {
        response = "<h1>you are not "+userid+"<br></h1>"
    }
	fmt.Fprintf(*s.w,response)
}
