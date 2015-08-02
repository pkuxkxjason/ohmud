package main


import (
     "os"
     "fmt"
     "strings"
     "bytes"
)

func check_user_pass(userid string, passwd string) *user {
	//read database to valid the userid and passwd
    file, err := os.Open("userdb.txt")
    defer file.Close()

    if err == nil {
        data := make([]byte, 1024)
        count, _ := file.Read(data)
        if count > 0 {
            fmt.Printf("read %d bytes: %q\n", count, data[:count])
            var b bytes.Buffer
            b.Write(data[:count])
            lines := strings.Split(b.String(),"\n")
            for lineno, line := range lines {
                fmt.Println(lineno)
                tokens := strings.Split(line,":")
                if len(tokens) == 2 {
                    fmt.Printf("%s,%s\n",tokens[0],tokens[1])
                    fmt.Printf("%s,%s\n",userid,passwd)
                    if strings.EqualFold(tokens[0],userid) && strings.EqualFold(strings.TrimSpace(tokens[1]),strings.TrimSpace(passwd)) { 
                        u := find_user(userid)
                        if u == nil {
                            u = &user{userid,"",0}
                            add_user(u)
                        }
                        u.incr_visit_cnt()
                        return u;
                    }
                }
            }
        }
    }
    return nil

}