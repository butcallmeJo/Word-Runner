package main

import (
    "os"
    "fmt"
    "time"
    "sync"
    "regexp"
    "net/http"
    "io/ioutil"
    "strings"
)

func main() {
    Start_l := os.Args[1]
    End_l := os.Args[2]
    reFindLink, err := regexp.Compile("href=\"(/wiki/[^\"/ :#]*)\"")

    if err != nil {
        fmt.Println(err)
        return
    }

    mutex := new(sync.Mutex)
    visited := make(map[string]bool)
    base_link := "https://en.wikipedia.org"
    http_start := []string{"/wiki/" + Start_l}
    link_goal := "/wiki/" + End_l
    done := false
    level := 1
    u := 1
    thread_count := 200
    var solution []string
    queue := make(chan []string, 100000000)
    queue <- http_start

    for i := 0; i < thread_count; i++ {
        go func (id int) {
            defer func() {
                thread_count--
            } ()
            for L := range queue {
                if len(L) > level {
                    level++
                }
                req, err := http.Get(base_link + L[len(L)-1])
                if err == nil {
                    b, _ := ioutil.ReadAll(req.Body)
                    links := reFindLink.FindAllStringSubmatch(string(b), -1)
                    mutex.Lock()
                    if !done {
                        for t := range links {
                            _, v := visited[links[t][1]]
                            if links[t][1] == link_goal {
                                solution = append(L, links[t][1])
                                done = true
                                mutex.Unlock()
                                return
                            }
                            if !v {
                                visited[links[t][1]] = true
                                queue <- append(L, links[t][1])
                            }
                        }
                        if (u < 1000) {
                            u = u+1
                        } else {
                            mutex.Unlock()
                            done = true
                            return
                        }
                    }
                    mutex.Unlock()
                } 
                if done {
                    return
                }
                time.Sleep(10);
            }
            return;
        }(i)
    }
    for !done || thread_count > 0 {
        time.Sleep(1)
    }
    if ( u < 1000) {
        stringfinal := strings.Join(solution, ", ")
        fmt.Println(stringfinal)
    } else {
        fmt.Println("Error - bad data or could not find a path.")
    }
}
