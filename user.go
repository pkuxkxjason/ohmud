package main

type user struct {
	id string
	session string	
    cnt int
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

func (u* user) incr_visit_cnt() {
    u.cnt++
}

