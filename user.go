package main

import "fmt"

type user struct {
	id string
	session string	
    cnt int
    ch chan string
    res_ch chan string

    response_buffer string
}

func (u *user) init(id string) {
    u.id = id
    u.session = ""
    u.cnt = 0
    u.ch = nil
    u.response_buffer = ""
}

var users_list []*user

func find_user(userid string) *user {
    if users_list == nil {
        return nil
    }
    for _, u := range users_list {
        if u.id == userid {
            return u
        }
    }
    return nil
}

func add_user(u *user) bool {
    if users_list == nil {
        users_list = make([]*user,0,10)
    }   
    users_list = append(users_list, u)
    return true
}

func user_routine(u *user) {
    for true {
        cmd := <-u.ch 
        switch cmd {
        case "hello":
            u.response("welcome, " + u.id)
        default:
            u.response("whatever, " + u.id)
        }
    }
}

func (u* user) incr_visit_cnt() {
    u.cnt++
}

func (u* user) init_user() {
    if u.ch == nil {
        u.ch = make(chan string)
        u.res_ch = make(chan string)
        go user_routine(u)
    }
}

func (u* user) process_command(cmd string) string {
    u.ch <- cmd    
    res := <- u.res_ch
    fmt.Println("process_command")
    return res
}

func (u* user) response(res string) {
    u.res_ch <- res
}