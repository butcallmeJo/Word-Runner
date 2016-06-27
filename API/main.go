package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "os/exec"
)

func main() {
    r := mux.NewRouter()
    r.Handle("/wiki/", FindHandler).Methods("POST") 
    http.ListenAndServe(":8091", handlers.LoggingHandler(os.Stdout, r))
}

var FindHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ 
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers",
        "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    r.ParseForm()
    Start_l := r.FormValue("first")
    End_l := r.FormValue("second")
    if ((len(Start_l) > 0) && (len(End_l) > 0)) {
        fmt.Println("Path find: " + Start_l + " " + End_l)
        out, err := exec.Command("/api/pathfinder", Start_l, End_l).Output()
        if err == nil {
            fmt.Println(out)
            w.Write([]byte(out))
        } else {
            fmt.Println(err)
            w.Write([]byte("Error - pathfinder did not run."))
        }
    } else {
        fmt.Println("Error - missing data.")
        w.Write([]byte("Error - missing data."))
    }
})
